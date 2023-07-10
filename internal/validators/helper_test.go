package validators

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_validator_ValidateUnknownParams(t *testing.T) {
	type args struct {
		reqBody interface{}
		ctx     *gin.Context
	}
	tests := []struct {
		name    string
		v       *validator
		args    args
		wantErr bool
	}{
		{
			name: "happy case",
			args: args{ctx: &gin.Context{
				Request: &http.Request{Body: io.NopCloser(bytes.NewBuffer([]byte(`{
				"status": "APPROVED"
			}`)))},
				Params: gin.Params{gin.Param{Key: "id", Value: "123"}}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			if err := v.ValidateUnknownParams(tt.args.reqBody, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("validator.ValidateUnknownParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
