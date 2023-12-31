package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
	posts         []*Post
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {

		return err
	}
	if err := initPostIndexMap(filePath); err != nil {

		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		if "" == text {
			return nil
		}
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {

		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)

	for scanner.Scan() {
		text := scanner.Text()
		if "" == text {
			return nil
		}
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {

			return err
		}

		posts = append(posts, &post)

	}
	for i := 0; i < len(posts); i++ {

		childPosts, ok := postTmpMap[posts[i].ParentId]
		if !ok {
			postTmpMap[posts[i].ParentId] = []*Post{posts[i]}
			continue
		}
		childPosts = append(childPosts, posts[i])
		postTmpMap[posts[i].ParentId] = posts
	}

	postIndexMap = postTmpMap
	return nil
}
