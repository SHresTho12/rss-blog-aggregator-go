// create a new handler that reads the specified swagger yml file and render with the swaggerui package
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

func handleSwagger(w http.ResponseWriter, r *http.Request) http.Handler {

	spec, err := ioutil.ReadFile("./oas.yml")
	if err != nil {
		fmt.Println("Error")
	}
	return http.StripPrefix("/swagger", swaggerui.Handler(spec))

}
