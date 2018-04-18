// Sample app

package main

import (
   t "time"
   "fmt"
   "net/http"
)


func main(){
   fmt.Println(t.Second)

   now := t.Now();
   fmt.Println(now)
   http.Get("news.sina.com.cn")
   fmt.Println(t.Hour + t.Since(now))
}
