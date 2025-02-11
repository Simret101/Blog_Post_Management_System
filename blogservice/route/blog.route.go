package routes

import (
	"log"
	"os"

	"auth/domain"
	"auth/middleware"
	tokenservice "auth/token_service"
	"blog/controller"
	blog1 "blog/database"
	blog "blog/repository"
	blog3 "blog/usecase"
	user1 "user/database"
	user "user/repository"
	user3 "user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewBlogRoutes(group *gin.RouterGroup, blog_collection blog1.CollectionInterface, user_collection user1.CollectionInterface) {
	repo := blog.NewBlogRepository(blog_collection)
	blog_usecase := blog3.NewBlogUsecase(repo)

	user_repo := user.NewUserRepository(user_collection)
	user_usecase := user3.NewUserUseCase(user_repo)
	ctrl := controller.NewBlogController(blog_usecase, user_usecase)

	//load middlewares
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == "" {
		log.Panic("No accesstoken")
	}

	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == "" {
		log.Panic("No refreshtoken")
	}
	TokenSvc := *tokenservice.NewTokenService(access_secret, refresh_secret)

	LoggedInmiddleWare := middleware.LoggedIn(TokenSvc)

	mustOwn := middleware.RoleBasedAccess(access_secret, []domain.Role{domain.Role("creator")}) // Only Admin has access

	group.GET("api/allblog", ctrl.GetAllBlogs())
	group.GET("api/search-blog/", ctrl.FilterBlogs())

	group.GET("api/blog/:id", LoggedInmiddleWare, ctrl.GetOneBlog())
	group.POST("api/blog/", LoggedInmiddleWare, ctrl.CreateBlog())

	group.PUT("api/blog/:id", LoggedInmiddleWare, ctrl.UpdateBlog())
	group.DELETE("api/blog/:id", LoggedInmiddleWare, mustOwn, ctrl.DeleteBlog())

}
