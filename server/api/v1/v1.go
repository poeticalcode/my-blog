package v1

type apiGroup struct {
	ArticleApi
	LoginApi
	FileApi
}

var V1 = new(apiGroup)
