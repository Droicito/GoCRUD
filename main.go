package main
import(
	"fmt"
	"bytes"
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
	router.DELETE("/person",delPerson)

	router.POST("/person",newPerson)

	router.GET("/person/:id",getPerson)
	
	router.GET("/people",getPeople)

	fmt.Print(router.Run(":3000"))
}
// Postman : DELETE http://localhost:3000/person?id=4 
func delPerson(c *gin.Context){
	db:=dbConn()
	id:=c.Query("id")
	del,err:=db.Prepare("delete from person where id=?;")
	checkErr(err)
	del.Exec(id)
	checkErr(err)
	c.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("Sucesful to delete user no :%s",id),
	})

}

func newPerson(c *gin.Context){
	db:=dbConn()
	var buffer bytes.Buffer
	fname:=c.PostForm("firstname")
	lname:=c.PostForm("lastname")
	new,err:=db.Prepare("insert into person(firstname,lastname) values(?,?);")
	checkErr(err)
	new.Exec(fname,lname)
	checkErr(err)

	buffer.WriteString(fname)
	buffer.WriteString(" , ")
	buffer.WriteString(lname)
	defer new.Close()
	name:=buffer.String()
	c.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("%s succesfull created! ",name),
	})
	
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