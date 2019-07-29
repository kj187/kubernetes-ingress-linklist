package app

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Init() {
	fmt.Println("Initializing k8s ingress linklist")
	initConfiguration()
}

func initConfiguration() {
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
