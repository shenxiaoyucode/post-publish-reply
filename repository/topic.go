package repository

import (
	"encoding/json"
	"math"
	"os"
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) CreateTopic(id int64, title string, content string, createTime int64) (topic Topic, err error) {
	newTopic := Topic{
		Id:         id,
		Title:      title,
		Content:    content,
		CreateTime: createTime,
	}
	topicIndexMap[newTopic.Id] = &newTopic
	f, err := os.OpenFile("./data/topic", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return newTopic, err
	}
	defer f.Close()
	marshal, _ := json.Marshal(newTopic)
	if _, err = f.WriteString(string(marshal) + "\n"); err != nil {
		return newTopic, err
	}
	return newTopic, nil
}

func (*TopicDao) FindMaxId() int64 {

	id := int64(math.MinInt64)
	for k, _ := range topicIndexMap {
		if k > id {
			id = k
		}
	}
	return id
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
