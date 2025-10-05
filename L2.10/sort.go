package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sort(input *input) error {

	scnrs := make([]*bufio.Scanner, len(input.chunks))
	files := make([]*os.File, len(input.chunks))
	lines := make([]string, len(input.chunks))
	alive := make([]bool, len(input.chunks))

	if err := scanChunks(scnrs, files, lines, alive, input.chunks); err != nil {
		return err
	}

	mainSort(scnrs, lines, alive, input.flags)
	return cleanup(files, input.chunks)

}

func scanChunks(scnrs []*bufio.Scanner, files []*os.File, lines []string, alive []bool, chunks []string) error {
	for i, file := range chunks {
		currentFile, err := os.Open(file)
		if err != nil {
			return err
		}
		files[i] = currentFile
		scnrs[i] = bufio.NewScanner(currentFile)
		if scnrs[i].Scan() {
			lines[i] = scnrs[i].Text()
			alive[i] = true
		} else {
			lines[i] = ""
			alive[i] = false
		}
	}
	return nil
}

func mainSort(scnrs []*bufio.Scanner, lines []string, alive []bool, flags *Flags) {
	var prevLine string
	for {
		minIdx := -1
		for currentIdx := range lines {
			if !alive[currentIdx] {
				continue
			}
			minIdx = chooseMinIndex(minIdx, currentIdx, lines, alive, flags)
		}
		if minIdx == -1 { // input file is sorted, all file chunks at EOF
			break
		}
		if flags.u {
			if flags.n {
				if _, err := strconv.Atoi(lines[minIdx]); err == nil {
					prevLine = printUnique(lines[minIdx], prevLine)
				}
			} else {
				prevLine = printUnique(lines[minIdx], prevLine)
			}
		} else {
			fmt.Println(lines[minIdx])
		}
		if scnrs[minIdx].Scan() {
			lines[minIdx] = scnrs[minIdx].Text()
			alive[minIdx] = true
		} else {
			lines[minIdx] = ""
			alive[minIdx] = false
		}
	}
	if flags.printLast != "" {
		fmt.Println(flags.printLast)
	}
}

func chooseMinIndex(minIdx int, currentIdx int, lines []string, alive []bool, flags *Flags) int {
	if minIdx == -1 {
		return currentIdx
	}
	if !alive[minIdx] {
		return currentIdx
	}
	if compareLines(lines[minIdx], lines[currentIdx], flags) > 0 {
		return currentIdx
	}
	return minIdx
}

func compareLines(line1, line2 string, flags *Flags) int {
	var cmpRes int
	if flags.n {
		cmpRes = numericCompare(line1, line2)
	} else {
		cmpRes = stringCompare(line1, line2)
	}
	if flags.r {
		cmpRes = -cmpRes
	}
	return cmpRes
}

func numericCompare(line1 string, line2 string) int { // strings go up and the bigger the int the lower it goes
	intLine1, err1 := strconv.Atoi(line1)
	intLine2, err2 := strconv.Atoi(line2)
	if err1 != nil && err2 != nil {
		return stringCompare(line1, line2)
	}
	if intLine1 < intLine2 || (err1 != nil && err2 == nil) { // str < int
		return -1
	}
	if intLine1 > intLine2 || (err1 == nil && err2 != nil) { // int > str
		return 1
	}
	return 0 // int == int
}

func stringCompare(line1 string, line2 string) int {
	if line1 < line2 {
		return -1
	} else if line1 > line2 {
		return 1
	}
	return 0
}

func printUnique(line string, prevLine string) string {
	if line != prevLine {
		fmt.Println(line)
	}
	return line
}

func cleanup(files []*os.File, chunks []string) error {
	for _, file := range files {
		if err := file.Close(); err != nil {
			return err
		}
	}
	if err := deleteChunks(chunks); err != nil {
		return err
	}
	return nil
}

func deleteChunks(chunks []string) error {
	for _, chunk := range chunks {
		if err := os.Remove(chunk); err != nil {
			return err
		}
	}
	return nil
}

func sortChunk(lines []string, flags *Flags) {
	checkRNU(lines, flags)
	slices.SortStableFunc(lines, func(line1 string, line2 string) int {
		if flags.k {
			fields1 := strings.Fields(line1)
			fields2 := strings.Fields(line2)
			var column1, column2 string
			if len(fields1) > flags.clmnToSort {
				column1 = fields1[flags.clmnToSort]
			} else {
				column1 = ""
			}
			if len(fields2) > flags.clmnToSort {
				column2 = fields2[flags.clmnToSort]
			} else {
				column2 = ""
			}
			if res := compareLines(column1, column2, flags); res != 0 {
				return res
			}
		}
		return compareLines(line1, line2, flags)
	})
}

func checkRNU(lines []string, flags *Flags) {
	if flags.n && flags.u && !flags.doOnce {
		for _, line := range lines {
			if _, err := strconv.Atoi(line); err != nil {
				flags.doOnce = true
				if !flags.r {
					fmt.Println(line)
				} else {
					flags.printLast = line
				}
				break
			}
		}
	}
}
