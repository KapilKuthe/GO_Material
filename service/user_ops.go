package service

import (
	"fmt"
	"goLogin/database"
	"goLogin/models"
	"goLogin/security"
	"goLogin/utility"

	"github.com/kataras/iris/v12"
)

func CreateUser(ctx iris.Context) {
	var user models.User

	//* validation for json
	err := ctx.ReadJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body!"})
		return
	}

	//* execute in DB
	user, err = database.CreateUser(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to save data!", "error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(user)
}

// ! Request json email,password
func UserLogin(ctx iris.Context) {
	var user models.User

	//* validation for json
	err := ctx.ReadJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body!"})
		return
	}
	loginPassword := user.Password
	//* fetch user details
	user, err = database.GetUser(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "No Record found!", "error": err.Error()})
		return
	}

	//* compare the password
	if !utility.ComparePassward(loginPassword, user.Password) {
		return
	}

	//* generation of token
	token, err := security.GenerateToken(user.Email, uint64(user.ID))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Unable to generate token!", "error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"status": "success", "token": token})
}

func Getlanging(ctx iris.Context) {

}
