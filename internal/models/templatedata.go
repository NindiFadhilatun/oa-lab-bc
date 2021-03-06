package models

import "github.com/okiprakasa/oa-lab-bc/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Quote     string
	Author    string
	Error     string
	Form      *forms.Form
}
