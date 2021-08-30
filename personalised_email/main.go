package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"sync"
)

type User struct {
	FullName string
	Name     string
	Email    string
}

type Message struct {
	Subject  string
	HtmlFile string `yaml:"html_file"`
}

func main() {
	_ = godotenv.Load()
	sendMessage := setupSMTP()
	users := getReceivers()

	var wg sync.WaitGroup
	wg.Add(len(users))
	for _, user := range users {
		go func(user User) {
			defer wg.Done()
			msg := getMessage(user)
			if err := sendMessage(user.Email, msg); err != nil {
				log.Panic(err)
			}
			log.Println("Sent to " + user.Name + " (" + user.Email + ")")
		}(user)
	}

	wg.Wait()
	log.Println("All emails sent!")
}

func getMessage(user User) []byte {
	var msg Message
	unmarshalYaml("message.yaml", &msg)
	html, err := ioutil.ReadFile(msg.HtmlFile)
	if err != nil {
		log.Panic(err)
	}
	greet := "Hello " + user.Name + ","
	return []byte(fmt.Sprintf(
		"MIME-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\";\r\n"+
			"From :%s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n%s%s\r\n",
		os.Getenv("FROM"),
		user.Email,
		msg.Subject,
		greet,
		html,
	))
}

func setupSMTP() func(to string, msg []byte) error {
	// Load env variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("FROM")
	serverName := smtpHost + ":" + smtpPort
	// Setup SMTP
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	return func(to string, msg []byte) error {
		return smtp.SendMail(serverName, auth, from, []string{to}, msg)
	}
}

func getReceivers() []User {
	if os.Getenv("TEST") != "false" {
		email := os.Getenv("SMTP_USERNAME")
		return []User{
			{"full name", "tester", email},
		}
	}
	var users []User
	unmarshalYaml("users.yaml", &users)
	return users
}

func unmarshalYaml(fName string, i interface{}) {
	f, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Panic(err)
	}
	if err := yaml.Unmarshal(f, i); err != nil {
		log.Panic(err)
	}
}
