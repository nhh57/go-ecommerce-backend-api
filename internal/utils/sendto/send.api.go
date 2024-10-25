package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaByAPI(otp string, email string, purpose string) error {
	// URL API
	fmt.Println("SendEmailToJavaByAPI")
	postURL := "http://10.56.66.54:8082/email/send_text"
	//Data Json
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP is " + otp,
		Subject:     "VERIFY OTP " + purpose,
		Attachment:  "path/to/email",
	}

	//convert struct to json
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	// create a request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	//PUT header
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	fmt.Sprintln("Response status:: ", res.Status)
	return nil
}
