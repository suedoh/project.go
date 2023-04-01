package main

import "log"

func CheckFatalln(msg string, e error) {
	if e != nil {
		log.Fatalln(msg, e)
		panic(e)
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
