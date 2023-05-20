package model

type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[string]int
	FloatMap     map[string]float32
	TemplateData map[string]interface{}
	CSRFToken    string
	Flash        string
	Warning      string
	Errors       string
}
