package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds the application config
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	Inproduction  bool
	Session       *scs.SessionManager
}
