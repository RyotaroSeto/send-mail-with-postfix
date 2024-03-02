package main

import (
	"log"
	"net/smtp"
)

func main() {
	// SMTPサーバーの設定
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "myusername"
	smtpPassword := "mypassword"
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// メールの内容
	from := "from@example.com"
	to := []string{"example@gmail.com"}
	msg := []byte("To: example@gmail.com\r\n" +
		"Subject: Test email\r\n" +
		"\r\n" +
		"This is a test email.")

	// メールの送信
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mail sent successfully")
}
