package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"quizGo/constants"

	"github.com/go-redis/redis/v8"
)

func Cache() *redis.Client {
	redisURI := constants.EnvConstant("REDISURI")

	addr, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(addr)

	fmt.Println("Cache Server connected")

	return rdb
}

var rdb *redis.Client = Cache()

func SetCache(ctx context.Context, key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func CheckKeyCache(ctx context.Context, key string) bool {
	fmt.Println(key)
	exist, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println("Cache Key Error: ", err.Error())
		return false
	}

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func GetCache(ctx context.Context, key string, dest interface{}) error {
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), dest)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCacheKey(ctx context.Context, key string) error {
	fmt.Println("Deleting Cache Key: ", key)
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}

func UpdateDocumentCountCache(ctx context.Context, key string, delata int64) error {
	var existingDocument int64
	if exist := CheckKeyCache(ctx, key); exist {
		err := GetCache(ctx, key, &existingDocument)
		if err != nil {
			return err
		}

		err = SetCache(ctx, key, existingDocument+delata)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unable to get key")
	}

	return nil
}
