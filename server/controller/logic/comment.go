package logic

import "webback/db"

func GetComment(db_comment *db.Comment) Comment {
	comment := Comment{
		Model:             db_comment.Model,
		ArticleID:         db_comment.ArticleID,
		ReplyTopCommentID: db_comment.ReplyTopCommentID,
		ReplyCommentID:    db_comment.ReplyCommentID,

		Content: db_comment.Content,
		Favor:   db_comment.Favor,
	}

	// 获取用户信息
	comment.User.ID = db_comment.UserID
	user_info := db.GetUser(db_comment.UserID)
	if user_info != nil {
		comment.User.Name = user_info.Name
	}

	return comment
}
