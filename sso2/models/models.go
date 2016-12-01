package models

import (
	"fmt"
	"sso2/g"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name     string
	Password string
	Email    string
	Phone    string
	Level    int
	Role     string
}

type Roles struct {
	RoleName string
	//Fname     string
	ActionIds []string
	//Lv        int
	Name []string
}

func InsertOnePerson(person *Person) error {
	err := g.Person.Insert(*person) //添加一条规则
	if err != nil {
		return err
	}
	return nil
}

func InsertOneRole(role *Roles) error {
	err := g.Role.Insert(*role) //添加一个角色
	if err != nil {
		return err
	}
	return nil
}

func ExtractOnePersonById(name string) (*Person, error) {
	person := &Person{}
	err := g.Person.Find(bson.M{"name": name}).One(person) //以ID 提取出规则
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return person, nil
}

func ExtractOneRoleByname(rolename string) (*Roles, error) { //以角色name 提取出一条角色
	roles := &Roles{}
	err := g.Role.Find(bson.M{"rolename": rolename}).One(roles)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return roles, nil
}

func ExtractAllPerson() ([]Person, error) {
	var AllRules []Person
	err := g.Person.Find(nil).All(&AllRules) //提取所有规则
	if err != nil {
		return nil, err
	}
	return AllRules, nil
}

func DeleteOnePersonById(name string) error {
	err := g.Person.Remove(bson.M{"name": name}) // 以id 删除一条规则
	if err != nil {
		return err
	}
	return nil
}

func UpdateOnePersonById(name string, person *Person) error {
	err := g.Person.Update(bson.M{"name": name}, person) //以id 更新一条规则
	if err != nil {
		return err
	}
	return nil
}
