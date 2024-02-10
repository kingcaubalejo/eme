package users

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"fmt"
)

func GetUserBy(w http.ResponseWriter, r *http.Request) {
	param 	:= mux.Vars(r)
	by 		:= param["by"]
	query 	:= r.URL.Query()
	
	roleId, _ := strconv.Atoi(query.Get("roleId"))
	user := models.Users{
		RoleId: uint(roleId),
	}
	fmt.Println(user)

	switch by {
	case "role":
		roleId, _ := strconv.Atoi(query.Get("roleId"))
		user = models.Users {
			RoleId: uint(roleId),
		}
	case "name":
		name := query.Get("name")
		fmt.Println(name)
	default:
		fmt.Println("Invalid day!")
	}

	fmt.Println(&user)

	users, total, err := user.FindUserByRole(r)
	fmt.Println(users, total, err)
	if err != nil {

		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}

	// utils.Response(map[string]interface{}{
	// 	"statusCode": 200,
	// 	"devMessage": users,
	// 	"paginate":   utils.Paginate(rows, page, int(total)),
	// }, utils.Code.OK, w)

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": user,
	}, utils.Code.OK, w)
}
