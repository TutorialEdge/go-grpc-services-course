package main

import "log"

func Run() error {
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err.Error())
	}
}
