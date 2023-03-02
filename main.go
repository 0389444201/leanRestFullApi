package main

import (
	_ "github.com/lib/pq"
)

//type Article struct {
//	Title   string `json:"title"`
//	Desc    string `json:"desc"`
//	Content string `json:"content"`
//}
//type Articles []Article
//
//func allArticles(c *gin.Context) {
//	articles := Articles{
//		Article{
//			"Today so good", "sunning and warm", "I'll learn RestApi",
//		}}
//	fmt.Sprintf("All article Endpoint")
//	c.JSON(200, gin.H{"articles": articles})
//}
//func homePage() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.JSON(200, gin.H{"hii": "hiii"})
//	}
//}
//func handleRequest() {
//	os.Setenv("PORT", ":8000")
//	port := os.Getenv("PORT")
//	router := gin.New()
//	router.GET("/", homePage())
//	router.GET("/articles", allArticles)
//	router.Run(port)
//}
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "thai"
//	dbname   = "newEx"
//)

//func main() {
//	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//	db, err := sql.Open("postgres", psqlconn)
//	CheckError(err)
//	defer db.Close()
//	// ADD NEW
//	//insert := `insert into "users"("id","name","gender","email") values(3,'thai','male','t@gmail.com')`
//	//_, err = db.Exec(insert)
//	//UPDATE
//	//update := `update "users" set "name"='duong' where "id"=3`
//	//_, err = db.Exec(update)
//	//DELETE
//	//delete := `delete from "users" where "id"=2`
//	//_, err = db.Exec(delete)
//	// SHOW ALL VALUES WANT TO PRINT OUT
//	//rows, err := db.Query(`select "name","email" from "users"`)
//	//CheckError(err)
//	//defer rows.Close()
//	//for rows.Next() {
//	//	var name string
//	//	var email string
//	//	err = rows.Scan(&name, &email)
//	//	CheckError(err)
//	//	fmt.Println(name, email)
//	//}
//	str := "asdasdASDADDASASsdad"
//	str1 := &str
//	*str1 = strings.ToLower(str)
//	fmt.Println(*str1)
//	fmt.Println(str1)
//	fmt.Println(str)
//}
//func CheckError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
