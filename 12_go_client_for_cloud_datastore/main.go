package main
func main() {
	handle("/",
		func(w response, r request) {
			var array []base
			getall(query("base").Filter("Area=","rice").Order("-TimeBorn").Limit(10),&array)//limitは取り出す数
			writetemplate(w, "index.html", &array)
		})
	handle("/form",
		func(w response, r request) {
			put(&base{
				Area: "rice",
				Name: r.URL.Query().Get("name"),
				Text: r.URL.Query().Get("text"),
			})
			redirect(w,r,"/")
		})
	serve("crediantial.json")
}