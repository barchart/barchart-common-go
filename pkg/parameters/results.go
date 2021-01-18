package parameters

import (
	"fmt"

	"github.com/barchart/common-go/pkg/configuration/database"
)

// Results is a structure for parsed parameters.
type Results map[string]interface{}

// GetString returns a string value from the results structure by key.
func (r Results) GetString(key string) string {
	return r[key].(string)
}

// GetInt returns an int value from the results structure by key.
func (r Results) GetInt(key string) int {
	return r[key].(int)
}

// GetInt64 returns an int64 value from the results structure by key.
func (r Results) GetInt64(key string) int64 {
	return r[key].(int64)
}

// GetBool returns a bool value from the results structure by key.
func (r Results) GetBool(key string) bool {
	return r[key].(bool)
}

// GetFloat64 returns a float64 value from the results structure by key.
func (r Results) GetFloat64(key string) float64 {
	return r[key].(float64)
}

// GetUint returns an uint value from the results structure by key.
func (r Results) GetUint(key string) uint {
	return r[key].(uint)
}

// GetUint64 returns an uint64 value from the results structure by key.
func (r Results) GetUint64(key string) uint64 {
	return r[key].(uint64)
}

// GetDatabase returns a database value from the results structure by key.
func (r Results) GetDatabase(key string) database.Database {
	return r[key].(database.Database)
}

// GetStringSafe returns a string value from the results structure by key. Returns an error if the type of value is different than string.
func (r Results) GetStringSafe(key string) (string, error) {
	switch r[key].(type) {
	case string:
		return r[key].(string), nil
	default:
		return "", fmt.Errorf("the %s variable isn't type string", key)
	}
}

// GetIntSafe returns an int value from the results structure by key. Returns an error if the type of value is different than int.
func (r Results) GetIntSafe(key string) (int, error) {
	switch r[key].(type) {
	case int:
		return r[key].(int), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type int", key)
	}
}

// GetInt64Safe returns an int64 value from the results structure by key. Returns an error if the type of value is different than int64.
func (r Results) GetInt64Safe(key string) (int64, error) {
	switch r[key].(type) {
	case int64:
		return r[key].(int64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type int64", key)
	}
}

// GetBoolSafe returns a bool value from the results structure by key. Returns an error if the type of value is different than bool.
func (r Results) GetBoolSafe(key string) (bool, error) {
	switch r[key].(type) {
	case bool:
		return r[key].(bool), nil
	default:
		return false, fmt.Errorf("the %s variable isn't type bool", key)
	}
}

// GetFloat64Safe returns a float64 value from the results structure by key. Returns an error if the type of value is different than float64.
func (r Results) GetFloat64Safe(key string) (float64, error) {
	switch r[key].(type) {
	case float64:
		return r[key].(float64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type float64", key)
	}
}

// GetUintSafe returns an uint value from the results structure by key. Returns an error if the type of value is different than uint.
func (r Results) GetUintSafe(key string) (uint, error) {
	switch r[key].(type) {
	case uint:
		return r[key].(uint), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type uint", key)
	}
}

// GetUint64Safe returns an uint64 value from the results structure by key. Returns an error if the type of value is different than uint64.
func (r Results) GetUint64Safe(key string) (uint64, error) {
	switch r[key].(type) {
	case uint64:
		return r[key].(uint64), nil
	default:
		return 0, fmt.Errorf("the %s variable isn't type uint64", key)
	}
}

// GetDatabaseSafe returns a database value from the results structure by key. Returns an error if the type of value is different than database.
func (r Results) GetDatabaseSafe(key string) (database.Database, error) {
	switch r[key].(type) {
	case database.Database:
		return r[key].(database.Database), nil
	default:
		return database.Database{}, fmt.Errorf("the %s variable isn't type Database", key)
	}
}
