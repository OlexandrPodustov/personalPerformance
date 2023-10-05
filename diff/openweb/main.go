package main

import "fmt"

type Conversation struct {
	Comments []Comment
}

type Comment struct {
	messageID string
	content   string

	Children []Comment
}

var storage Conversation

func AddComent(parentID, messageID, content string) error {
	if parentID == "" {
		storage.Comments = append(storage.Comments, Comment{
			messageID: messageID,
			content:   content,
		})
		return nil
	}
	for i := 0; i < len(storage.Comments); i++ {
		if findParentAndAddcomment(&storage.Comments[i], parentID, messageID, content) {
			break
		}
	}

	return nil
}

func findParentAndAddcomment(comment *Comment, parentID, messageID, content string) bool {
	if comment.messageID == parentID {
		comment.Children = append(comment.Children, Comment{
			messageID: messageID,
			content:   content,
		})
		return true
	}
	if len(comment.Children) == 0 {
		return false
	}
	for i := 0; i < len(comment.Children); i++ {
		if findParentAndAddcomment(&comment.Children[i], parentID, messageID, content) {
			return true
		}
	}

	return false
}

func ReadConversation() (*Conversation, error) {
	return &storage, nil
}

func main() {
	AddComent("", "parenttttID1", "con con cntent")
	AddComent("parenttttID1", "ch1", "1.1")
	fmt.Println(ReadConversation())
}

// root comment 1
// 	child comment 1.1
// 	child comment 1.2
// root comment 2
