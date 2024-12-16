package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	Post "github.com/phn00dev/go-web-temp-crud/controller/Post"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/", Post.Post{}.IndexPost)
	// add post
	r.GET("/posts/create", Post.Post{}.CreatePost)
	r.POST("/posts/store", Post.Post{}.StorePost)
	// update post
	r.GET("/posts/edit/:id", Post.Post{}.EditPost)
	r.POST("/posts/update/:id", Post.Post{}.UpdatePost)
	//delete post
	r.GET("/posts/delete/:id", Post.Post{}.DeletePost)
	//server files
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/assets/*filepath", http.Dir("assets"))
	return r

}
