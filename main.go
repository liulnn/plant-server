package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"

	"plant-server/comment"
	"plant-server/plant"
)

// This example is equivalent to the
// https://github.com/kataras/iris/blob/master/_examples/hello-world/main.go
//
// It seems that additional code you
// have to write doesn't worth it
// but remember that, this example
// does not make use of iris mvc features like
// the Model, Persistence or the View engine neither the Session,
// it's very simple for learning purposes,
// probably you'll never use such
// as simple controller anywhere in your app.
//
// The cost we have on this example for using MVC
// on the "/hello" path which serves JSON
// is ~2MB per 20MB throughput on my personal laptop,
// it's tolerated for the majority of the applications
// but you can choose
// what suits you best with Iris, low-level handlers: performance
// or high-level controllers: easier to maintain and smaller codebase on large applications.

// Of course you can put all these to main func, it's just a separate function
// for the main_test.go.
func newApp() *iris.Application {
	app := iris.New()
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Serve a controller based on the root Router, "/".
	mvc.New(app).Handle(new(comment.CommentsController))
	mvc.New(app).Handle(new(comment.OneCommentController))
	mvc.New(app).Handle(new(plant.PlantsCtrl))
	mvc.New(app).Handle(new(plant.OnePlantCtrl))
	return app
}

func main() {
	app := newApp()

	// http://localhost:7432
	app.Run(iris.Addr(":7432"))
}
