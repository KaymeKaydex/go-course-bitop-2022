package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Введите один аргумент - название файла")
		os.Exit(2)
	}
	resp, err := http.Get("https://api.www.bmstu.ru/members?&offset=60")
	if err != nil { //Если err неравен nil(нулю), то resp.Body точно не пустой
		log.Println("Не получилось совершить запрос")
		os.Exit(2)
	}
	var rowResp []byte
	rowResp, err = ioutil.ReadAll(resp.Body)
	if err != nil { //Если err неравен nil(нулю), то resp.Body точно не пустой
		log.Println("Не получилось достать информацию из запроса")
		os.Exit(2)
	}
	text_request := string(rowResp)
	re := regexp.MustCompile("[0-9]+")
	numbers_request := re.FindAllString(text_request, -1)
	//log.Println(numbers_request) - можно посмотреть числа, которые взялись из входного json
	var sum int = 0
	var composition int = 1
	for _, number_request := range numbers_request {
		number, err := strconv.Atoi(number_request) //переводим строку в число
		if err != nil {
			log.Println("Не удалось перевести в число")
			os.Exit(2)
		}
		sum = sum + number
		composition = composition * number
	}
	sum_str := strconv.Itoa(sum)
	composition_str := strconv.Itoa(composition)
	mydata := []byte(sum_str + " " + composition_str)
	err = ioutil.WriteFile(os.Args[1], mydata, 0777)
	if err != nil {
		log.Println("Не удалось записать в файл")
		log.Println(err)
	}
}

