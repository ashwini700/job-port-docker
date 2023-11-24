package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_handler_AcceptApplicants(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *handler
		args args
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.AcceptApplicants(tt.args.c)
		})
	}
}
