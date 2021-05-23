package main

import (
  "github.com/galdor/go-cmdline"
  "os"
  // local packages
  "github.com/gpenaud/needys-output-producer/internal/models"
  "github.com/gpenaud/needys-output-producer/internal/utils"
)

func main() {
  cmdline := cmdline.New()

  cmdline.AddOption("", "rabbitmq.host", "HOST", "host of the rabbitMQ server")
  cmdline.SetOptionDefault("rabbitmq.host", "127.0.0.1")

  cmdline.AddOption("", "rabbitmq.port", "PORT", "port of the rabbitMQ server")
  cmdline.SetOptionDefault("rabbitmq.port", "5672")

  cmdline.AddOption("", "rabbitmq.username", "USERNAME", "username of rabbitMQ server")
  cmdline.SetOptionDefault("rabbitmq.username", "guest")

  cmdline.AddOption("", "rabbitmq.password", "PASSWORD", "password of the rabbitMQ user")
  cmdline.SetOptionDefault("rabbitmq.password", "guest")

  cmdline.AddFlag("v", "verbose", "log more information")
  cmdline.Parse(os.Args)

  models.Cfg.Rabbitmq.Host = cmdline.OptionValue("rabbitmq.host")
  models.Cfg.Rabbitmq.Port = cmdline.OptionValue("rabbitmq.port")
  models.Cfg.Rabbitmq.Username = cmdline.OptionValue("rabbitmq.username")
  models.Cfg.Rabbitmq.Password = cmdline.OptionValue("rabbitmq.password")

  utils.WaitForAmqpMessages()
}
