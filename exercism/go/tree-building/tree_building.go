package tree

import "sort"

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
		return nil, nil
	}
	sort.SliceStable(nrecords, func(i, j int) bool {
		return nrecords[i].ID < nrecords[j].ID
	})
	currentNode := root
	for i := range nrecords {
		cur := nrecords[i]
		if cur.Parent == currentNode.ID {
			child := &Node{
				ID: cur.ID,
			}
			currentNode.Children = append(currentNode.Children, child)
		} else {
			lookAndAttach(root.Children, cur)
		}
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

func lookAndAttach(children []*Node, record Record) {
	for i := range children {
		if children[i].ID == record.Parent {
			children[i].Children = append(children[i].Children, &Node{ID: record.ID})
		} else {
			lookAndAttach(children[i].Children, record)
		}
	}
}
