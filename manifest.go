package util

import (
	"encoding/json"
	"os"
)

func (manifest Manifest) New(file string) *Manifest {
	data, _ := os.ReadFile(file)
	json.Unmarshal(data, &manifest)
	return &manifest
}

type Manifest struct {
	BackgroundColor           string                `json:"bakground_color"`
	Categories                []string              `json:"categroies"`
	Description               string                `json:"description"`
	Display                   string                `json:"display"`
	DisplayOverride           string                `json:"display_override"`
	Icons                     []ManifestIcon        `json:"icons"`
	Id                        string                `json:"id"`
	Name                      string                `json:"name"`
	Orientation               string                `json:"orientation"`
	PreferRelatedApplications bool                  `json:"prefer_related_applications"`
	ProtocolHandlers          []ManifestProtocol    `json:"protocol_handlers"`
	RelatedApplications       []ManifestRelated     `json:"related_applications"`
	Scope                     string                `json:"scope"`
	Screenshots               []ManifestScreenshots `json:"screenshots"`
	ShareTarget               ManifestShareTarget   `json:"share_target"`
	ShortName                 string                `json:"short_name"`
	Shortcuts                 string                `json:"shortcuts"`
	StartUrl                  string                `json:"start_url"`
	ThemeColor                string                `json:"theme_color"`
}

type ManifestIcon struct {
	Purpose string `json:"purpose"`
	Src     string `json:"src"`
	Sizes   string `json:"sizes"`
	Type    string `json:"type"`
}

type ManifestProtocol struct {
	Protocol string `json:"protocol"`
	URL      string `json:"url"`
}

type ManifestRelated struct {
	Id       string `json:"id"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

type ManifestScreenshots struct {
	Src      string `json:"url"`
	Sizes    string `json:"sizes"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Label    string `json:"label"`
}

type ManifestShareTarget struct {
	Action  string                    `json:"action"`
	Enctype string                    `json:"enctype"`
	Method  string                    `json:"method"`
	Params  ManifestShareTargetParams `json:"params"`
}

type ManifestShareTargetParams struct {
	Title string                          `json:"type"`
	Text  string                          `json:"text"`
	URL   string                          `json:"url"`
	Files []ManifestShareTargetParamsFile `json:"files"`
}

type ManifestShareTargetParamsFile struct {
	Name   string   `json:"name"`
	Accept []string `json:"accept"`
}
