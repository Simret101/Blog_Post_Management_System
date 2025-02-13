package controller

import (
	"net/http"
	"strconv"

	"auth/config" // Import the config package for error handling
	blog "blog/domain"
	user "user/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase blog.Blog_Usecase_interface
	UserUsecase user.User_Usecase_interface
}

func NewBlogController(uc blog.Blog_Usecase_interface, uu user.User_Usecase_interface) *BlogController {
	return &BlogController{
		BlogUsecase: uc,
		UserUsecase: uu,
	}
}

// CreateBlog creates a new blog post
//
//	@Summary		Create a new blog post
//	@Description	Create a new blog post
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			blog	body		blog.PostBlog	true	"Blog Post"
//	@Success		200		{object}	blog.Blog
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/api/blog [post]
func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var blog blog.PostBlog
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request: " + err.Error()})
			return
		}

		// Create a new blog post
		createdBlog, err := bc.BlogUsecase.CreateBlog(blog)
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post created successfully!", "blog": createdBlog})
	}
}

// GetAllBlogs retrieves all blogs with pagination
//
//	@Summary		Get all blogs with pagination
//	@Description	Get all blogs with pagination
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page number"
//	@Param			limit	query		int	false	"Limit"
//	@Success		200		{array}		blog.Blog
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/api/allblog [get]
func (bc *BlogController) GetAllBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 5
		}

		posts, err := bc.BlogUsecase.GetBlogs(page, limit)
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": posts,
			"meta": gin.H{
				"limit": limit,
				"page":  page,
			},
		})
	}
}

// GetOneBlog retrieves a single blog post by ID
//
//	@Summary		Get a single blog post by ID
//	@Description	Get a single blog post by ID
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Blog ID"
//	@Success		200	{object}	blog.Blog
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/blog/{id} [get]
func (bc *BlogController) GetOneBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": config.ErrInvalidToken.Error()})
			return
		}

		blog, err := bc.BlogUsecase.GetOneBlog(id.Hex())
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post retrieved successfully!", "blog": blog})
	}
}

// UpdateBlog updates a blog post by ID
//
//	@Summary		Update a blog post by ID
//	@Description	Update a blog post by ID
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Blog ID"
//	@Param			blog	body		blog.Blog	true	"Blog Post"
//	@Success		200		{object}	blog.Blog
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/api/blog/{id} [put]
func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": config.ErrInvalidToken.Error()})
			return
		}

		var blog blog.Blog
		if err := c.BindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request: " + err.Error()})
			return
		}

		blog.ID = id

		updatedBlog, err := bc.BlogUsecase.UpdateBlog(id.Hex(), blog)
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully!", "blog": updatedBlog})
	}
}

// DeleteBlog deletes a blog post by ID
//
//	@Summary		Delete a blog post by ID
//	@Description	Delete a blog post by ID
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Blog ID"
//	@Success		200	{string}	string	"Blog post deleted successfully!"
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/blog/{id} [delete]
func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": config.ErrInvalidToken.Error()})
			return
		}

		err = bc.BlogUsecase.DeleteBlog(id.Hex())
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully!"})
	}
}

// FilterBlogs filters blog posts based on query parameters
//
//	@Summary		Filter blog posts
//	@Description	Filter blog posts based on query parameters
//	@Tags			Blogs
//	@Accept			json
//	@Produce		json
//	@Param			title		query		string		false	"Title"
//	@Param			author		query		string		false	"Author"
//	@Param			tags		query		[]string	false	"Tags"
//	@Param			start_date	query		string		false	"Start Date"
//	@Param			end_date	query		string		false	"End Date"
//	@Param			popularity	query		int			false	"Popularity"
//	@Success		200			{array}		blog.Blog
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		500			{object}	httputil.HTTPError
//	@Router			/api/search-blog/ [get]
func (bc *BlogController) FilterBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		author := c.Query("author")

		tags := c.QueryArray("tags")
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")
		popularity := c.Query("popularity")

		filters := make(map[string]interface{})

		if title != "" {
			filters["title"] = title
		}
		if author != "" {
			filters["owner.username"] = author
		}

		if len(tags) > 0 {
			filters["tag"] = tags
		}
		if startDate != "" {
			filters["created_at"] = bson.M{"$gte": startDate}
		}
		if endDate != "" {
			filters["created_at"] = bson.M{"$lte": endDate}
		}

		if popularity != "" {
			popularityValue, err := strconv.Atoi(popularity)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid popularity value."})
				return
			}
			filters["like_count"] = bson.M{"$gte": popularityValue}
		}

		blogs, err := bc.BlogUsecase.FilterBlog(filters)
		if err != nil {
			c.JSON(config.GetStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Filtered blog posts retrieved successfully!", "blogs": blogs})
	}
}
