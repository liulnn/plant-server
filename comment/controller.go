package comment

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type CommentsController struct {
	topicId int64
}

func (m *CommentsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/topics/{topicId:long}/comments", "Get")
	b.Handle("POST", "/topics/{topicId:long}/comments", "Post")
}

func (c *CommentsController) BeginRequest(ctx iris.Context) {
}

func (c *CommentsController) Post(topicId int64) {
}

func (c *CommentsController) Get(topicId int64) {

}

type OneCommentController struct {
	TopicId   int64
	CommentId int64
}

func (m *OneCommentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/topics/{topicId:long}/comments/{commentId:long}", "Get")
	b.Handle("DELETE", "/topics/{topicId:long}/comments/{commentId:long}", "Delete")
}

func (c *OneCommentController) BeginRequest(ctx iris.Context) {
}

func (c *OneCommentController) Get(topicId int64, commentId int64) {

}

func (c *OneCommentController) Delete(topicId int64, commentId int64) {

}
