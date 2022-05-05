package handler

type ItemForCreatingShortURLRequest struct {
	OriginURL string `json:"originURL" binding:"required"`
}

type ItemForCreatingShortURLResponse struct {
	ShortURl string `json:"shortURL" binding:"required"`
}

type ItemForGettingOriginURL struct {
	ShortURL string `form:"shortURL" binding:"required"`
}
