package main

import (
	"github.com/Pelegrinetti/trellenge-go/internal/config"
	"github.com/Pelegrinetti/trellenge-go/pkg/http"
)

func main() {
	config := config.New()
	s := http.New()

	s.Run(config.Port)
}
