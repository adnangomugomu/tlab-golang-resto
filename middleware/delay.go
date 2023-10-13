package middleware

import (
	"github.com/gin-gonic/gin"
)

func DelayApi(c *gin.Context) {
	// durasiPenundaan := 2 * time.Second
	// time.Sleep(durasiPenundaan)
	c.Next()
}
