package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/XMPlusDev/XMPlusv1/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
