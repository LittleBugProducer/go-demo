package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stdout)
	log.Info().Str("content", "hello world").Int("count", 3).Msg("print something")

	log = log.With().Caller().Str("foo", "bar").Logger()
	log.Info().Msg("hello world")

}
