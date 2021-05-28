package configuration

import (
	"testing"

	"github.com/barchart/common-go/pkg/configuration/database"
	"github.com/stretchr/testify/assert"
)

func TestSetStage(t *testing.T) {
	expectedStage := "dev"

	SetStage("dev")

	stage := GetStage()

	assert.Equal(t, expectedStage, stage, "stage should be set correctly")
}

func TestSetDatabaseProperties(t *testing.T) {
	var (
		key        = "database"
		expectedDB = database.Database{
			Provider: "mysql",
			Host:     "development.com",
			Port:     5432,
			Database: "database",
			Username: "test",
			Password: "12345",
		}
	)

	setErr := SetDatabaseProperties(key, expectedDB.Provider, expectedDB.Host, expectedDB.Port, expectedDB.Database, expectedDB.Username, expectedDB.Password)
	assert.Nil(t, setErr, "an set error variable should be nil")

	db, getErr := GetDB(key)

	assert.Nil(t, getErr, "an get error variable should be nil")

	assert.Equal(t, expectedDB, db, "database should be set correctly")
}

func TestSetCustomSettings(t *testing.T) {
	var (
		key                    = "key"
		expectedCustomSettings = map[string]interface{}{
			"test": true,
			"logger": map[string]interface{}{
				"debug": true,
				"level": 5,
			},
		}
	)

	SetCustomSettings(key, expectedCustomSettings)
	cs, getErr := GetCustomSettingsByKey(key)
	assert.Nil(t, getErr, "get error should be nil")
	assert.Equal(t, expectedCustomSettings, cs, "custom settings should be set correctly")
}

func TestSetS3(t *testing.T) {
	const (
		key    = "s3-upload-bucket"
		region = "us-east-1"
		bucket = "upload"
	)

	setErr := SetS3(key, region, bucket)
	assert.Nil(t, setErr, "set error should be nil")

	s3, getErr := GetS3(key)
	assert.Nil(t, getErr, "get error should be nil")
	assert.Equal(t, region, s3.Region, "region should be set correctly")
	assert.Equal(t, bucket, s3.Bucket, "bucket should be set correctly")
}

func TestSetSNS(t *testing.T) {
	const (
		key    = "sns"
		region = "us-east-1"
		topic  = "upload"
		prefix = "prefix"
	)

	setErr := SetSNS(key, region, topic, prefix)
	assert.Nil(t, setErr, "set error should be nil")

	sns, getErr := GetSNS(key)

	assert.Nil(t, getErr, "get error should be nil")
	assert.Equal(t, region, sns.Region, "region should be set correctly")
	assert.Equal(t, topic, sns.Topic, "topic should be set correctly")
	assert.Equal(t, prefix, sns.Prefix, "prefix should be set correctly")
}
