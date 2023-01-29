package main

import (
	"fmt"
	"log"
	"net/http"
)

//Buidling a web sever with Go Lang

func main() {
	//we are telling golang to look for static folder and it automatically looks for index.html file
	fileServer := http.FileServer(http.Dir("./static"))

	//Handling the routes
	//when use hits the root route send the file Server to it that is index.html
	http.Handle("/", fileServer)

	//It means when the route is /hello then call the helloHandlerfunction
	http.HandleFunc("/hello", helloHandler)

	//Handling the /form route
	http.HandleFunc("/form", formHandler)

	//starting a Sever on 8080 ListenAndServe will create the server and can return error or nil
	fmt.Printf("Starting Sever at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	/*
		err := http.ListenAndServe(":8080")
		if err != nil {
		log.Fatal(err)
		}
	*/

}

// With every route we have two things response and a request (request(r) - user) and (response(w) - server)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successfull")
	//fectching names from the form
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

/*
Http package functions used
 http.FileServer
 http.Handle
 http.HandleFunc
 http.ListenAndServe
 http.Error
*/
