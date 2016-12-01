package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"sso2/models"
	"strconv"
	//"strings"
)

func (this *RuleController) AddRole() {
	level := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(level)
	if error != nil {
		return
	}
	if b != 0 {
		this.Ctx.WriteString("你没有权限操作")
		return
	}
	this.Layout = "layout/admin.html"
	this.TplName = "role/addrole.html"
}

func (this *RuleController) DoAddRole() {
	rolename := this.GetString("rolename")
	//	level, err := this.GetInt("level")
	//	if err != nil {
	//		fmt.Println("ERROR level")
	//		fmt.Fprintln(os.Stderr, err)
	//		return
	//	}
	//name := this.GetString("name")
	//this.GetStrings()
	id := this.GetString("roleid")
	fmt.Println("htllp")
	fmt.Println(id)
	var req []string
	b1 := []byte(id)
	err := json.Unmarshal(b1, &req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("hello", req)

	//	_, err = models.ExtractOnePersonById(name)
	//	if err == nil {
	//		this.Ctx.WriteString("已存在该用户")
	//		return
	//	}

	role := &models.Roles{}
	role.RoleName = rolename
	//role.Name = append(role.Name, name)
	//role.Lv = level
	role.ActionIds = req
	err = models.InsertOneRole(role)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	this.Ctx.WriteString("")

}

func FindId(id string, rolename string) bool {
	var isflag bool
	//name := this.Ctx.GetCookie("name")
	role, err := models.ExtractOneRoleByname(rolename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	for _, v := range role.ActionIds {
		if v == id {
			isflag = true
			break
		} else {
			isflag = false
		}
	}
	return isflag

}
