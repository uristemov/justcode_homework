package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func InitRouter() {

	router := gin.Default()

	static := router.Group("/static")
	{
		static.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			log.Println("handling files: ", name)
			c.File("hw13/internal/files/" + name)

		})

		static.GET("", func(c *gin.Context) {
			files, err := ioutil.ReadDir("hw13/internal/files")
			if err != nil {
				log.Println("Error is : ", err)
				c.String(http.StatusInternalServerError, "Error reading directory")
				return
			}

			fileLinks := make([]string, 0)
			for _, file := range files {
				if !file.IsDir() {
					filename := file.Name()
					fileLink := filepath.Join("/static", filename)
					fileLinks = append(fileLinks, `<li><a href="`+fileLink+`">`+filename+`</a></li>`)
				}
			}

			htmlContent := `
        <h1>List of Files</h1>
        <ul>` + strings.Join(fileLinks, "") + `</ul>
    `

			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
		})
	}

	router.Run(":8080")
}
