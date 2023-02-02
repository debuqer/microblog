package routes

import (
	"fmt"
	"net/http"

	c "github.com/debuqer/microblog/controllers"
	"github.com/julienschmidt/httprouter"
)

func ApiRouter() *httprouter.Router {
	router := httprouter.New()

	// Tags
	// get tags
	router.GET("/tag/all", c.TagAll)
	// store
	router.POST("/tag/store", c.TagStore)
	// show tag
	router.GET("/tag/:tgid", c.TagShow)

	/** Timeline **/
	// show timeline
	router.GET("/timeline/:slug", c.TimelineShow)
	// store
	router.POST("/timeline/store", c.TimelineStore)
	// update
	router.PUT("/timeline/update", c.TimelineUpdate)
	// clone into
	router.POST("/timeline/clone", c.TimelineClone)
	// merge
	router.PUT("/timeline/merge", c.TimelineMerge)
	// star
	router.POST("/timeline/reaction", c.TimelineReaction)
	// load events
	router.GET("/timeline/events/:scale", c.TimelineEvents)
	// latest updated
	router.GET("/timeline/latest", c.TimelineLatest)

	/** Event **/
	// show event
	router.GET("/timeline/event/:eid", c.EventShow)
	// store
	router.POST("/timeline/event/store", c.EventStore)
	// update
	router.PUT("/timeline/event/update", c.EventUpdate)
	// clone into
	router.POST("/timeline/event/clone", c.EventClone)
	// star
	router.POST("/timeline/event/reaction", c.EventReaction)
	// comment
	router.POST("/timeline/event/comment", c.EventComment)
	// load comments
	router.GET("/timeline/event/comments", c.EventComments)

	/** User **/
	// show user
	router.GET("/user/:uid", c.UserShow)
	// update user
	router.PUT("/user/update", c.UserUpdate)

	/** Home **/
	// sign up
	router.POST("/signup", hello)
	// login
	router.POST("/login", hello)
	// forget password
	router.GET("/forget-password", c.UserForgetPassword)
	// verify link
	router.GET("/verify-link", c.UserVerifyLink)
	// update password
	router.POST("/update-password", c.UserUpdatePassword)

	return router
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
