package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string    `json:"app_name"`
	AppMode  string    `json:"app_mode"`
	AppHost  string    `json:"app_host"`
	AppPort  string    `json:"app_port"`
	Sms      SmsConfig `json:"sms"`
	Database DbConfig  `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	Appsecret    string `json:"app_secret"`
}

type DbConfig struct {
	Driver    string `json:"driver"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	DbName    string `json:"db_name"`
	CharSet   string `json:"char_set"`
	IsShowsql bool   `json:"isShowsql"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)              //返回reader指针
	decoder := json.NewDecoder(reader)           //NewDecoder解析器
	if err = decoder.Decode(&_cfg); err != nil { //判断返回状态
		return nil, err
	}

	return _cfg, nil
}
