package lxc

import (
	"github.com/caglar10ur/lxc"
)

type Container struct {
	Name  string
	State string
}

type ContainersReply struct {
	Containers []*Container
}

func containers() ([]*Container, error) {
	containers := make([]*Container, 0)
	for _, v := range lxc.Containers() {
		c := &Container{
			Name:  v.Name(),
			State: v.State().String(),
		}
		containers = append(containers, c)
	}
	return containers, nil
}

func (l *LXC) Containers(args int, reply *ContainersReply) error {
	var err error
	reply.Containers, err = containers()
	if err != nil {
		return err
	}
	return nil
}
