package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type server struct {
	Addr              string        `yaml:"addr"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
}

type modem struct {
	Host			string	`yaml:"host"`
	PathHome		string	`yaml:"pathHome"`
	PathReboot		string	`yaml:"pathReboot"`
	BodyReboot		string	`yaml:"bodyReboot"`
	CheckHost		string	`yaml:"checkHost"`
}

type Cfg struct {
	Server	server
	Modem	modem
}

func ReadConfig(confPath string) *Cfg {
	f, err := os.Open(confPath)
	if err != nil {
		// TODO: Сделать лшгирование ошибок.
		log.Fatalln(err)
	}
	defer f.Close()

	var cfg Cfg
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		// TODO: Сделать лшгирование ошибок.
		log.Fatalln(err)
	}

	return &cfg
}
