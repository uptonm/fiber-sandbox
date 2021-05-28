package main

import (
	"flag"
	"fmt"
	"github.com/uptonm/fiber-sandbox/src/www"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/uptonm/fiber-sandbox/src/config"
)

var (
	env config.Environment
)

func init() {
	envFlag := flag.Bool("prod", false, "The env flag is used to specify the runtime mode")

	env = config.Development
	if *envFlag == true {
		env = config.Production
	}
}

func main() {
	appConfig := config.ReadConfig(env)

	log.Printf("%s:%s - %s", appConfig.Ingress.Host, appConfig.Ingress.Port, env)
	r := fiber.New(fiber.Config{
		Prefork: env == config.Production,
	})

	www.WireHandlers(r, appConfig, env == config.Production)

	// graceful shutdown with SIGINT | SIGTERM and others will hard kill
	// credit for this lovely method https://github.com/dechristopher/dchr.host
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		_ = r.Shutdown()
	}()

	// listen for connections on primary listening port
	if err := r.Listen(fmt.Sprintf(":%s", appConfig.Ingress.Port)); err != nil {
		log.Println(err.Error())
	}

	// exit cleanly
	os.Exit(0)
}
