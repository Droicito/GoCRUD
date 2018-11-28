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
	router.GET("/people",getPeople)

	fmt.Print(router.Run(":3000"))
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