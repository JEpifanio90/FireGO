package models

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type FireAuth struct {
	cli          *auth.Client
	unAuthorized func(c *gin.Context)
}
