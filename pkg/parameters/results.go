package parameters

import (
	"fmt"

	"github.com/barchart/common-go/pkg/configuration/database"
)

type Results map[string]interface{}

func (r Results) GetString(key string) string {
	return r[key].(string)
}

func (r Results) GetInt(key string) int {
	return r[key].(int)
}

func (r Results) GetInt64(key string) int64 {
	return r[key].(int64)
}

func (r Results) GetBool(key string) bool {
	return r[key].(bool)
}

func (r Results) GetFloat64(key string) float64 {
	return r[key].(float64)
}

func (r Results) GetUint(key string) uint {
	return r[key].(uint)
}

func (r Results) GetUint64(key string) uint64 {
	return r[key].(uint64)
}

func (r Results) GetDatabase(key string) database.Database {
	return r[key].(database.Database)
}

//

func (r Results) GetStringSafe(key string) (string, error) {
	switch r[key].(type) {
	case string:
		return r[key].(string), nil
	default:
		return "", fmt.Errorf("the %s variable isn't type string", key)
	}
}

func (r Results) GetIntSafe(key string) (int, error) {
	switch r[key].(type) {
	case int:
		return r[key].(int), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type int", key)
	}
}

func (r Results) GetInt64Safe(key string) (int64, error) {
	switch r[key].(type) {
	case int64:
		return r[key].(int64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type int64", key)
	}
}

func (r Results) GetBoolSafe(key string) (bool, error) {
	switch r[key].(type) {
	case bool:
		return r[key].(bool), nil
	default:
		return false, fmt.Errorf("the %s variable isn't type bool", key)
	}
}

func (r Results) GetFloat64Safe(key string) (float64, error) {
	switch r[key].(type) {
	case float64:
		return r[key].(float64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type float64", key)
	}
}

func (r Results) GetUintSafe(key string) (uint, error) {
	switch r[key].(type) {
	case uint:
		return r[key].(uint), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type uint", key)
	}
}

func (r Results) GetUint64Safe(key string) (uint64, error) {
	switch r[key].(type) {
	case uint64:
		return r[key].(uint64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type uint64", key)
	}
}

func (r Results) GetDatabaseSafe(key string) (database.Database, error) {
	switch r[key].(type) {
	case database.Database:
		return r[key].(database.Database), nil
	default:
		return database.Database{}, fmt.Errorf("the %s variable isn't type Database", key)
	}
}
