package main

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type Flags struct {
	k          bool
	clmnToSort int
	kRaw       string
	n          bool
	r          bool
	u          bool
	M          bool
	b          bool
	c          bool
	h          bool
	doOnce     bool
	printLast  string
}

func scanFlags() *Flags {
	flags := new(Flags)
	pflag.StringVarP(&flags.kRaw, "key=KEYDEF", "k", "", "sort via a key; KEYDEF gives location and type")
	pflag.BoolVarP(&flags.n, "numeric-sort", "n", false, "compare according to string numerical value")
	pflag.BoolVarP(&flags.r, "reverse", "r", false, "reverse the result of comparisons")
	pflag.BoolVarP(&flags.u, "unique", "u", false, "with -c, check for strict ordering;\n  without -c, output only the first of an equal run")
	pflag.BoolVarP(&flags.M, "month-sort", "M", false, "compare (unknown) < 'JAN' < ... < 'DEC'")
	pflag.BoolVarP(&flags.b, "ignore-leading-blanks", "b", false, "ignore leading blanks")
	pflag.BoolVarP(&flags.c, "check", "c", false, "check for sorted input; do not sort")
	pflag.BoolVarP(&flags.h, "human-numeric-sort", "h", false, "compare human readable numbers (e.g., 2K 1G)")
	pflag.Parse()
	kCheck(&flags.k, flags.kRaw, &flags.clmnToSort)
	return flags
}

func kCheck(flagK *bool, kString string, clmnToSort *int) error {
	if kString != "" {
		*flagK = true
		fields := strings.FieldsSeq(kString)
		for number := range fields {
			if kInt, err := strconv.Atoi(number); err != nil {
				return err
			} else {
				*clmnToSort = kInt - 1
				break
			}
		}
	}
	return nil
}
