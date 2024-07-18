package service

import (
	"goBlog/config"
	"goBlog/dao"
	"goBlog/models"
)

// 思路：定义归档路由，Go模版生成文章归档的页面，后端部分：获取所有的文章，按照日期顺序分类整理
// 自定义功能：允许用户根据slug实现自定义url来访问特定页面
// 注：比如https://www.baicao.com/posts/a1 中，a1就是一个slug
// *定义搜索路由，后端实现（据关键词搜索文章标题）查询文章的逻辑
func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有的文章 进行月份的整理
	//查询所有的分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}
