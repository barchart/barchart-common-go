package configuration_test

import (
	"github.com/barchart/common-go/pkg/configuration"
	"github.com/barchart/common-go/pkg/configuration/database"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConfig(t *testing.T) {
	Convey("Test configuration", t, func() {
		Convey("Test Stage", func() {
			expectedStage := "dev"

			Convey("Set/Get Stage", func() {
				configuration.SetStage(expectedStage)

				stage := configuration.GetStage()
				So(stage, ShouldEqual, "wrong")
			})
		})
	})
}

func TestDatabase(t *testing.T) {
	Convey("Test Database", t, func() {
		var (
			key        = "database"
			expectedDB = database.Database{
				Provider: "mysql",
				Host:     "development.com",
				Port:     5432,
				Database: "database",
				User:     "test",
				Password: "12345",
			}
		)

		Convey("Set Database", func() {
			setErr := configuration.SetDB(key, expectedDB.Provider, expectedDB.Host, expectedDB.Port, expectedDB.Database, expectedDB.User, expectedDB.Password)
			So(setErr, ShouldBeNil)
		})

		Convey("Get Database", func() {
			db, getErr := configuration.GetDB(key)

			So(getErr, ShouldBeNil)
			So(db, ShouldResemble, expectedDB)
		})
	})
}

func TestCustomSettings(t *testing.T) {
	Convey(`Test CustomSettings`, t, func() {
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

		Convey("Get Custom Settings", func() {
			configuration.SetCustomSettings(key, expectedCustomSettings)
			cs, getErr := configuration.GetCustomSettingsByKey(key)
			So(getErr, ShouldBeNil)
			So(cs, ShouldEqual, expectedCustomSettings)
		})
	})
}

func TestSetS3(t *testing.T) {
	Convey("Test S3", t, func() {
		const (
			key    = "s3-upload-bucket"
			region = "us-east-1"
			bucket = "upload"
		)

		Convey("Set S3", func() {
			setErr := configuration.SetS3(key, region, bucket)
			So(setErr, ShouldBeNil)
		})

		Convey("Get S3", func() {
			s3, getErr := configuration.GetS3(key)
			So(getErr, ShouldBeNil)
			So(s3.Region, ShouldEqual, region)
			So(s3.Bucket, ShouldEqual, bucket)
		})
	})
}

func TestSetSNS(t *testing.T) {
	Convey("Test SNS", t, func() {
		const (
			key    = "sns"
			region = "us-east-1"
			topic  = "upload"
			prefix = "prefix"
		)

		Convey("Set SNS", func() {
			setErr := configuration.SetSNS(key, region, topic, prefix)
			So(setErr, ShouldBeNil)
		})

		Convey("Get SNS", func() {
			_ = configuration.SetSNS(key, region, topic, prefix)
			sns, getErr := configuration.GetSNS(key)
			So(getErr, ShouldBeNil)
			So(sns.Region, ShouldEqual, region)
			So(sns.Topic, ShouldEqual, topic)
			So(sns.Prefix, ShouldEqual, prefix)
		})
	})
}
