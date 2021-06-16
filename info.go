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
)

var info = `
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
            <form method="POST" class="box" live-submit="logoutformulari">
              <figure class="image is-128x128">
                <img class=is-rounded src="https://media-exp1.licdn.com/dms/image/C5603AQEGWFeheGOWyA/profile-displayphoto-shrink_200_200/0/1576758118970?e=1629331200&v=beta&t=bbgd5RKV15grOdy2KbGMG36ilYwiRi1DpEdcI_Wq0PU">
              </figure>
              <div class="field" >
                <label>Nom:</label>
                <div class="control">
                  <input type="text" name="nom" class="input">
                </div>
              </div>
              <div class="field" >
                <label>Cognoms:</label>
                <div class="control">
                  <input type="text" name="cognoms" class="input">
                </div>
              </div>
              <div class="field" >
                <label>Email:</label>
                <div class="control">
                  <input type="text" name="nom" class="input">
                </div>
              </div>      
              <div class="field" >
                <label>Grup:</label>
                <div class="control">
                  <input type="text" name="nom" class="input">
                </div>  
              </div>
              <div class="field is-grouped">
                <p class="control">
                  <button class="button is-success" type="submit">
                    Continuar treballant
                  </button>
                </p>
                <p class="control">
                  <button class="button is-link" type="submit">
                    A parir panteras
                  </button>
                </p>
                <p class="control">
                  <button class="button is-danger" live-click="logoutformulari">
                  LogOut
                  </button>
                </p>
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
type User struct {
  Nom string
  Cognoms string
  Email string
  Grup string
  
}

func NouUsuari(s *live.Socket) *User {
  m, ok := s.Assigns().(*User)
  if !ok {
    return &User{}
  }
  return m
}


func miInformacion() *live.Handler {

  h, _ := live.NewHandler(cookieStore)

  h.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    
    _, ok := usuaris[s.Session.ID]
    if !ok {
      ul, _ :=url.Parse("/login")
      s.Redirect(ul)
      return nil, nil
    }

    m := NouUsuari(s)
    return m, nil
  }

  h.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    var buf bytes.Buffer
    t, err := template.New("info").Parse(info)
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

  /*h.HandleEvent("infoformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NouUsuari(s)
    m.Nom = p.String("nom")
    m.Cognoms = p.String("cognoms")
    m.Email = p.String("email")
    m.Grup = p.String("grup")
  
    return m, nil
  })
  
  return h*/

  h.HandleEvent("logoutformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NouLogout(s)
    
    u,_ := url.Parse("/logout")
    s.Redirect(u)

    uLock.Lock()
    usuaris[s.Session.ID] = m.Sortir
    uLock.Unlock()
    
    return m, nil
  })

  return h


}

