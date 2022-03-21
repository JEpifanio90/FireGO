package FireGO

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

const fireTokenID = "FIREBASE_ID_TOKEN"

func AuthMiddleware() gin.HandlerFunc {
	cli := setup()

	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Missing Authorization header"})

			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Missing token"})

			return
		}

		idToken, err := cli.VerifyIDToken(context.Background(), token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)

			return
		}

		ctx.Set("Authorization", idToken)
		ctx.Next()
	}
}

func ExtractClaims(ctx *gin.Context) *auth.Token {
	idToken, ok := ctx.Get(fireTokenID)
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
