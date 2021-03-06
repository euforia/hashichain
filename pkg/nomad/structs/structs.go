// Package structs contains structs to write out a nomad HCL file
package structs

type Spread struct {
	Attribute string          `hcl:"attribute" hcle:"omitempty"`
	Weight    float64         `hcl:"weight" hcle:"omitempty"`
	Targets   []*SpreadTarget `hcl:"target" hcle:"omitempty"`
}

type SpreadTarget struct {
	ID      string  `hcl:",key"`
	Percent float64 `hcl:"percent"`
}

// Vault is nomad vault block
type Vault struct {
	ChangeMode   string `hcl:"change_mode" hcle:"omitempty"`
	ChangeSignal string `hcl:"change_signal" hcle:"omitempty"`
	// Specifies if the VAULT_TOKEN environment variable should be set when starting the task.
	Env      bool     `hcl:"env" hcle:"omitempty"`
	Policies []string `hcl:"policies" hcle:"omitempty"`
}

// ReschedulePolicy is a nomad group ReschedulePolicy
type ReschedulePolicy struct {
	Attempts      int    `hcl:"attempts" hcle:"omitempty"`
	Interval      string `hcl:"interval" hcle:"omitempty"`
	Delay         string `hcl:"delay" hcle:"omitempty"`
	MaxDelay      string `hcl:"max_delay" hcle:"omitempty"`
	DelayFunction string `hcl:"delay_function" hcle:"omitempty"`
	Unlimited     bool   `hcl:"unlimited" hcle:"omitempty"`
}

// RestartPolicy is a nomad restart policy
type RestartPolicy struct {
	Attempts int    `hcl:"attempts"`
	Interval string `hcl:"interval"`
	Delay    string `hcl:"delay"`
	Mode     string `hcl:"mode"`
}

// MigrateStrategy holds the nomad migrate block
type MigrateStrategy struct {
	MaxParallel     int    `hcl:"max_parallel" hcle:"omitempty"`
	HealthCheck     string `hcl:"health_check" hcle:"omitempty"`
	MinHealthyTime  string `hcl:"min_healthy_time" hcle:"omitempty"`
	HealthyDeadline string `hcl:"healthy_deadline" hcle:"omitempty"`
}

// UpdateStrategy holds the nomad update block
type UpdateStrategy struct {
	AutoRevert       bool   `hcl:"auto_revert" hcle:"omitempty"`
	Canary           int    `hcl:"canary" hcle:"omitempty"`
	Stagger          string `hcl:"stagger" hcle:"omitempty"`
	MaxParallel      int    `hcl:"max_parallel" hcle:"omitempty"`
	HealthCheck      string `hcl:"health_check" hcle:"omitempty"`
	MinHealthyTime   string `hcl:"min_healthy_time" hcle:"omitempty"`
	HealthyDeadline  string `hcl:"healthy_deadline" hcle:"omitempty"`
	ProgressDeadline string `hcl:"progress_deadline" hcle:"omitempty"`
}

// Constraint is a nomad constraint
type Constraint struct {
	Attribute string `hcl:"attribute" hcle:"omitempty"`
	Value     string `hcl:"value" hcle:"omitempty"`
	Operator  string `hcl:"operator" hcle:"omitempty"`
}

type Affinity struct {
	Attribute string  `hcl:"attribute" hcle:"omitempty"`
	Value     string  `hcl:"value" hcle:"omitempty"`
	Operator  string  `hcl:"operator" hcle:"omitempty"`
	Weight    float64 `hcl:"weight" hcle:"omitempty"`
}

// EphemeralDisk is a nomad task EphemeralDisk
type EphemeralDisk struct {
	Sticky  bool `hcl:"sticky" hcle:"omitempty"`
	Migrate bool `hcl:"migrate" hcle:"omitempty"`
	SizeMB  int  `hcl:"size"`
}

// Group is a nomad TaskGroup
type Group struct {
	Affinities    []*Affinity       `hcl:"affinity" hcle:"omitempty"`
	Count         int               `hcl:"count" hcle:"omitempty"`
	Constraints   []*Constraint     `hcl:"constraint" hcle:"omitempty"`
	EphemeralDisk *EphemeralDisk    `hcl:"ephemeral_disk" hcle:"omitempty"`
	Meta          map[string]string `hcl:"meta" hcle:"omitempty"`
	Migrate       *MigrateStrategy  `hcl:"migrate" hcle:"omitempty"`
	Name          string            `hcl:",key"`
	Restart       *RestartPolicy    `hcl:"restart" hcle:"omitempty"`
	Reschedule    *ReschedulePolicy `hcl:"reschedule" hcle:"omitempty"`
	Spread        *Spread           `hcl:"spread" hcle:"omitempty"`
	Tasks         []*Task           `hcl:"task"`
	Update        *UpdateStrategy   `hcl:"update" hcle:"omitempty"`
	Vault         *Vault            `hcl:"vault" hcle:"omitempty"`
}

// Job is the actual nomad job
type Job struct {
	Name        string            `hcle:"omit"`
	Region      string            `hcl:"region"`
	Datacenters []string          `hcl:"datacenters"`
	Type        string            `hcl:"type"`
	AllAtOnce   bool              `hcl:"all_at_once" hcle:"omitempty"`
	Meta        map[string]string `hcl:"meta" hcle:"omitempty"`
	Update      *UpdateStrategy   `hcl:"update" hcle:"omitempty"`
	Migrate     *MigrateStrategy  `hcl:"migrate" hcle:"omitempty"`
	Constraints []*Constraint     `hcl:"constraint" hcle:"omitempty"`
	Affinities  []*Affinity       `hcl:"affinity" hcle:"omitempty"`
	Spread      *Spread           `hcl:"spread" hcle:"omitempty"`
	Groups      []*Group          `hcl:"group"`
}
