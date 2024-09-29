//retrieve multiple users with pagination
package main
import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

func main(){
  resp, err := http.Get("https://dummyjson.com/users")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }
  defer resp.Body.Close()

  body,_ := ioutil.ReadAll(resp.Body)
  var users map[string] interface{}
  json.Unmarshal(body, &users)
  prettyJSON, err := json.MarshalIndent(users,"","  ")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }
  fmt.Println(string(prettyJSON))
}
