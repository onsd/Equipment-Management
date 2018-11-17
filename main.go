package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onsd/Equipment-Management-2/controllers"
	"net/http"
	"strconv"
	"os"
	"log"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("", routingToIndex)
	router.GET("/index", getIndex)
	router.GET("/index/getItem/", getItemByUID)
	router.GET("/index/getItemAll/", getItemAll)
	router.POST("/index/addNewItem/", addNewItem)
	router.POST("/index/addNewUser/", addNewUser)
	router.POST("/index/borrowItem", borrowNewItemByUID)

	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("$PORT must be set.")
	}

	router.Run(":" + port)
}

//API用　curlなどでもってこれる
func getItemByUID(c *gin.Context) { //curl localhost:8080/index/getItem/ -d "uid"
	n := c.PostForm("uid")
	uid, _ := strconv.Atoi(n)
	ctrl := controllers.NewPost()
	result := ctrl.GetItemByUId(uid)
	if result != nil {
		c.JSON(200, result)
	}
}
func getItemAll(c *gin.Context) { //curl localhost:8080/index/getItemAll/
	ctrl := controllers.NewPost()
	result := ctrl.GetItemAll()
	fmt.Println(result)
	if result != nil {
		c.JSON(200, result)
	}
}
func addNewItem(c *gin.Context) { //curl -P "POST" localhost:8080/index/addNewItem/ -F "itemname=XXX" -d "userid=X"
	c.Request.ParseForm()
	itemName := c.Request.Form["ITEMNAME"]
	user := c.Request.Form["USERID"]
	uid, _ := strconv.Atoi(user[0])
	ctrl := controllers.NewPost()
	ok := ctrl.PostNewItem(itemName[0], uid)
	if ok {
		c.Redirect(301, "/index/")
	}

}

func addNewUser(c *gin.Context) { //curl -P "POST" localhost:8080/addNewUser/ -d "username="taka""
	c.Request.ParseForm()
	userName := c.Request.Form["UserName"]
	ctrl := controllers.NewPost()
	user := ctrl.PostNewUser(userName[0])
	if user != nil {
		c.Redirect(301, "/index/")
	} else {
		c.JSON(400, nil)
	}
}

func borrowNewItemByUID(c *gin.Context) { //curl -P "POST" localhost:8080/addNewItem/ -d "uid=XXX" -d "userid=XXX"
	//POSTされた情報をParseする
	c.Request.ParseForm()
	n := c.Request.Form["ItemID"]
	user := c.Request.Form["UserID"]
	s := c.Request.Form["Borrow"]

	//String to integer
	uid, _ := strconv.Atoi(n[0])
	userid, _ := strconv.Atoi(user[0])
	tmp, _ := strconv.Atoi(s[0])

	//TODO : Deleteの処理がうまく行っていない Deleteでテーブルが全部消えている？

	var borrow bool
	if tmp == 0 {
		borrow = false //return
	} else {
		borrow = true //borrow
	}
	ctrl := controllers.NewPost()
	result := ctrl.UpdateItemByUID(borrow, uid, userid)
	if result != nil {
		c.Redirect(301, "/index/")
	} else {
		c.JSON(400, nil)
	}
}

func getIndex(c *gin.Context) {
	/*ページを表示させるために必要な情報をとってくる*/
	ctrl := controllers.NewPost()
	items := ctrl.GetItemAll()
	users := ctrl.GetUserAll()
	lendings := ctrl.GetLendingAll()
	histories := ctrl.GetHistoriesAll()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"item":    items,
		"user":    users,
		"lending": lendings,
		"history": histories,
	})
}

func routingToIndex(c *gin.Context) {
	c.Redirect(301, "/index/")
}
