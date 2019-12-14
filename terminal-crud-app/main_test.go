package main

import (
  "Create"
  "Delete"
  "fmt"
  "Post"
  "Read"
  "testing"
  "Update"
  )

var posts []Post.NewPost


func TestCreatePost(t *testing.T) {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)

  posts = append(posts, Post.NewPost{Id:1, Content: "Hello World!", Author: "Sherlock"})
  posts = append(posts, Post.NewPost{Id:1, Content: "Bonjour Madam!", Author: "M. Pierre"})
  posts = append(posts, Post.NewPost{Id:3, Content: "Hola Amigos", Author: "Prof. Sergio"})
  posts = append(posts, Post.NewPost{Id:4, Content: "Boring...", Author: "Sherlock"})

  for _,post := range posts {
    var err1 error
    err1 = Create.AddPost(post,PostById)
    if err1==nil {
      // Check if the entry really exists and all values are same
      if (PostById[post.Id].Id!=post.Id) || (PostById[post.Id].Author!=post.Author) || (PostById[post.Id].Content!=post.Content) {
        t.Errorf("Post was not added!")
      }
    } else {
      fmt.Println(post,err1)
    }
  }
}


func TestReadPost(t *testing.T) {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)

  posts = append(posts, Post.NewPost{Id:1, Content: "Hello World!", Author: "Sherlock"})
  posts = append(posts, Post.NewPost{Id:1, Content: "Bonjour Madam!", Author: "M. Pierre"})
  posts = append(posts, Post.NewPost{Id:3, Content: "Hola Amigos", Author: "Prof. Sergio"})
  posts = append(posts, Post.NewPost{Id:4, Content: "Boring...", Author: "Sherlock"})

  for _,post := range posts {
    var err1 error
    err1 = Create.AddPost(post,PostById)
    if err1==nil {
      // Check if the entry is being read
      ReadPost,err2 := Read.ReadPostById(post.Id, PostById)
      if (err2==nil) {
        if(ReadPost.Id!=post.Id) || (ReadPost.Author!=post.Author) || (ReadPost.Content!=post.Content){
          t.Errorf("Post added was not read.")
        }
      }else {
        fmt.Println(post, err2)
      }
    }
  }
}


func TestReadAllPosts(t *testing.T) {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)
  posts = append(posts, Post.NewPost{Id:1, Content: "Hello World!", Author: "Sherlock"})
  posts = append(posts, Post.NewPost{Id:1, Content: "Bonjour Madam!", Author: "M. Pierre"})
  posts = append(posts, Post.NewPost{Id:3, Content: "Hola Amigos", Author: "Prof. Sergio"})
  posts = append(posts, Post.NewPost{Id:4, Content: "Boring...", Author: "Sherlock"})

  for _,post := range posts {
    _ = Create.AddPost(post,PostById)
  }

  num, err := Read.ReadAllPosts(PostById)
  if err!=nil{
    if num!=len(posts)-1 {
      t.Errorf("Total number of records are not equal to number of posts added.")
    }
  }else {
    fmt.Println(err)
  }
}


func TestUpdatePost(t *testing.T) {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)
  posts = append(posts, Post.NewPost{Id:1, Content: "Hello World!", Author: "Sherlock"})
  posts = append(posts, Post.NewPost{Id:3, Content: "Hola Amigos", Author: "Prof. Sergio"})
  posts = append(posts, Post.NewPost{Id:4, Content: "Boring...", Author: "Sherlock"})

  for _,post := range posts {
    _ = Create.AddPost(post,PostById)
  }
   var err error
  s:="A plan needs atleast 5 months."
  ids:=[2]int{2,3}
  for _,id := range ids {
    err = Update.UpdatePostById(id,s,PostById)
    if err==nil {
      if(PostById[id].Content != s) {
        t.Errorf("Post was not updated.")
      }
    } else {
      fmt.Println(err)
    }
  }
}


func TestDeletePost(t *testing.T) {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)
  posts = append(posts, Post.NewPost{Id:1, Content: "Hello World!", Author: "Sherlock"})
  posts = append(posts, Post.NewPost{Id:3, Content: "Hola Amigos", Author: "Prof. Sergio"})
  posts = append(posts, Post.NewPost{Id:4, Content: "Boring...", Author: "Sherlock"})

  for _,post := range posts {
    _= Create.AddPost(post,PostById)
  }
  var err error

  ids:=[2]int{2,3}
  for _,id := range ids {
    err = Delete.DeletePostById(id,PostById)
    if err==nil {
      _,v := PostById[id]
      if v {
        t.Errorf("Post was not deleted.")
      }
    } else {
      fmt.Println(err)
    }
  }
}
