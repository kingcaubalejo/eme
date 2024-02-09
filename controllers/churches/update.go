package churches

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
	churchId, isValidInt := strconv.Atoi(param["churchId"])

	if isValidInt != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "[UPD0001]Unable to update church.",
		}, 200, w)
		return
	}

	church := models.Churches{}
	if err := utils.DecodeJSONRequestBody(r, &church); err != nil {
		http.Error(w, "[UPD0010]Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	validate := validation.Validate()
	err := validate.Struct(church)

	if err != nil {
		errs := validation.GetErrors(err)
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": errs,
		}, 200, w)
		return
	}

	if err := church.Update(churchId); err != nil {
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
