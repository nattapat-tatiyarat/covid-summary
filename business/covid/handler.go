package covid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CovidHandler struct {
	svc Service
}

func NewCovidHandler(covidService Service) *CovidHandler {
	return &CovidHandler{
		svc: covidService,
	}
}

func (h *CovidHandler) GetCovidSummary(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := h.svc.GetCovidSummary(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
