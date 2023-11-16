package http

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (http *Http) Run() {

	r := gin.Default()

	r.Use(validateHeaders())

	r.POST("/Save", http.Save)

	port := ":8989"

	log.Printf("Running in port%v", port)

	r.Run(port)

}

// ValidateHeaders validate corp headers
func validateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		headers := strings.Split(os.Getenv("headers"), ",")

		log.Printf("%v headers", len(headers))
		log.Printf("estos son los headers: %v", headers)

		for _, value := range headers {

			if _, m := c.Request.Header[value]; m != true {

				log.Printf("Invalid header: %v", c.Request.Header)
				log.Printf("Este header está webeando: %v", c.Request.Header[value])

				c.AbortWithStatus(http.StatusBadRequest)

				break
			}
		}
	}
}
