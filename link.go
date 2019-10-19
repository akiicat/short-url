// Package p contains an HTTP Cloud Function.
package link

import (
        "context"
        "fmt"
        "log"
        "os"
        "net/http"

        "cloud.google.com/go/firestore"
      )


// GCLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GCLOUD_PROJECT")

func Link(w http.ResponseWriter, r *http.Request) {

  // https://github.com/golang/go/issues/15867#issuecomment-223748637
  cs := w.Header().Get("Set-Cookie")
  cs += "; SameSite=None; Secure"
  w.Header().Set("Set-Cookie", cs)

  url := r.URL.Path
  ctx := context.Background()

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

  linkTo, ok := m["link"].(string)
  if !ok {
    log.Fatalf("404 - Data Not Found %s -> %s", url, linkTo)
    http.Error(w, "404 - Data Not Found " + url + " -> " + linkTo, http.StatusNotFound)
    return
  }

  log.Printf("Redirect to: %s -> %s", url, linkTo)
  http.Redirect(w, r, linkTo, http.StatusFound)
}
