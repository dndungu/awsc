package main

import (
	"context"
	"log"
	"os"

	"sire.run/awsc/pkg/app"
)

func main() {
	if err := do(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func do(ctx context.Context) error {
	return app.App{
		Args:   os.Args,
		Input:  "in.yaml",
		Output: "out.yaml",
	}.Do(ctx)
}
