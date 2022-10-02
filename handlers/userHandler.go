package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	user := models.User{}

	if err := bodyJson.Decode(&user); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		user.Save()
		utils.CreatedResponse(w, utils.RespOk{
			Message: fmt.Sprintf("Usuario creado id: %d", user.Id),
			Data:    user,
		})
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.ListUsers()
	if err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "Hubo un error",
			StatusCode: 400,
		})
	} else {
		utils.SendResponse(w, utils.RespOk{
			Message: "Usuarios encontrados",
			Data:    users,
		})
	}
}

func GetUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	user, err2 := models.GetUserId(id)
	if (err != nil && err2 != nil) || user.Id == 0 {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		utils.SendResponse(w, utils.RespOk{
			Message: "Usuarios encontrados",
			Data:    user,
		})
	}
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	bodyJson := json.NewDecoder(r.Body)
	user := models.User{}

	if err := bodyJson.Decode(&user); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	}
	if id != int(user.Id) {
		utils.BadResponse(w, utils.RespBad{
			Message:    "Los ids no coinciden",
			StatusCode: http.StatusBadRequest,
		})
	} else {
		user.Save()
		utils.SendResponse(w, utils.RespOk{
			Message: "Usuario modificado",
			Data:    user,
		})
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "Los id no se pudo procesar",
			StatusCode: http.StatusBadRequest,
		})
	}
	models.DeleteUser(id)
	utils.SendResponse(w, utils.RespOk{
		Message: "Usuario modificado",
		Data:    models.NewUser("", "", ""),
	})
}
