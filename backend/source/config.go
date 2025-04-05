package source

import (
	"fmt"
	"io/ioutil"
	"os"

	"nyauth_backed/source/logger"

	"gopkg.in/yaml.v2"
)

// ServerConfig 结构体定义服务器配置项
type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	BaseURL string `yaml:"base_url"`
}

// DatabaseConfig 结构体定义数据库配置项
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type turnstileConfig struct {
	SiteKey   string `yaml:"sitekey"`
	SecretKey string `yaml:"secretkey"`
}

type smtpConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Config 结构体定义配置项
type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Database  DatabaseConfig  `yaml:"database"`
	Turnstile turnstileConfig `yaml:"turnstile"`
	SMTP      smtpConfig      `yaml:"smtp"`
}

// 全局变量保存配置
var AppConfig *Config

// defaultConfig 返回默认配置
func defaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:    "0.0.0.0",
			Port:    8080,
			BaseURL: "http://localhost:8080",
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     27017,
			Username: "root",
			Password: "123456",
		},
		Turnstile: turnstileConfig{
			SiteKey:   "sitekey",
			SecretKey: "secretkey",
		},
		SMTP: smtpConfig{
			Host:     "smtp.example.com",
			Port:     587,
			From:     "",
			Username: "",
			Password: "your-email-password",
		},
	}
}

// LoadConfig 加载配置文件
func LoadConfig() error {
	const configFile = "config.yaml"

	// 检查配置文件是否存在
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		config := defaultConfig()
		data, err := yaml.Marshal(config)
		if err != nil {
			return fmt.Errorf("error marshaling default config: %w", err)
		}
		if err := ioutil.WriteFile(configFile, data, 0644); err != nil {
			return fmt.Errorf("error writing default config file: %w", err)
		}
		logger.Info("Default config file created.")
		os.Exit(0)
	}

	// 读取配置文件
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// 解析配置文件
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("error unmarshaling config file: %w", err)
	}

	AppConfig = &config
	return nil
}
