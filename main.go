package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

// exitCode wraps a return value for the application
type exitCode struct{ Code int }

func main() {
	configPath := flag.String("config-dir", ".", "Directory that contains the configuration file")
	flag.Parse()

	viper.SetConfigName("sendConf")
	viper.AddConfigPath(*configPath)

	_, _ = fmt.Fprintln(os.Stderr, "Reading configuration from", *configPath)

	err := viper.ReadInConfig()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed reading configuration:", err.Error())
		panic(exitCode{1})
	}

	fmt.Println(viper.GetString("smtp_conf.host"))

	//err = core.Mail()
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(time.Now().Format("2006-01-02"))

}
