package middleware

import "github.com/gin-gonic/gin"

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "max-age=86400, stale-while-revalidate=172800") // 48小时
	c.Next()
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST")
		c.Header("Allow", "GET,POST")
		//c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		//c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//一些安全设置
// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
