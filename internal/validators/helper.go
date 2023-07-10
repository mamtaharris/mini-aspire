package validators

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type validator struct{}

func NewValidator() ValidatorInterface {
	return &validator{}
}

//go:generate mockgen -package mocks -source=helper.go -destination=mocks/helper_mocks.go
type ValidatorInterface interface {
	ValidateUnknownParams(reqBody interface{}, ctx *gin.Context) error
}

func (v *validator) ValidateUnknownParams(reqBody interface{}, ctx *gin.Context) error {
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
