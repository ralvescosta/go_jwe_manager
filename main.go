package main

import (
	"jwemanager/cmd"
	"log"
)

func main() {
	log.Fatal(cmd.HttpServer())
}
