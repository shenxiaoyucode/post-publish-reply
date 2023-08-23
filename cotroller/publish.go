package cotroller

import (
	"pro/repository"
	"time"
)

func PublishTopic(title string, content string) (topic repository.Topic, err error) {

	id := repository.NewTopicDaoInstance().FindMaxId() + 1
	time := time.Now().Unix()
	topic, err = repository.NewTopicDaoInstance().CreateTopic(id, title, content, time)
	if err != nil {
		return topic, err
	}
	return topic, nil
}
func PublishPost(pid int64, content string) (post repository.Post, err error) {

	id := repository.NewPostDaoInstance().FindMaxId() + 1
	time := time.Now().Unix()
	post, err = repository.NewPostDaoInstance().CreatePost(id, pid, content, time)
	if err != nil {
		return post, err
	}
	return post, nil
}
