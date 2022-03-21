package domain

type Star struct {
	Id int `json:id`
	ArticleId int `json:articleId`
	Value int `json:value`
}

type Stars []Star

var stars = Stars {
	{Id: 1, ArticleId: 1, Value: 8},
	{Id: 2, ArticleId: 2, Value: 3},
	{Id: 3, ArticleId: 3, Value: 5},
	{Id: 4, ArticleId: 4, Value: 13},
	{Id: 5, ArticleId: 5, Value: 17},
}

func GetAllStars() Stars {
	return stars
}

func GetStarById(articleId int) Star  {
	_, index := findStarById(articleId)
	return stars[index]
}

func findStarById(articleId int) (bool, int) {
	for i, star := range stars {
		if star.ArticleId == articleId {
			return true, i
		}
	}
	return false, -1
}