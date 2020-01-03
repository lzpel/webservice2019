package main

func main() {
	Handle("/", mainHandler)
	Listen(map[string]string{})
}
func mainHandler(w Response, r Request) {
	Writef(w, nil, "index.html")
}
