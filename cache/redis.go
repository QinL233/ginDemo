package cache

import (
	"encoding/json"
	"ginDemo/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Cache *redis.Pool

func Init() {
	Cache = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

type Prefix struct {
	Name string
}

func (p *Prefix) Set(key string, value interface{}, time int) error {
	c := Cache.Get()
	//函数调用完毕释放
	defer c.Close()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.Do("SET", p.Name+key, json)
	if err != nil {
		return err
	}
	_, err = c.Do("EXPIRE", p.Name+key, time)
	if err != nil {
		return err
	}
	return nil
}

func (p *Prefix) Exists(key string) bool {
	c := Cache.Get()
	defer c.Close()

	exists, err := redis.Bool(c.Do("EXISTS", p.Name+key))
	if err != nil {
		return false
	}

	return exists
}

func (p *Prefix) Get(key string) ([]byte, error) {
	c := Cache.Get()
	defer c.Close()

	reply, err := redis.Bytes(c.Do("GET", p.Name+key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (p *Prefix) Delete(key string) (bool, error) {
	c := Cache.Get()
	defer c.Close()

	return redis.Bool(c.Do("DEL", p.Name+key))
}
