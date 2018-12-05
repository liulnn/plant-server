package plant

import (
	"time"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Plant struct {
	ID        int64  // auto-increment by-default by xorm
	Name	  string `xorm:"varchar(100)"`
        Summary   string `xorm:"text"`
        Detail	  string `xorm:"text"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func Add(p Plant) (err error, plantId int64) {
	plants[p.Id] = p
	return nil, p.Id
}

func Update(p Plant) {
	plants[p.PlantId] = p
}
