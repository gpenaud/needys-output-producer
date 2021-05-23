package models

type Config struct {
  Rabbitmq struct {
    Port string `yaml:"port"`
    Host string `yaml:"host"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
  } `yaml:"rabbitmq"`
}
