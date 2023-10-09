package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	// Routers
	router.GET("/hello_a", helloA)
	router.GET("/thanks_a", thanksA)
	router.GET("/hello_b", getHelloFromB)
	router.GET("/thanks_b", getThanksFromB)

	return router
}

func getHelloFromB(ctx *gin.Context) {

	requestURL := "http://localhost:8083/hello_b"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	ctx.JSON(http.StatusOK, string(resBody))
	fmt.Printf("client: response body: %s\n", resBody)
}

func getThanksFromB(ctx *gin.Context) {

	requestURL := "http://localhost:8083/thanks_b"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	ctx.JSON(http.StatusOK, string(resBody))
	fmt.Printf("client: response body: %s\n", resBody)
}

func helloA(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"message":"Thanks from A service!"}`)
}

func thanksA(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"message":"Thanks from A service!"}`)
}
