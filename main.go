package main

import (
	"os"

	"github.com/liampulles/countdown/pkg/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
