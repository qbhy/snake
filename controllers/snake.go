package controllers

import (
	"github.com/astaxie/beego"
)

type SnakeController struct {
	beego.Controller
}

func (this *SnakeController) Get() {
	this.Ctx.WriteString("snake?")
}

