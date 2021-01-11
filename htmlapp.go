package main
import(
	"fmt"
	"net/http"
	"html/template"
)
type Values struct{
	Name string
}

func Dashboard(w http.ResponseWriter,r *http.Request){
	val := Values{"Go"}
	tmpl,err := template.ParseFiles("index.php")
	if err!=nil{
		panic(err)
	}
	// fmt.Fprint(w,"<h1>Hello from Go</h1>")
	tmpl.Execute(w, val)
}
func test(w http.ResponseWriter,r *http.Request){
	val := Values{"Padmesh"}
	tmpl,err := template.ParseFiles("index.php")
	if err!=nil{
		panic(err)
	}
	fmt.Fprint(w,"<h1>Hello from Test</h1>")
	tmpl.Execute(w, val)
}
func test1(w http.ResponseWriter,r *http.Request){
	val := Values{"Manish"}
	tmpl,err := template.ParseFiles("index.php")
	if err!=nil{
		panic(err)
	}
	fmt.Fprint(w,"<h1>Hello from Test</h1>")
	tmpl.Execute(w, val)
}
func main(){
	http.HandleFunc("/",Dashboard)
	http.HandleFunc("/padmesh",test)
	http.HandleFunc("/manish",test1)
	http.ListenAndServe("localhost:4000",nil)
}