package dbase
import(
	"database/sql"
	_ "mysql"
	"fmt"
	// "strconv"
)
type User struct{
	name string
	pass string
}
type Students struct{
	Text string	`json:"text"`
	Name []string `json:"name"`
	Roll_no []string `json:"roll_no"`
	Gender []string `json:"gender"`
	Location []string `json:"location"`
	Email_id []string `json:"email_id"`
}
func DbGet(username,password string) (bool,string){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT name FROM users WHERE username=? AND password=?",username,password)
	if err!=nil{throwError(err)}
	defer results.Close()
	var user User
	for results.Next(){
		results.Scan(&user.name)
	}
	if user.name != ""{
		return true,user.name
	}
	return false,user.name
}

func DbGetStud() (bool,Students){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	results,err := db.Query("SELECT name,roll_no,gender,location,email_id FROM students")
	if err!=nil{throwError(err)}
	defer results.Close()
	var student Students
	// var strStud string
	i := 0
	for results.Next() {
		var name string
		var roll_no string
		var gender string
		var location string
		var email_id string
		node := "Node 1"
		results.Scan(&name,&roll_no,&gender,&location,&email_id)
		student.Text = node
		student.Name = append(student.Name,name)
		student.Roll_no = append(student.Roll_no,roll_no)
		student.Gender = append(student.Gender,gender)
		student.Location = append(student.Location,location)
		student.Email_id = append(student.Email_id,email_id)
		i++
	}
	if i > 0{
		return true,student
	}
	return false,student
}
func AddStud(Name,Gender,Location,Email string) bool{
	fmt.Println(Name,Gender,Location,Email)
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{throwError(err)}
	defer db.Close()
	insert,err := db.Prepare("INSERT INTO Students(name,gender,location,email_id) VALUES (?,?,?,?)")
	if err!=nil{throwError(err)}
	defer insert.Close()
	_,err1 := insert.Exec(Name,Gender,Location,Email)
	if err1!=nil{throwError(err)}else{
		fmt.Println(true)
		return true
	}
	return false
}
func throwError(err error){
	panic(err) //panic(..interface{})
}