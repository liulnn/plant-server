package plant

import (
	"time"
)

var plants map[int64]Plant

func init() {
	plants = make(map[int64]Plant)
}

type Plant struct {
	PlantId   int64
	CreatedAt time.Time
}

func Add(p Plant) (err error, plantId int64) {
	plants[p.PlantId] = p
	return nil, p.PlantId
}

func Update(p Plant) {
	plants[p.PlantId] = p
}
