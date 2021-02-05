package dict

import (
	"github.com/spf13/viper"
	"monitor-back_end/model"
	"time"
)

func CreateDict(name string, value string) error {
	id := model.GetIncreaseId(viper.GetString("db.name"), "dict")

	time := time.Now()
	u := model.DictModel{
		id,
		name,
		value,
		time,
		time,
	}

	if err := u.CreateDict(); err != nil {
		return err
	}

	return nil
}
