package main

import (
	"github.com/urfave/negroni"
)

func main() {
	//invoke new router function
	r := NewRouter()

	//use negroni middleware for logging http request and response
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8088")
	//log.Fatal(http.ListenAndServe(":8088", r))
}
