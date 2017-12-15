package main

import (
	"regexp"
	"strings"
)

func parseAccountInfo(input string) []string {
	var ACINFO = regexp.MustCompile(`[0-9]+.[0-9]+`)
	var out []string
	out = make([]string, 1)

	switch ercheck := input; {
	case strings.Contains(ercheck, "not enough credits"):
		out[0] = "No Credit"

	case strings.Contains(ercheck, "Authentication error"):
		out[0] = "Authentication Error"

	default:
		ACObj := ACINFO.FindAllStringSubmatch(input, -1)

		out = make([]string, len(ACObj))
		for i := range out {
			out[i] = ACObj[i][0]
		}
	}

	return out
}
