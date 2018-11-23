package main
import (
	"fmt"
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
	dbname:="goblog"
	db,err:=sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbname)
	checkErr(err)
	return db
}
func main(){
	db:=dbConn()
	fmt.Println("Sucesful WebApp - Db Mysql Connection!")
	rows,err:=db.Query("select * from employee;")
	checkErr(err)
	
	for rows.Next(){
		var eid int
		var ename string
		var ecity string
		err=rows.Scan(&eid,&ename,&ecity)
		checkErr(err)
		fmt.Println(eid)
		fmt.Println(ename)
		fmt.Println(ecity)
	}
	defer db.Close()
}

