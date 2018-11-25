package main
import (
//	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"text/template"
)
type Employee struct{
	Id int
	Name string
	City string
}

func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}
func dbConn()(db *sql.DB){
	dbDriver:="mysql"
	dbUser:="root"
	dbPass:="root"
	dbName:="goblog"
	db,err:=sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	checkErr(err)
	return db
}

func main(){
	http.HandleFunc("/",Index)
	http.ListenAndServe(":8080",nil)
}
var tmpl=template.Must(template.ParseGlob("form/*"))
func Index(w http.ResponseWriter,r *http.Request){
	db:=dbConn()
	selDB,err:=db.Query("select * from employee order by id desc;")
	checkErr(err)
	emp:=Employee{}
	res:=[]Employee{}
	for selDB.Next(){
		var id int
		var name,city string
		err=selDB.Scan(&id,&name,&city)
		checkErr(err)
		emp.Id=id
		emp.Name=name
		emp.City=city
		res=append(res,emp)
	}
	tmpl.ExecuteTemplate(w,"index",res)
	defer db.Close()
}

