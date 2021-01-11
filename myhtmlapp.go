package main
import(
	"net/http"
	"htmlapp/funcs"
)

func main(){
	http.HandleFunc("/",funcs.Index)
	// http.HandleFunc("/login",funcs.Login)
	// http.HandleFunc("/dashboard",funcs.Dashboard)
	// http.HandleFunc("/logout",funcs.Logout)
	// http.HandleFunc("/ViewStudent",funcs.ViewStudent)
	http.HandleFunc("/AddStudent",funcs.AddStudent)
	// http.HandleFunc("/addnew",funcs.Addnew)
	http.Handle("/css/",http.StripPrefix("/css/",http.FileServer(http.Dir("css/"))))
	http.Handle("/js/",http.StripPrefix("/js/",http.FileServer(http.Dir("js/"))))
	http.ListenAndServe(":4000",nil)
}