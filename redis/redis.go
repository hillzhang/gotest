package redis

import (
	"sys_monitor/garyburd/redigo/redis"
	"time"
	"fmt"
)

var RedisConnPool *redis.Pool


func InitRedisConnPool() {
	RedisConnPool = &redis.Pool{
		MaxIdle:     50,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "172.17.0.2:39300")
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH","D0bq117");err != nil{
				c.Close()
				return nil, err
			}
			c.Do("SELECT", 8)
			return c, err
		},
	}
}

func Redis(){
	InitRedisConnPool()

	rd := RedisConnPool.Get()
	defer rd.Close()

	reply,err := rd.Do("set","test","hshshs")
	if err != nil {
		fmt.Println(err)
	}

	reply,err = redis.String(rd.Do("get","test"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}

