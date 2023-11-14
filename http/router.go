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

	port := ":8989"

	log.Printf("Running in port%v", port)

	r.Run(port)

}

// ValidateHeaders validate corp headers
func validateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		headers := strings.Split(os.Getenv("Headers"), ",")

		log.Printf("%v headers", len(headers))

		for _, value := range headers {

			if _, m := c.Request.Header[value]; m != true {

				log.Printf("Invalid header: %v", c.Request.Header)

				c.AbortWithStatus(http.StatusBadRequest)

				break
			}
		}
	}
}
