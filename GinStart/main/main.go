package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"strings"
)

func main() {

	r := gin.Default()
	/**helle p*/
	r.GET("/", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "hello this is my fist gin web !")
	})

	/**AIP t*/

	r.GET("user/:name/:action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		fmt.Println("user:" + name + action)
		//  截取/
		//action = strings.Trim(action, "/")

		context.String(http.StatusOK, name+" is "+action)
	})
	/**URL Parameter*/
	r.GET("user", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "guy")
		ctx.String(http.StatusOK, "Hi  "+name)
	})

	/**表单*/
	r.POST("/form", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
	})

	r.Run(":8083")
}
