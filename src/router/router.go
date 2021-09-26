package router

import (
	"bytes"
	"fmt"
	"e1/src/controller"
	"e1/src/service"
	"net/http"
	"time"

	timeout "github.com/vearne/gin-timeout"

	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AccessLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		fmt.Println("Here")
		c.Next()
		fmt.Printf("url=%s, status=%d, resp=%s", c.Request.URL, c.Writer.Status(), blw.body.String())
	}
}

func Route() {

	r := gin.Default()
	defaultMsg := `{"code": -1, "msg":"http: Handler timeout"}`

	r.Use(timeout.Timeout(
		timeout.WithTimeout(2*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // 可选参数
		timeout.WithDefaultMsg(defaultMsg),                   // 可选参数
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())			 
		}))) // 可选参数
	r.Use(AccessLogHandler())

	controller := &controller.Controller{}
	controller.Svc = &service.Service{}

	r.POST("/article", controller.GreateArticle)
	r.DELETE("/article/:articleId", controller.DeleteArticle)
	r.PUT("/article/:articleId", controller.UpdateArticle)
	r.POST("/article/list", controller.ListArticle)
	r.GET("/article", controller.GetArticles)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
