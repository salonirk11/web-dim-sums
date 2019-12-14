package Delete

import(
  "fmt"
  "Post"
)

func DeletePostById(id int, PostById map[int]*Post.NewPost) error {
  _,v := PostById[id]
  if v==false {
    return fmt.Errorf("ID does not exist in record, can't delete.")
  }
  delete(PostById,id)
  return nil
}


func DeletePost(PostById map[int]*Post.NewPost)  {
  var id int
  fmt.Println("Delete post")
  fmt.Println("Enter ID: ")
  fmt.Scanf("%d", &id)

  err := DeletePostById(id, PostById)
  if err!=nil {
    fmt.Println(err)
  }
}
