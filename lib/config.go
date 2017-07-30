package sqscheck

// CheckSpec definition
type CheckSpec struct {
	Queues             []string `json:"queues"`
	WarningThreashold  int      `json:"warning_threashold"`
	CriticalThreashold int      `json:"critical_threashold"`
	ContactGroup       []string `json:"contact_group"`
}

// Config definition
type Config struct {
	AwsAccessID      string      `json:"aws_access_id"`
	AwsAccessSecret  string      `json:"aws_access_secret"`
	AwsRegion        string      `json:"aws_region"`
	AwsAccountNum    string      `json:"aws_account_num"`
	Queues           []string    `json:"queues"`
	DefaultCheckSpec CheckSpec   `json:"default_check_spec"`
	CustomChecks     []CheckSpec `json:"custom_checks"`
}
