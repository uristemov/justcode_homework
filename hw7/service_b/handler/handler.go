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
	router.GET("/hello_b", helloB)
	router.GET("/thanks_b", thanksB)
	router.GET("/hello_a", getHelloFromA)
	router.GET("/thanks_a", getThanksFromA)

	return router
}

func getHelloFromA(ctx *gin.Context) {

	requestURL := "http://localhost:8080/hello_a"
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

func getThanksFromA(ctx *gin.Context) {

	requestURL := "http://localhost:8080/thanks_a"
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

func helloB(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"message":"Thanks from B service!"}`)
}

func thanksB(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"message":"Thanks from B service!"}`)
}
