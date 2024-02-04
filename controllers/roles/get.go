package roles

import (
	"api/models"
	"api/utils"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	page := utils.PagerTernary(r.FormValue("page"), 1)
	rows := utils.PagerTernary(r.FormValue("rows"), 10)

	Roles := models.Roles{}

	roles, total, err := Roles.Get(r)

	if err != nil {

		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": roles,
		"paginate":   utils.Paginate(rows, page, int(total)),
	}, utils.Code.OK, w)
}