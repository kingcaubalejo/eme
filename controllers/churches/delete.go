package churches

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	churchId, isValidInt := strconv.Atoi(param["churchId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[DEL0001]Unable to delete church.",
		}, 200, w)
		return
	}

	church := models.Churches{}
	if err := church.Delete(churchId); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}
	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": "success",
	}, 200, w)

}
