package layout

type Settings struct {
	SiteName    string
	SubTitle    string
	HeaderLinks []Link
}

type Link struct {
	Name string
	Href string
}
