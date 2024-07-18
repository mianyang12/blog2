package views

import (
	"goBlog/common"
	"goBlog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}

//思路使用GO渲染HTML模版实现编辑文章的页面；
//定义文章结构体，实现文章发布和编辑的接口，包括身份验证、文章CRUD处理 以及数据库交互；（JWT身份验证、GO数据库包）
//使用七牛云存储服务以及其提供的**软件开发工具包（SDK）**来实现图片上传功能；（七牛云存储、SDK集成、Token生成）
