package controllers

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
	"GO-JWT-Auth/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mail.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SignUp(c *gin.Context, logger *zap.SugaredLogger) {

	var body struct {
		Email    string
		Password string
	}

	logger.Info("Email: ", body.Email)

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Info("Error: ", err.Error())
		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		logger.Info("Email and password are required")
		return
	}

	var existingUser models.Auth
	isUser := initializers.DB.Where("email = ?", body.Email).First(&existingUser)

	if isUser.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
		logger.Info("A user with this email already exists")
		return
	}

	user := models.Auth{Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		logger.Info("Failed to create user")
		return
	}

	//Generate Token
	tokenString, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		logger.Info("Failed to create token")
		return
	}

	logger.Info("User created successfully")

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	})

	logger.Info("id: ", user.ID, " email: ", user.Email)

}

func SignIn(c *gin.Context, logger *zap.SugaredLogger) {

	var body struct {
		Email    string
		Password string
	}

	logger.Info("Email: ", body.Email)

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Info("Error: ", err.Error())
		return
	}

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		logger.Info("Email and password are required")
		return
	}

	var user models.Auth
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid credentials"})
		logger.Info("Invalid credentials")
		return
	}

	pwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if pwErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid password"})
		logger.Info("Invalid password")
		return
	}

	//Generate Token
	tokenString, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		logger.Info("Failed to create token")
		return
	}

	logger.Info("Signed in successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed in successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"token": tokenString,
		},
	})

	logger.Info("id: ", user.ID, " email: ", user.Email)

}

func SendOtp(c *gin.Context, logger *zap.SugaredLogger) {
	var body struct {
		Email string
	}

	logger.Info("Email: ", body.Email)

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Info("Error: ", err.Error())
		return
	}

	if body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required!"})
		logger.Info("Email is required!")
		return
	}

	// Load html file
	htmlContent, err := ioutil.ReadFile("templates/otp_template.html")
	if err != nil {
		logger.Info("Error: ", err.Error())
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

	var user models.Otp
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {

		result := initializers.DB.Create(&otp)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed send otp",
			})
			logger.Info("Failed send otp")
			return
		}
	} else {
		result := initializers.DB.Model(&user).Updates(models.Otp{Otp: strconv.Itoa(otpValue), ExpiresAt: time.Now().Add(time.Minute * 3)})

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed send otp",
			})
			logger.Info("Failed send otp")
			return
		}
	}

	logger.Info("Email sent successfully!")

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully!",
	})

}

func ResetPassword(c *gin.Context, logger *zap.SugaredLogger) {
	var body struct {
		Email    string
		Otp      string
		Password string
	}

	logger.Info("Email: ", body.Email)
	logger.Info("OTP: ", body.Otp)

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Info("Error: ", err.Error())
		return
	}

	if body.Email == "" || body.Otp == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email, otp and password are required"})
		logger.Info("Email, otp and password are required")
		return
	}

	var otp models.Otp
	initializers.DB.First(&otp, "email = ?", body.Email)

	if otp.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email"})
		logger.Info("Invalid email")
		return
	}

	if otp.Otp != body.Otp || time.Now().After(otp.ExpiresAt) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid OTP"})
		logger.Info("Invalid OTP")
		return
	}

	//Update password

	var user models.Auth
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email"})
		logger.Info("Invalid email")
		return
	}

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	initializers.DB.Model(&user).Updates(models.Auth{Password: string(hashedPassword)})
	logger.Info("Password updated successfully")

	//Delete from OTP table
	initializers.DB.Unscoped().Delete(&otp)

	logger.Info("OTP deleted successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "Password reset successfully",
	})

}
