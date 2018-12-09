package main

import (
	"bufio"
	"io"
)

type Node struct {
	MetaData []int
	Children []Node
}

func ReadNode(reader io.Reader) (Node, error) {
	buffered := bufio.NewReader(reader)
	node := Node{}
	header, err := ReadHeader(buffered)
	if err != nil {
		return node, err
	}
	node.MetaData = make([]int, 0, header.MetaDataCount)
	node.Children = make([]Node, 0, header.ChildCount)

	for i := 0; i < header.ChildCount; i++ {
		child, err := ReadNode(buffered)
		if err != nil {
			return node, err
		}
		node.Children = append(node.Children, child)
	}

	for i := 0; i < header.MetaDataCount; i++ {
		metaData, err := ReadInteger(buffered)
		if err != nil {
			return node, err
		}
		node.MetaData = append(node.MetaData, metaData)
	}
	return node, nil
}

func (n *Node) GetValue() int {
	if len(n.Children) == 0 {
		return n.SumMetaData()
	}
	sum := 0
	for _, m := range n.MetaData {
		index := m - 1
		if index >= 0 && index < len(n.Children) {
			sum = sum + n.Children[index].GetValue()
		}
	}
	return sum
}

func (n *Node) SumMetaData() int {
	sum := 0
	n.sumMetaData(&sum)
	return sum
}

func (n *Node) sumMetaData(sum *int) {
	for _, m := range n.MetaData {
		*sum = *sum + m
	}
	for _, c := range n.Children {
		c.sumMetaData(sum)
	}
}
