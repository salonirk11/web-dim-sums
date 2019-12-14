package main

import(
  "Create"
  "Delete"
  "fmt"
  "Post"
  "Read"
  "Update"
  )


func menu() int {
  var input int
  fmt.Println("Discussion Forum")
  fmt.Println("1. Create new post.")
  fmt.Println("2. Read a post.")
  fmt.Println("3. Update a post.")
  fmt.Println("4. Delete a post.")
  fmt.Println("5. Show all posts.")
  fmt.Println("6. Exit")
  fmt.Println("Choose option: ")
  fmt.Scanf("%d", &input)
  return input
}


func main() {
  var PostById map[int]*Post.NewPost
  PostById = make(map[int]*Post.NewPost)
  var s string

  for {
    choice := menu()
    switch choice{
    case 1:
      fmt.Println("\033[2J")
      var err error
      err = Create.CreatePost(PostById)
      if(err==nil){
        fmt.Println("Entry Added!")
      } else {
        fmt.Println(err)
      }
      fmt.Println("Press any character to go back.")
      fmt.Scanf("%s",&s)
      fmt.Println("\033[2J")

    case 2:
      fmt.Println("\033[2J")
      Read.ReadPost(PostById)
      fmt.Println("Press any character to go back.")
      fmt.Scanf("%s",&s)
      fmt.Println("\033[2J")

    case 3:
      fmt.Println("\033[2J")
      Update.UpdatePost(PostById)
      fmt.Println("Press any character to go back.")
      fmt.Scanf("%s",&s)
      fmt.Println("\033[2J")

    case 4:
      fmt.Println("\033[2J")
      Delete.DeletePost(PostById)
      fmt.Println("Press any character to go back.")
      fmt.Scanf("%s",&s)
      fmt.Println("\033[2J")

    case 5:
      fmt.Println("\033[2J")
      _, err := Read.ReadAllPosts(PostById)
      if err!=nil{
        fmt.Println(err)
      }
      fmt.Println("Press any character to go back.")
      fmt.Scanf("%s",&s)
      fmt.Println("\033[2J")

    case 6:
      return
    }
  }
}
