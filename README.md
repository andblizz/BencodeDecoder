# Bencode Decoder

## Overview

The Bencode Decoder is a simple Go program that decodes bencoded data, commonly used in torrent files, and saves the output in a JSON file. This tool is useful for extracting and reading the information contained in torrent files.

## Features

- Decodes bencoded integers, strings, lists, and dictionaries.
- Reads data from a specified torrent file.
- Outputs the decoded data to a JSON file.
- Includes unit tests to ensure the accuracy of the decoding process.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

### Clone the repository

```sh
git clone https://github.com/yourusername/bencode-decoder.git
cd bencode-decoder
```

### Initialize the project

```sh
go mod init bencode-decoder
```

## Usage

### Running the Program

To run the Bencode Decoder, use the following command:

```sh
go run main.go <path_to_torrent_file> <output_json_file>
```

#### Example

```sh
go run main.go sample.torrent output.json
```

### Running Tests

To run the unit tests, use the following command:

```sh
go test -v
```


## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.
