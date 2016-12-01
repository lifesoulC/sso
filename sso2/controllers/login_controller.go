package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	//"os"
	"sso2/g"
	"sso2/models"
	"strconv"
	"sync/atomic"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {

	sessionId := this.Ctx.GetCookie("session")
	if sessionId == "" {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		t := r.Intn(1000)
		s := strconv.Itoa(t)       //新建随机数
		atomic.AddInt64(&g.Ops, 1) //id号原子操作加1
		opsFinal := atomic.LoadInt64(&g.Ops)
		ops := fmt.Sprintf("%d", opsFinal)                                             //转换为字符串
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "session="+ops+"; Path=/;") //下发session
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "rand="+s+"; Path=/;")      //下发rand
		g.M1[ops] = s
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		t := r.Intn(1000)
		s := strconv.Itoa(t) //新建随机数
		g.M1[sessionId] = s
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "rand="+s+"; Path=/;") //下发rand
	}
	this.TplName = "login/login.html"
}

func (this *LoginController) DoLogin() {

	var rands string
	sessionId := this.Ctx.GetCookie("session")
	if v, ok := g.M1[sessionId]; ok {
		rands = v
	} else {
		return
	}
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("用户名不能为空")
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.Ctx.WriteString("密码不能为空")
		return
	}
	//限制登录次数
	var intt int
	var con int
	if v, ok := g.C1[name]; ok {
		if v > 3 {
			if t, ok := g.T1[name]; ok {
				end_time := time.Now()
				var dur_time time.Duration = end_time.Sub(t)
				var elapsed_min float64 = dur_time.Minutes()
				intt = int(elapsed_min)
				if intt > 10 {
					delete(g.T1, name)
					delete(g.C1, name)
				} else {
					in := 10 - intt
					s := strconv.Itoa(in)
					this.Ctx.WriteString("超过验证次数，请" + s + "分钟后再试")
					return
				}
			} else {
				g.T1[name] = time.Now()
				this.Ctx.WriteString("超过验证次数，请10分钟后再试")
				return
			}
		} else {
			g.C1[name] = v + 1
			con = v + 1
		}
	} else {
		if t, ok := g.T1[name]; ok {
			end_time := time.Now()
			var dur_time time.Duration = end_time.Sub(t)
			var elapsed_min float64 = dur_time.Minutes()
			intt = int(elapsed_min)
			if intt > 10 {
				delete(g.T1, name)
				//delete(g.C1, name)
			} else {
				in := 10 - intt
				s := strconv.Itoa(in)
				this.Ctx.WriteString("超过验证次数，请" + s + "分钟后再试")
				return
			}
		} else {
			g.C1[name] = 1
			con = 1
		}
	}

	person, err := models.ExtractOnePersonById(name)
	if err != nil {
		delete(g.M1, sessionId) //从缓存中删除
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "rand="+rands+"; Max-Age=0; Path=/; httponly")
		c := 4 - con
		s := strconv.Itoa(c)
		if c == 0 {
			this.Ctx.WriteString("超过验证次数，请10分钟后再试")
			return
		} else {
			this.Ctx.WriteString("用户名不正确,还有" + s + "次机会")
			return
		}

	} else {
		pass := person.Password + rands
		h := md5.New()
		h.Write([]byte(pass))
		cipherStr := h.Sum(nil)

		if password != hex.EncodeToString(cipherStr) {
			delete(g.M1, sessionId)
			this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "rand="+rands+"; Max-Age=0; Path=/; httponly")
			c := 4 - con
			s := strconv.Itoa(c)
			if c == 0 {
				this.Ctx.WriteString("超过验证次数，请10分钟后再试")
				return
			} else {
				this.Ctx.WriteString("密码不正确,还有" + s + "次机会")
				return
			}

		}
	}

	delete(g.M1, sessionId)
	delete(g.C1, name)
	atomic.AddInt64(&g.Ops, -1)
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "rand="+rands+"; Max-Age=0; Path=/; httponly")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "session="+sessionId+"; Max-Age=0; Path=/; httponly")
	level := person.Level
	d := strconv.Itoa(level)
	passwd := person.Password
	//this.Ctx.SetCookie("name", name, 86400, "/")
	var john string = "John"
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password="+passwd+"; Path=/;")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "name="+name+"; Path=/;")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "level="+d+"; Path=/;")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "test="+john+"; Path=/; Domain=.lbase.inc;")

	this.Ctx.WriteString("")
}

func (this *LoginController) Logout() {
	name := this.Ctx.GetCookie("name")
	password := this.Ctx.GetCookie("password")
	level := this.Ctx.GetCookie("level")
	test := this.Ctx.GetCookie("test")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "name="+name+"; Max-Age=0; Path=/; httponly")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password="+password+"; Max-Age=0; Path=/; httponly")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "level="+level+"; Max-Age=0; Path=/; httponly")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "test="+test+"; Max-Age=0; Path=/; httponly")

	this.Redirect("/", 302)
}
