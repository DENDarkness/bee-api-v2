package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type http struct {
	Server server
	Client client
}

type server struct {
	Addr              string        `yaml:"addr"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
}

type client struct {
	Timeout	time.Duration	`yaml:"timeout"`
}

type modem struct {
	Host       string `yaml:"host"`
	PathHome   string `yaml:"pathHome"`
	PathReboot string `yaml:"pathReboot"`
	PathToken  string `yaml:"pathToken"`
	BodyReboot string `yaml:"bodyReboot"`
	CheckHost  string `yaml:"checkHost"`
}

type url struct {
	GetIP string `yaml:"getIP"`
}

type Cfg struct {
	HTTP  http
	Modem modem
	URL url
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
