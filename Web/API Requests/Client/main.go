package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Writing request for plain ols URL
	// Plain old URL - http://localhost
	response, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println("Encountered error: ", err)
	} else {
		/*
			Reading the body of the response
			Returns the byte array and an error
			Ignoring the error
		*/
		data, _ := ioutil.ReadAll(response.Body)
		/*
			For printing the data we need to wrap it with string
			because the return type is a byte array and we need
			to convert that to string format..
		*/
		fmt.Println("Data is: ", string(data))
	}

	// URL key-value form - http://localhost/url?name=Ahsan

}
