package core

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"time"
)

func Mail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("mail_conf.from"))
	toArray := viper.GetStringSlice("mail_conf.to")
	m.SetHeader("To", toArray[0])
	m.SetHeader("Subject", fmt.Sprintf("%s 日报", time.Now().Format("2006-04-14")))
	m.SetBody("text/html", "Hello <b>xushiyin</b> and <i>Cora</i>!")
	d := gomail.NewDialer(viper.GetString("smtp_conf.host"),
		viper.GetInt("smtp_conf.port"),
		viper.GetString("smtp_conf.user"),
		viper.GetString("smtp_conf.password"))

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	} else {
		return nil
	}
}
