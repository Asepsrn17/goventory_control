package middleware

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/config"
	"go_inven_ctrl/entity"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("admin_ic")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		
		claims := token.Claims.(jwt.MapClaims)
		c.Set("claims", claims)
		c.Next()
	}
}

func Login(c *gin.Context) {
	var adminIc entity.AdminIc

	if err := c.ShouldBindJSON(&adminIc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbAdminIc, err := getEmailAdminIc(adminIc.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unregister admin ic"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbAdminIc.Password), []byte(adminIc.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = dbAdminIc.Email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


func getEmailAdminIc( email string) (*entity.AdminIc, error) {
	var adminIc entity.AdminIc

	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()


	err = db.QueryRow("SELECT id, name, email, password, phone, photo FROM ic_team WHERE email= $1", email).Scan(&adminIc.ID, &adminIc.Name, &adminIc.Email, &adminIc.Password, &adminIc.Phone, &adminIc.Photo)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Admin ic not found in database with email:", email)
			return nil, fmt.Errorf("admin ic with email %s not found", email)
		}
		fmt.Println("Error retrieving admin IC from database:", err)
		return nil, err
	}
	fmt.Println("Admin IC retrieved from database:", adminIc)
	return &adminIc, nil
}

func Profile(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)

	dbAdminIc, err := getEmailAdminIc(email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get admin IC info"})
		return
	}

	response := map[string]interface{} {
		"message" : "Welcome Admin Ic",
		"email" : dbAdminIc.Email,
		"name": dbAdminIc.Name,
		"phone": dbAdminIc.Phone,
		"photo": dbAdminIc.Photo,
	}
	c.JSON(http.StatusOK, response)
}