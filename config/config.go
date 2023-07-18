package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Name  string
	Debug bool

	RPC struct {
		Url string
	}

	HTTP struct {
		Port int
	}
	Redis struct {
		Host string
		Port int
	}
}

var (
	// Configuration instance
	Configuration config
)

// Read config file
func Read() {

	viper.SetConfigFile(`./config.json`)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Configuration); err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")

		prettyJSON, err := json.Marshal(Configuration)
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}

		fmt.Printf("%s\n", string(prettyJSON))
	}
}
