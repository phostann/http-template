package main

import (
	"fmt"
	"os"

	"github.com/phostann/template/http/app"
	_ "github.com/phostann/template/http/docs"
)

func main() {
	err := app.NewApp().Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to start server: %v", err)
		os.Exit(1)
	}
}
