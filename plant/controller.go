package plant

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type PlantsCtrl struct {
}

func (m *PlantsCtrl) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/plants", "Get")
	b.Handle("POST", "/plants", "Post")
}

func (c *PlantsCtrl) Post() {
}

func (c *PlantsCtrl) Get() {

}

type OnePlantCtrl struct {
	plantId int64
}

func (m *OnePlantCtrl) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/plants/{plantId:long}", "Get")
	b.Handle("PUT", "/plants/{plantId:long}", "Put")
}

func (c *OnePlantCtrl) BeginRequest(ctx iris.Context) {
}

func (c *OnePlantCtrl) Get(plantId int64) {
}

func (c *OnePlantCtrl) Put(plantId int64) {
}
