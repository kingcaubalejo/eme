package roles

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	roleId, isValidInt := strconv.Atoi(param["roleId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[DEL0001]Unable to delete role.",
		}, 200, w)
		return
	}

	role := models.Roles{}
	if err := role.Delete(roleId); err != nil {
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
