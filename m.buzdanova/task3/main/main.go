package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*type HabrResponse struct {
	Items []struct {
		Title      string   `json:"title"`
		ItemUrl    string   `json:"itemUrl"`
		ImageUrl   string   `json:"imageUrl"`
		Properties []string `json:"properties"`
	} `json:"items"`
}*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Введите аргумент: название файла")
		os.Exit(2)
	}
	log.Println("Выходной файл будет записан в \"./output.json\"")
	log.Println("Делаю запрос")
	begin := time.Now()
	resp, err := http.Get("https://habr.com/kek/v2/inset/courses?hubs=hosting%2Cpay_system%2Cs_admin%2Cbilling%2Cbrowsers%2Cprogramming%2Creactjs%2Cpopular_science%2Castronomy%2Cpostgresql&tags=vps%2Fvds%2C%D0%B2%D1%8B%D0%B4%D0%B5%D0%BB%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9%20%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80%2Cswift%2Cvisa%2Cmastercard%2C%D0%BF%D0%BE%D0%B8%D1%81%D0%BA%D0%B2%D0%BF%D1%81%2CBitB%2C%D0%B1%D1%80%D0%B0%D1%83%D0%B7%D0%B5%D1%80%20%D0%B2%D0%BD%D1%83%D1%82%D1%80%D0%B8%20%D0%B1%D1%80%D0%B0%D1%83%D0%B7%D0%B5%D1%80%D0%B0%2CURL%2C%D0%B1%D0%BB%D0%BE%D0%BA%D0%B8%D1%80%D0%BE%D0%B2%D1%89%D0%B8%D0%BA%20%D1%80%D0%B5%D0%BA%D0%BB%D0%B0%D0%BC%D1%8B&hl=ru")
	if err != nil {
		log.Println("Ошибка совершения запроса ", err)
		os.Exit(2)
	} else {
		log.Println("Запрос совершен успешно")
		log.Println("Запрос длился: ", time.Since(begin))
	}

	var rawResp []byte
	rawResp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка считывания файла ", err)
		os.Exit(2)
	}

	/*var habrResp HabrResponse

	err = json.Unmarshal(rawResp, &habrResp)
	if err != nil {
		log.Println("error ", err)
		os.Exit(2)
	}*/
	var sum, mult int32 = 0, 1
	for _, value := range string(rawResp) {
		if (value >= 0) && (value <= 100000) {
			sum += value
			mult *= value
		}
	}
	result := fmt.Sprintf("Сумма цен: %v, Произведение цен: %v", sum, mult)
	err = ioutil.WriteFile(os.Args[1], []byte(result), os.FileMode(0755))
	if err != nil {
		log.Println("Не удалось записать в json-файл ", err)
		os.Exit(2)
	} else {
		log.Println("Скрипт выполнен успешно")
	}

}
