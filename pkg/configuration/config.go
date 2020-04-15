package configuration

import (
	"errors"
	"fmt"
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
	"os"
)

var config Config
var stage string
var validate = validation.GetValidator()

func init() {
	config = Config{
		CustomSettings: map[string]interface{}{},
	}
}

// InitConfigFromFile reads configuration from file and return config
func InitConfigFromFile(path string, name string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	stage = os.Getenv("APP_ENV")

	if stage == "" {
		stage = "dev"
	}

	err := viper.ReadInConfig()

	if err != nil {
		errText := fmt.Sprintf("unable to read config file, %v", err)
		return nil, errors.New(errText)
	}

	err = viper.UnmarshalKey(stage, &config)

	if err != nil {
		errText := fmt.Sprintf("unable to decode into config struct, %v", err)
		return nil, errors.New(errText)
	}

	config.Stage = stage

	return &config, nil
}

// region Getters

// GetConfig returns instance of the Config
func GetConfig() *Config {
	return &config
}

// GetCustomSettings returns the Custom Settings
func (cfg Config) GetCustomSettingsByKey(key string) (interface{}, error) {
	if cfg.CustomSettings == nil {
		cfg.CustomSettings = map[string]interface{}{}
	}

	cs, ok := cfg.CustomSettings[key]

	if !ok {
		return nil, errors.New("custom settings [ " + key + " ] not found")
	}

	return cs, nil
}

// GetDB returns Database configuration by key
func (cfg Config) GetDB(key string) (Database, error) {
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

// GetDynamo returns the Dynamo configuration by key
func (cfg Config) GetDynamo(key string) (Dynamo, error) {
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

// GetS3 returns the S3 configuration by key
func (cfg Config) GetS3(key string) (S3, error) {
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

// GetSES returns the SES configuration by key
func (cfg Config) GetSES(key string) (SES, error) {
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

// GetSNS returns the SNS configuration by key
func (cfg Config) GetSNS(key string) (SNS, error) {
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

// GetSQS returns the SQS configuration by key
func (cfg Config) GetSQS(key string) (SQS, error) {
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

// GetSecretsManager returns SecretManager configuration
func (cfg *Config) GetSecretsManager() (secretsmanager.SecretsManager, error) {
	if cfg.AWS == nil || cfg.AWS.SecretsManager == nil {
		return secretsmanager.SecretsManager{}, errors.New("secrets manager configuration hasn't been set")
	}

	return *cfg.AWS.SecretsManager, nil
}

// GetStage returns current stage
func (cfg Config) GetStage() string {
	return cfg.Stage
}

// endregion Getters

// region Setters

// SetCustomSettings sets the Custom Setting
func (cfg *Config) SetCustomSettings(key string, cs interface{}) {
	if cfg.CustomSettings == nil {
		cfg.CustomSettings = map[string]interface{}{}
	}

	cfg.CustomSettings[key] = cs
}

// SetDB sets the Database configuration
func (cfg *Config) SetDB(key string, provider string, host string, port int, database string, user string, password string) error {
	if cfg.Databases == nil {
		cfg.Databases = Databases{}
	}

	db := Database{
		Provider: provider,
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
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

// SetDynamo sets the Dynamo configuration
func (cfg *Config) SetDynamo(key string, region string, prefix string) error {
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

// SetS3 sets the S3 configuration
func (cfg *Config) SetS3(key string, region string, bucket string) error {
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

// SetSES sets the SES configuration
func (cfg *Config) SetSES(key string, region string, from string, domain string) error {
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

// SetSNS sets the SNS configuration
func (cfg *Config) SetSNS(key string, region string, topic string, prefix string) error {
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

// SetSQS sets the SQS configuration
func (cfg *Config) SetSQS(key string, region string, prefix string, queue string) error {
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

// SetSecretsManager creates a Secrets Manager instance and sets it into the instance of the configuration
func (cfg *Config) SetSecretsManager(region string) {
	if cfg.AWS == nil {
		cfg.AWS = &AWS{}
	}

	cfg.AWS.SecretsManager = secretsmanager.New(region)
}

// SetStage sets the current stage
func (cfg *Config) SetStage(stage string) {
	cfg.Stage = stage
}

// endregion Setters
