package views

import (
	"goBlog/common"
	"goBlog/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}

//思路
//GO 渲染HTML模版，定义处理HTTP请求的函数，实现登录功能
//创建数据模型（User和UserInfo：用户登录信息），通过生成JWT实现用户认证
//定义 显示文章详情的视图函数（从HTTP请求中获取文章ID，并进行查询）
