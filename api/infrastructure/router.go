package infrastructure

import (
	"log"

	"github.com/finder-app/finder-backend/api/interface/controller"

	"github.com/finder-app/finder-backend/api/infrastructure/env"
	"github.com/finder-app/finder-backend/api/infrastructure/firebase"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	engine := gin.Default()

	// NOTE: Cors()を先に設定すること。Auth()で401の場合は処理を中断させるため
	engine.Use(Cors())

	// NOTE: midllewareとしてfirebaseで認証させるように
	firebaseApp := firebase.NewFirebaseApp()
	firebaseClient := firebase.NewClient(firebaseApp)
	engine.Use(Auth(firebaseClient))

	return &Router{
		Engine: engine,
	}
}

// TODO: GraphQLは今後導入予定のため、一旦コメントアウトで使わないようにしておく
// func (r *Router) GraphQL(server *handler.Server, playGroundHandler http.Handler) {
// 	r.Engine.POST("/query", func(c *gin.Context) {
// 		server.ServeHTTP(c.Writer, c.Request)
// 	})
// 	r.Engine.GET("/", func(c *gin.Context) {
// 		playGroundHandler.ServeHTTP(c.Writer, c.Request)
// 	})
// }

func (r *Router) Run() {
	log.Print("http server start")
	r.Engine.Run(":" + env.PORT)
}

// NOTE: routingのテストをするため、router配下に書くこと
func (r *Router) Users(userController *controller.UserController) {
	r.Engine.GET("/users", userController.Index)
	r.Engine.POST("/users", userController.Create)
	r.Engine.GET("/users/:uid", userController.Show)
}

func (r *Router) Profile(profileController *controller.ProfileController) {
	r.Engine.GET("/profile", profileController.Index)
	r.Engine.PUT("/profile", profileController.Update)
}

func (r *Router) FootPrints(footPrintController *controller.FootPrintController) {
	r.Engine.GET("/foot_prints", footPrintController.Index)
	r.Engine.GET("/foot_prints/unread_count", footPrintController.UnreadCount)
}

func (r *Router) Likes(likeController *controller.LikeController) {
	r.Engine.POST("/users/:uid/likes", likeController.Create)
	r.Engine.GET("/likes", likeController.Index)
	r.Engine.PUT("/likes/:sent_uesr_uid/skip", likeController.Skip)
	r.Engine.PUT("/likes/:sent_uesr_uid/consent", likeController.Consent)
	// router.Engine.GET("/likes/recieved", func(c *gin.Context) { likeController.Recieved(c) })
	// router.Engine.GET("/likes/sent", func(c *gin.Context) { likeController.Sent(c) })
}

func (r *Router) Rooms(roomController *controller.RoomController) {
	r.Engine.GET("/rooms", roomController.Index)
}

func (r *Router) Messages(messageController *controller.MessageController) {
	r.Engine.GET("/rooms/:id/messages", messageController.Index)
}
