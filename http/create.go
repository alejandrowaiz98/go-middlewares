package http

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"github.com/gin-gonic/gin"
)

func (h *Http) Save(ctx *gin.Context) {

	var reqBody []entitys.Billet

	log.Println(ctx.Request.Body)

	err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody)

	if err != nil {

		err = fmt.Errorf("[ Http | Save | Decode ] error: %v", err)

		log.Println(err)

		ctx.JSON(400, gin.H{"error": err.Error()})

		return

	}

	errors := h.service.Save(reqBody)

	if len(errors) > 0 {

		err = fmt.Errorf("[ Http | Save | Servuce ] errors: %v", errors)

		ctx.JSON(400, gin.H{"errors": errors})

	}

}
