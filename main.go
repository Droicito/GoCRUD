package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main(){

	fmt.Println("GO - MSQL Server Connection!")
	db,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/goblog")
	if err !=nil{
		panic(err.Error())
	}
	fmt.Println("Connected Succesfull!")
	defer db.Close()
}