package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/behummble/temporary_pass/internal/external_service/db"
	"github.com/behummble/temporary_pass/internal/external_service/office_service"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	connect     *redis.Client
	sub *redis.PubSub
}

type Cookie struct {
	// some struct from redis or string
	value string
}

type UserMessage struct {
	// some struct from redis
	corelationId float64
	requestTable string
	requestTask string
	meta interface{}
	args interface{}
	successTasks interface{}
	errorTasks interface{}
	office officeservice.Office
}

func Connect() Redis {
	options := &redis.Options{
		Addr: os.Getenv("REDIS_ADDRES"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 1,
	}

	client := Redis{
		connect: redis.NewClient(options),
	}

	_, err := client.connect.Ping(context.Background()).Result()

	if err != nil {
		log.Println(err)
		return client
	}

	client.subscripe()
	return client
}

func (redis *Redis) subscripe() {
	redis.sub = redis.connect.Subscribe(context.Background(), os.Getenv("REDIS_CHANNEL"))
}

func (redis Redis) GetCookie() db.Cookie {
	//get data from connect
	return Cookie{}
}

func (redis Redis) GetMessages(messages chan<- db.UserMessage) {
	//get data from connect
	queue := redis.sub.Channel()
	for {
		select {
		case redisMsg := <- queue :
			messages<- convertRedisMessage(redisMsg)
		}
	}
}

func (cookie Cookie) String() string {
	return cookie.value
} 

func (msg UserMessage) JsonString() string {
	return ""
}

func convertRedisMessage(message *redis.Message) UserMessage {
	usrMsg := &UserMessage{}
	
	json.Unmarshal(bytes.NewBufferString(message.Payload).Bytes(), usrMsg)

	return *usrMsg
}

func (userMessage UserMessage) GetOffice() officeservice.Office {
	return userMessage.office
}

func (userMessage UserMessage) GetUserName() string {
	return userMessage.args.(string)
}

func (userMessage UserMessage) GetPhoneNumber() string {
	return userMessage.args.(string)
}