package main

import (
	"os"

	"github.com/Sathishwar/resque_exporter"
)

func main() {
	resqueExporter.Run(os.Args[1:])
}
