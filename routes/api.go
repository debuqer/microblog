package routes

import (
	"fmt"
	"net/http"

	"github.com/debuqer/microblog/controllers"
	"github.com/julienschmidt/httprouter"
)

func ApiRouter() *httprouter.Router {
	router := httprouter.New()

	// Tags
	// get tags
	router.GET("/tag/all", controllers.TagAll)
	// store
	router.POST("/tag/store", controllers.TagStore)
	// show tag
	router.GET("/tag/:tgid", controllers.TagShow)

	/** Timeline **/
	// show timeline
	router.GET("/timeline/:slug", controllers.TimelineShow)
	// store
	router.POST("/timeline/store", controllers.TimelineStore)
	// update
	router.PUT("/timeline/update", controllers.TimelineUpdate)
	// clone into
	router.POST("/timeline/clone", controllers.TimelineClone)
	// merge
	router.PUT("/timeline/merge", controllers.TimelineMerge)
	// star
	router.POST("/timeline/reaction", controllers.TimelineReaction)
	// load events
	router.GET("/timeline/events/:scale", controllers.TimelineEvents)
	// latest updated
	router.GET("/timeline/latest", controllers.TimelineLatest)

	/** Event **/
	// show event
	router.GET("/timeline/event/:eid", controllers.EventShow)
	// store
	router.POST("/timeline/event/store", controllers.EventStore)
	// update
	router.PUT("/timeline/event/update", controllers.EventUpdate)
	// clone into
	router.POST("/timeline/event/clone", controllers.EventClone)
	// star
	router.POST("/timeline/event/reaction", controllers.EventReaction)
	// comment
	router.POST("/timeline/event/comment", controllers.EventComment)
	// load comments
	router.GET("/timeline/event/comments", controllers.EventComments)

	/** User **/
	// show user
	router.GET("/user/:uid", controllers.UserShow)
	// update user
	router.PUT("/user/update", controllers.UserUpdate)

	/** Home **/
	// sign up
	router.POST("/signup", hello)
	// login
	router.POST("/login", hello)
	// forget password
	router.GET("/forget-password", controllers.UserForgetPassword)
	// verify link
	router.GET("/verify-link", controllers.UserVerifyLink)
	// update password
	router.POST("/update-password", controllers.UserUpdatePassword)

	return router
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
