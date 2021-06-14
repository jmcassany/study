package main

import (
  "github.com/jfyne/live"
  "net/http"
  "fmt"
  "context"
  "bytes"
  "io"
  "html/template"
)

var page = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>I WILL BE A LEGEND!</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.2/css/bulma.min.css">
  </head>
  <body>
    <section class="section">
      <div class="container-fluid">
        <div class="columns">
          <div class="column">
            <p class="subtitle">
              Quiero registrarme en:<strong>PadelJM</strong>!
            </p>
            <span class="tag is-success">Astonishing</span>
          </div>
          <div class="column">
            <form method="POST" live-change="elmeuformulari">
              <div class="field">
                <label>Nom:</label>
                <div class="control">
                  <input type="text" name="nom" class="input">
                </div>
              </div>
              <div class="field">
                <label>Cognoms:</label>
                <div class="control">
                  <input type="text" name="cognoms" class="input">
                </div>
              </div>
              <div class="field">
                <label>Posicio::</label>
                <div class="control">
                  <input type="text" name="posicio" class="input">
                </div>
              </div>
              <div class="field">
                <label>Company:</label>
                <div class="control">
                  <input type="text" name="company" class="input">
                </div>
              </div>
              <div class="field">
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
                  <input type="submit" class="button">
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

var login = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>I WILL BE A LEGEND!</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.2/css/bulma.min.css">
  </head>
  <body>
    <section class="section">
      <div class="container-fluid">
        <div class="columns">
          <div class="column">
            <p class="subtitle">
              Login en:<strong>PadelJM</strong>!
            </p>
            <span class="tag is-success">Astonishing</span>
          </div>
          <div class="column">
            <form method="POST" live-change="loginformulari">
      
              <div class="field">
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
              
            </form>
          </div>
        </div>
      </div>
    </section>
    <section class="section" id="section-{{.Contrasenya}}">
      <div class="container">
        <div class="notification is-danger">
          <a class="button is-succes">aceptar</a>
        </div>
      </div>
    </section>
    <script src="/live.js"></script>
  </body>
</html>
`

type MevaAplicacio struct {
  Nom string
  Cognoms string
  Posicio string
  Company string
  Usuari string
  Contrasenya string

}

func NovaAplicacio(s *live.Socket) *MevaAplicacio {
  m, ok := s.Assigns().(*MevaAplicacio)
  if !ok {
    m = &MevaAplicacio{ }
  }
  return m
}

func main() {
  
  h, _ := live.NewHandler(live.NewCookieStore("lamevaaplicacio", []byte("elmeusecret")))
  j, _ := live.NewHandler(live.NewCookieStore("lamevaaplicacio", []byte("elmeusecret")))

  h.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    m := NovaAplicacio(s)
    return m, nil
  }

  j.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    m := NovaAplicacio(s)
    return m, nil
  }

  h.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    var buf bytes.Buffer
    t, err := template.New("blablabla").Parse(page)
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

  j.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    var buf bytes.Buffer
    t, err := template.New("blablabla2").Parse(login)
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

  h.HandleEvent("elmeuformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NovaAplicacio(s)
    nom := p.String("nom")
    cognoms := p.String("cognoms")
    posicio := p.String("posicio")
    company := p.String("company")
    usuari := p.String("usuari")
    contrasenya := p.String("contrasenya")
    m.Nom = nom
    m.Cognoms = cognoms
    m.Posicio = posicio
    m.Company = company
    m.Usuari = usuari
    m.Contrasenya = contrasenya
    

    return m, nil
  })

  j.HandleEvent("loginformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NovaAplicacio(s)
    usuari := p.String("usuari")
    contrasenya := p.String("contrasenya")
    m.Usuari = usuari
    m.Contrasenya = contrasenya
    
    if m.Contrasenya != "12345" {
      fmt.Println("ERROR")
      /*c.Redirect(302, "/signin")*/
    }else{
      fmt.Println("OK")
    }
    fmt.Println(m.Contrasenya)
    return m, nil
  
  })

/*  j.HandleEvent("verificar", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NovaAplicacio(s)
    // accio := p.String("laverificacio")
    usuari := p.String("usuari")
    contrasenya := p.String("contrasenya")
    m.Usuari = usuari
    m.Contrasenya = contrasenya

    if m.Contrasenya != "12345" {
      fmt.Println("ERROR")
    }else{
      fmt.Println("OK")
    }
    fmt.Println(m.Contrasenya)
    return m, nil
  
  })*/



  http.Handle("/register", h)
  http.Handle("/login", j)
  http.Handle("/live.js", live.Javascript{})
  err := http.ListenAndServe(":8081", nil)
  if err != nil {
    fmt.Println(err)
  }

}
