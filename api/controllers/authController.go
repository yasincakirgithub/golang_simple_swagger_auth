package controllers

import (
	"GOLANG/api/initializers"
	"GOLANG/api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Tags users
// @Summary Register for users
// @Schemes
// @Accept json
// @Produce json
// @Param register body models.RegisterInput true "User Register"
// @Success 200 {object} models.User
// @Router /users/register [post]
func Register(c *gin.Context) {

	var RegisterInput models.RegisterInput

	if err := c.ShouldBindJSON(&RegisterInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", RegisterInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(RegisterInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: RegisterInput.Username,
		Password: string(passwordHash),
		Email:    RegisterInput.Email,
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})

}

// @Tags users
// @Summary Auth for users
// @Schemes
// @Accept json
// @Produce json
// @Param auth body models.AuthInput true "User Auth"
// @Success 200 {object} models.TokenResponse
// @Router /users/auth [post]
func Auth(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// @Tags users
// @Summary User Detail
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users/me [get]
// @Security BearerAuth
func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"user": user,
	})
}
