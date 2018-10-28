package compose

import (
	"fmt"

	"github.com/docker/cli/cli/compose/types"
	"github.com/euforia/cchain/pkg/compose"
	"github.com/hashicorp/nomad/api"
)

// NewJob converts the compose config to a nomad job spec
func NewJob(workdir string, env map[string]string, filelist ...string) (*api.Job, error) {
	c, err := compose.NewCompose(workdir, env, filelist...)
	if err != nil {
		return nil, err
	}

	return translate(c.Config())
}

func translate(c *types.Config) (*api.Job, error) {
	job := api.NewServiceJob("test", "test", "us-west-2", 50)

	group := api.NewTaskGroup("default", 1)

	for _, service := range c.Services {
		imgName, _, err := splitImageNameTag(service.Image)
		if err != nil {
			return nil, err
		}

		task, err := translateService(service)
		if err != nil {
			return nil, err
		}

		count := 1
		if service.Deploy.Replicas != nil {
			count = int(*service.Deploy.Replicas)
		}

		updateConf := translateDeployUpdateConfig(service.Deploy.UpdateConfig)

		if Datastores.IsSupported(imgName) {
			// Add datastores in there own TaskGroup
			dsGroup := api.NewTaskGroup("datastore", count)
			dsGroup = dsGroup.AddTask(task)
			dsGroup.Update = updateConf
			job = job.AddTaskGroup(dsGroup)
		} else {
			// Add to default group
			group = group.AddTask(task)
			if count > 1 {
				group.Count = &count
			}
			if updateConf != nil {
				group.Update = updateConf
			}
		}

	}

	return job.AddTaskGroup(group), nil
}

func translateDeployUpdateConfig(c *types.UpdateConfig) *api.UpdateStrategy {
	if c == nil {
		return nil
	}

	stagger := c.Delay
	us := &api.UpdateStrategy{
		Stagger: &stagger,
	}

	if c.Parallelism != nil {
		p := int(*c.Parallelism)
		us.MaxParallel = &p
	}

	return us
}

func makeTaskEnv(e map[string]*string) map[string]string {
	if len(e) == 0 {
		return nil
	}

	env := make(map[string]string)
	for key, val := range e {
		if val != nil {
			env[key] = *val
		} else {
			env[key] = ""
		}
	}
	return env
}

func makeTaskConfigPortMap(ports map[string]types.ServicePortConfig) map[string]interface{} {
	portMap := make(map[string]interface{})
	for label, port := range ports {
		portMap[label] = int(port.Target)
	}
	return portMap
}

func makeTaskNetworkResourcePorts(ports map[string]types.ServicePortConfig) *api.NetworkResource {
	l := len(ports)
	netRsrc := &api.NetworkResource{
		ReservedPorts: make([]api.Port, 0, l),
		DynamicPorts:  make([]api.Port, 0, l),
	}
	for label, port := range ports {
		if port.Published > 0 {
			netRsrc.ReservedPorts = append(netRsrc.ReservedPorts, api.Port{
				Label: label,
				Value: int(port.Published),
			})
		} else {
			netRsrc.DynamicPorts = append(netRsrc.DynamicPorts, api.Port{
				Label: label,
			})
		}
	}
	return netRsrc
}

func labelPorts(ports []types.ServicePortConfig) map[string]types.ServicePortConfig {
	l := len(ports)
	switch l {
	case 0:
		return nil
	case 1:
		return map[string]types.ServicePortConfig{
			"default": ports[0],
		}
	}

	out := make(map[string]types.ServicePortConfig)
	for _, port := range ports {
		key := fmt.Sprintf("port%d", port.Target)
		out[key] = port
	}
	return out
}

func translateService(conf types.ServiceConfig) (*api.Task, error) {
	task := &api.Task{
		Driver: "docker",
		Name:   conf.Name,
		Env:    makeTaskEnv(conf.Environment),
		Config: map[string]interface{}{
			"image": conf.Image,
		},
	}

	clen := len(conf.Command)
	switch clen {
	case 0:
	case 1:
		task = task.SetConfig("command", conf.Command[0])

	default:
		task = task.SetConfig("command", conf.Command[0])
		task = task.SetConfig("args", conf.Command[1:])

	}

	// This should be the last thing we do
	ports := labelPorts(conf.Ports)
	if len(ports) == 0 {
		return task, nil
	}

	portmap := makeTaskConfigPortMap(ports)
	task = task.SetConfig("port_map", []map[string]interface{}{portmap})

	task.Resources = &api.Resources{
		Networks: []*api.NetworkResource{
			makeTaskNetworkResourcePorts(ports),
		},
	}

	task.Services = makeTaskServices(conf.Name, ports)

	// for _, vol := range conf.Volumes {

	// }
	// for _, sec := range conf.Secrets {

	// }
	// for _, network := range conf.Networks {

	// }

	return task, nil
}

func makeTaskServices(name string, ports map[string]types.ServicePortConfig) []*api.Service {
	out := make([]*api.Service, 0, len(ports))

	for k := range ports {
		service := &api.Service{
			Name:      name,
			PortLabel: k,
		}
		out = append(out, service)
	}

	return out
}
