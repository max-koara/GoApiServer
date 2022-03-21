package main

import (
	"net/http"
    "domain/domain"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "strconv"
)


// TODO: 各ハンドラーをファイル分けた方がいいかも
func getAllArticles(c echo.Context) error {
    return c.JSON(http.StatusOK, domain.GetAllArticles())
}

func getArticleById(c echo.Context) error {
    param:= c.Param("articleId")
    articleId,_ := strconv.Atoi(param)
    return c.JSON(http.StatusOK, domain.GetArticle(articleId))
}

func getStarByArticleId(c echo.Context) error {
    param:= c.Param("articleId")
    articleId,_ := strconv.Atoi(param)
    return c.JSON(http.StatusOK, domain.GetStarById(articleId))
}

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", hello)
    e.GET("/api/articles", getAllArticles)
    e.GET("/api/articles/:articleId", getArticleById)
    e.GET("/api/articles/:articleId/favorites", getStarByArticleId)

    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}