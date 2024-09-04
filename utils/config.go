package utils

var AppConfig *Config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	Database          PgDatabase `yaml:"database"`
	DatabaseWarehouse PgDatabase `yaml:"database-warehouse"`
	Auth              struct {
		JwtSecretKey string `yaml:"jwt-secret-key"`
	} `yaml:"auth"`
	AI struct {
		ChatGpt struct {
			ApiKey string `yaml:"api-key"`
		} `yaml:"chat-gpt"`
	} `yaml:"ai"`
	Kafka struct {
		Producer KafkaProducerConfig `yaml:"producer"`
		Consumer KafkaConsumerConfig `yaml:"consumer"`
	} `yaml:"kafka"`
	Workers []Worker `yaml:"workers"`
}

type PgDatabase struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DbName      string `yaml:"dbname"`
	SslMode     string `yaml:"sslmode"`
	MaxOpenConn int    `yaml:"max-open-conn"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
}

type KafkaProducerConfig struct {
	BootstrapServer  string     `yaml:"bootstrap-servers"`
	SecurityProtocol string     `yaml:"security-protocol"`
	SaslMechanisms   string     `yaml:"sasl-mechanisms"`
	SaslUsername     string     `yaml:"sasl-username"`
	SaslPassword     string     `yaml:"sasl-password"`
	Acks             string     `yaml:"acks"`
	FooTopic         KafkaTopic `yaml:"topics-foo1"`
}

type KafkaConsumerConfig struct {
	BootstrapServer  string     `yaml:"bootstrap-servers"`
	SecurityProtocol string     `yaml:"security-protocol"`
	SaslMechanisms   string     `yaml:"sasl-mechanisms"`
	SaslUsername     string     `yaml:"sasl-username"`
	SaslPassword     string     `yaml:"sasl-password"`
	GroupId          string     `yaml:"group_id"`
	FooTopic         KafkaTopic `yaml:"topics-foo1"`
}

type KafkaTopic struct {
	Name      string `yaml:"name"`
	Partition int    `yaml:"partition"`
}

type Worker struct {
	Name     string `yaml:"name"`
	TimeRun  string `yaml:"time_run"`
	Function string `yaml:"function"`
	Enable   bool   `yaml:"enable"`
}
