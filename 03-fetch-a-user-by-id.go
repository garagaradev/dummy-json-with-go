//retrieve user information by their ID
package main
import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
  resp, err := http.Get("https://dummyjson.com/users/1")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }

  defer resp.Body.Close()

  body,_ := ioutil.ReadAll(resp.Body)
  var user map[string]interface{}
  json.Unmarshal(body, &user)
  prettyJSON, err := json.MarshalIndent(user, "","  ")
  if err != nil {
    fmt.Println("Error:",err)
    return
  }

  fmt.Println(string(prettyJSON))
}
