package config

import (
	"flag"
	"os"
)

type Config struct {
	username     string
	password     string
	instanceName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.username, "username", os.Getenv("USERNAME"), "serviceNow Username")
	flag.StringVar(&conf.password, "password", os.Getenv("PASSWORD"), "serviceNow Password")
	flag.StringVar(&conf.instanceName, "instanceName", os.Getenv("INSTANCE_NAME"), "serviceNow Instance Name")

	flag.Parse()

	return conf
}

func (c *Config) GetUsername() string {
	return c.username
}

func (c *Config) GetPassword() string {
	return c.password
}

func (c *Config) GetURL() string {
	return "https://" + c.instanceName + ".service-now.com/"
}
