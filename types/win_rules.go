package types

type WindowsRuleFormat struct {
	RuleNum  int    `json:"rule_num" toml:"rule_num"`
	EventID  int    `json:"event_id" toml:"event_id"`
	Name     string `json:"name" toml:"name"`
	Desc     string `json:"desc" toml:"desc"`
	LogLevel string `json:"log_level" toml:"log_level"`
	LogType  string `json:"log_type" toml:"log_type"`
}
