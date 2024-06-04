package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BencodeDecoder is a struct that holds the bencoded data and the current position in the data
type BencodeDecoder struct {
	data string
	pos  int
}

// decode is the main function that determines the type of the next bencoded value and decodes it.
func (bd *BencodeDecoder) decode() (interface{}, error) {
	if bd.pos >= len(bd.data) {
		return nil, errors.New("end of data")
	}

	switch bd.data[bd.pos] {
	case 'i':
		return bd.decodeInteger()
	case 'l':
		return bd.decodeList()
	case 'd':
		return bd.decodeDictionary()
	default:
		if bd.data[bd.pos] >= '0' && bd.data[bd.pos] <= '9' {
			return bd.decodeString()
		}
		return nil, fmt.Errorf("invalid character at position %d", bd.pos)
	}
}

func (bd *BencodeDecoder) decodeInteger() (int, error) {
	if bd.data[bd.pos] != 'i' {
		return 0, errors.New("expected 'i' at the start of an integer")
	}
	endPos := strings.Index(bd.data[bd.pos:], "e")
	if endPos == -1 {
		return 0, errors.New("unterminated integer")
	}
	endPos += bd.pos
	numStr := bd.data[bd.pos+1 : endPos]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	bd.pos = endPos + 1
	return num, nil
}

func (bd *BencodeDecoder) decodeString() (string, error) {
	colonPos := strings.Index(bd.data[bd.pos:], ":")
	if colonPos == -1 {
		return "", errors.New("invalid string length")
	}
	colonPos += bd.pos
	length, err := strconv.Atoi(bd.data[bd.pos:colonPos])
	if err != nil {
		return "", err
	}
	start := colonPos + 1
	end := start + length
	if end > len(bd.data) {
		return "", errors.New("string extends past end of data")
	}
	str := bd.data[start:end]
	bd.pos = end
	return str, nil
}

func (bd *BencodeDecoder) decodeList() ([]interface{}, error) {
	if bd.data[bd.pos] != 'l' {
		return nil, errors.New("expected 'l' at the start of a list")
	}
	bd.pos++
	var list []interface{}
	for bd.pos < len(bd.data) && bd.data[bd.pos] != 'e' {
		elem, err := bd.decode()
		if err != nil {
			return nil, err
		}
		list = append(list, elem)
	}
	if bd.pos >= len(bd.data) || bd.data[bd.pos] != 'e' {
		return nil, errors.New("unterminated list")
	}
	bd.pos++
	return list, nil
}

func (bd *BencodeDecoder) decodeDictionary() (map[string]interface{}, error) {
	if bd.data[bd.pos] != 'd' {
		return nil, errors.New("expected 'd' at the start of a dictionary")
	}
	bd.pos++
	dict := make(map[string]interface{})
	for bd.pos < len(bd.data) && bd.data[bd.pos] != 'e' {
		key, err := bd.decodeString()
		if err != nil {
			return nil, err
		}
		value, err := bd.decode()
		if err != nil {
			return nil, err
		}
		dict[key] = value
	}
	if bd.pos >= len(bd.data) || bd.data[bd.pos] != 'e' {
		return nil, errors.New("unterminated dictionary")
	}
	bd.pos++
	return dict, nil
}

// readFile reads the contents of a file and returns it as a string.
func readFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// saveToJSON saves the given data to a JSON file.
func saveToJSON(data interface{}, filepath string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path_to_torrent_file> <output_json_file>")
		return
	}

	filepath := os.Args[1]
	outputJSON := os.Args[2]
	data, err := readFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	decoder := BencodeDecoder{data: data}
	result, err := decoder.decode()
	if err != nil {
		fmt.Println("Error decoding bencoded data:", err)
		return
	}

	err = saveToJSON(result, outputJSON)
	if err != nil {
		fmt.Println("Error saving to JSON:", err)
		return
	}

	fmt.Printf("Decoded data saved to %s\n", outputJSON)
}
