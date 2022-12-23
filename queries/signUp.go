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

func SignUp(account *models.UserAccount) string {
	fmt.Print("sign up: ")
	url := generate_data.Host + "/auth/sign-up"
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

	log.Println(res.Status)
	cookie := res.Header.Get("Set-Cookie")
	UploadImage(account.UserType, cookie)

	return cookie
}
