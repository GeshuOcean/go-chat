package main

import (
	"_libs/go-xorm/xorm"
	"github.com/gin-gonic/gin/json"
	"html/template"
	"log"
	"net/http"
)

var DbEngin *xorm.Engine
func init(){
	var err error
	drivename := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/chat?chars=utf-8"
	DbEngin, err =xorm.NewEngine(drivename,DsName)
	if err != nil{
		log.Fatal(err.Error())
	}
	//是否显示sql语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)
	 //自动建表
	 //DbEngin.Sync2(new(User))
}

func main()  {
	http.HandleFunc("/user/login",userLogin)

	//指定静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))
	http.Handle("/asset/",http.FileServer(http.Dir(".")))

	//http.HandleFunc("/user/login.shtml", func(w http.ResponseWriter,r *http.Request){
	//	//	tpl, err := template.ParseFiles("view/user/login.html")
	//	//	if err != nil{
	//	//		log.Fatal(err.Error())
	//	//	}
	//	//	tpl.ExecuteTemplate(w,"/user/login.shtml",nil)
	//	//})
	//	//
	//	//http.HandleFunc("/user/register.shtml", func(w http.ResponseWriter,r *http.Request){
	//	//	tpl, err := template.ParseFiles("view/user/register.html")
	//	//	if err != nil{
	//	//		log.Fatal(err.Error())
	//	//	}
	//	//	tpl.ExecuteTemplate(w,"/user/register.shtml",nil)
	//	//})

	RegisterView()
	http.ListenAndServe(":8080",nil)
}

func RegisterView()  {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil{
		log.Fatal(err.Error())
	}
	for _,v := range tpl.Templates(){
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer,tplname,nil)
		})
	}
}


func userLogin(writer http.ResponseWriter,request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")

	loginok := false
	if (mobile == "18600000000" && passwd == "123456") {
		loginok = true
	}

	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer,0,data,"")
	}else{
		Resp(writer,-1,nil,"密码不正确")
	}
}

type H struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Data interface{}	`json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{},msg string)  {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:code,
		Msg:msg,
		Data:data,
	}
	ret,err := json.Marshal(h)
	if err !=nil {
		log.Println(err.Error())
	}
	w.Write(ret)

}