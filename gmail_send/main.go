package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("Subject", "Hello!")
	m.SetHeader("From", "tumbalulangan2@gmail.com")
    m.SetHeader("To", "mamangrakai@gmail.com")
	m.SetBody("text/html", "<h1>Hello <b>Rakai</b></h1>")

	// Send the email to Bob
	d := gomail.NewDialer("smtp.gmail.com", 587, "tumbalulangan2@gmail.com", "eforkhypcsuftjmh")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")
}
