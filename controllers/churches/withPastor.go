package churches

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
)

func WithPastor(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	churchId, isValidInt := strconv.Atoi(param["churchId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[WP0001]Unable to find church.",
		}, 200, w)
		return
	}

	church := models.Churches{}

	if err := church.GetPastor(churchId); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": church,
	}, 200, w)
}