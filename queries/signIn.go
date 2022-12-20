package queries

import (
	"encoding/json"
	"generateScript/models"
	"log"
	"net/http"
	"strings"
)

func SignUp(account *models.UserAccount) string {
	url := "http://localhost:8080/auth/sign-up"
	method := "POST"
	user, err := json.Marshal(account)
	if err != nil {
		log.Println(err)
		return ""
	}

	payload := strings.NewReader(string(user))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	return res.Header.Get("Set-Cookie")
}
