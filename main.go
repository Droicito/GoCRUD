package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Emp struct{
	Name string `json:"name"`
}
func main(){

	fmt.Println("GO - MSQL Server Connection!")
	db,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/goblog")
	if err !=nil{
		panic(err.Error())
	}
	fmt.Println("Connected Succesfull!")

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
func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}