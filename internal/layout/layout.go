package layout

import (
	"encoding/json"
	"log"
)

type Settings struct {
	SiteName    string
	SubTitle    string
	HeaderLinks []Link
	URL         string
	SEO         SEO
}

type SEO struct {
	// SiteTitle is used on the <title> tag
	SiteTitle string
	// SiteDescription is used on <meta name="description" ... />
	SiteDescription string
	SchemaOrg       SchemaOrg
}

type SchemaOrg struct {
	Context    string   `json:"@context"`
	Type       string   `json:"@type"`
	Name       string   `json:"name"`
	URL        string   `json:"url"`
	JobTitle   string   `json:"jobTitle"`
	KnowsAbout []string `json:"knowsAbout"`
	SameAs     []string `json:"sameAs"`
}

func NewSchemaOrg(name, url, jobTitle string, knowsAbout, sameAs []string) SchemaOrg {
	return SchemaOrg{
		Context:    "https://schema.org",
		Type:       "Person",
		Name:       name,
		URL:        url,
		JobTitle:   jobTitle,
		KnowsAbout: knowsAbout,
		SameAs:     sameAs,
	}
}

func (s SchemaOrg) ToJSON() string {
	json, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Error marshalling SchemaOrg to JSON: ", err)
	}
	return string(json)
}

type Link struct {
	Name string
	Href string
}
