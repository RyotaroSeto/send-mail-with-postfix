package main

import (
	"log"
	"net/smtp"
)

func main() {
	// SMTPサーバーの設定
	// smtpHost := "127.0.0.1"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// smtpUsername := "username" // ご利用のSMTPサーバーの認証ユーザー名
	// smtpPassword := "password" // ご利用のSMTPサーバーの認証パスワード
	// auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	auth := smtp.PlainAuth("", "myusername", "mypassword", smtpHost)

	// メールの内容
	from := "from@example.com"
	to := []string{"hahahahawaiwai321@gmail.com"}
	msg := []byte("To: hahahahawaiwai321@gmail.com\r\n" +
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
