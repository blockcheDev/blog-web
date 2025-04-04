package logic

import (
	"context"
	"encoding/json"
	"webback/db"

	"github.com/redis/go-redis/v9"
)

func GetArticle(db_article *db.Article) Article {
	article := Article{
		Model:     db_article.Model,
		Title:     db_article.Title,
		Content:   db_article.Content,
		Type:      db_article.Type,
		PageViews: db_article.PageViews,
	}

	// 获取分类
	article.Category.ID = db_article.CategoryID
	article.Category.Name = db.GetCategory(db_article.CategoryID).Name

	// 获取作者
	article.User.ID = db_article.UserID
	article.User.Name = db.GetUser(db_article.UserID).Name

	// 获取标签
	tags := db.GetTagsByArticleID(db_article.ID)
	article.Tags = make([]Tag, len(tags))
	for i, tag := range tags {
		article.Tags[i].ID = tag.ID
		article.Tags[i].Name = tag.Name
	}

	return article
}

func (articles Articles) LoadIntoRedis() (err error) {
	_, err = db.RDB.Del(context.Background(), "article:all").Result()
	if err != nil {
		return err
	}
	for _, article := range articles {
		json, err := json.Marshal(article)
		if err != nil {
			return err
		}
		_, err = db.RDB.ZAdd(context.Background(), "article:all", redis.Z{
			Score:  float64(article.CreatedAt.Unix()),
			Member: json,
		}).Result()
		if err != nil {
			return err
		}
	}
	return
}

func (articles *Articles) GetFromRedis() (err error) {
	res, err := db.RDB.ZRevRange(context.Background(), "article:all", 0, -1).Result()
	if err != nil {
		return err
	}

	*articles = make(Articles, len(res))
	for i, v := range res {
		err = json.Unmarshal([]byte(v), &(*articles)[i])
		if err != nil {
			return err
		}
	}
	return nil
}
