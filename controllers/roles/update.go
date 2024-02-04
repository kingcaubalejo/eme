package roles

import (
	"api/models"
	"api/utils"
	"api/validation"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	roleId, isValidInt := strconv.Atoi(param["roleId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[UPD0001]Unable to update role.",
		}, 200, w)
		return
	}

	role := models.Roles{}
	if err := utils.DecodeJSONRequestBody(r, &role); err != nil {
		http.Error(w, "[UPD0010]Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	validate := validation.Validate()
	err := validate.Struct(role)

	if err != nil {
		errs := validation.GetErrors(err)
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": errs,
		}, 200, w)
		return
	}

	if err := role.Update(roleId); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": role,
	}, 200, w)

}
