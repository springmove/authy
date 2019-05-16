package core

import (
	"github.com/linshenqi/authy/src/base"
	"github.com/linshenqi/authy/src/services/wechat"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfigFile struct {
	Http   HttpConfig          `yaml:"http"`
	WeChat wechat.WeChatConfig `yaml:"wechat"`
}

type Config struct {
	confPath string
	cfg      ConfigFile
}

func (s *Config) Init(service base.BaseService) error {

	f, err := os.Open(s.confPath)
	defer f.Close()

	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &s.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Config) Release() {

}

func (s *Config) SetConf(conf string) {
	s.confPath = conf
}

func (s *Config) Config() ConfigFile {
	return s.Config()
}

func (s *Config) ServiceName() string {
	return "config"
}
