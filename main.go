package liquidmarket

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.Handle("/", corsWrapper(http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	c := appengine.NewContext(r)
	
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	
	fmt.Fprintf(w, "Hello, %v!", u)
	
}
