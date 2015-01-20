package cluster

import "github.com/samalba/dockerclient"

type Cluster interface {
	CreateContainer(config *dockerclient.ContainerConfig, name string) (*Container, error)
	RemoveContainer(container *Container, force bool) error

	Images() []*Image
	Image(IdOrName string) *Image
	Containers() []*Container
	Container(IdOrName string) *Container

	Pull(name string, begin, end func(string))

	Info() [][2]string
}
