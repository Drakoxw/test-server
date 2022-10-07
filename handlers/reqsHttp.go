package handlers

import (
	"encoding/json"
	"net/http"
	httpclient "server/httpClient"
	"server/models/interfaces"
	"server/utils"
)

func GetPokes(w http.ResponseWriter, r *http.Request) {
	Url := "https://rickandmortyapi.com/api/character"
	httpclient.GetHttp(Url, w)
}

func ListBanksAve(w http.ResponseWriter, r *http.Request) {
	Url := "https://aveonline.co/api/comunes/v1.0/bancos.php"
	bodyJson := json.NewDecoder(r.Body)
	bodyAve := interfaces.GettBancosAve{}
	if err := bodyJson.Decode(&bodyAve); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	}
	httpclient.PostHttpAve(Url, &bodyAve, w)
}
