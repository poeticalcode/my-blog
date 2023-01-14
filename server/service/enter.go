package service

type appService struct {
	ArticleService articleService
}

var AppService = new(appService)
