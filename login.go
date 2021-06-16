package main

import (
  "github.com/jfyne/live"
  "net/http"
  "fmt"
  "context"
  "bytes"
  "io"
  //"log"
  "html/template"
  "sync"
  "net/url"
  //"time"
  //"go.mongodb.org/mongo-driver/mongo"
  //"go.mongodb.org/mongo-driver/mongo/options"
  //"os"
  //"github.com/gin-contrib/sessions"
  //"go.mongodb.org/mongo-driver/bson"
  //"github.com/gin-gonic/gin"
  //"github.com/gin-contrib/sessions/cookie"
)

var usuaris = map[string]string{}
var uLock sync.Mutex


var login = `
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
              Login en:<strong> JM stadium </strong>!
            </p>
            <figure class="image is-250x250">
              <img class=image src="https://www.imim.cat/media/comu/mobile/2.jpg">
            </figure>
          </div>
          <div class="column">
            {{ if .Errors }} 
            <article class="message is-danger">
              <div class="message-header">
                <p> {{ .Errors }} </p>  
              </div>
            </article>
            {{ end }}
            <form method="POST" class="box" live-submit="loginformulari">
              <div class="field" >
                <label>Usuari:</label>
                <div class="control">
                  <input type="text" name="usuari" class="input">
                </div>
              </div>
              <div class="field">
                <label>Contrasenya:</label>
                <div class="control">
                  <input type="text" name="contrasenya" class="input">
                </div>
              </div>
              <div class="field">
                <div class="control">
                  <button class="button is-success" type="submit">Login</button>                  
                </div>
              </div>              
            </form>
          </div>
        </div>
      </div>
    </section>
    <script src="/live.js"></script>
  </body>
</html>
`
type Login struct {
  Usuari string
  Contrasenya string
  Errors string
}

func NouLogin(s *live.Socket) *Login {
  m, ok := s.Assigns().(*Login)
  if !ok {
    m = &Login{ }
  }
  return m
}

func NewLoginHandler() *live.Handler {

  h, _ := live.NewHandler(cookieStore)

  h.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    m := NouLogin(s)
    return m, nil
  }

  h.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    var buf bytes.Buffer
    t, err := template.New("login").Parse(login)
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

  h.HandleEvent("loginformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NouLogin(s)
    
    m.Usuari = p.String("usuari")
    m.Contrasenya =  p.String("contrasenya")
    
    if m.Usuari == "jm" && m.Contrasenya == "pimpam"{
      fmt.Println("OK")
      
      u, _ := url.Parse("/info")
      s.Redirect(u)
      
      uLock.Lock()
      usuaris[s.Session.ID] = m.Usuari
      uLock.Unlock()

    }else{
      fmt.Println("ERROR")
      m.Errors = "Error en l'usuari o la contrasenya"
    }
    
    fmt.Println(m.Usuari)
    fmt.Println(m.Contrasenya)
    return m, nil
  })

  return h
}
