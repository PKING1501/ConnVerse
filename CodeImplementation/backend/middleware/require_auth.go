package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Naman-K-Jaiswal/ConnVerse/authn"
	"github.com/Naman-K-Jaiswal/ConnVerse/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RequireAuth(c *gin.Context) {
	tokenStr, err := c.Cookie("Authorization")
	fmt.Println(tokenStr)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	fmt.Println(token, err, "yaha pe")

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user authn.LoginDetails
		collection := database.DB.Collection("login_details")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		err := collection.FindOne(ctx, bson.M{"email": claims["sub"]}).Decode(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
