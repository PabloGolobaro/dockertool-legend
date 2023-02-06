package main

import (
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"log"
	"sync"
)

func main() {
	builder, err := app.NewAppBuilder()
	if err != nil {
		log.Fatal("Get builder error:", err)
	}

	apllication := builder.Build()

	var wg sync.WaitGroup

	wg.Add(1)
	go func(application app.Apllication) {
		application.Start()

		defer wg.Done()
	}(apllication)

	wg.Wait()

}
