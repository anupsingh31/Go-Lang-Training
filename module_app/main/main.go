package main

import (
	"module_app/greeting"
)
import "module_app/greeting/english"

func main() {
	greeting.SayHello()
	greeting.SayInHindi()
	english.SayInEnglish()
}
