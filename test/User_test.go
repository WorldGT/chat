package test

import (
	"chat_socket/server/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

type Userclaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

type Room struct {
	ID       int64
	Name     string
	Players  []int64
	MsgIndex int64
}

func (r *Room) AddMsg() int64 {
	r.MsgIndex++
	return r.MsgIndex
}

type ChatMessage struct {
	ID       int64
	Sender   string
	RoomName string
	Content  string
	CreateTs int64
}

var messageKeyPre string = "M:"

// AddMessage 存入消息
func AddMessage(client *redis.Client, msg ChatMessage) bool {
	key := messageKeyPre + msg.RoomName + strconv.FormatInt(msg.ID, 10)
	key = strings.TrimSpace(key)
	fmt.Println("key:", key)
	_, err := client.HSet(key, "id", msg.ID).Result()
	fmt.Println("HSet:", key)
	if err != nil {
		fmt.Println("redisop AddMessage HSet msg Error:", err.Error())
		return false
	}
	detail, err1 := json.Marshal(msg)
	if err1 != nil {
		return false
	}
	client.HSet(key, "detail", detail)
	return true

}

func NowTs() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}
func TestUser(t *testing.T) {

	msg := ChatMessage{
		ID:       1,
		Sender:   "gt钩",
		RoomName: "room1",
		Content:  "测试消息cesxiaoxi",
		CreateTs: NowTs(),
	}
	fmt.Println(msg)
	AddMessage(models.Rds, msg)
}
