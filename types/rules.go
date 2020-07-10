package types

type RuleType string

const (
	WindowsRule RuleType = "WindowsRule"
	LinuxRule   RuleType = "LinuxRule"
	NetflowRule RuleType = "NetflowRule"
)

type Ruleset struct {
	Type   RuleType
	reload bool
	Rules  []interface{}
}
