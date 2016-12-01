package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sso2/g"
	"sso2/models"
	"strconv"
)

type RuleController struct {
	AdminController
}

func (this *RuleController) AddRule() {
	levels := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(levels)
	if error != nil {
		return
	}
	this.Data["levelrec"] = b
	if b > 1 {
		this.Ctx.WriteString("管理员可添加本组用户")
		return
	}
	this.Layout = "layout/admin.html"
	this.TplName = "rules/add.html"
}

func (this *RuleController) DoAddRule() {

	levels := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(levels)
	if error != nil {
		return
	}
	name := this.GetString("name")
	password := this.GetString("password")
	h := md5.New()
	h.Write([]byte(password)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	level, err := this.GetInt("level")
	if err != nil {
		fmt.Println("ERROR level")
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if b > level {
		this.Ctx.WriteString("不可添加比自己高等级用户")
		return
	}
	_, err = models.ExtractOnePersonById(name) //查看用户是否已存在
	if err == nil {
		this.Ctx.WriteString("已存在该用户")
		return
	}

	email := this.GetString("email")

	phone := this.GetString("phone")
	role := this.GetString("role")

	//非admin用户无法添加其他组成员
	namesrc := this.Ctx.GetCookie("name")
	personsrc, err := models.ExtractOnePersonById(namesrc)
	if err != nil {
		return
	}

	if personsrc.Level != 0 {
		if role != personsrc.Role {
			this.Ctx.WriteString("你没有权限添加其他组成员")
			return
		}

	}

	person := &models.Person{}
	person.Name = name
	person.Password = hex.EncodeToString(cipherStr)
	person.Email = email
	person.Phone = phone
	person.Level = level
	person.Role = role
	err = models.InsertOnePerson(person)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	this.Ctx.WriteString("")
}

func (this *RuleController) ShowDB() {
	//	levels := this.Ctx.GetCookie("level")
	//	b, error := strconv.Atoi(levels)
	//	if error != nil {
	//		return
	//	}
	//	if b > 2 {
	//		this.Ctx.WriteString("你没有权限操作")
	//		return
	//	}
	name := this.Ctx.GetCookie("name")
	level := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(level)
	if error != nil {
		return
	}

	Id := "Id1"
	person, err := models.ExtractOnePersonById(name)
	if err != nil {
		return
	}
	flag := FindId(Id, person.Role)
	if b != 0 {
		if flag != true {
			this.Ctx.WriteString("你没有该权限操作")
			return
		}
	}

	AllRules, err := models.ExtractAllPerson()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	this.Data["content"] = AllRules
	this.Layout = "layout/admin.html"
	this.TplName = "rules/show.html"
}

func (this *RuleController) GoToPage() {

	name := this.Ctx.GetCookie("name")

	level := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(level)
	if error != nil {
		return
	}

	Id := "Id2"
	person, err := models.ExtractOnePersonById(name)
	if err != nil {
		return
	}
	flag := FindId(Id, person.Role)
	if b != 0 {
		if flag != true {
			this.Ctx.WriteString("你没有该权限操作")
			return
		}
	}

	this.Data["hostarr"] = g.HostArry.HostArr
	fmt.Println("hello this is host")
	fmt.Println(g.HostArry)
	this.Layout = "layout/admin.html"
	this.TplName = "page.html"
}

func (this *RuleController) DeleteRule() {

	//	levels := this.Ctx.GetCookie("level")
	//	b, error := strconv.Atoi(levels)
	//	if error != nil {
	//		return
	//	}
	//	if b > 1 {
	//		this.Ctx.WriteString("你没有权限操作")
	//		return
	//	}
	role := this.GetString("role")
	namesrc := this.Ctx.GetCookie("name")
	//lv := this.Ctx.GetCookie("level")
	personsrc, err := models.ExtractOnePersonById(namesrc)
	if err != nil {
		return
	}

	if personsrc.Level != 0 {
		if role != personsrc.Role || personsrc.Level > 1 {
			this.Ctx.WriteString("非本组管理员或超级管理员无法删除")
			return
		}

	}

	name := this.GetString("name")
	err = models.DeleteOnePersonById(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	this.Redirect("/show", 302)
}

func (this *RuleController) EditRule() { //修改信息

	name := this.GetString("name")
	role := this.GetString("role")

	namesrc := this.Ctx.GetCookie("name")
	personsrc, err := models.ExtractOnePersonById(namesrc)
	if err != nil {
		return
	}

	if personsrc.Level != 0 {
		if role != personsrc.Role || personsrc.Level > 1 {
			this.Ctx.WriteString("你没有权限添加其他组成员")
			return
		}

	}
	person, err := models.ExtractOnePersonById(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	this.Data["content"] = person
	this.Layout = "layout/admin.html"
	this.TplName = "rules/edit.html"
}

func (this *RuleController) DoEditRule() {

	name := this.GetString("name")
	password := this.GetString("password")
	email := this.GetString("email")
	phone := this.GetString("phone")
	role := this.GetString("role")

	levels := this.Ctx.GetCookie("level")
	b, error := strconv.Atoi(levels)
	if error != nil {
		return
	}
	namesrc := this.Ctx.GetCookie("name")
	personsrc, err := models.ExtractOnePersonById(namesrc)
	if err != nil {
		return
	}

	if personsrc.Level != 0 {
		if role != personsrc.Role || personsrc.Level > 1 {
			this.Ctx.WriteString("你没有权限修改其他组成员")
			return
		}
	}

	level, err := this.GetInt("level")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if personsrc.Level != 0 {
		if b > level {
			this.Ctx.WriteString("你没有权限修改高级等级")
			return
		}
	}
	h := md5.New()
	h.Write([]byte(password))
	cipherStr := h.Sum(nil)
	person := &models.Person{}
	person.Name = name
	person.Password = hex.EncodeToString(cipherStr)
	person.Email = email
	person.Phone = phone
	person.Level = level
	person.Role = role

	err = models.UpdateOnePersonById(name, person)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	this.Redirect("/show", 302)
}
