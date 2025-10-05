package main

import (
	"bufio"
	"os"

	"github.com/spf13/pflag"
)

type input struct {
	chunks []string
	flags  *Flags
}

func processInput() (*input, error) {

	inp := &input{flags: scanFlags()}
	var scanner *bufio.Scanner

	files := pflag.Args()
	if len(files) > 0 {
		file, err := os.Open(files[0])
		if err != nil {
			return inp, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if err := splitToChunks(scanner, inp); err != nil {
		return inp, err
	}

	return inp, nil

}

func splitToChunks(scanner *bufio.Scanner, inp *input) error {
	var chunk []string
	i := 0
	for scanner.Scan() {
		chunk = append(chunk, scanner.Text()) // add line to a chunk
		i++
		if i == 100 { // if chunk is full, sort it and save it's name to input's schunks slice
			sortChunk(chunk, inp.flags)
			inp.chunks = append(inp.chunks, saveChunk(chunk))
			chunk = chunk[:0]
			i = 0
		}
	}
	if len(chunk) > 0 { // if the last chunk has less than 100 lines, it will be sorted and saved here, after the cycle
		sortChunk(chunk, inp.flags)
		inp.chunks = append(inp.chunks, saveChunk(chunk))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// saveChunk saves sorted chunk to a temp file
func saveChunk(chunk []string) (chankName string) {
	tempFile, _ := os.CreateTemp(".", "chunk_")
	defer tempFile.Close()
	writer := bufio.NewWriter(tempFile)
	for _, line := range chunk {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	return tempFile.Name()
}
