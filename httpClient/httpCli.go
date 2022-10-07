package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/models/interfaces"
	"server/utils"
)

func GetHttp(url string, w http.ResponseWriter) {
	bodyBites, _ := json.Marshal(``)
	response, err := processReq("GET", url, bodyBites, w)
	if err == nil {
		utils.ResponseFile(w, response)
	}
}

func PostHttp(url string, body any, w http.ResponseWriter) {
	bodyBites, _ := json.Marshal(body)
	response, err := processReq("POST", url, bodyBites, w)
	if err == nil {
		utils.ResponseFile(w, response)
	}
}

func PostHttpAve(url string, body any, w http.ResponseWriter) {
	bodyBites, _ := json.Marshal(body)
	response, err := processReq("POST", url, bodyBites, w)
	if err == nil {
		fmt.Fprintln(w, response)
	}
}

func processReq(method string, url string, body []byte, w http.ResponseWriter) (string, error) {
	peticion, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		msj := fmt.Sprintf("No se pudo realizar la peticion, Error: %s", err)
		utils.BadResponse(w, utils.RespBad{
			StatusCode: http.StatusBadRequest,
			Message:    msj,
		})
	}
	peticion.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(peticion)
	if err != nil {
		msj := fmt.Sprintf("Peticion fallida, Error: %s", err)
		utils.BadResponse(w, utils.RespBad{StatusCode: http.StatusNotFound, Message: msj})
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		msjError := ""
		badResp := interfaces.ResponseAveGen{}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			msjError = "error inesperado"
		} else {
			json.Unmarshal(body, &badResp)
			if badResp.Message == "" {
				msjError = "No hubo mensaje del error"
			} else {
				msjError = badResp.Message
			}
		}
		utils.BadResponse(w, utils.RespBad{StatusCode: resp.StatusCode, Message: msjError})
		return "", errors.New(msjError)
	} else {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			msj := fmt.Sprintf("No se pudo procesar la respuesta, Error: %s", err)
			utils.BadResponse(w, utils.RespBad{StatusCode: http.StatusBadRequest, Message: msj})
		}
		return string(respBytes), nil
	}

}
