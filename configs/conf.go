package configuration

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Addr            string   `yaml:"addr"`
	Port            string   `yaml:"port"`
	SmsMmsProv      []string `yaml:"smsprov"`
	VoiceProv       []string `yaml:"vprov"`
	EmailProv       []string `yaml:"eprov"`
	PathToRead      string   `yaml:"pathtoread"`
	SmsFileName     string   `yaml:"smsfilename"`
	VoiceFileName   string   `yaml:"voicefilename"`
	EmailFileName   string   `yaml:"emailfilename"`
	BillingFileName string   `yaml:"billingfilename"`
	MmsURL          string   `yaml:"mmsurl"`
	SupportURL      string   `yaml:"supporturl"`
	IncidentURL     string   `yaml:"incidenturl"`
}

func CheckCfg() *Configuration {
	var c *Configuration

	yamlFile, err := ioutil.ReadFile("./configs/conf.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
