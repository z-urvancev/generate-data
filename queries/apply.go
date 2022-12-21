package queries

import (
	"fmt"
	generate_data "generateScript"
	"log"
	"net/http"
	"strings"
)

func Apply(vacancyId, resumeId uint, cookie string) {
	url := fmt.Sprintf("%s/api/vacancy/apply/%d", generate_data.Host, resumeId)
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
    "id": %d
}`, vacancyId))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	log.Println(res.Status)
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
