package main

import (
	"github.com/cucumberjaye/softline/internal/pkg/app"
	"github.com/rs/zerolog/log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Err(err)
	}
}
