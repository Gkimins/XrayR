package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Gkimins/XrayR/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
