package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates carrega os templates HTML da aplicação
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate executa um template HTML com os dados fornecidos
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
