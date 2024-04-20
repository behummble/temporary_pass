package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/behummble/temporary_pass/internal/external_service/db"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	connect     *redis.Client
}

type Cookie struct {
	// some struct from redis or string
	value string
}

type UserMessage struct {
	correlationId float64
	RequestTable string
	RequestTask string
	meta struct{}
	Args struct {
		Visitor string
		StartDate string
		EndDate string
		Phone string
	}
	successTasks []string
	errorTasks []string
}

var (
	Conn *redis.Client
)

func NewClient() db.DB {
	if Conn == nil {
		Connect()
	}
	return Redis{
		connect: Conn,
	}
}

func Connect() {
	options := &redis.Options{
		Addr: os.Getenv("REDIS_ADDRES"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	Conn = redis.NewClient(options)

	_, err := Conn.Ping(context.Background()).Result()

	if err != nil {
		log.Println(err)
	}
}

func (redis Redis) GetCookie() (db.Cookie, error) {
	cookie, err := redis.connect.Get(
		context.Background(), 
		"TEOREMA_COOKIE").Result()

	return Cookie{value: cookie}, err
}

func (redis Redis) GetMessages(messages chan<- db.UserMessage, queue string) {
	
	for {
		ctx := context.WithValue(context.Background(), "queue", queue)
		passMessages, err := redis.connect.BLPop(ctx, time.Second * 5, queue).Result()
		if err != nil {
			log.Println(err)
			continue
		}
		for _, message :=  range passMessages {
			messages<- convertRedisMessage(message)
		}
	}
}

func (cookie Cookie) String() string {
	return cookie.value
} 

func (msg UserMessage) JsonString() string {
	return ""
}

func convertRedisMessage(message string) UserMessage {
	usrMsg := &UserMessage{}
	
	json.Unmarshal(bytes.NewBufferString(message).Bytes(), usrMsg)

	return *usrMsg
}

func (userMessage UserMessage) GetOfficeName() string {
	return userMessage.RequestTable
}

func (userMessage UserMessage) GetUserName() string {
	return userMessage.Args.Visitor
}

func (userMessage UserMessage) GetPhoneNumber() string {
	return userMessage.Args.Phone
}

func (redis Redis) SetCookie(value string) {
	success := redis.connect.Set(
		context.Background(), 
		"TEOREMA_COOKIE",
		value,
		time.Second * 5)
	if success.Err() != nil {
		log.Println(success.Err())
	}
}