package service

import (
	"goLogin/database"
	"goLogin/models"

	"github.com/kataras/iris/v12"
)

func CreateCustomer(ctx iris.Context) {
	var customer models.Customer

	//* validation for json
	err := ctx.ReadJSON(&customer)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body!"})
		return
	}

	//* execute in DB
	customer, err = database.CreateCustomer(customer)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to save data!", "error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(customer)
}

func GetAllCustomer(ctx iris.Context) {
	// var customer []models.Customer

	//* fetch all customers from DB
	customer, err := database.GetAllCustomer()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to fetch data!", "error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(customer)
}

func UpdateCustomer(ctx iris.Context) {
	var customer models.Customer

	//* validation for Id param
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid customer Id!", "error": err.Error()})
		return
	}

	//* assignment to var
	customer.Id = uint64(id)
	err = ctx.ReadJSON(&customer)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body!"})
		return
	}

	//* request update to DB
	database.UpdateCustomer(customer)

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(customer)
}

func DeleteCustomer(ctx iris.Context) {
	//* validation for Id param
	id, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid customer Id!", "error": err.Error()})
		return
	}

	//* request delete to DB
	err = database.DeleteCustomer(uint64(id))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Deletion Failed!", "error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "customer Deleted succussfully!"})
}

// func verifyCustomer(customer models.Customer) models.Customer {
// 	var existingCustomer models.Customer

// 	if customer.Id != 0 {
// 		existingCustomer.Id = customer.Id
// 	}
// 	if customer.Name != "" {
// 		existingCustomer.Name = customer.Name
// 	}
// 	if customer.Dob != "" {
// 		existingCustomer.Dob = customer.Dob
// 	}
// 	if customer.Mobile != 0 {
// 		existingCustomer.Mobile = customer.Mobile
// 	}
// 	if customer.Email != "" {
// 		existingCustomer.Email = customer.Email
// 	}

// 	return existingCustomer
// }
