package hcl

import (
	"github.com/euforia/hashichain/pkg/nomad/structs"
	"github.com/hashicorp/nomad/api"
)

func translateGroup(g *api.TaskGroup) *structs.Group {
	group := &structs.Group{
		Name:  *g.Name,
		Count: *g.Count,
		Meta:  translateMapStringString(g.Meta),
		Tasks: make([]*structs.Task, len(g.Tasks)),
	}

	group.Restart = translateRestart(g.RestartPolicy)
	group.EphemeralDisk = translateEphemeralDisk(g.EphemeralDisk)
	group.Constraints = translateConstraints(g.Constraints)
	group.Reschedule = translateReschedule(g.ReschedulePolicy)
	group.EphemeralDisk = translateEphemeralDisk(g.EphemeralDisk)
	group.Update = translateUpdateStrategy(g.Update)

	for i, task := range g.Tasks {
		group.Tasks[i] = translateTask(task)
	}

	return group
}

func translateEphemeralDisk(d *api.EphemeralDisk) *structs.EphemeralDisk {
	if d == nil {
		return nil
	}

	disk := &structs.EphemeralDisk{}
	if d.Sticky != nil {
		disk.Sticky = *d.Sticky
	}
	if d.SizeMB != nil {
		disk.SizeMB = *d.SizeMB
	}
	if d.Migrate != nil {
		disk.Migrate = *d.Migrate
	}

	return disk
}
