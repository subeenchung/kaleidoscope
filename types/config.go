package types

type Config struct {
	General GeneralConfig `toml:"general" yaml:"general"`
	Kafka KafkaConfig `toml:"kafka" yaml:"kafka"`
	Smtp SmtpConfig `toml:"smtp" yaml:"smtp"`
	Alerts AlertConfig `toml:"alerts" yaml:"alerts"`
}

type GeneralConfig struct {
        RulesPath string `toml:"rules_path" yaml:"rules_path" json:"rules_path"`
        LogPath string `toml: "log_path" yaml:"log_path" json:"log_path"`
}

type KafkaConfig struct {
	Brokers string `toml:"brokers" yaml:"brokers" json:"brokers"`
	GroupId string `toml:"group_id" yaml:"group_id" json:"group_id"`
	AutoOffsetReset string `toml:"auto_offset_reset yaml:"auto_offset_reset" json:"auto_offset_reset"`
	//Timeout int `toml:"timeout" yaml:"timeout" json:"timeout"`
}

type SmtpConfig struct {
	SmtpServer1 string `toml:"smtp_server1" yaml:"smtp_server1"`
	SmtpServer2 string `toml:"smtp_server2" yaml:"smtp_server2"`
}

type AlertConfig struct {
	AlertTimelimit string `toml:"alert_timelimit" yaml:"alert_timelimit"`
}
