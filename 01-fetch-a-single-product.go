//fetch a single product by its ID from https://dummyjson.com's product endpoint 
package main
import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
  resp, err := http.Get("https://dummyjson.com/products/1")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }

  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  var product map[string]interface{}
  json.Unmarshal(body, &product)

  // Pretty-print the JSON 
  prettyJSON, err := json.MarshalIndent(product, "","  ")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }
  
  fmt.Println(string(prettyJSON))
}
