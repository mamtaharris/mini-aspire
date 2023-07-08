package validators

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func ValidateUnknownParams(reqBody interface{}, ctx *gin.Context) error {
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	if err != nil {
		return err
	}
	payloadBS, err := json.Marshal(&reqBody)
	if err != nil {
		return err
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(payloadBS))
	return nil
}
