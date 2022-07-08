package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var rPool *redis.Pool

func nPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", ":6379")

			// Connection error handling
			if err != nil {
				log.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func set(c redis.Conn) error {
	_, err := c.Do("SET", "name", "rommel")
	if err != nil {
		return err
	}

	_, err = c.Do("SET", "country", "Philippines")
	if err != nil {
		return err
	}

	return nil
}

func get(c redis.Conn) error {
	key := "name"
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %s\n", key, s)

	key = "country"
	i, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}

	fmt.Printf("%s = %s\n", key, i)

	// Example where key was not set
	key = "Nonexistent Key"
	s, err = redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Printf("%s : Alert! this Key does not exist\n", key)
	} else if err != nil {
		return err
	} else {
		fmt.Printf("%s = %s\n", key, s)
	}

	return nil
}

func ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)

	return nil
}

func main() {
	rPool = nPool()
	conn, err := rPool.Dial()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn)
	set(conn)
	get(conn)
}
