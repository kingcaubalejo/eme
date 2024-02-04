package users

import (
	"api/models"
	"api/utils"
	"api/validation"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}
	if err := utils.DecodeJSONRequestBody(r, &user); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	validate := validation.Validate()
	err := validate.Struct(user)

	if err != nil {
		errs := validation.GetErrors(err)
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": errs,
		}, 200, w)
		return
	}

	user.Password = utils.MakePassword(user.Password)
	if err := user.Create(); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": user.AccountId,
	}, 200, w)

}
