package main
import (
    "net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host + req.URL.Path 
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    http.Redirect(w, req, target,
            // see @andreiavrammsd comment: often 307 > 301
            http.StatusTemporaryRedirect)
}

func main() {
    // redirect every http request to https
    go http.ListenAndServe(":80", http.HandlerFunc(redirect))
}