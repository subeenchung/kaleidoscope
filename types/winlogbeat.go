package main

import "fmt"

type Winlogbeat struct {
	Event struct {
		Action  string    `json:"action"`
		Kind    string    `json:"kind"`
		Code    int       `json:"code"`
		Created time.Time `json:"created"`
	} `json:"event"`
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Timestamp time.Time `json:"@timestamp"`
	Agent     struct {
		EphemeralID string `json:"ephemeral_id"`
		Hostname    string `json:"hostname"`
		Type        string `json:"type"`
		ID          string `json:"id"`
		Version     string `json:"version"`
	} `json:"agent"`
	Kafka struct {
		Topic         string      `json:"topic"`
		Partition     int         `json:"partition"`
		Key           interface{} `json:"key"`
		ConsumerGroup string      `json:"consumer_group"`
		Offset        int         `json:"offset"`
	} `json:"kafka"`
	Host struct {
		Architecture string `json:"architecture"`
		Hostname     string `json:"hostname"`
		Os           struct {
			Kernel   string `json:"kernel"`
			Name     string `json:"name"`
			Family   string `json:"family"`
			Platform string `json:"platform"`
			Version  string `json:"version"`
			Build    string `json:"build"`
		} `json:"os"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"host"`
	Ecs struct {
		Version string `json:"version"`
	} `json:"ecs"`
	Message string `json:"message"`
	Version string `json:"@version"`
	Winlog  struct {
		Task    string `json:"task"`
		API     string `json:"api"`
		Process struct {
			Pid    int `json:"pid"`
			Thread struct {
				ID int `json:"id"`
			} `json:"thread"`
		} `json:"process"`
		EventID      int      `json:"event_id"`
		RecordID     int      `json:"record_id"`
		Version      int      `json:"version"`
		Opcode       string   `json:"opcode"`
		ComputerName string   `json:"computer_name"`
		Keywords     []string `json:"keywords"`
		ProviderName string   `json:"provider_name"`
		ProviderGUID string   `json:"provider_guid"`
		EventData    struct {
			ProcessID          string `json:"ProcessId"`
			NewProcessID       string `json:"NewProcessId"`
			TokenElevationType string `json:"TokenElevationType"`
			NewProcessName     string `json:"NewProcessName"`
			SubjectUserName    string `json:"SubjectUserName"`
			SubjectDomainName  string `json:"SubjectDomainName"`
			SubjectUserSid     string `json:"SubjectUserSid"`
			SubjectLogonID     string `json:"SubjectLogonId"`
		} `json:"event_data"`
		Channel string `json:"channel"`
	} `json:"winlog"`
}
