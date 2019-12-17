package main

import(
  "net/http"
  "html/template"
  "fmt"
  "Post"
  "Create"
  "strconv"
  "Read"
  "Delete"
)

var tpl *template.Template
var PostById map[int]*Post.NewPost

func init() {
  tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

  PostById = make(map[int]*Post.NewPost)
  http.Handle("/assets/", http.StripPrefix("/assets/",http.FileServer(http.Dir("assets"))))
  http.HandleFunc("/", index)
  http.HandleFunc("/form", form)
  http.HandleFunc("/addPost", addPost)
  http.HandleFunc("/readRequest", readRequest)
  http.HandleFunc("/readPost", readPost)
  http.HandleFunc("/updateRequest", updateRequest)
  http.HandleFunc("/editPost", editPost)
  http.HandleFunc("/deleteRequest", deleteRequest)
  http.HandleFunc("/deletePost", deletePost)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}


func form(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "form.html", nil)
}


func addPost(w http.ResponseWriter, r *http.Request) {
  var post Post.NewPost
  post.Id,_ = strconv.Atoi(r.FormValue("id"))
  post.Title = r.FormValue("title")
  post.Author = r.FormValue("author")
  post.Content = r.FormValue("content")

  fmt.Println(post)

  err := Create.AddPost(post, PostById)
  if err!=nil {
    fmt.Println("Post not added. Error: ", err)
  }
  tpl.ExecuteTemplate(w, "addPost.html", PostById[post.Id])
}

func readRequest(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "readRequest.html", nil)
}

func readPost(w http.ResponseWriter, r *http.Request) {
  id,_ := strconv.Atoi(r.FormValue("id"))
  post, err := Read.ReadPostById(id,PostById)

  fmt.Println(post)

  if err!=nil {
    fmt.Println("Post not read. Error: ", err)
  }
  tpl.ExecuteTemplate(w, "readPost.html", PostById[post.Id])
}

func updateRequest(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "updateRequest.html", nil)
}

func editPost(w http.ResponseWriter, r *http.Request) {
  id,_ := strconv.Atoi(r.FormValue("id"))
  post, err := Read.ReadPostById(id,PostById)
  err = Delete.DeletePostById(id,PostById)
  if err==nil {
    tpl.ExecuteTemplate(w,"editPost.html", post)
  }
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "deleteRequest.html", nil)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
  id,_ := strconv.Atoi(r.FormValue("id"))
  err := Delete.DeletePostById(id,PostById)
  tpl.ExecuteTemplate(w,"deletePost.html", err)
}
