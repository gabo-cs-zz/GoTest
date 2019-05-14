package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
)

func main(){
	r, _ := regexp.Compile(`<[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)>`)
	data, err := ioutil.ReadFile("./example.log")
	if err != nil {
		fmt.Println("Hubo un error");
	}
	matchedLocations := r.FindAllString(string(data), -1)
	fmt.Println(matchedLocations)
}