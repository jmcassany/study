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
              Login en:<strong> JM stadium </strong>!
            </p>
          </div>
        </div>
      </div>
    </section>
    <script src="/live.js"></script>
  </body>
</html>
`

func informacion(){

  j, _ := live.NewHandler(live.NewCookieStore("lamevaaplicacio", []byte("elmeusecret")))

  j.Mount = func(c context.Context, r *http.Request, s *live.Socket) (interface{}, error) {
    m := NovaAplicacio(s)
    return m, nil
  }

  j.Render = func(c context.Context, data interface{}) (io.Reader, error) {
    var buf bytes.Buffer
    t, err := template.New("blablabla22").Parse(info)
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

  j.HandleEvent("infoformulari", func(c context.Context, s *live.Socket, p live.Params) (interface{}, error) {
    m := NovaAplicacio(s)
    m.Usuari = p.String("usuari")
    
    return m, nil
  })
  

  
  http.Handle("/info", j)
  http.Handle("/live.js", live.Javascript{})
  err := http.ListenAndServe(":8081", nil)
  if err != nil {
    fmt.Println(err)
  }
  
}