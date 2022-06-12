package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	log.Println("Starting script")

	resp, err := http.Get("https://vkpay.com/merchants/data/aliexpress.json")
	if err != nil {
		log.Println("Can't create request to VKPAY site", err)
		os.Exit(2)
	}

	var rawResp []byte
	rawResp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	log.Println("Searching numbers in request")
	strResp := string(rawResp)
	re, _ := regexp.Compile(`\d+?`)
	res := re.FindAllString(strResp, -1)

	log.Println("Convertating numbers to int")
	var sum int = 0
	for _, strNumb := range res {
		nmb, err := strconv.Atoi(strNumb)
		if err != nil {
			log.Println("Convertation error")
			os.Exit(2)
		}
		sum += nmb
	}

	log.Println("Creating JSON file")
	strSum := strconv.Itoa(sum)

	jsonSum, jerr := json.Marshal(strSum)
	if jerr != nil {
		log.Println("Can't create JSON file")
		os.Exit(2)
	}
	ioutil.WriteFile("summary.json", jsonSum, os.ModePerm)
	log.Println("Script completed successifull")
}
