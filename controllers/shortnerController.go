package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/satheeshds/shortly/interfaces"
)

type ShortnerController struct {
	Service interfaces.IShortnerService
}

// @Summary Create a short url for the given url
// @Description Create a short url for the given url
// @ID create-short
// @Accept x-www-form-urlencoded
// @Param req_url formData string true "url"
// @Produce json
// @Success 201 {string} string "new url"
// @Router /shorten [post]
func (s *ShortnerController) Short(ctx *gin.Context) {
	originalUrl := ctx.PostForm("req_url")
	if _, err := url.Parse(originalUrl); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result, err := s.Service.ShortURL(originalUrl); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"short_url": fmt.Sprintf("http://%s/%s", ctx.Request.Host, result)})
	}

	return
}

func (s *ShortnerController) Redirect(ctx *gin.Context) {
	shortId := ctx.Param("shortId")

	longUrl, err := s.Service.GetRedirectURL(shortId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the long URL
	ctx.Redirect(http.StatusMovedPermanently, longUrl)

}

// @Summary Get top domains shortened
// @Description Get top domains shortened
// @Produce json
// @Router /shorten/topDomains [get]
func (s *ShortnerController) GetTopDomains(ctx *gin.Context) {

	res, err := s.Service.GetTopShortedDomains()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
