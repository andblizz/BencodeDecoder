package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDecodeInteger(t *testing.T) {
	decoder := BencodeDecoder{data: "i123e"}
	result, err := decoder.decode()
	if err != nil {
		t.Fatalf("Failed to decode integer: %v", err)
	}
	if result != 123 {
		t.Fatalf("Expected 123, got %v", result)
	}
}

func TestDecodeString(t *testing.T) {
	decoder := BencodeDecoder{data: "4:spam"}
	result, err := decoder.decode()
	if err != nil {
		t.Fatalf("Failed to decode string: %v", err)
	}
	if result != "spam" {
		t.Fatalf("Expected 'spam', got %v", result)
	}
}

func TestDecodeList(t *testing.T) {
	decoder := BencodeDecoder{data: "l4:spam4:eggse"}
	result, err := decoder.decode()
	if err != nil {
		t.Fatalf("Failed to decode list: %v", err)
	}
	expected := []interface{}{"spam", "eggs"}
	if !compareSlices(result.([]interface{}), expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestDecodeDictionary(t *testing.T) {
	decoder := BencodeDecoder{data: "d3:cow3:moo4:spam4:eggse"}
	result, err := decoder.decode()
	if err != nil {
		t.Fatalf("Failed to decode dictionary: %v", err)
	}
	expected := map[string]interface{}{
		"cow":  "moo",
		"spam": "eggs",
	}
	if !compareMaps(result.(map[string]interface{}), expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestSaveToJSON(t *testing.T) {
	data := map[string]interface{}{
		"cow":  "moo",
		"spam": "eggs",
	}
	filepath := "test_output.json"
	err := saveToJSON(data, filepath)
	if err != nil {
		t.Fatalf("Failed to save to JSON: %v", err)
	}
	defer os.Remove(filepath)

	fileData, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(fileData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if !compareMaps(result, data) {
		t.Fatalf("Expected %v, got %v", data, result)
	}
}

func compareSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
