package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func accountinfo(url, user, pass, info string) string {
	var out string
	response, err := http.Get(url + "api/send_sms.cfm?uname=" + user + "&pword=" + pass + "&sacc=" + user + "&action=balance")
	if err != nil {
		log.Println("Error while downloading", url, ":", err)
	}

	// Verify if the response was ok
	if response.StatusCode != http.StatusOK {
		log.Printf("Server return non-200 status: %v\n", response.Status)
	}

	defer response.Body.Close()

	message, _ := ioutil.ReadAll(response.Body)

	parsed := parseAccountInfo(string(message))

	if len(parsed) > 0 {
		switch info {
		case "credit":
			out = parsed[0]
		case "tax":
			out = parsed[2]
		default:
			for i, aObj := range parsed {
				if i == len(parsed)-1 {
					out += aObj
				} else {
					out += aObj + ","
				}
			}
		}
	} else {
		fmt.Println("Improper Result")
		fmt.Println(string(message))
	}

	return out
}
