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

	m.Get("/", func(ctx *macaron.Context) {
		//ctx.Data["Name"] = "jeremy"
		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Run()
}
