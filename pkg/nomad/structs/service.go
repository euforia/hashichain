package structs

// Service holds a nomad service block
type Service struct {
	Name        string          `hcl:"name"`
	PortLabel   string          `hcl:"port" hcle:"omitempty"`
	Tags        []string        `hcl:"tags" hcle:"omitempty"`
	CanaryTags  []string        `hcl:"canary_tags" hcle:"omitempty"`
	AddressMode string          `hcl:"address_mode" hcle:"omitempty"`
	Checks      []*ServiceCheck `hcl:"check" hcle:"omitempty"`
}

// CheckRestart holds a nomad check_restart block
type CheckRestart struct {
	Limit          int    `hcl:"limit" hcle:"omitempty"`
	Grace          string `hcl:"grace" hcle:"omitempty"`
	IgnoreWarnings bool   `hcl:"ignore_warnings" hcle:"omitempty"`
}

// ServiceCheck holds a service check block
type ServiceCheck struct {
	AddressMode   string              `hcl:"address_mode" hcle:"omitempty"`
	Args          []string            `hcl:"args" hcle:"omitempty"`
	CheckRestart  *CheckRestart       `hcl:"check_restart" hcle:"omitempty"`
	Command       string              `hcl:"command" hcle:"omitempty"`
	GRPCService   string              `hcl:"grpc_service" hcle:"omitempty"`
	GRPCUseTLS    bool                `hcl:"grpc_user_tls" hcle:"omitempty"`
	InitialStatus string              `hcl:"initial_status" hcle:"omitempty"`
	Interval      string              `hcl:"interval" hcle:"omitempty"`
	Method        string              `hcl:"method" hcle:"omitempty"`
	Path          string              `hcl:"path" hcle:"omitempty"`
	Protocol      string              `hcl:"protocol" hcle:"omitempty"`
	Name          string              `hcl:"name" hcle:"omitempty"`
	PortLabel     string              `hcl:"port" hcle:"omitempty"`
	Header        map[string][]string `hcl:"header" hcle:"omitempty"`
	Timeout       string              `hcl:"timeout" hcle:"omitempty"`
	Type          string              `hcl:"type" hcle:"omitempty"`
	TLSSkipVerify bool                `hcl:"tls_skip_verify" hcle:"omitempty"`
}
