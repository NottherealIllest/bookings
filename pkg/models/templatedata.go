 package models

 // TemplateData holds data send from handlers to temolate
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FlaotMap  map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}
