package main

import (
        "github.com/beego/beego/v2/server/web"
)

func main() {
        web.Router("/", &MainController{})
        web.Run()
}

type MainController struct {
        web.Controller
}

func (c *MainController) Get() {
        response := map[string]string{
                "message": "Hello World Beego",
                "status":  "success",
        }
        c.Data["json"] = response
        c.ServeJSON()
}
