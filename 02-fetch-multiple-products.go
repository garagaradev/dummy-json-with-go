//fetch a list of products with pagination support
package main
import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
  resp, err := http.Get("https://dummyjson.com/products")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  var products map[string]interface{}
  json.Unmarshal(body, &products)
  prettyJSON,err := json.MarshalIndent(products,""," ")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }

  fmt.Println(string(prettyJSON))
}


