package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiRouter() *httprouter.Router {
	router := httprouter.New()

	// Tags
	// get tags
	router.GET("/tag/all", hello)
	// store
	router.POST("/tag/store", hello)
	// show tag
	router.GET("/tag/:tgid", hello)

	/** Timeline **/
	// show timeline
	router.GET("/timeline/:slug", hello)
	// store
	router.POST("/timeline/store", hello)
	// update
	router.PUT("/timeline/update", hello)
	// clone into
	router.POST("/timeline/clone", hello)
	// merge
	router.PUT("/timeline/merge", hello)
	// star
	router.POST("/timeline/reaction", hello)
	// load events
	router.GET("/timeline/events/:scale", hello)
	// latest updated
	router.GET("/timeline/latest", hello)

	/** Event **/
	// show event
	router.GET("/timeline/event/:eid", hello)
	// store
	router.POST("/timeline/event/store", hello)
	// update
	router.PUT("/timeline/event/update", hello)
	// clone into
	router.POST("/timeline/event/clone", hello)
	// star
	router.POST("/timeline/event/reaction", hello)
	// comment
	router.POST("/timeline/event/comment", hello)
	// load comments
	router.GET("/timeline/event/comments", hello)

	/** User **/
	// show user
	router.GET("/user/:uid", hello)
	// update user
	router.PUT("/user/update", hello)

	/** Home **/
	// sign up
	router.POST("/signup", hello)
	// login
	router.POST("/login", hello)
	// forget password
	router.GET("/forget-password", hello)
	// verify link
	router.GET("/verify-password", hello)
	// update password
	router.POST("/update-password", hello)

	return router
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
