package tree

import (
	"errors"
	"fmt"
	"log"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(inputDataRecords []Record) (*Node, error) {
	if len(inputDataRecords) == 0 {
		return nil, nil
	}
	rootNodePointer := &Node{}
	todoSliceCurrentIteration := []*Node{rootNodePointer}

	n := 1
	for {
		log.Println("n", n)
		log.Println("rootNodePointer", rootNodePointer)
		log.Printf("%v - todoSliceCurrentIteration\n", todoSliceCurrentIteration)
		log.Println(len(todoSliceCurrentIteration), "len(todoSliceCurrentIteration)")
		if len(todoSliceCurrentIteration) == 0 {
			log.Println("break\n\n")
			break
		}
		newTodo := []*Node(nil)
		for _, currentSliceValue := range todoSliceCurrentIteration {
			for _, currentRecordValue := range inputDataRecords {
				if currentRecordValue.Parent == currentSliceValue.ID {
					if currentRecordValue.ID < currentSliceValue.ID {
						return nil, errors.New("a")
					} else if currentRecordValue.ID == currentSliceValue.ID {
						if currentRecordValue.ID != 0 {
							return nil, fmt.Errorf("b")
						}
					} else {
						n++
						switch len(currentSliceValue.Children) {
						case 0:
							nn := &Node{ID: currentRecordValue.ID}
							currentSliceValue.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							nn := &Node{ID: currentRecordValue.ID}
							if currentSliceValue.Children[0].ID < currentRecordValue.ID {
								currentSliceValue.Children = []*Node{currentSliceValue.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								currentSliceValue.Children = []*Node{nn, currentSliceValue.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							nn := &Node{ID: currentRecordValue.ID}
							newTodo = append(newTodo, nn)
						breakpoint:
							for range []bool{false} {
								for i, cc := range currentSliceValue.Children {
									if cc.ID > currentRecordValue.ID {
										//try to create a slice outside all loops with a redundant capasity
										a := make([]*Node, len(currentSliceValue.Children)+1)
										copy(a, currentSliceValue.Children[:i])
										copy(a[i+1:], currentSliceValue.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										currentSliceValue.Children = a
										break breakpoint
									}
								}
								currentSliceValue.Children = append(currentSliceValue.Children, nn)
							}
						}
					}
				}
			}
		}
		todoSliceCurrentIteration = newTodo
	}
	if n != len(inputDataRecords) {
		return nil, Mismatch{}
	}
	if err := chk(rootNodePointer, len(inputDataRecords)); err != nil {
		return nil, err
	}
	return rootNodePointer, nil
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
