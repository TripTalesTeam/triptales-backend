package route

import (
	"github.com/breezjirasak/triptales/internal/handler"
	"github.com/breezjirasak/triptales/internal/middleware"
	"github.com/breezjirasak/triptales/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authService *service.AuthService, userService *service.UserService, countryService *service.CountryService, tripService *service.TripService,
	tripCompanionService *service.TripCompanionService, friendService *service.FriendService, bookmarkService *service.BookmarkService) *gin.Engine {

	r := gin.Default()

	// Auth routes
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	countryHandler := handler.NewCountryHandler(countryService)
	tripHandler := handler.NewTripHandler(tripService)
	tripCompanionHandler := handler.NewTripCompanionHandler(tripCompanionService)
	friendHandler := handler.NewFriendHandler(friendService)
	bookmarkHandler := handler.NewBookmarkHandler(bookmarkService)
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)

		// Protected auth routes
		protected := authGroup.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("/me", authHandler.GetMe)
		}
	}

	// API routes that require authentication
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware())
	{
		user := api.Group("/users")
		{
			user.GET("/", userHandler.GetUsers)
			user.DELETE("/:id", userHandler.DeleteUser)
			user.PUT("/update", userHandler.UpdateUser)
		}

		country := api.Group("/countries")
		{
			country.POST("/", countryHandler.CreateCountry)
			country.GET("/", countryHandler.GetAllCountries)
			country.GET("/by-name", countryHandler.GetCountryByName)
			country.GET("/:id", countryHandler.GetCountryByID)
			country.PUT("/", countryHandler.UpdateCountry)
			country.DELETE("/:id", countryHandler.DeleteCountry)
		}

		trip := api.Group("/trips")
		{
			trip.POST("/", tripHandler.CreateTrip)
			trip.GET("/", tripHandler.GetAllTrips)
			trip.GET("/:id", tripHandler.GetTripByID)
			trip.GET("/friend", tripHandler.GetFriendTrip)
			trip.GET("/bookmark", tripHandler.GetBookmarkTrip)
			trip.PUT("/:id", tripHandler.UpdateTrip)
			trip.DELETE("/:id", tripHandler.DeleteTrip)
		}

		tripCompanion := api.Group("/trip-companions")
		{
			tripCompanion.POST("", tripCompanionHandler.AddCompanion)
			tripCompanion.DELETE("/:tripId/:userId", tripCompanionHandler.RemoveCompanion)
			tripCompanion.GET("/:tripId", tripCompanionHandler.GetCompanions)
		}
		friend := api.Group("/friends")
		{
			friend.POST("/", friendHandler.AddFriend)
			friend.GET("/", friendHandler.GetFriends)
			friend.DELETE("/:friend_id", friendHandler.RemoveFriend)
		}

		bookmark := api.Group("/bookmarks")
		{
			bookmark.POST("/", bookmarkHandler.AddBookmark)
			bookmark.DELETE("/:id", bookmarkHandler.RemoveBookmark)
			bookmark.GET("/", bookmarkHandler.GetBookmarks)
		}
	}

	return r
}
