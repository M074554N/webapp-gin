package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", apiRoot)

		products := v1.Group("/products")
		{
			products.GET("/", productsIndex)
			products.GET("/:id", productsGet)
			products.POST("/", productsStore)
			products.PUT("/:id", productsUpdate)
			products.DELETE("/:id", productsDelete)
		}
	}

	r.Run(":80")
}

func apiRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "v1 home",
	})
}

func productsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Products",
	})
}

func productsGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get product with id: " + c.Param("id"),
	})
}

func productsStore(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "New product added",
	})
}

func productsUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Product updated",
	})
}

func productsDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Product deleted",
	})
}
