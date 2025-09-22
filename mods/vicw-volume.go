package mods

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/The-Viccyware-Group/wired-viccyware/vars"
)

var VolumeFile = "/data/vicw-volume"
var DefaultVolume = "MEDIUM"

type VolumeControl struct {
	vars.Modification
}

func NewVolumeControl() *VolumeControl {
	return &VolumeControl{}
}

func (modu *VolumeControl) Name() string {
	return "VolumeControl"
}

func (modu *VolumeControl) Description() string {
	return "Control system volume level."
}

func (modu *VolumeControl) Load() error {
	return nil
}

func (m *VolumeControl) HTTP(w http.ResponseWriter, r *http.Request) {
	if vars.IsEndpoint(r, "set") {
		level := r.FormValue("level")
		if level == "" {
			vars.HTTPError(w, r, "empty volume level")
			return
		}
		
		if !isValidVolume(level) {
			vars.HTTPError(w, r, "invalid volume level")
			return
		}
		
		err := setVolume(level)
		if err != nil {
			vars.HTTPError(w, r, "failed to set volume: "+err.Error())
			return
		}
		
		vars.HTTPSuccess(w, r)
		return
	} else if vars.IsEndpoint(r, "get") {
		volume, err := getVolume()
		if err != nil {
			volume = DefaultVolume
		}
		w.Write([]byte(volume))
		return
	} else {
		vars.HTTPError(w, r, "404 not found")
	}
}

func isValidVolume(level string) bool {
	validLevels := []string{"MUTE", "LOW", "MEDIUM_LOW", "MEDIUM", "MEDIUM_HIGH", "HIGH"}
	for _, valid := range validLevels {
		if level == valid {
			return true
		}
	}
	return false
}

func setVolume(level string) error {
	err := vars.SaveFile(level, VolumeFile)
	if err != nil {
		return fmt.Errorf("failed to write volume file: %v", err)
	}
	cmd := exec.Command("/etc/initscripts/anki-audio-init")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute audio init: %v", err)
	}
	
	fmt.Printf("Volume set to: %s\n", level)
	return nil
}

func getVolume() (string, error) {
	content, err := vars.ReadFile(VolumeFile)
	if err != nil {
		return "", err
	}
	
	volume := strings.TrimSpace(content)
	if !isValidVolume(volume) {
		return DefaultVolume, nil
	}
	
	return volume, nil
}