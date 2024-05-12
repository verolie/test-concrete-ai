package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Authentication Empty" , "")))
            return 
        }
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("API_SECRET")), nil
    })
    println(err.Error())

        if err != nil || !token.Valid {
            c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Token Invalid" , "")))
            return
        }

        println(err.Error())
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if emailClaim, ok := claims["email"]; ok {
                if email, ok := emailClaim.(string); ok {

                    c.Set("email", email)
                    println(email)
                    c.Next()
                    return
                }
            }
        }
    }
}