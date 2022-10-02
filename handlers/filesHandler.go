package handlers

import (
	"fmt"
	"net/http"
	"server/models/interfaces"
	"server/utils"
	"text/template"
)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	file, err := utils.OpenFile("static/hola.json")
	if err != nil {
		fmt.Println(err)
	} else {
		utils.ResponseFile(w, r, file)
	}
}

func TemplateFile(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("web/holaMundo.html", "web/base.html")
	if err != nil {
		errors := utils.RespBad{
			Message:    err.Error(),
			StatusCode: 404,
		}
		utils.BadResponse(w, errors)
	} else {
		user := interfaces.UserTemplt{
			Nombre: "Lord Drako",
			Edad:   33,
		}
		templ.Execute(w, user)
	}
}
