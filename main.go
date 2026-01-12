package main

import (
	"fmt"
	"net/http"

	"github.com/The-Viccyware-Group/wired-viccyware/mods"
	"github.com/The-Viccyware-Group/wired-viccyware/vars"
)

var EnabledMods []vars.Modification = []vars.Modification{
	mods.NewFreqChange(),
	mods.NewWakeWordPV(),
	mods.NewAutoUpdate(),
	mods.NewSensitivityPV(),
	mods.NewJdocSettings(),
	mods.NewVolumeControl(),
	mods.NewFaces(),
}

func main() {
	vars.EnabledMods = EnabledMods
	vars.InitMods()
	startweb()
}

func startweb() {
	fmt.Println("starting web at port 8080")
	fs := http.FileServer(http.Dir("/etc/wired/webroot"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// no mno non o caching
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		fs.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
