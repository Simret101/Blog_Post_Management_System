package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	blog "practice/blog_service/domain"
	user "practice/user_service/domain"

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post created successfully!", "blog": createdBlog})
	}
}

func (bc *BlogController) GetMyBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		iuser, _ := c.Get("user")
		user := iuser.(user.ResponseUser)

		id := user.ID
		filter := make(map[string]interface{})
		filter["owner._id"] = id

		blogs, err := bc.BlogUsecase.FilterBlog(filter)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve your blog posts. Please try again later: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Your blog posts retrieved successfully!", "blogs": blogs})
	}
}
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
func (bc *BlogController) GetOneBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		blog, err := bc.BlogUsecase.GetOneBlog(id.Hex())
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post retrieved successfully!", "blog": blog})
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post updated successfully!", "blog": updatedBlog})
	}
}
func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Please provide a valid blog ID."})
			return
		}

		err = bc.BlogUsecase.DeleteBlog(id.Hex())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete blog post: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully!"})
	}
}

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter blog posts. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Filtered blog posts retrieved successfully!", "blogs": blogs})
	}
}
func (bc *BlogController) GetUniqueBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve cookie values
		title, err := c.Cookie("title")
		if err != nil && err != http.ErrNoCookie {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve title cookie"})
			return
		}

		author, err := c.Cookie("author")
		if err != nil && err != http.ErrNoCookie {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve author cookie"})
			return
		}

		tagsCookie, err := c.Cookie("tags")
		if err != nil && err != http.ErrNoCookie {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve tags cookie"})
			return
		}

		var tags []string
		if tagsCookie != "" {
			if err := json.Unmarshal([]byte(tagsCookie), &tags); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse tags from cookie"})
				return
			}
		}

		// Prepare filters based on cookies
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

		// Retrieve unique blog posts
		var posts []blog.Blog
		if err := bc.BlogUsecase.GetUniqueBlog(filters, &posts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve unique posts"})
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}
