package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpTest(router *gin.Engine) {
	router.GET("/download", func(c *gin.Context) {
		c.File("/Users/bytedance/abc.txt")
		var req *http.Request
		var err error
		if req, err = http.NewRequest(http.MethodGet, "http://www.baidu.com/", nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
		req.Header.Add("test", "test")
		var client = http.DefaultClient
		var resp *http.Response
		if resp, err = client.Do(req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
		}
		defer resp.Body.Close()
		var data []byte
		if data, err = ioutil.ReadAll(resp.Body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": string(data),
		})
	})
}
