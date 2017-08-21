package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "wensongluo@qq.com")
	m.SetHeader("To", "89066570@qq.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>World</b> and <i>aha</i>!")

	d := gomail.NewDialer("smtp.qq.com", 587, "wensongluo@qq.com", "vgpyrdgzpwphhbeb")

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
