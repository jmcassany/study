package main

import (
  "github.com/jfyne/live"
  "net/http"
  //"fmt"
  "context"
  "bytes"
  "io"
  //"log"
  "html/template"
  //"sync"
  "net/url"
  //"time"
  //"go.mongodb.org/mongo-driver/mongo"
  //"go.mongodb.org/mongo-driver/mongo/options"
  //"os"
  //"github.com/gin-contrib/sessions"
  //"go.mongodb.org/mongo-driver/bson"
  //"github.com/gin-gonic/gin"
  //"github.com/gin-contrib/sessions/cookie"
  //"github.com/kataras/iris"
)

var logout = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta nom="viewport" content="width=device-width, initial-scale=1">
    <title>YOU WILL BE A LEGEND!</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.2/css/bulma.min.css">
  </head>
  <body>
    <section class="section">
      <div class="container-fluid">
        <div class="columns">
          <div class="column">
            <p class="subtitle">
              Informació usuari:<strong> Pepito de los palotes </strong>!
            </p>
            <figure class="image is-250x250">
              <img class=image src="https://www.imim.cat/media/comu/mobile/2.jpg">
            </figure>
          </div>
        </div>
      </div>
    </section>
    <script src="/live.js"></script>
  </body>
</html>
`

type Logout struct {
  Sortir string
}

func NouLogout(s *live.Socket) *Logout {
  m, ok := s.Assigns().(*Logout)
  if !ok {
    m = &Logout{ }
  }
  return m
}

func NewLogoutHandler() *live.Handler {

  h, _ := live.NewHandler(cookieStore)

  h.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    
    /*m := NouLogout(s)
    
    _, ok := usuaris[s.Session.ID]
    if ok {
      delete(usuaris, s.Session.ID)
    } else {
      u, _ := url.Parse("/login")
      s.Redirect(u)
      return nil, nil
    }
    return m, nil*/

    ul, _ := url.Parse("/login")
    s.Redirect(ul)

    m := NouLogout(s)
    return m, nil
  }

  h.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    
    var buf bytes.Buffer
    t, err := template.New("logout").Parse(logout)
    if err != nil {
      buf.WriteString(err.Error())
      return &buf, nil
    }

    err = t.Execute(&buf, data)
    if err != nil {
      buf.WriteString(err.Error())
      return &buf, nil
    }

    return &buf, nil
  }
  return h
  
}
