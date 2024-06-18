package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKeyRequestID struct{}

func setCookie(ctx *gin.Context) {
	var sessionID string
	sessionID, err := ctx.Cookie(cookieSessionID)
	if err != nil {
		u, _ := uuid.NewRandom()
		sessionID = u.String()

	}
	// fmt.Printf("setCookie key: %s, value:%s \n", cookieSessionID, sessionID)
	ctx.SetCookie(cookieSessionID, sessionID, cookieMaxAge, "/", "localhost", false, true)
	ctx.Next()
}
