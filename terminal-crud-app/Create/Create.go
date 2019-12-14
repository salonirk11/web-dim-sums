package Create

import(
  "bufio"
  "fmt"
  "os"
  "Post"
  "reflect"
)

func AddPost(post Post.NewPost, PostById map[int]*Post.NewPost) (error) {
  _,v := PostById[post.Id]
  if (reflect.TypeOf(post.Id).Kind()!=reflect.TypeOf(2).Kind()) || (v) {
    return fmt.Errorf("Invalid ID")
  }
  PostById[post.Id]=&post
  return nil
}


func CreatePost(PostById map[int]*Post.NewPost) (error){
  var post Post.NewPost
  input := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter ID: ")
  fmt.Scanf("%d",&post.Id)

  fmt.Println("\nEnter Author name: ")
  input.Scan()
  post.Author = input.Text()

  fmt.Println("\nEnter Content:")
  input.Scan()
  post.Content = input.Text()


  return AddPost(post, PostById)
}
