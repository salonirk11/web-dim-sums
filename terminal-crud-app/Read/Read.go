package Read

import(
  "fmt"
  "Post"
)


func ReadPostById(id int, PostById map[int]*Post.NewPost) (*Post.NewPost,error) {
  _,v := PostById[id]
  if v {
      return PostById[id], nil
  }
  return PostById[id], fmt.Errorf("Invalid ID")
}


func ReadPost(PostById map[int]*Post.NewPost) {
  var id int
  fmt.Println("Enter ID: ")
  fmt.Scanf("%d", &id)

  post,err := ReadPostById(id, PostById)
  if(err!=nil){
    fmt.Println(err)
  }else {
    fmt.Printf("ID: %d\n", post.Id)
    fmt.Printf("Author: %s\n", post.Author)
    fmt.Printf("%s\n", post.Content)
  }

}


func ReadAllPosts(PostById map[int]*Post.NewPost) (num int, err error){
  if len(PostById)==0{
    return 0,fmt.Errorf("No elements to show.")
  }
  count:=0
  for _,post := range PostById {
    fmt.Printf("ID: %d\n", post.Id)
    fmt.Printf("Author: %s\n", post.Author)
    fmt.Printf("%s\n", post.Content)
    count++
  }
  return count, nil
}
