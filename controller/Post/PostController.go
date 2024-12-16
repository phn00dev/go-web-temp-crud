package Post

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"github.com/phn00dev/go-web-temp-crud/helpers"
	"github.com/phn00dev/go-web-temp-crud/models"
)

type Post struct{}

func (post Post) IndexPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("postIndex").Funcs(template.FuncMap{
		"getDate": func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d", t.Day(), int(t.Month()), t.Year())
		},
	}).ParseFiles(helpers.Include("Post/index")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["posts"] = models.Post{}.GetAllPost()
	data["alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "postIndex", data)
}

func (post Post) CreatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("Post/create")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "createPost", nil)
}

func (post Post) StorePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("postTitle")
	slug := slug.Make(title)
	desc := r.FormValue("postDesc")
	status := r.FormValue("postStatus")
	// uploads begins
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("postImage")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/postImages/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	//uploads end
	models.Post{
		Title:    title,
		Slug:     slug,
		Desc:     desc,
		Status:   status,
		ImageUrl: "uploads/postImages/" + header.Filename,
	}.Create()
	helpers.SetAlert(w, r, "Post Döredildi !")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (post Post) EditPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("Post/update")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["post"] = models.Post{}.GetPost(params.ByName("id"))
	view.ExecuteTemplate(w, "editPost", data)
}

func (post Post) UpdatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post_update := models.Post{}.GetPost(params.ByName("id"))
	title := r.FormValue("postTitle")
	slug := slug.Make(title)
	desc := r.FormValue("postDesc")
	status := r.FormValue("postStatus")
	is_selected := r.FormValue("isSelected")

	var image_url string

	if is_selected == "1" {
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("postImage")
		if err != nil {
			fmt.Println(err)
		}
		f, err := os.OpenFile("uploads/postImages/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
		}
		image_url = "uploads/postImages/" + header.Filename
		os.Remove(post_update.ImageUrl)
	} else {
		image_url = post_update.ImageUrl
	}
	post_update.Updates(models.Post{
		Title:    title,
		Slug:     slug,
		Desc:     desc,
		Status:   status,
		ImageUrl: image_url,
	})
	helpers.SetAlert(w, r, "Post üýtgedildi !")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (post Post) DeletePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post_delete := models.Post{}.GetPost(params.ByName("id"))
	// Remove the file
	err := os.Remove(post_delete.ImageUrl)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	post_delete.Delete()

	helpers.SetAlert(w, r, "Post öçürildi !")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
