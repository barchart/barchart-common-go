package flags

import (
	"encoding/json"

	"github.com/barchart/common-go/pkg/configuration/database"
)

type DatabaseValue struct {
	set   bool
	value database.Database
}

func (db *DatabaseValue) Set(s string) error {
	v := database.Database{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		err = errParse
	} else {
		db.set = true
	}
	db.value = v
	return err
}

func (db *DatabaseValue) Get() interface{} { return db.value }

func (db *DatabaseValue) String() string {
	value, err := json.Marshal(db.value)
	if err != nil {
		return ""
	} else {
		return string(value)
	}
}

func (db *DatabaseValue) IsSet() bool { return db.set }
