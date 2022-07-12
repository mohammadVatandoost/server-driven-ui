package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func (s *Server) GetDemoPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()


	APIResponse(c, http.StatusOK, nil, &iamV1.IngPage{
		OrganizationId: ing.OrganizationID,
		Name:           ing.Name,
		Id:             ing.ID,
	}, nil, "", nil)
}