# GoFASTA

## Index
1. [Overview](#overview)  
2. [Features](#features)  
3. [Installation](#installation)  
   - [Linux](#linux)  
   - [macOS](#macos)  
   - [Windows](#windows)  
4. [Usage](#usage)  
   - [Basic Command](#basic-command)  
   - [Options](#options)  
   - [Examples](#examples)  
5. [License](#license)  


## Overview
**GoFASTA** is a command-line tool for analyzing FASTA files. It supports extracting headers and calculating the GC content of sequences in a FASTA file.

## Features
- **Print Headers**: Extract and display all headers from a FASTA file.
- **Calculate GC Content**: Compute the GC-content of sequences in the FASTA file.


## Installation

Download the appropriate binary for your operating system from the [Releases](https://github.com/rodjc/gofasta/releases) page.

### Linux
```bash
wget https://github.com/rodjc/gofasta/releases/download/v1.0.0/gofasta-linux
chmod +x gofasta-linux
sudo mv gofasta-linux /usr/local/bin/gofasta
```

### macOS
```bash
wget https://github.com/rodjc/gofasta/releases/download/v1.0.0/gofasta-macos
chmod +x gofasta-macos
sudo mv gofasta-macos /usr/local/bin/gofasta
```

### Windows
Download `gofasta-windows.exe` and place it in a folder in your `PATH`.


## Usage

### Basic Command
```bash
gofasta -f <file_path> [options]
```

### Options
- `-f <file_path>`: Path to the FASTA file (required).
- `-hs`: Print the headers from the FASTA file.
- `-gc`: Calculate and display the GC content.

### Examples
1. **Print Headers**:
   ```bash
   gofasta -f data/example.fa -hs
   ```

2. **Calculate GC Content**:
   ```bash
   gofasta -f data/example.fa -gc
   ```

## License
This tool is open-source and available under the [Apache License 2.0](LICENSE).