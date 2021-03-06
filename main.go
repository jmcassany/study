package main

import (
	"net/http"
	"fmt"
	"github.com/jfyne/live"
)

var cookieStore = live.NewCookieStore("lamevaaplicacio", []byte("elmeusecret"))


func main(){
  
  logoutHandler := NewLogoutHandler()
  loginHandler := NewLoginHandler()
  infoHandler := miInformacion()
  http.Handle("/info", infoHandler)
  http.Handle("/logout", logoutHandler)
  http.Handle("/login", loginHandler)
  http.Handle("/live.js", live.Javascript{})
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println(err)
  }

}
