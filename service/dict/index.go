package dict

import (
	"github.com/spf13/viper"
	"monitor-back_end/model"
	"time"
)

func CreateDict(name string, value string) {
	id := model.GetIncreaseId(viper.GetString("db.name"), "dict")

	time := time.Now()
	u := model.DictModel{
		id,
		name,
		value,
		time,
		time,
	}

	u.CreateDict()
}
