package setup

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutConfig() *gin.Engine {
	router := gin.Default()
	router.Use(
		recoveryMiddleware(),
	)
	rootRouters(router)

	return router

}

func SendErrorWithMessage(w http.ResponseWriter, status int, msg string, err error) {
	resp := map[string]interface{}{
		"msg": msg,
		"err": err.Error(),
	}
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(resp)
}

func recoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//Add Loggers
				SendErrorWithMessage(c.Writer, http.StatusInternalServerError, "Some thing went wrong", errors.New("unhandled Panics"))
				return
			}
		}()
		c.Next()
	}
}

func organizeController(handler func(context.Context, http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c.Copy(), c.Writer, c.Request)
		c.Writer.Flush()
	}
}

func organizeControllerwithGin(handler func(context.Context, http.ResponseWriter, *http.Request, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c.Copy(), c.Writer, c.Request, c)
		c.Writer.Flush()
	}
}
