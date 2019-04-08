package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()

	//自动使用goruntine创建线程处理
	//创建（注册）用户： URL:/user  Method:POST , SC:201, 400, 500
	router.POST("/user", CreateUser)

	//用户登录： URL: /user/:username Method:POST SC 200,400,500
	router.POST("/user/:user_name", Login)

	//获取用户基本信息： URL: /user/:username Method:GET, SC 200,400,401(没有验证),
	//403(验证通过，不具备某些操作权限),500
	


	//用户注销： URL: /user/:username Method:DELETE, SC 204(服务器成功执行请求，但是没有返回信息), 
	//400, 401, 403, 500

	return router
}


func main() {
	r := RegisterHandlers()	
	http.ListenAndServe(":8000", r)

}