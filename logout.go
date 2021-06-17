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
  //"net/url"
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

//var usuaris = map[string]string{}
//var uLock sync.Mutex


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
          <div class="column">
            <form method="POST" class="box" live-click="logoutformulari">
              <figure class="image is-128x128">
                <img class=is-rounded src="https://media-exp1.licdn.com/dms/image/C5603AQEGWFeheGOWyA/profile-displayphoto-shrink_200_200/0/1576758118970?e=1629331200&v=beta&t=bbgd5RKV15grOdy2KbGMG36ilYwiRi1DpEdcI_Wq0PU">
              </figure>                    
              <article class="message is-danger">
                <div class="message-header">
                  <p> Has tancat la sessió. Fins aviat =) </p>  
                </div>
              </article>
            </form>
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
    
    //var usuaris = map[string]string{}
    //_, ok := usuaris[s.Session.ID]
    
    delete(usuaris, s.Session.ID)
  
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

  h.HandleEvent("logoutformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NouLogout(s)
    
    
    return m, nil
  })

  return h
}
