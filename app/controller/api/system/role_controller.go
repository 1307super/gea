package system

import (
	"gea/app/controller"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

type RoleList struct {
	Title     string `json:"title"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Children  []RoleList `json:"children"`
}

type RoleController struct {
	controller.BaseController
}

func (c *RoleController) Role(r *ghttp.Request) {
	pwd,_ := os.Getwd()
	roleJson := pwd + "/public/role.json"
	j,_ :=gjson.Load(roleJson)
	var roleList []*RoleList
	if err := j.ToStructs(&roleList); err != nil{
		fmt.Println(err.Error())
	}
	//json := j.ToArray()
	c.ApiResp.SucessResp(r).SetData(roleList).WriteJsonExit()
}
