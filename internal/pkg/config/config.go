package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Config struct {
	Env       string
	Http      HttpConfig
	Common    CommonConfig
	Redis     RedisConfig
	Endpoints map[string]map[string]string
}

var cfg *Config

func Instance() *Config {
	if cfg == nil {
		cfg = &Config{
			Env:       "",
			Http:      HttpConfig{},
			Common:    CommonConfig{},
			Redis:     RedisConfig{},
			Endpoints: make(map[string]map[string]string),
		}
	}
	return cfg
}

// Load loads configurations from file and env
func Load(configFile string) error {
	// Default Config values
	c := Instance()
	defaults.SetDefaults(c)

	// --- hacking to load reflect structure Config into env ----//
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Read Config file failed. ", err)

		configBuffer, err := json.Marshal(c)

		if err != nil {
			return err
		}

		err = viper.ReadConfig(bytes.NewBuffer(configBuffer))
		if err != nil {
			return err
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// -- end of hacking --//

	fmt.Println(viper.GetString("ENV"))
	viper.AutomaticEnv()
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
