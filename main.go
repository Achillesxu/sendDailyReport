package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
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

	fmt.Println(viper.GetString("smtp_config.host"))

}
