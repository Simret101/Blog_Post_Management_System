package route

import (
	"practice/blog_service/controller"
	blog "practice/blog_service/repository"
	bloguse "practice/blog_service/usecase"
	"practice/database"
	user "practice/user_service/repository"
	useruse "practice/user_service/usecase"

	"github.com/gin-gonic/gin"
)

func NewBlogRoutes(group *gin.RouterGroup, blogCollection database.CollectionInterface, userCollection database.CollectionInterface) {
	blogRepository := blog.NewBlogRepository(blogCollection)
	blogUseCase := bloguse.NewBlogUsecase(blogRepository)

	userRepository := user.NewUserRepository(userCollection)
	userUseCase := useruse.NewUserUseCase(userRepository)
	blogController := controller.NewBlogController(blogUseCase, userUseCase)

	group.GET("api/allblog", blogController.GetAllBlogs())
	//group.GET("api/search-blog/", blogController.FilterBlogs())
	group.GET("api/blog/:id", blogController.GetOneBlog())
	group.POST("api/blog/", blogController.CreateBlog())
	group.PUT("api/blog/:id", blogController.UpdateBlog())
	group.DELETE("api/blog/:id", blogController.DeleteBlog())
}
