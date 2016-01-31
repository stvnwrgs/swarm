package strategy

import (
	"fmt"
	"testing"

	"github.com/docker/swarm/scheduler/node"
	"github.com/stretchr/testify/assert"
)

func TestBalancedOnTwoSameNodes(t *testing.T) {
	s := &SpreadPlacementStrategy{}

	nodes := []*node.Node{
		createNode(fmt.Sprintf("node-0"), 128, 42),
		createNode(fmt.Sprintf("node-1"), 128, 42),
	}

	// add 25 containers
	for i := 0; i < 25; i++ {
		config := createConfig(0, 0)
		node := selectTopNode(t, s, config, nodes)
		assert.NoError(t, node.AddContainer(createContainer(fmt.Sprintf("c%d", i), config)))
	}

	assert.Equal(t, len(nodes[0].Containers), 13)
	assert.Equal(t, len(nodes[1].Containers), 12)

	// add 3 more containers (28)
	for i := 25; i < 28; i++ {
		config := createConfig(0, 0)
		node := selectTopNode(t, s, config, nodes)
		assert.NoError(t, node.AddContainer(createContainer(fmt.Sprintf("c%d", i), config)))
	}

	assert.Equal(t, len(nodes[0].Containers), 14)
	assert.Equal(t, len(nodes[1].Containers), 14)
}

func TestBalancedOnThreeDifferentNodes(t *testing.T) {
	s := &SpreadPlacementStrategy{}

	nodes := []*node.Node{
		createNode(fmt.Sprintf("node-0"), 128, 42),
		createNode(fmt.Sprintf("node-1"), 256, 21),
		createNode(fmt.Sprintf("node-2"), 64, 84),
	}

	// add 25 containers
	for i := 0; i < 25; i++ {
		config := createConfig(0, 0)
		node := selectTopNode(t, s, config, nodes)
		assert.NoError(t, node.AddContainer(createContainer(fmt.Sprintf("c%d", i), config)))
	}

	assert.Equal(t, len(nodes[0].Containers), 9)
	assert.Equal(t, len(nodes[1].Containers), 8)
	assert.Equal(t, len(nodes[2].Containers), 8)

	// add 5 more containers (30)
	for i := 25; i < 30; i++ {
		config := createConfig(0, 0)
		node := selectTopNode(t, s, config, nodes)
		assert.NoError(t, node.AddContainer(createContainer(fmt.Sprintf("c%d", i), config)))
	}

	assert.Equal(t, len(nodes[0].Containers), 10)
	assert.Equal(t, len(nodes[1].Containers), 10)
	assert.Equal(t, len(nodes[2].Containers), 10)
}
