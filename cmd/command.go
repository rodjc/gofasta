package cmd

import (
	"flag"
	"fmt"
	"os"
	"rodjc/bioinfotools/gofasta/fasta"
)

type CmdFlags struct {
	input     string
	headers   bool
	gcContent bool
}

func NewCmdFlags() *CmdFlags {
	cf := new(CmdFlags)

	flag.StringVar(&cf.input, "f", "", "Path of the FASTA file (required)")
	flag.BoolVar(&cf.headers, "hs", false, "Print headers of the FASTA file")
	flag.BoolVar(&cf.gcContent, "gc", false, "Print GC-content of the FASTA file")

	flag.Usage = func() {
		fmt.Println("Usage: gofasta -f <filepath> [options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}

	flag.Parse()

	return cf
}

func (cf *CmdFlags) Execute() {
	if cf.input == "" {
		fmt.Println("Error: Please provide a FASTA file path with the -f flag.")
		flag.Usage()
		return
	}

	if _, err := os.Stat(cf.input); os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist.\n", cf.input)
		return
	}

	if cf.headers && cf.gcContent {
		fmt.Println("Error: Please specify only one of -hs or -gc.")
		return
	}

	fasta := fasta.New(cf.input)

	switch {
	case cf.headers:
		fmt.Println("Printing FASTA headers...")
		fasta.PrintHeaders()
	case cf.gcContent:
		fmt.Println("Calculating GC content...")
		fasta.PrintGCContent()
	default:
		fmt.Println("Error: Please specify an action (e.g., -hs or -gc).")
		flag.Usage()
	}
}
