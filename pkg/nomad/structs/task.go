package structs

type Task struct {
	Name            string                 `hcl:",key"`
	Artifacts       []*Artifact            `hcl:"artifact" hcle:"omitempty"`
	Config          map[string]interface{} `hcl:"config" hcle:"omitempty"`
	Constraints     []*Constraint          `hcl:"constraint" hcle:"omitempty"`
	Affinities      []*Affinity            `hcl:"affinity" hcle:"omitempty"`
	Devices         []*Device              `hcl:"device" hcle:"omitempty"`
	DispatchPayload *DispatchPayload       `hcl:"dispatch_payload" hcle:"omitempty"`
	Driver          string                 `hcl:"driver"`
	Env             map[string]string      `hcl:"env" hcle:"omitempty"`
	KillTimeout     string                 `hcl:"kill_timeout" hcle:"omitempty"`
	KillSignal      string                 `hcl:"kill_signal" hcle:"omitempty"`
	Leader          bool                   `hcl:"leader" hcle:"omitempty"`
	Logs            *LogConfig             `hcl:"logs" hcle:"omitempty"`
	Meta            map[string]string      `hcl:"meta" hcle:"omitempty"`
	Resources       *Resources             `hcl:"resources" hcle:"omitempty"`
	Services        []*Service             `hcl:"service" hcle:"omitempty"`
	ShutdownDelay   string                 `hcl:"shutdown_delay" hcle:"omitempty"`
	User            string                 `hcl:"user" hcle:"omitempty"`
	Templates       []*Template            `hcl:"template" hcle:"omitempty"`
	Vault           *Vault                 `hcl:"vault" hcle:"omitempty"`
}

type DispatchPayload struct {
	File string `hcl:"file"`
}

type Artifact struct {
	Destination string            `hcl:"destination" hcle:"omitempty"`
	Mode        string            `hcl:"mode" hcle:"omitempty"`
	Options     map[string]string `hcl:"options" hcle:"omitempty"`
	Source      string            `hcl:"source" hcle:"omitempty"`
}

type Device struct {
	ID string `hcl:",key"`
}

type LogConfig struct {
	MaxFiles      int `hcl:"max_files" hcle:"omitempty"`
	MaxFileSizeMB int `hcl:"max_file_size" hcle:"omitempty"`
}
type Template struct {
	Data           string `hcl:"data" hcle:"omitempty"`
	Destination    string `hcl:"destination" hcle:"omitempty"`
	ChangeMode     string `hcl:"change_mode" hcle:"omitempty"`
	ChangeSignal   string `hcl:"change_signal" hcle:"omitempty"`
	Env            bool   `hcl:"env" hcle:"omitempty"`
	Source         string `hcl:"source" hcle:"omitempty"`
	Splay          string `hcl:"splay" hcle:"omitempty"`
	VaultGrace     string `hcl:"vault_grace" hcle:"omitempty"`
	Perms          string `hcl:"perms" hcle:"omitempty"`
	RightDelimiter string `hcl:"right_delimiter" hcle:"omitempty"`
	LeftDelimiter  string `hcl:"left_delimiter" hcle:"omitempty"`
}

type Port struct {
	Label  string `hcl:",key"`
	Static int    `hcl:"static" hcle:"omitempty"`
}

type NetworkResource struct {
	MBits int     `hcl:"mbits" hcle:"omitempty"`
	Ports []*Port `hcl:"port" hcle:"omitempty"`
}

type Resources struct {
	CPU      int                `hcl:"cpu" hcle:"omitempty"`
	MemoryMB int                `hcl:"memory" hcle:"omitempty"`
	IOPS     int                `hcl:"iops" hcle:"omitempty"`
	Networks []*NetworkResource `hcl:"network" hcle:"omitempty"`
}
