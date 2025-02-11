package controller

import (
	"net/http"

	"auth/config"
	"blog/domain"

	"github.com/gin-gonic/gin"
)

type BookmarkController struct {
	BookmarkUsecase domain.BookmarkUseCaseInterface
}

func NewBookmarkController(usecase domain.BookmarkUseCaseInterface) *BookmarkController {
	return &BookmarkController{
		BookmarkUsecase: usecase,
	}
}

// AddBookmark adds a bookmark for a user
//
//	@Summary		Add a bookmark
//	@Description	Add a bookmark for a user
//	@Tags			Bookmarks
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"User ID"
//	@Param			blogID	path		string	true	"Blog ID"
//	@Success		200		{string}	string	"Bookmark added successfully"
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/add/{userID}/{blogID} [post]
func (bc *BookmarkController) AddBookmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		blogID := c.Param("blogID")

		if err := bc.BookmarkUsecase.AddBookmark(userID, blogID); err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bookmark added successfully"})
	}
}

// RemoveBookmark removes a bookmark for a user
//
//	@Summary		Remove a bookmark
//	@Description	Remove a bookmark for a user
//	@Tags			Bookmarks
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"User ID"
//	@Param			blogID	path		string	true	"Blog ID"
//	@Success		200		{string}	string	"Bookmark removed successfully"
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/remove/{userID}/{blogID} [delete]
func (bc *BookmarkController) RemoveBookmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		blogID := c.Param("blogID")

		if err := bc.BookmarkUsecase.RemoveBookmark(userID, blogID); err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bookmark removed successfully"})
	}
}

// GetUserBookmarks retrieves all bookmarks for a user
//
//	@Summary		Get user bookmarks
//	@Description	Get all bookmarks for a user
//	@Tags			Bookmarks
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		string	true	"User ID"
//	@Success		200		{array}		domain.Bookmark
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/getbook/{userID} [get]
func (bc *BookmarkController) GetUserBookmarks() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")

		bookmarks, err := bc.BookmarkUsecase.GetUserBookmarks(userID)
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, bookmarks)
	}
}
