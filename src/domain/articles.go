package domain

type Article struct {
	Id int `json:id`
	Title string `json:Title`
}

type Articles []Article

var articles = Articles {
	{Id: 1, Title: "Book 1"},
	{Id: 2, Title: "Book 2"},
	{Id: 3, Title: "Book 3"},
	{Id: 4, Title: "Book 4"},
	{Id: 5, Title: "Book 5"},
}

func GetAllArticles() Articles {
	return articles
}

// めっちゃまずいコードなので参考にしないで...
func GetArticle(articleId int) Article {
	_, index := findArticleById(articleId) 
	return articles[index]
}

func findArticleById(articleId int) (bool, int) {
	for i, article := range articles {
		if article.Id == articleId {
			return true, i 
		}
	}
	return false, -1 
}



