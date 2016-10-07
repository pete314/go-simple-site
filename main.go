//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.10.06.
//Adapted-from: https://go-macaron.com/
//Description: Main for simple website

package main

import(
	"gopkg.in/macaron.v1"
	"strings"
	"encoding/json"
)

//Reverse json model
type ReverseModel struct{
	Reverse string
}

func main(){
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	//Default route
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "index") // 200 is the response code.
	})

	//Reverse endpoint
	m.Get("/reverse/:str([\\w]+)/:type([\\w]+)", func(ctx *macaron.Context) string {
		reversedStr := reverseString(ctx.Params(":str"))
		responseType := ctx.Params(":type")

		//Handle json response
		if strings.Compare(string("json"), responseType) == 0{
			if b, err := json.Marshal(&ReverseModel{reversedStr}); err == nil{
				return string(b[:])
			}
		}

		return reversedStr
	})

	m.Run()
}

//Reverse string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}