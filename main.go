package main
import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
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
func Index(w http.ResponseWriter,r *http.Request){
	db:=dbConn()
	selDB,err:=db.Query("select * from employee order by id desc;")
	checkErr(err)
	for selDB.Next(){
		var id int
		var name,city string
		err=selDB.Scan(&id,&name,&city)
		checkErr(err)
		fmt.Fprintln(w,id)
		fmt.Fprintln(w,name)
		fmt.Fprintln(w,city)
	}
	defer db.Close()
}

