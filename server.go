// Go application with hot/dymnamic reloading

// Is the first thing that runs when we start the application
package main 

// Import packages fmt - printing things out , net/http -create a web server 
import (
	"fmt"
	"net/http"
   )
  
func index(w http.ResponseWriter, r *http.Request) {  // Response writer and a pointer to a request
	fmt.Fprintf(w, "home")

}

func about(w http.ResponseWriter, r *http.Request) {
	   fmt.Fprintf(w, "about")
   
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact")

}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	
	// Start server and listen on port 8000
	http.ListenAndServe(":8000", nil) //nill is the default serve mux
}
   