package main

import (
	"fmt"
	"net/http"

	"github.com/os-vector/wired/mods"
	"github.com/os-vector/wired/vars"
)

var EnabledMods []vars.Modification = []vars.Modification{
	mods.NewFreqChange(),
	mods.NewWakeWordPV(),
	mods.NewAutoUpdate(),
}

func main() {
	vars.EnabledMods = EnabledMods
	vars.InitMods()
	startweb()
}

func startweb() {
	fmt.Println("starting web at port 8080")
	http.Handle("/", http.FileServer(http.Dir("/etc/wired/webroot")))
	http.ListenAndServe(":8080", nil)
}
