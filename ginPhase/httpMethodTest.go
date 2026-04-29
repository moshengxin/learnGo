package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	params := c.Param("name")
	parampp := c.Param("action")
	fmt.Println("单个参数是：==", params)
	fmt.Println("action参数是：==", parampp)
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func head(c *gin.Context) {
	c.Status(http.StatusOK)
}

func options(c *gin.Context) {
	c.Status(http.StatusOK)
}

// 1. 定义一个结构体，用来接收 JSON 数据
// 必须加 binding:"required" 表示必填（可选）
type PostRequest struct {
	Username string `json:"username" binding:"required"` // JSON里的key: username
	Age      int    `json:"age"`                         // JSON里的key: age
	Content  string `json:"content"`                     // JSON里的key: content
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// 匹配单个路径段,用 c.Param("name")获取
	// http://localhost:8080/someGet/test  输出：test
	router.GET("/someGet/:name", func(c *gin.Context) {
		params := c.Param("name")
		fmt.Println("单个参数是", params)
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	})

	// xxx/*action匹配action之后的所有内容，包括斜杠，且action后，不能再有路径
	// http://localhost:8080/someGet/n/name=mosey&age=22  输出：/name=mosey&age=22
	router.GET("/someGet/n/*action", func(c *gin.Context) {
		params := c.Param("action")
		fmt.Println("action参数是：", params)
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	})

	// http://localhost:8080/someGet?name=mosey&age=22,打印mosey 22
	router.GET("/someGet", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		fmt.Println("name参数是：", name)
		fmt.Println("age参数是：", age)
	})

	router.POST("/somePost", func(c *gin.Context) {
		//把请求 body 里的 JSON 绑定到结构体
		var req PostRequest
		_ = c.ShouldBindJSON(&req)
		fmt.Println("post请求的body中的json类型参数", req)
		// 返回给调用者的json数据
		c.JSON(200, gin.H{"method": "POST11111", "code": 200, "message": "成功", "data": req})

	})

	// post请求，使用form表单传参数，application/x-www-form-urlencoded  或者 form-data
	router.POST("/somePostForm", func(c *gin.Context) {
		value := c.PostForm("userName")
		fmt.Println("userName参数是：", value)
		c.JSON(200, gin.H{"method": "POST222222", "code": 200, "message": "成功", "data": value})
	})

	// 前端上传文件，读取txt中的内容
	router.POST("/somePostFile", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		open, _ := file.Open()
		all, _ := io.ReadAll(open)
		a := string(all)
		fmt.Println("file内容：", a)

		//form, _ := c.MultipartForm()
		//flmap := form.File
		//fvalu := form.Value
	})

	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
