package fasta

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FASTA struct {
	filepath string
}

func New(filepath string) *FASTA {
	return &FASTA{
		filepath: filepath,
	}
}

func (f *FASTA) GetFilePath() string {
	return f.filepath
}

func (f *FASTA) GetHeaders() ([]string, error) {
	headers := make([]string, 0)

	file, err := os.Open(f.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			headers = append(headers, strings.TrimPrefix(line, ">"))
		}
	}
	return headers, nil

}

func (f *FASTA) PrintHeaders() {
	headers, err := f.GetHeaders()
	if err != nil {
		fmt.Println("")
	}
	for _, h := range headers {
		fmt.Println(h)
	}
}

func (f *FASTA) GetGCContent() (map[string]float64, error) {
	counts := make(map[string]int)

	file, err := os.Open(f.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total float64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			continue
		}
		for _, n := range line {
			total++
			switch {
			case n == 'A' || n == 'T':
				counts["AT"]++
			case n == 'C' || n == 'G':
				counts["GC"]++
			}
		}
	}

	gcContent := make(map[string]float64)

	gcContent["AT"] = float64(counts["AT"]) / total
	gcContent["GC"] = float64(counts["GC"]) / total

	return gcContent, nil
}

func (f *FASTA) PrintGCContent() {
	gcContent, err := f.GetGCContent()
	if err != nil {
		fmt.Println("")
	}

	for n, c := range gcContent {
		fmt.Printf("%s: %.2f\n", n, c)
	}
}
