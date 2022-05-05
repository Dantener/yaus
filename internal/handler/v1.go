package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaus/internal/entity"
)

// @Summary		 Метод для генерации сокращённого URL
// @Description	 Метод для сокращения полного URL в короткий
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        input body ItemForCreatingShortURLRequest	true "Исходный URL"
// @Success      200  {object} ItemForCreatingShortURLResponse
// @Router       /v1/short/url [post]
func (h *Handler) createShortURL(c *gin.Context) {
	var input ItemForCreatingShortURLRequest
	var scheme string

	if c.Request.TLS != nil {
		scheme = "https"
	} else {
		scheme = "http"
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := entity.Item{Id: 0, OriginURL: input.OriginURL, ShortURL: ""}
	data.ShortURL = scheme + "://" + c.Request.Host + "/"

	res, err := h.services.Convertation.CreateShortURL(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := ItemForCreatingShortURLResponse{ShortURl: res}
	c.JSON(http.StatusOK, response)
}

// @Summary		 Метод для получения полного URL
// @Description	 Метод для преобразования короткого URL в полный
// @Tags         v1
// @Accept       json
// @Produce		 html
// @Param        input query ItemForGettingOriginURL	true "Сокращённый URL"
// @Success      301
// @Router       /v1/short/url [get]
func (h *Handler) getURLByShortURL(c *gin.Context) {
	var input ItemForGettingOriginURL
	if err := c.BindQuery(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data := entity.Item{Id: 0, OriginURL: "", ShortURL: input.ShortURL}
	res, err := h.services.Convertation.GetOriginURLByShortURL(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, res)
}
