package Update

import(
  "bufio"
  "fmt"
  "os"
  "Post"
)

func UpdatePostById(id int, s string, PostById map[int]*Post.NewPost) (error) {
  _,v := PostById[id]
  if !v {
    return fmt.Errorf("ID does not exist in record, can't update")
  }
  PostById[id].Content = s
  return nil
}


func UpdatePost(PostById map[int]*Post.NewPost)  {
  var id int
  var err error
  fmt.Println("Enter ID: ")
  fmt.Scanf("%d", &id)
  
  fmt.Println("Enter new content: \n")
  input := bufio.NewScanner(os.Stdin)
  input.Scan()
  err = UpdatePostById(id, input.Text(), PostById)

  if err!=nil{
    fmt.Println(err)
  }
}
