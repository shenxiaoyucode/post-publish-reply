package repository

import (
	"encoding/json"
	"math"
	"os"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}

func (*PostDao) FindMaxId() int64 {

	id := int64(math.MinInt64)

	for i := 0; i < len(posts); i++ {
		if posts[i].Id > id {
			id = posts[i].Id
		}
	}
	return id
}
func (*PostDao) CreatePost(id int64, pid int64, content string, time int64) (post Post, err error) {
	post = Post{
		Id:         id,
		ParentId:   pid,
		Content:    content,
		CreateTime: time,
	}

	posts = append(posts, &post)

	//重新设置map
	childPosts := postIndexMap[post.ParentId]
	childPosts = append(childPosts, &post)
	postIndexMap[post.ParentId] = childPosts
	f, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return post, err
	}
	defer f.Close()
	marshal, _ := json.Marshal(post)
	if _, err = f.WriteString(string(marshal) + "\n"); err != nil {
		return post, err
	}
	return post, nil

}
