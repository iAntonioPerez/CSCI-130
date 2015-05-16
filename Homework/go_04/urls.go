//Antonio Perez
package urls

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/info", info)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Look at this amazing website!!")
	fmt.Fprintln(w, "There is much more if you go to /info")

}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I lied.")
	fmt.Fprintln(w, "There isn't much here at all.")

}
