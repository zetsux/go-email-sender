package main

import (
	"bytes"
	"fmt"
	"net/mail"
	"os"
	"text/template"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

const (
	DEFAULT_EMAIL_FILE_NAME          = "email_template.html"
	DEFAULT_EMAIL_TEMPLATE_DIRECTORY = "template"
)

type EmailConfig struct {
	Host         string `mapstructure:"SMTP_HOST"`
	Port         int    `mapstructure:"SMTP_PORT"`
	SenderName   string `mapstructure:"SMTP_SENDER_NAME"`
	AuthEmail    string `mapstructure:"SMTP_AUTH_EMAIL"`
	AuthPassword string `mapstructure:"SMTP_AUTH_PASSWORD"`
}

func NewEmailConfig() (*EmailConfig, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config EmailConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SendEmail(target string, subject string, body string, fileName string) error {
	config, err := NewEmailConfig()
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SenderName)
	mailer.SetHeader("To", target)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	if fileName != "" {
		mailer.Attach(fileName)
	}

	dialer := gomail.NewDialer(
		config.Host,
		config.Port,
		config.AuthEmail,
		config.AuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func MakeEmailBody(name string) (string, error) {
	readHtml, err := os.ReadFile(
		fmt.Sprintf("%s/%s", DEFAULT_EMAIL_TEMPLATE_DIRECTORY, DEFAULT_EMAIL_FILE_NAME),
	)
	if err != nil {
		return "", err
	}

	data := struct {
		Name string
	}{
		Name: name,
	}

	template, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return "", err
	}

	var mailRes bytes.Buffer
	err = template.Execute(&mailRes, data)
	if err != nil {
		return "", err
	}

	return mailRes.String(), nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return (err == nil)
}
