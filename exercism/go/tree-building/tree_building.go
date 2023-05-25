package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nrecords, root := lookForRoot(records)
	if root == nil {
		return nil, fmt.Errorf("no root")
	}
	sort.SliceStable(nrecords, func(i, j int) bool {
		return nrecords[i].ID < nrecords[j].ID
	})
	currentNode := root
	for i := range nrecords {
		cur := nrecords[i]
		fmt.Println("processing record ", i, cur)
		fmt.Println("tree before", root)

		if cur.Parent == currentNode.ID {
			child := &Node{
				ID: cur.ID,
			}
			currentNode.Children = append(currentNode.Children, child)
		} else {
			if cur.ID == cur.Parent {
				return nil, fmt.Errorf("direct cycle ID == Parent: %v %v", cur.ID, cur.Parent)
			}

			if err := lookAndAttach(root.Children, cur); err != nil {
				return nil, err
			}
		}
		fmt.Println("tree after", root)
		fmt.Println()
	}

	return root, nil
}

func lookForRoot(records []Record) ([]Record, *Node) {
	rootID := -1
	for i, v := range records {
		if isRoot(v) {
			rootID = v.ID
			if i == len(records) {
				records = records[:i]
			}
			records = append(records[:i], records[i+1:]...)
			break
		}
	}
	if rootID == -1 {
		return records, nil
	}
	return records, &Node{ID: rootID}
}

func isRoot(record Record) bool {
	return record.ID == record.Parent
}

func lookAndAttach(children []*Node, record Record) error {
	for i := range children {
		if record.ID == record.Parent {
			return fmt.Errorf("lookAndAttach direct cycle ID == Parent: %v %v", record.ID, record.Parent)
		}
		if children[i].ID == record.Parent {
			children[i].Children = append(children[i].Children, &Node{ID: record.ID})
		} else {
			if err := lookAndAttach(children[i].Children, record); err != nil {
				return err
			}
		}
	}

	return nil
}
