package main

import (
  "fmt"
  "github.com/gojektech/heimdall"
  "github.com/gojektech/heimdall/httpclient"
  "io/ioutil"
  "time"
  "net/http"
  // "Post"
)

func main() {
  // First set a backoff mechanism. Constant backoff increases the backoff at a constant rate
  backoffInterval := 2 * time.Millisecond
  maximumJitterInterval := 5 * time.Millisecond
  backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)
  retrier := heimdall.NewRetrier(backoff)

  timeout := 1000 * time.Millisecond
  client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout), httpclient.WithRetrier(retrier),httpclient.WithRetryCount(4),)
  // var post Post.NewPost
  // post.Id = 2
  // post.Title = "Test"
  // post.Author = "Author"
  // post.Content = "Content"
  // req, err := http.NewRequest("GET", "http://localhost:8080/form", bytes.NewBufferString(post))
  req, err := http.NewRequest("GET", "http://localhost:8080/", nil)

  // Use the clients GET method to create and execute the request
  res, err := client.Do(req)
  if err != nil{
  	panic(err)
  }

  // Heimdall returns the standard *http.Response object
  body, err := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
}
