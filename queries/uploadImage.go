package queries

import (
	"bytes"
	"fmt"
	generate_data "generateScript"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadImage(userType, cookie string) {
	fmt.Print("upload image: ")
	url := fmt.Sprintf("%s/api/user/image", generate_data.Host)
	method := "POST"
	imagePath := fmt.Sprintf("%s_images/%d.jpg", userType, 1+rand.Int()%10)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(imagePath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("avatar", filepath.Base(imagePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		log.Println("err with copy file:", errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		log.Println("close writer:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println("request err: ", err)
		return
	}
	req.Header.Add("Cookie", cookie)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Println("request do err: ", err)
		return
	}
	log.Println(res.Status)
	defer res.Body.Close()
}
