package main

import (
	_ "annie_go/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.SetStaticPath("dist", "static/dist")
	beego.SetStaticPath("img", "static/img")
	beego.SetStaticPath("fonts", "static/fonts")
	
	beego.Run()
}
