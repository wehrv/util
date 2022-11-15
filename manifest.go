package util

import (
	"encoding/json"
	"os"
)

type Manifest struct {
	BackgroundColor           string
	Categories                []string
	Description               string
	Display                   string
	DisplayOverride           string
	Icons                     []ManifestIcon
	Id                        string
	Name                      string
	Orientation               string
	PreferRelatedApplications bool
	ProtocolHandlers          []ManifestProtocol
	RelatedApplications       []ManifestRelated
	Scope                     string
	Screenshots               []ManifestScreenshots
	ShareTarget               ManifestShareTarget
	ShortName                 string
	Shortcuts                 string
	StartUrl                  string
	ThemeColor                string
}

type ManifestIcon struct {
	Src   string
	Sizes string
	Type  string
}

type ManifestProtocol struct {
	Protocol string
	URL      string
}

type ManifestRelated struct {
	Id       string
	Platform string
	URL      string
}

type ManifestScreenshots struct {
	Src      string
	Sizes    string
	Type     string
	Platform string
	Label    string
}

type ManifestShareTarget struct {
	Action  string
	Enctype string
	Method  string
	Params  ManifestShareTargetParams
}

type ManifestShareTargetParams struct {
	Title string
	Text  string
	URL   string
	Files []ManifestShareTargetParamsFile
}

type ManifestShareTargetParamsFile struct {
	Name   string
	Accept []string
}

func (manifest Manifest) New(file string) Manifest {
	data, _ := os.ReadFile(file)
	json.Unmarshal(data, &manifest)
	return manifest
}
