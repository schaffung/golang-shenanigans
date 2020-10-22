package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/* Let's analyze what we have here...
	 * 1. A route /hello which when reached, will invoke our function
	 * which takes in...w and r
     * 2. w being as the type states a responsewrite so basically what
     * we'll send as a response.
     * 3. r being the request so basically what was requsted..
	 */
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("What was requested...oh yeah r -> ", *r)
		fmt.Fprintf(w, "Hello!")
	})

	// We are listening!!!! Shoot your requests...
	fmt.Printf("Starting webserver at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
