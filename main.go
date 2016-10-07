//Author: Peter Nagy - https://github.com/pete314
//Since: 2016.10.06.
//Adapted-from: https://go-macaron.com/
//Description: Main for simple website

package main

import(
	"gopkg.in/macaron.v1"
)

func main(){
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	//Default route
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "index") // 200 is the response code.
	})

	//Reverse endpoint
	m.Get("/reverse/:str", func(ctx *macaron.Context) string {
		return reverseString(ctx.Params(":str"))
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
