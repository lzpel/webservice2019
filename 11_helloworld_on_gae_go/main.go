package main

func main() {
	handle("/", mainHandler)
	serve("")
}
func mainHandler(w response, r request) {
	writetemplate(w, "index.html", nil)
}