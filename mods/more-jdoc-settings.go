package mods

import (
	"net/http"
	"strings"

	"github.com/The-Viccyware-Group/wired-viccyware/vars"
)

type MoreJdocSettings struct {
	vars.Modification
}

func NewMoreJdocSettings() *MoreJdocSettings {
	return &MoreJdocSettings{}
}

func (modu *MoreJdocSettings) Name() string {
	return "MoreJdocSettings"
}

func (modu *MoreJdocSettings) Description() string {
	return "More Jdoc Settings"
}

func (modu *MoreJdocSettings) Load() error {
	return nil
}

type eyeColorResponse struct {
	IsCustom         bool    `json:"iscustom"`
	Preset           int     `json:"preset"`
	CustomHue        float32 `json:"hue"`
	CustomSaturation float32 `json:"saturation"`
}

func (m *MoreJdocSettings) HTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api/mods/"+m.Name()) {
		return
	}
	switch strings.TrimPrefix(r.URL.Path, "/api/mods/"+m.Name()+"/") {
	case "getEyeColor":
	case "setEyeColor":
	case "getVolume":
	case "setVolume":
	// this could be where names go
	default:
		vars.HTTPError(w, r, "404 not found")
		return
	}
}
