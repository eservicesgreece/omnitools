package main

import (
	"log"
	"net/http"

	"unicode/utf8"
)

//https://www.omnivoice.eu/api/send_sms.cfm?uname=&pword=&sacc=&message=&destid=&action=test
func sendSMS(url, message, user, pass, cid, sender, sendtype, recepient string) bool {
	sendtype = "test"

	response, err := http.Get(url + "api/send_sms.cfm?uname=" + user + "&pword=" + pass + "&sacc=" + user + "&cid=" + cid + "&message=" + message + "&destid=" + recepient + "&action=" + sendtype)
	if err != nil {
		log.Println("Error while downloading", url, ":", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("Server return non-200 status: %v\n", response.Status)
	}

	defer response.Body.Close()

	return true
}

func checkNumbers(numbers []string) bool {
	for num := range numbers {
		if numbers[num][:2] == "30" {
			println("ok")
		}
	}
	return true
}

func checkMessage(message string) bool {
	var mlen = false
	var menc = false

	if len(message) <= 159 {
		mlen = true
	}

	if utf8.Valid([]byte(message)) {
		menc = true
	}
	return mlen && menc
}
