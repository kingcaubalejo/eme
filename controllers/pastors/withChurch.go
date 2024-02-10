package pastors

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WithChurch(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	pastorId, isValidInt := strconv.Atoi(param["pastorId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[WC0001]Unable to find pastor .",
		}, 200, w)
		return
	}

	pastors := models.Pastors{}

	if err := pastors.GetChurch(pastorId); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": pastors,
	}, 200, w)
}