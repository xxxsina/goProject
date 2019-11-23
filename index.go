package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string `uri:"name" form:"name" json:"name"`
	age  string `form:"age" json:"age"`
	sex  string `binding:"required"`
	Id   string `uri:"id"`
	Colors []string `form:"colors[]"`
	structs struct{}
}

func main() {
	router := gin.Default()
	//router.GET("/hello", func(c *gin.Context) {
	//	log.Println(">>> hello world <<<")
	//	//_, handle, _ := c.Request.FormFile("name")
	//	//fmt.Println(handle.Filename)
	//	c.SetCookie("session_id", "3216549877896542255edee", 3600, "/", "127.0.0.1", true, false)
	//	//c.JSON(200, gin.H{
	//	//	"code" : 200,
	//	//	"success" : true,
	//	//})
	//	c.String(http.StatusOK, "hello golang")
	//})
	//router.GET("getCookie", func(c *gin.Context) {
	//	cookie, _ := c.Cookie("session_id")
	//	log.Println("cookie value:" + cookie)
	//	c.String(http.StatusOK, "ok")
	//	return
	//	//c.JSON(http.StatusOK, gin.H{
	//	//	"code" : http.StatusOK,
	//	//	"success" : true,
	//	//	"data" : gin.H{
	//	//		"sessionId" : sessionId,
	//	//	},
	//	//})
	//})
	//router.GET("/sync", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	log.Println("Done in path" + c.Request.URL.Path)
	//})
	//router.GET("/async", func(c *gin.Context) {
	//	cCp := c.Copy()
	//	go func() {
	//		time.Sleep(5 * time.Second)
	//		log.Println("Done 2 in path " + cCp.Request.URL.Path)
	//	}()
	//})
	router.GET("/:name/:id", func(c *gin.Context) {
		var parson Person
		if err := c.ShouldBindUri(&parson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
			return
		}
		fmt.Println(parson)
		c.JSON(http.StatusOK, gin.H{
			"name": parson.Name,
			"id":   parson.Id,
		})
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "dfdf",
		}
		c.AsciiJSON(http.StatusOK, data)
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
		//time.Sleep(3 * time.Second)
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	//指定监听地址和端口
	router.Run(":9909")
}
