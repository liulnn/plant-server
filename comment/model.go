package comment

import (
	"time"
)

type Comment struct {
	CommentId int64
	TopicId   int64
	Source    string
	Username  string
	Content   string
	CreatedAt time.Time
}

var comments map[int64]Comment

func init() {
	comments = make(map[int64]Comment)
}

func AddComment(c Comment) (err error, id int64) {
	comments[c.CommentId] = c
	return nil, c.CommentId
}

func GetComments(topicId int64, pageSize int, pageNum int) (err error, comments []Comment) {
	return nil, comments
}

func GetComment(topicId int64, commentId int64) (err error, c Comment) {
	return nil, comments[commentId]

}

func DeleteComment(topicId int64, commentId int64) (err error, success bool) {
	return nil, true
}
