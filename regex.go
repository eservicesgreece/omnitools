package main

import (
	"regexp"
	"strings"
)

func parseAccountInfo(input string) []string {
	var ACINFO = regexp.MustCompile(`[0-9]+.[0-9]+`)
	var out []string

	if strings.Contains(input, "not enough credits") {
		out = make([]string, 1)
		out[0] = "No Credit"
	} else {

		ACObj := ACINFO.FindAllStringSubmatch(input, -1)

		out = make([]string, len(ACObj))
		for i := range out {
			out[i] = ACObj[i][0]
		}
	}

	return out
}
