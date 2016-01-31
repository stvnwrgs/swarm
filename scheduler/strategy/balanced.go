package strategy

import (
	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/scheduler/node"
	"sort"
)

type byContainersOnNode []*node.Node

func (n byContainersOnNode) Len() int {
	return len(n)
}

func (n byContainersOnNode) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// Compare container count on two nodes
func (n byContainersOnNode) Less(i, j int) bool {
	return len(n[i].Containers) < len(n[j].Containers)
}

// BalancedPlacementStrategy places a container to the cluster.
type BalancedPlacementStrategy struct {
}

// Initialize a BalancedPlacementStrategy.
func (p *BalancedPlacementStrategy) Initialize() error {
	return nil
}

// Name returns the name of the strategy.
func (p *BalancedPlacementStrategy) Name() string {
	return "balanced"
}

// RankAndSort sorts the list of nodes by containers running on each node.
func (p *BalancedPlacementStrategy) RankAndSort(config *cluster.ContainerConfig, nodes []*node.Node) ([]*node.Node, error) {
	sort.Sort(byContainersOnNode(nodes))
	return nodes, nil
}
