// cmd/nftmetadataindexerstudionext/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexerstudionext/internal/nftmetadataindexerstudionext"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexerstudionext.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
