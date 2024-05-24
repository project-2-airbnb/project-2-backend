package handler

import (
	"net/http"
	"project-2/app/middlewares"
	"project-2/features/review"
	"project-2/utils/responses"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	ReviewService review.ReviewService
}

func New(rh review.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		ReviewService: rh,
	}
}

func (rh *ReviewHandler) AddReview(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse("Unauthorized", nil))
	} 

	newReview := ReviewRequest{}
	errBind := c.Bind(&newReview)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error binding data: "+errBind.Error(), nil))
	}

	review := review.Review{
		UserID: uint(userID),
		RoomID: newReview.RoomID,
		Rating: newReview.Rating,
		Comment: newReview.Comment,
	}

	err := rh.ReviewService.AddReview(review)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("gagal membuat review: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("membuat review berhasil", nil))
}
