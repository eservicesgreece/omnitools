package main

import "regexp"

func parseAccountInfo(input string) []string {
	var ACINFO = regexp.MustCompile(`[0-9]+.[0-9]+`)

	ACObj := ACINFO.FindAllStringSubmatch(input, -1)

	out := make([]string, len(ACObj))
	for i := range out {
		out[i] = ACObj[i][0]
	}

	return out

}
