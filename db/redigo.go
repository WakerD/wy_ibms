package db

import (
	"github.com/garyburd/redigo/redis"
	//"src/github.com/kylelemons/go-gypsy/yaml"
	//"demo/src/config"
	"time"
	"fmt"
)

var (
	//定义常量
	RedisClient *redis.Pool
	REDIS_HOST string
	REDIS_PORT string
	REDIS_DB int
)

func RedisInit() {
	//从配置文件中获取redis的ip以及db
	//config,err := yaml.ReadFile("dbConfig.yaml")
	REDIS_HOST = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_DB = 0
	//建立连接池
	RedisClient = &redis.Pool{
		//从配置文件获取maxidle以及maxactive，娶不到则用后面的默认值
		MaxIdle:100,
		MaxActive:1024,
		IdleTimeout:180 * time.Second,

		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",REDIS_HOST+":"+REDIS_PORT)
			if err !=nil{
				return nil,err
			}
			//选择db
			//c.Do("SELECT",REDIS_DB)
			return c,nil
		},
	}
}

//GetDB ...
func GetRedisPool() *redis.Pool {
	return RedisClient
}

/*redis订阅信息*/
func Subscribe()  {
	c:=RedisClient.Get()
	psc:=redis.PubSubConn{c}
	psc.Subscribe("redChatRoom")
	defer func() {
		c.Close()
		psc.Unsubscribe("redChatRoom")	//取消订阅
	}()
	for {
		switch v:=psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s:messages:%s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			continue
		case error:
			fmt.Println(v)
			return
		}
	}
}

/*redis发布信息*/
func Pubscribe(s string)  {
	c:=RedisClient.Get()
	defer c.Close()

	_,err:=c.Do("PUBLISH","redChatRoom",s)
	if err!=nil {
		fmt.Println("pub err:",err)
		return
	}
}

func test()  {
	//从池里获取连接
	rc:=RedisClient.Get()
	//用完后将连接放回连接池
	defer rc.Close()
	//rc.Do()
	//n,_:=rc.Do("EXPIRE",key,24*3600)
	//value,err:=redis.String(rs.Do("GET",key))
	return
}