package pastors

import (
	"api/models"
	"api/utils"
	"api/validation"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	pastor := models.Pastors{}
	if err := utils.DecodeJSONRequestBody(r, &pastor); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
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

	if err := pastor.Create(); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": pastor.PastorId,
	}, 200, w)

}
