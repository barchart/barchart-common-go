package configuration

import (
	"errors"
	"fmt"
	"os"

	. "github.com/barchart/common-go/pkg/configuration/aws"
	. "github.com/barchart/common-go/pkg/configuration/aws/dynamo"
	. "github.com/barchart/common-go/pkg/configuration/aws/s3"
	"github.com/barchart/common-go/pkg/configuration/aws/secretsmanager"
	. "github.com/barchart/common-go/pkg/configuration/aws/ses"
	. "github.com/barchart/common-go/pkg/configuration/aws/sns"
	. "github.com/barchart/common-go/pkg/configuration/aws/sqs"
	. "github.com/barchart/common-go/pkg/configuration/database"
	"github.com/barchart/common-go/pkg/validation"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var stage string
var validate = validation.GetValidator()

var config Config

func init() {
	config = *newConfig()
}

func newConfig() *Config {
	return &Config{
		Databases:      nil,
		AWS:            nil,
		CustomSettings: map[string]interface{}{},
		Stage:          "",
	}
}

// InitConfigFromFile reads configuration from file and return config
func InitConfigFromFile(path string, name string) (*Config, error) {
	return config.initConfigFromFile(path, name)
}

func (cfg *Config) initConfigFromFile(path string, name string) (*Config, error) {
	vip := viper.New()
	vip.AddConfigPath(path)
	vip.SetConfigName(name)

	stage = os.Getenv("APP_ENV")

	if stage == "" {
		stage = "dev"
	}

	err := vip.ReadInConfig()

	if err != nil {
		errText := fmt.Sprintf("unable to read config file, %v", err)
		return nil, errors.New(errText)
	}

	err = vip.UnmarshalKey(stage, &cfg)

	if err != nil {
		errText := fmt.Sprintf("unable to decode into config struct, %v", err)
		return nil, errors.New(errText)
	}

	cfg.Stage = stage

	return cfg, nil
}

// region Singleton Getters

// GetCustomSettings returns the Custom Settings
func GetCustomSettingsByKey(key string) (interface{}, error) {
	return config.getCustomSettingsByKey(key)
}

// GetDB returns Database configuration by key
func GetDB(key string) (Database, error) {
	return config.getDB(key)
}

// GetDynamo returns the Dynamo configuration by key
func GetDynamo(key string) (Dynamo, error) {
	return config.getDynamo(key)
}

// GetS3 returns the S3 configuration by key
func GetS3(key string) (S3, error) {
	return config.getS3(key)
}

// GetSes returns the SES configuration by key
func GetSES(key string) (SES, error) {
	return config.getSes(key)
}

// GetSNS returns the SNS configuration by key
func GetSNS(key string) (SNS, error) {
	return config.getSNS(key)
}

// GetSQS returns the SQS configuration by key
func GetSQS(key string) (SQS, error) {
	return config.getSQS(key)
}

// GetSecretsManager returns SecretManager configuration
func GetSecretsManager() (secretsmanager.SecretsManager, error) {
	return config.getSecretsManager()
}

// GetStage returns current stage
func GetStage() string {
	return config.getStage()
}

// endregion Getters

// region Singleton setters

// SetCustomSettings sets the Custom Setting
func SetCustomSettings(key string, cs interface{}) {
	config.setCustomSettings(key, cs)
}

// SetDatabaseProperties sets the Database configuration by providing parameters
func SetDatabaseProperties(key string, provider string, host string, port int, database string, user string, password string) error {
	return config.setDatabaseProperties(key, provider, host, port, database, user, password)
}

// SetDatabaseObject sets the Database configuration
func SetDatabaseObject(key string, database Database) error {
	return config.setDatabaseObject(key, database)
}

// SetDatabase sets the Database configuration
func SetDatabase(key string, database Database) error {
	return SetDatabaseObject(key, database)
}

// SetDynamo sets the Dynamo configuration
func SetDynamo(key string, region string, prefix string) error {
	return config.setDynamo(key, region, prefix)
}

// SetS3 sets the S3 configuration
func SetS3(key string, region string, bucket string) error {
	return config.setS3(key, region, bucket)
}

// SetSES sets the SES configuration
func SetSES(key string, region string, from string, domain string) error {
	return config.setSES(key, region, from, domain)
}

// SetSNS sets the SNS configuration
func SetSNS(key string, region string, topic string, prefix string) error {
	return config.setSNS(key, region, topic, prefix)
}

// SetSQS sets the SQS configuration
func SetSQS(key string, region string, prefix string, queue string) error {
	return config.setSQS(key, region, prefix, queue)
}

// SetSecretsManager creates a Secrets Manager instance and sets it into the instance of the configuration
func SetSecretsManager(region string) {
	config.setSecretsManager(region)
}

// SetStage sets the current stage
func SetStage(stage string) {
	config.setStage(stage)
}

// endregion Setters

// region Getters

func (cfg Config) getCustomSettingsByKey(key string) (interface{}, error) {
	if cfg.CustomSettings == nil {
		cfg.CustomSettings = map[string]interface{}{}
	}

	cs, ok := cfg.CustomSettings[key]

	if !ok {
		return nil, errors.New("custom settings [ " + key + " ] not found")
	}

	return cs, nil
}

func (cfg Config) getDB(key string) (Database, error) {
	if cfg.Databases == nil {
		cfg.Databases = Databases{}
	}

	Databases := cfg.Databases

	if db, ok := Databases[key]; ok {
		return db, nil
	} else {
		err := fmt.Sprintf("database [ %v ] configuration not found", key)
		return Database{}, errors.New(err)
	}
}

func (cfg Config) getDynamo(key string) (Dynamo, error) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.Dynamo == nil {
		cfg.AWS.Dynamo = &map[string]Dynamo{}
	}

	dynamoList := *cfg.AWS.Dynamo

	if dynamo, ok := dynamoList[key]; ok {
		return dynamo, nil
	} else {
		err := fmt.Sprintf("AWS dynamo [ %v ] configuration not found", key)
		return Dynamo{}, errors.New(err)
	}
}

func (cfg Config) getS3(key string) (S3, error) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.S3 == nil {
		cfg.AWS.S3 = &map[string]S3{}
	}

	s3List := *cfg.AWS.S3

	if s3, ok := s3List[key]; ok {
		return s3, nil
	} else {
		err := fmt.Sprintf("AWS S3 [ %v ] configuration not found", key)
		return S3{}, errors.New(err)
	}
}

func (cfg Config) getSes(key string) (SES, error) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SES == nil {
		cfg.AWS.SES = &map[string]SES{}
	}

	sesList := *cfg.AWS.SES

	if ses, ok := sesList[key]; ok {
		return ses, nil
	} else {
		err := fmt.Sprintf("AWS SES [ %v ] configuration not found", key)
		return SES{}, errors.New(err)
	}
}

func (cfg Config) getSNS(key string) (SNS, error) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SNS == nil {
		cfg.AWS.SNS = &map[string]SNS{}
	}

	snsList := *cfg.AWS.SNS

	if sns, ok := snsList[key]; ok {
		return sns, nil
	} else {
		err := fmt.Sprintf("AWS SNS [ %v ] configuration not found", key)
		return SNS{}, errors.New(err)
	}
}

func (cfg Config) getSQS(key string) (SQS, error) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SQS == nil {
		cfg.AWS.SQS = &map[string]SQS{}
	}

	sqsList := *cfg.AWS.SQS

	if sqs, ok := sqsList[key]; ok {
		return sqs, nil
	} else {
		err := fmt.Sprintf("AWS SQS [ %v ] configuration not found", key)
		return SQS{}, errors.New(err)
	}
}

func (cfg *Config) getSecretsManager() (secretsmanager.SecretsManager, error) {
	if cfg.AWS == nil || cfg.AWS.SecretsManager == nil {
		return secretsmanager.SecretsManager{}, errors.New("secrets manager configuration hasn't been set")
	}

	return *cfg.AWS.SecretsManager, nil
}

func (cfg Config) getStage() string {
	return cfg.Stage
}

// endregion Getters

// region Setters

func (cfg *Config) setCustomSettings(key string, cs interface{}) {
	if cfg.CustomSettings == nil {
		cfg.CustomSettings = map[string]interface{}{}
	}

	cfg.CustomSettings[key] = cs
}

func (cfg *Config) setDatabaseProperties(key string, provider string, host string, port int, database string, user string, password string) error {
	if cfg.Databases == nil {
		cfg.Databases = Databases{}
	}

	db := Database{
		Provider: provider,
		Host:     host,
		Port:     port,
		Database: database,
		Username: user,
		Password: password,
	}

	err := validate.Struct(db)
	if err != nil {
		return err
	}

	Databases := cfg.Databases

	Databases[key] = db

	return nil
}

func (cfg *Config) setDatabaseObject(key string, database Database) error {
	if cfg.Databases == nil {
		cfg.Databases = Databases{}
	}

	err := validate.Struct(database)
	if err != nil {
		return err
	}

	Databases := cfg.Databases

	Databases[key] = database

	return nil
}

func (cfg *Config) setDynamo(key string, region string, prefix string) error {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.Dynamo == nil {
		cfg.AWS.Dynamo = &map[string]Dynamo{}
	}

	dynamo := Dynamo{
		Prefix: prefix,
		Region: region,
	}

	err := validate.Struct(dynamo)

	if err != nil {
		return err
	}

	dynamoList := *cfg.AWS.Dynamo

	dynamoList[key] = dynamo

	return nil
}

func (cfg *Config) setS3(key string, region string, bucket string) error {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.S3 == nil {
		cfg.AWS.S3 = &map[string]S3{}
	}

	s3 := S3{
		Region: region,
		Bucket: bucket,
	}

	err := validate.Struct(s3)

	if err != nil {
		return err
	}

	s3List := *cfg.AWS.S3

	s3List[key] = s3

	return nil
}

func (cfg *Config) setSES(key string, region string, from string, domain string) error {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SES == nil {
		cfg.AWS.SES = &map[string]SES{}
	}

	ses := SES{
		From:   from,
		Region: region,
		Domain: domain,
	}

	err := validate.Struct(ses)

	if err != nil {
		return err
	}

	sesList := *cfg.AWS.SES

	sesList[key] = ses

	return nil
}

func (cfg *Config) setSNS(key string, region string, topic string, prefix string) error {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SNS == nil {
		cfg.AWS.SNS = &map[string]SNS{}
	}

	sns := SNS{
		Region: region,
		Topic:  topic,
		Prefix: prefix,
	}

	err := validate.Struct(sns)

	if err != nil {
		return err
	}

	snsList := *cfg.AWS.SNS

	snsList[key] = sns

	return nil
}

func (cfg *Config) setSQS(key string, region string, prefix string, queue string) error {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	if cfg.AWS.SQS == nil {
		cfg.AWS.SQS = &map[string]SQS{}
	}

	sqs := SQS{
		Prefix: prefix,
		Region: region,
		Queue:  queue,
	}

	err := validate.Struct(sqs)

	if err != nil {
		return err
	}

	sqsList := *cfg.AWS.SQS

	sqsList[key] = sqs

	return nil
}

func (cfg *Config) setSecretsManager(region string) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	cfg.AWS.SecretsManager = secretsmanager.New(region)
}

func (cfg *Config) setStage(stage string) {
	cfg.Stage = stage
}

// endregion Setters
