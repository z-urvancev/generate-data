package queries

import (
	"encoding/json"
	"fmt"
	generate_data "generateScript"
	"generateScript/models"
	"log"
	"net/http"
	"strings"
)

func CreateVacancy(vacancy *models.Vacancy, cookie string) {
	fmt.Print("create vacancy: ")
	url := generate_data.Host + "/api/vacancy"
	method := "POST"

	vacancyJson, err := json.Marshal(vacancy)
	if err != nil {
		log.Println(err)
		return
	}
	payload := strings.NewReader(string(vacancyJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println(res.Status)
	defer res.Body.Close()

}
