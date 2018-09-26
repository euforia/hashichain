// Package hcl translates a nomad job to hcl encodable nomad job
package hcl

import (
	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/hashicorp/nomad/api"
)

// NewJob translates a native nomad job into one that can be written
// out as HCL
func NewJob(job *api.Job) *structs.Job {
	js := &structs.Job{
		Name:        *job.Name,
		Region:      *job.Region,
		Datacenters: job.Datacenters,
		Type:        *job.Type,
		Constraints: make([]*structs.Constraint, len(job.Constraints)),
		Groups:      make([]*structs.Group, len(job.TaskGroups)),
	}
	if job.AllAtOnce != nil {
		js.AllAtOnce = *job.AllAtOnce
	}

	js.Meta = translateMapStringString(job.Meta)
	js.Constraints = translateConstraints(job.Constraints)
	js.Update = translateUpdateStrategy(job.Update)
	js.Migrate = translateMigrateStrategy(job.Migrate)

	for i, g := range job.TaskGroups {
		js.Groups[i] = translateGroup(g)
	}

	return js
}

func translateNetworkResource(n *api.NetworkResource) *structs.NetworkResource {
	o := &structs.NetworkResource{}
	if n.MBits != nil {
		o.MBits = *n.MBits
	}

	ports := make([]*structs.Port, 0, len(n.DynamicPorts)+len(n.ReservedPorts))

	if len(n.DynamicPorts) > 0 {
		for _, p := range n.DynamicPorts {
			ports = append(ports, &structs.Port{Label: p.Label})
		}
	}

	if len(n.ReservedPorts) > 0 {
		for _, p := range n.ReservedPorts {
			ports = append(ports, &structs.Port{Label: p.Label, Static: p.Value})
		}
	}

	o.Ports = ports

	return o
}

func translateNetworkResources(nn []*api.NetworkResource) []*structs.NetworkResource {
	out := make([]*structs.NetworkResource, len(nn))
	for i, n := range nn {
		out[i] = translateNetworkResource(n)
	}
	return out
}

func translateResources(nr *api.Resources) *structs.Resources {
	if nr == nil {
		return nil
	}

	r := &structs.Resources{}
	if nr.CPU != nil {
		r.CPU = *nr.CPU
	}
	if nr.MemoryMB != nil {
		r.MemoryMB = *nr.MemoryMB
	}

	r.Networks = translateNetworkResources(nr.Networks)

	return r
}

func translateReschedule(nr *api.ReschedulePolicy) *structs.ReschedulePolicy {
	if nr == nil {
		return nil
	}

	r := &structs.ReschedulePolicy{}
	if nr.Attempts != nil {
		r.Attempts = *nr.Attempts
	}
	if nr.Interval != nil {
		r.Interval = nr.Interval.String()
	}
	if nr.Delay != nil {
		r.Delay = nr.Delay.String()
	}
	if nr.MaxDelay != nil {
		r.MaxDelay = nr.MaxDelay.String()
	}
	if nr.DelayFunction != nil {
		r.DelayFunction = *nr.DelayFunction
	}
	if nr.Unlimited != nil {
		r.Unlimited = *nr.Unlimited
	}

	return r
}

func translateConstraints(cs []*api.Constraint) []*structs.Constraint {
	if cs == nil {
		return nil
	}
	out := make([]*structs.Constraint, len(cs))
	for i, c := range cs {
		out[i] = translateConstraint(c)
	}
	return out
}

func translateRestart(nr *api.RestartPolicy) *structs.RestartPolicy {
	if nr == nil {
		return nil
	}

	r := &structs.RestartPolicy{}
	if nr.Attempts != nil {
		r.Attempts = *nr.Attempts
	}
	if nr.Interval != nil {
		r.Interval = nr.Interval.String()
	}
	if nr.Delay != nil {
		r.Delay = nr.Delay.String()
	}
	if nr.Mode != nil {
		r.Mode = *nr.Mode
	}

	return r
}

func translateLogConfig(lc *api.LogConfig) *structs.LogConfig {
	if lc == nil {
		return nil
	}

	out := &structs.LogConfig{}
	if lc.MaxFiles != nil {
		out.MaxFiles = *lc.MaxFiles
	}
	if lc.MaxFileSizeMB != nil {
		out.MaxFileSizeMB = *lc.MaxFileSizeMB
	}

	return out
}

func translateVault(v *api.Vault) *structs.Vault {
	if v == nil {
		return nil
	}

	out := &structs.Vault{}
	if v.ChangeMode != nil {
		out.ChangeMode = *v.ChangeMode
	}
	if v.ChangeSignal != nil {
		out.ChangeSignal = *v.ChangeSignal
	}
	if v.Env != nil {
		out.Env = *v.Env
	}
	if v.Policies != nil {
		out.Policies = v.Policies
	}

	return out
}

func translateConstraint(c *api.Constraint) *structs.Constraint {
	nc := &structs.Constraint{
		Attribute: c.LTarget,
		Value:     c.RTarget,
	}
	if c.Operand != "=" {
		nc.Operator = c.Operand
	}
	return nc
}

func translateMigrateStrategy(in *api.MigrateStrategy) *structs.MigrateStrategy {
	if in == nil {
		return nil
	}

	us := &structs.MigrateStrategy{}
	if in.MaxParallel != nil {
		us.MaxParallel = *in.MaxParallel
	}
	if in.HealthCheck != nil {
		us.HealthCheck = *in.HealthCheck
	}
	if in.MinHealthyTime != nil {
		us.MinHealthyTime = in.MinHealthyTime.String()
	}

	if in.HealthyDeadline != nil {
		us.HealthyDeadline = in.HealthyDeadline.String()
	}

	return us
}

func translateUpdateStrategy(u *api.UpdateStrategy) *structs.UpdateStrategy {
	if u == nil {
		return nil
	}

	us := &structs.UpdateStrategy{}
	if u.Stagger != nil {
		us.Stagger = u.Stagger.String()
	}
	if u.AutoRevert != nil {
		us.AutoRevert = *u.AutoRevert
	}
	if u.Canary != nil {
		us.Canary = *u.Canary
	}
	if u.MaxParallel != nil {
		us.MaxParallel = *u.MaxParallel
	}
	if u.HealthCheck != nil {
		us.HealthCheck = *u.HealthCheck
	}
	if u.MinHealthyTime != nil {
		us.MinHealthyTime = u.MinHealthyTime.String()
	}
	if u.ProgressDeadline != nil {
		us.ProgressDeadline = u.ProgressDeadline.String()
	}
	if u.HealthyDeadline != nil {
		us.HealthyDeadline = u.HealthyDeadline.String()
	}

	return us
}

func translateMapStringString(m map[string]string) map[string]string {
	if m == nil {
		return m
	}

	meta := make(map[string]string)
	for k, v := range m {
		meta[k] = v
	}
	return meta
}
