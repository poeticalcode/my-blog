package v1

type apiGroup struct {
	ArticleApi
	LoginApi
}

var V1 = new(apiGroup)
