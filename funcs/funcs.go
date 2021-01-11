package funcs
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"html/template"
	"htmlapp/dbase"
	"encoding/json"
	"database/sql"
	_ "mysql"
)
type User struct{
	Name string
	Jsonstring string
}
type NewStud struct{
	Name string `json:"name"`
	Gender string `json:"gender"`
	Location string `json:"location"`
	Email string `json:"email"`
}
var currUser = User{}
var t *template.Template

//Login form - root
func Index(w http.ResponseWriter, r *http.Request){
	t,_ = template.ParseGlob("*.html")
	t.ExecuteTemplate(w,"index.html",nil)
}

//Login handler function
func Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	var b bool
	uname := r.FormValue("username")
	pass := r.FormValue("password")
	b, currUser.Name = dbase.DbGet(uname,pass)
	if b {
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
	}else{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	
}

//Homepage handler function
func Dashboard(w http.ResponseWriter, r *http.Request){
	t.ExecuteTemplate(w,"home.html",currUser)
}

//Logout handler function
func Logout(w http.ResponseWriter, r *http.Request){
	http.Redirect(w,r,"/",http.StatusSeeOther)
}

//View handler function
func ViewStudent(w http.ResponseWriter, r *http.Request){
	b,strStud := dbase.DbGetStud()
	// bytes,_ := json.Marshal(strStud)
	// currUser.Jsonstring = `[`
	// for i:=0;i<len(strStud.Name);i++{
		currUser.Jsonstring += fmt.Sprintf(`[{"text":"%s","nodes":
		[{"text":"Gender: %s"},
		 {"text":"Location: %s"},
		 {"text":"Email: %s"}
		]}]`,strStud.Name[0],strStud.Gender[0],strStud.Location[0],strStud.Email_id[0])
	// }
	fmt.Println(currUser.Jsonstring)
	fmt.Println(len(strStud.Name))
	Dashboard(w,r)
	if b{
		// fmt.Fprintf(w,`<table><tr><th>Name</th><th>Roll No.</th><th>Gender</th><th>Location</th><th>Email-id</th></tr>`)
		for i:=0;i<len(strStud.Name);i++{
			// fmt.Fprintf(w,`<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td>`,strStud.Name[i],strStud.Roll_no[i],strStud.Gender[i],strStud.Location[i],strStud.Email_id[i])
		}
		// fmt.Fprintf(w,`</table>`)
	}else{
		fmt.Fprintf(w,"<h3>No students found<h3>")
	}
}

//Add student form
func AddStudent(w http.ResponseWriter, r *http.Request){
	t.ExecuteTemplate(w,"addStud.html",nil)
}

//Adding student handler function
func Addnew(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(r.Body)
	var n NewStud
	err = json.Unmarshal(body,&n)		//Converting JSON to struct
	if err!= nil{
		fmt.Println(err)
	}
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{fmt.Println(err)}
	defer db.Close()
	insert,err := db.Prepare("INSERT INTO Students(name,gender,location,email_id) VALUES (?,?,?,?)")
	if err!=nil{fmt.Println(err)}
	defer insert.Close()
	_,err1 := insert.Exec(n.Name,n.Gender,n.Location,n.Email)
	if err1!=nil{fmt.Println(err)}else{
		// fmt.Println(w,r)
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
	}
}