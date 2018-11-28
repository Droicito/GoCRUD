package main
import(
	"fmt"
//	"bytes"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)
type Person struct{
	Id int
	Firstname string
	Lastname string
}

func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}
func dbConn()(db *sql.DB){
	db,err:=sql.Open("mysql","root:root@/goblog")
	checkErr(err)
	return db
}

func main(){
	router:=gin.Default()
	router.GET("/person/:id",getPerson)
	
	router.GET("/people",getPeople)

	fmt.Print(router.Run(":3000"))
}

func getPerson(c *gin.Context){
	var(
		person Person
		result gin.H
	)
	db:=dbConn()
	id:=c.Param("id")
	row:=db.QueryRow("select * from person where id=?;",id)
	err:=row.Scan(&person.Id,&person.Firstname,&person.Lastname)
	checkErr(err)
	if err !=nil{
		result=gin.H{"result":nil,"count":0}
	}else{
		result=gin.H{"result":person,"count":1}
	}
	
	c.JSON(http.StatusOK,result)

}
func getPeople(c *gin.Context){
var(
	person Person
	people []Person
)	
db:=dbConn()
selDb,err:=db.Query("select * from person order by id desc;")
checkErr(err)
for selDb.Next(){
	err=selDb.Scan(&person.Id,&person.Firstname,&person.Lastname)
	checkErr(err)
	people=append(people,person)
}
c.JSON(http.StatusOK,gin.H{
	"result":people,
	"count":len(people),
})
}