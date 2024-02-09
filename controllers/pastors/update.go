package pastors

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
	pastorId, isValidInt := strconv.Atoi(param["pastorId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[UPD0001]Unable to update pastor.",
		}, 200, w)
		return
	}

	pastor := models.Pastors{}
	if err := utils.DecodeJSONRequestBody(r, &pastor); err != nil {
		http.Error(w, "[UPD0010]Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	validate := validation.Validate()
	err := validate.Struct(pastor)

	if err != nil {
		errs := validation.GetErrors(err)
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": errs,
		}, 200, w)
		return
	}

	if err := pastor.Update(pastorId); err != nil {
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
