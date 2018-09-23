package translate

import (
	"strings"

	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/hashicorp/nomad/api"
)

func translateTask(t *api.Task) *structs.Task {

	task := &structs.Task{
		Name:       t.Name,
		Driver:     t.Driver,
		Leader:     t.Leader,
		Config:     t.Config,
		User:       t.User,
		KillSignal: t.KillSignal,
	}
	if t.ShutdownDelay > 0 {
		task.ShutdownDelay = t.ShutdownDelay.String()
	}
	if t.KillTimeout != nil {
		task.KillTimeout = t.KillTimeout.String()
	}

	task.Artifacts = translateArtifacts(t.Artifacts)
	task.Env = translateMapStringString(t.Env)
	task.Logs = translateLogConfig(t.LogConfig)
	task.Meta = translateMapStringString(t.Meta)
	task.Services = translateServices(t.Services)
	task.Resources = translateResources(t.Resources)
	task.Templates = translateTemplates(t.Templates)
	task.Vault = translateVault(t.Vault)

	return task
}

func translateChecks(cs []api.ServiceCheck) []*structs.ServiceCheck {
	out := make([]*structs.ServiceCheck, len(cs))
	for i, c := range cs {
		out[i] = translateCheck(c)
	}

	return out
}

func translateServices(ss []*api.Service) []*structs.Service {
	out := make([]*structs.Service, len(ss))
	for i, s := range ss {
		out[i] = translateService(s)
	}
	return out
}

func translateCheck(c api.ServiceCheck) *structs.ServiceCheck {
	return &structs.ServiceCheck{
		Name:          c.Name,
		Type:          c.Type,
		Timeout:       c.Timeout.String(),
		Interval:      c.Interval.String(),
		Command:       c.Command,
		Args:          c.Args,
		AddressMode:   c.AddressMode,
		GRPCService:   c.GRPCService,
		GRPCUseTLS:    c.GRPCUseTLS,
		InitialStatus: c.InitialStatus,
		Method:        c.Method,
		Path:          c.Path,
		Protocol:      c.Protocol,
		PortLabel:     c.PortLabel,
		TLSSkipVerify: c.TLSSkipVerify,
		CheckRestart:  translateCheckRestart(c.CheckRestart),
	}
}

func translateTemplates(ts []*api.Template) []*structs.Template {
	out := make([]*structs.Template, len(ts))
	for i, t := range ts {
		out[i] = translateTemplate(t)
	}
	return out
}

func translateTemplate(t *api.Template) *structs.Template {
	out := &structs.Template{}
	if t.DestPath != nil {
		out.Destination = *t.DestPath
	}
	if t.ChangeMode != nil {
		out.ChangeMode = *t.ChangeMode
	}
	if t.ChangeSignal != nil {
		out.ChangeSignal = *t.ChangeSignal
	}
	if t.EmbeddedTmpl != nil {
		// out.Data = *t.EmbeddedTmpl
		tmo := strings.Replace(*t.EmbeddedTmpl, `"`, `\"`, -1)
		out.Data = strings.Replace(tmo, "\n", "\\n", -1)
		// fmt.Println(out.Data)
	}
	if t.SourcePath != nil {
		out.Source = *t.SourcePath
	}
	if t.Envvars != nil {
		out.Env = *t.Envvars
	}

	return out
}

func translateService(s *api.Service) *structs.Service {
	if s == nil {
		return nil
	}

	out := &structs.Service{
		Name:      s.Name,
		PortLabel: s.PortLabel,
	}

	if s.Tags != nil {
		out.Tags = make([]string, len(s.Tags))
		copy(out.Tags, s.Tags)
	}

	out.Checks = translateChecks(s.Checks)

	return out
}

func translateCheckRestart(cr *api.CheckRestart) *structs.CheckRestart {
	if cr == nil {
		return nil
	}
	return &structs.CheckRestart{
		Limit:          cr.Limit,
		Grace:          cr.Grace.String(),
		IgnoreWarnings: cr.IgnoreWarnings,
	}
}

func translateArtifacts(arts []*api.TaskArtifact) []*structs.Artifact {
	out := make([]*structs.Artifact, len(arts))
	for i, art := range arts {
		out[i] = translateArtifact(art)
	}
	return out
}

func translateArtifact(a *api.TaskArtifact) *structs.Artifact {
	if a == nil {
		return nil
	}

	out := &structs.Artifact{}
	if a.GetterSource != nil {
		out.Source = *a.GetterSource
	}
	if a.GetterOptions != nil {
		out.Options = make(map[string]string)
		for k, v := range a.GetterOptions {
			out.Options[k] = v
		}
	}
	if a.GetterMode != nil {
		out.Mode = *a.GetterMode
	}

	return out
}
