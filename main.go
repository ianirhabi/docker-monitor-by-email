/**
* created by Andrian Latif on 02-18-2020
* for monitor docker services
**/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DataName dataname
}

type dataname struct {
	CONFIG_SMTP_HOST string
	CONFIG_SMTP_PORT int
	CONFIG_EMAIL     string
	CONFIG_PASSWORD  string
	SEND_TO          string
	SEND_CC1         string
	SEND_CC2         string
}

func main() {
	er := OpenFile()
	if er != nil {
		log.Println(er)
	}
}

func Mail(s string, h string) {
	var conf Config
	if _, err := toml.DecodeFile("/etc/dockermonitor/config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	to := []string{conf.DataName.SEND_TO, conf.DataName.SEND_TO}

	//if you want to send cc
	cc := []string{conf.DataName.SEND_CC1, conf.DataName.SEND_CC2}

	subject := "Service Docker " + s
	message := "service docker " + s + " on instance " + h + "\noffline time is " + __getTime().String()

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	var conf Config
	if _, err := toml.DecodeFile("/etc/dockermonitor/config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	fmt.Println("running .... ")
	body := "From: " + conf.DataName.CONFIG_EMAIL + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", conf.DataName.CONFIG_EMAIL, conf.DataName.CONFIG_PASSWORD, conf.DataName.CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", conf.DataName.CONFIG_SMTP_HOST, conf.DataName.CONFIG_SMTP_PORT)

	fmt.Println("check ", smtpAddr)
	err := smtp.SendMail(smtpAddr, auth, conf.DataName.CONFIG_EMAIL, append(to, cc...), []byte(body))

	if err != nil {
		return err
	}
	return nil
}

func OpenFile() (e error) {
	for true {
		time.Sleep(1 * time.Second)
		file, err := os.Open("/etc/dockermonitor/list_container_stop.txt")

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}

		file.Close()
		for _, eachline := range txtlines {
			fmt.Println(eachline)
			if eachline == "" {
				println("-------------")
			} else {
				fmt.Println(eachline)
				Mail(eachline, ___getHostname())
				time.Sleep(60 * time.Second)
			}
		}
	}
	return
}

// know current directory from an app
func ___getPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

// get hostname instance
func ___getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return hostname
}

// get time
func __getTime() time.Time {
	var time_not_utc = time.Now()
	return time_not_utc
}
