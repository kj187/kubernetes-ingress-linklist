package app

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Init ...
func Init() {
	fmt.Println("Initializing k8s ingress linklist")
	initConfiguration()
}

func initConfiguration() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("/kubernetes-ingress-linklist")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
