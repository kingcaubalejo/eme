package pastors

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	pastorId, isValidInt := strconv.Atoi(param["pastorId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[GI0001]Unable to find pastor.",
		}, 200, w)
		return
	}


	pastor := models.Pastors{}

	if err := pastor.GetInfo(pastorId); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": pastor,
	}, 200, w)

}