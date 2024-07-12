package controllers

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
	"GO-JWT-Auth/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	var existingUser models.Auth
	isUser := initializers.DB.Where("email = ?", body.Email).First(&existingUser)

	if isUser.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
		return
	}

	user := models.Auth{Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Generate Token
	tokenString, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	})

}

func SignIn(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	var user models.Auth
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid credentials"})
		return
	}

	pwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if pwErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid password"})
		return
	}

	//Generate Token
	tokenString, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed in successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	})

}

func ResetPassword(c *gin.Context) {
	var body struct {
		Email string
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required!"})
		return
	}

	// Load html file
	htmlContent, err := ioutil.ReadFile("templates/otp_template.html")
	if err != nil {
		log.Fatal(err)
	}

	// Generate a 4 digit number
	otpValue := rand.Intn(9000) + 1000

	htmlBody := strings.Replace(string(htmlContent), "{{OTP_CODE}}", strconv.Itoa(otpValue), -1)

	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SENDER_EMAIL")) // Replace with your email address
	m.SetHeader("To", body.Email)                  // Replace with recipient's email address
	m.SetHeader("Subject", "Reset Password - OTP Verification")
	m.SetBody("text/html", htmlBody)

	// Set up the SMTP server configuration
	d := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("SENDER_EMAIL"), os.Getenv("SENDER_PASSWORD"))

	// Use STARTTLS
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

	otp := models.Otp{Email: body.Email, Otp: strconv.Itoa(otpValue)}
	result := initializers.DB.Create(&otp)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed send otp",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully!",
	})

	log.Println("Email sent successfully!")
}
