package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseEmailServiceApiPath string = "http://18.158.138.59:8080"

type Request struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Type    int    `json:"type"`
	Email   string `json:"email"`
}

func SendEmail(fullName string, templateType int, email string) {
	body := &Request{
		Name:    fullName,
		Surname: "",
		Type:    templateType,
		Email:   email,
	}

	requestPath := BaseEmailServiceApiPath + "/v1/email"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, err := http.NewRequest("POST", requestPath, buf)
	fmt.Println(err)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return
	}

	defer res.Body.Close()
}
