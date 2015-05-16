package static

import (
    	"net/http"
	//"html/template"
	//"flag"
	//"fmt"
	//"log"
	//"code.google.com/p/google-api-go-client/googleapi/transport"
	//"code.google.com/p/google-api-go-client/youtube/v3"
	
)

func init() {
    http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "files/"+r.URL.Path)
}

/*func searchstats(w http.ResponseWriter, r *http.Request) {

	 key := r.*/

