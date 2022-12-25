// @author: wzmiiiiii
// @since: 2022/12/26 00:00:59
// @desc: TODO

package tool

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string     `yaml:"app_name"`
	AppMode string     `yaml:"app_mode"`
	AppHost string     `yaml:"app_host"`
	AppPort int        `yaml:"app_port"`
	Sms     *SmsConfig `yaml:"sms"`
}

type SmsConfig struct {
	SignName     string `yaml:"sign_name"`
	TemplateCode string `yaml:"template_code"`
	RegionID     string `yaml:"region_id"`
	AppKey       string `yaml:"app_key"`
	AppSecret    string `yaml:"app_secret"`
}

var _cfg *Config

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}

func (c *Config) Address() string {
	return c.AppHost + ":" + strconv.Itoa(c.AppPort)
}
