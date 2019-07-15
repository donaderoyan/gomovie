package main

import(
  "github.com/donaderoyan/gomovie/app"
  "github.com/donaderoyan/gomovie/config"
)

func main() {
  config := config.GetConfig()

  app := &app.App{}
  app.Initialize(config)
  app.Run(":7200")
}
