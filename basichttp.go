package main

import (
       "fmt"
       "net/http"
       "encoding/json"
		"log"
      )


// Controller Methods
func userHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var u User   
    err := decoder.Decode(&u)
	if err != nil {
		log.Println("Error received with request %s", err)
		}
     fmt.Fprintf(w, "%s", u.Username)	
	log.Println(u.Username + u.Password)
}


func likeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var u Like   
    err := decoder.Decode(&u)
	if err != nil {
		log.Println("Error received with request %s", err)
		}
     fmt.Fprintf(w, "%s", u.Sender)	
	log.Println(u.Sender + u.Receiver)
}



func main() {
	// Routes
     http.HandleFunc( "/user", userHandler)
	 http.HandleFunc( "/sender", likeHandler)
     http.ListenAndServe(":3000", nil)
}

