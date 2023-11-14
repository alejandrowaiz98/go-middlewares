package http

import (
	"encoding/json"
	"fmt"

	"github.com/alejandrowaiz98/Golang-Middlewares/entitys"
	"github.com/gin-gonic/gin"
)

func Save(ctx *gin.Context) {

	var reqBody []entitys.Billet

	err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody)

	if err != nil {

		err = fmt.Errorf("[ Http | Save | Decode ] error: %v", err)

		ctx.JSON(400, gin.H{"error": err.Error()})

		return

	}

}
