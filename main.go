package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// home handler function
func home(w http.ResponseWriter, r *http.Request) {
	//if not with html
	//w.Write([]byte("My name is Tamika Chen. I like to travel and I dislike spiders. My email is 2019120211@ub.edu.bz\n"))

	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Home</title>
				<style>
					p {
						color: red;
						font-family: verdana;
					}
				</style>
			</head>
			<body>
			<p>My name is Tamika Chen. I like to travel and I dislike spiders. My email is 2019120211@ub.edu.bz</p>			
			</body>
		</html>
	`
	fmt.Fprint(w, html)
}

// greeting handler function
func greeting(w http.ResponseWriter, r *http.Request) {
	currtime := time.Now()
	//if not with html
	//w.Write([]byte("I hope you're having a good day! The time right now is " + currtime.Format("15:04:05")))

	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Greeting</title>
				<style>
					p {
						color: red;
						font-family: verdana;
					}
				</style>
			</head>
			<body>
			<p>I hope you're having a good day! The time right now is {{time}}</p>			
			</body>
		</html>
	`
	html = strings.ReplaceAll(html, "{{time}}", currtime.Format("15:04:05"))
	fmt.Fprint(w, html)
}

// random handler function
func random(w http.ResponseWriter, r *http.Request) {
	quotes := make([]string, 0)
	quotes = append(quotes,
		"Adulting is soup and I am a fork.",
		"Roses are red, im going to bed.",
		"Dont be the bigger person today, be the problem.",
		"My socks aren't wet physically but emotionally I feel like my socks are wet.",
		"God knew I'd be too powerful if I knew how to code")

	//if not with html
	//w.Write([]byte(quotes[rand.Intn(len(quotes))]))

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))

	html := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Quote</title>
				<style>
					p {
						color: red;
						font-family: verdana;
					}
				</style>
			</head>
			<body>
			<p>`+ quotes[randomIndex] +`</p>			
			</body>
		</html>
	`
	fmt.Fprintf(w, html)
}

func main() {
	//multiplexer(map)
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux) //create a server
	log.Fatal(err)
}