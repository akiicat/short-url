package api

import (
        "context"
        "fmt"
        "log"
        "os"
        "net/http"

        "cloud.google.com/go/firestore"
        "github.com/julienschmidt/httprouter"
      )

var mux = newRouter()

func newRouter() *httprouter.Router {
  mux := httprouter.New()
  mux.GET("/:id", LinkController)
  mux.GET("/api/:id", LinkController)
	return mux
}

//F represents cloud function entry point
func Link(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

// GCLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("PROJECT_ID")

func LinkController(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

  // https://github.com/golang/go/issues/15867#issuecomment-223748637
  cs := w.Header().Get("Set-Cookie")
  cs += "; SameSite=None; Secure"
  w.Header().Set("Set-Cookie", cs)

  url := r.URL.Path
  ctx := context.Background()

  log.Println("ingress url:", url, "id:", params.ByName("id"))

  client, err := firestore.NewClient(ctx, projectID)
  if err != nil {
    log.Fatalln(err)
    http.Error(w, "500 - Internal Server Error " + url, http.StatusInternalServerError)
    return
  }

  dsnap, err := client.Collection("links").Doc(url[1:]).Get(ctx)
  if err != nil {
    log.Fatalln(err)
    http.Error(w, "404 - Not Found url", http.StatusNotFound)
    fmt.Fprintf(w, "Hello, World!")
    return
  }

  m := dsnap.Data()
  linkTo := m["link"].(string)

  log.Printf("Redirect to: %s -> %s", url, linkTo)
  http.Redirect(w, r, linkTo, http.StatusFound)
}
