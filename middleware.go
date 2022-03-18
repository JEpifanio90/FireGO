package FireGO

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"strings"
)

const fireTokenID = "FIREBASE_ID_TOKEN"

func AuthMiddleware() gin.HandlerFunc {
	cli := setup()

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := cli.VerifyIDToken(context.Background(), token)
		if err != nil {
			log.Println("FUUUUUCK")
			log.Fatalln(err.Error())
			//if fam.unAuthorized != nil {
			//	fam.unAuthorized(c)
			//} else {
			//	c.JSON(http.StatusUnauthorized, gin.H{
			//		"status":  http.StatusUnauthorized,
			//		"message": http.StatusText(http.StatusUnauthorized),
			//	})
			//}
			return
		}
		c.Set(fireTokenID, idToken)
		c.Next()
	}
}

func ExtractClaims(c *gin.Context) *auth.Token {
	idToken, ok := c.Get(fireTokenID)
	if !ok {
		return new(auth.Token)
	}

	return idToken.(*auth.Token)
}

func setup() *auth.Client {
	clientOption := option.WithCredentialsFile("./credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, clientOption)

	if err != nil {
		log.Fatalln(err.Error())
	}

	auth, err := app.Auth(context.Background())

	if err != nil {
		log.Fatalln(err.Error())
	}

	return auth
}
