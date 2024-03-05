package service

import (
	"bytes"
	"fmt"
	db "goNotification/database"
	"goNotification/model"
	"html/template"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func Sendemail(ctx iris.Context) {
	var jsonRequest model.JRequest

	//* validation for json
	err := ctx.ReadJSON(&jsonRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "invalid request body!"})
		return
	}

	//* copy to tbl struct
	comtmpl := model.MSCommunication{
		TmplID: jsonRequest.TemplateID,
	}

	//* fetching db template data
	comtmpl, err = db.GetTemplate(comtmpl)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to fetch data!", "error": err.Error()})
		return
	}

	//* Parse the HTML template
	htmlTemplate, err := template.New("emailTemplate").Parse(comtmpl.TmplMessage)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to generate html!", "error": err.Error()})
		return
	}

	//* Execute the template with dynamic values
	var result bytes.Buffer
	err = htmlTemplate.Execute(&result, jsonRequest.Variables)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "unable to render html!", "error": err.Error()})
		return
	}

	// Print or use the result as needed
	// fmt.Println(result.String())

	//* load smtp variables
	emailConfig, err := loadEnv()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "smtp failure!", "error": err.Error()})
		return
	}

	//* converting & binding Email struct
	email := convertToEmail(jsonRequest)
	email.Subject = comtmpl.Action
	email.Message = result.String()

	//* smtp call
	status := mailsender(email, emailConfig)
	if status {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{"message": "Email sent succussfully!"})
	} else {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Failure sending mail!"})
	}
}

func loadEnv() (model.EmailConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Unable to load env file %v", err)
		return model.EmailConfig{}, err
	}

	emailConfig := model.EmailConfig{
		UserName: os.Getenv("username"),
		Password: os.Getenv("password"),
		Host:     os.Getenv("host"),
	}

	return emailConfig, nil
}

// ? extraction request to email struct
func convertToEmail(jsonRequest model.JRequest) model.Email {
	email := model.Email{
		From:    jsonRequest.Data["from"],
		To:      []string{jsonRequest.Data["to"]}, // Assuming "to" is a single recipient, modify as needed
		Cc:      nil,                              // Add logic to populate Cc if available in your data
		Bcc:     nil,                              // Add logic to populate Bcc if available in your data
		Subject: "",                               // Assuming "subject" is a string, modify as needed
		Message: "",                               // Assuming "message" is a string, modify as needed
	}

	return email
}

// ? smtp email sender
func mailsender(email model.Email, emailConfig model.EmailConfig) bool {
	fmt.Println("email:", email)
	fmt.Println("emailConfig", emailConfig)
	auth := smtp.PlainAuth(
		"",
		emailConfig.UserName,
		emailConfig.Password,
		emailConfig.Host,
	)
	//? the mine line is responsible for renderin the html output
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	//? subject is captured in message itseft separated with \n 
	msg := "Subject:" + email.Subject + "\n" + mime + "\n" + email.Message

	err := smtp.SendMail(
		emailConfig.Host+":587",
		auth,
		email.From,
		email.To,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("Error sending email:", err)
		return false
	} else {
		fmt.Println("Email sent successfully!")
		return true
	}
}
