package main

type Config struct {
	Webhook string `yaml:"webhook"`
	Secret  string `yaml:"secret"`
}
