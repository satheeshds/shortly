package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var topDomainSortedSet string = "topDomains"

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() *RedisRepository {
	url := os.Getenv("REDIS_URL")
	fmt.Println("REDIS_URL:", url)
	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
	})

	return &RedisRepository{
		client: redisClient,
	}
}

func (r *RedisRepository) Store(shortUrl, original string) error {
	if err := r.isRepoInitialized(); err != nil {
		return err
	}

	res := r.client.Get(context.Background(), shortUrl)
	if res.Err() == nil {
		return fmt.Errorf("%s already mapped with %s value", shortUrl, res.Val())
	}

	//no expiration set, ideally we should be having some expiration for the short urls
	if err := r.client.Set(context.Background(), shortUrl, original, 0).Err(); err != nil {
		return err
	}

	if err := r.client.Set(context.Background(), original, shortUrl, 0).Err(); err != nil {
		return err
	}

	return nil

}

func (r *RedisRepository) Get(shortUrl string) (string, error) {
	if err := r.isRepoInitialized(); err != nil {
		return "", err
	}

	res := r.client.Get(context.Background(), shortUrl)
	return res.Val(), res.Err()
}

func (r *RedisRepository) GetPreviousShortenedIfExist(original string) (string, error) {
	return r.Get(original)
}

func (r *RedisRepository) AddDomain(domain string) error {

	if err := r.isRepoInitialized(); err != nil {
		return err
	}

	if err := r.client.ZIncrBy(context.Background(), topDomainSortedSet, 1, domain).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) GetTopShortedDomains() (map[string]int, error) {
	if err := r.isRepoInitialized(); err != nil {
		return nil, err
	}
	// get the top two elements with scores from the sorted set
	results, err := r.client.ZRevRangeWithScores(context.Background(), topDomainSortedSet, 0, 2).Result()
	if err != nil {
		return nil, err
	}

	// add the results to the map
	res := make(map[string]int, 3)
	for _, result := range results {
		res[result.Member.(string)] = int(result.Score)
	}

	return res, nil
}

func (r *RedisRepository) isRepoInitialized() error {
	if r == nil || r.client == nil {
		return fmt.Errorf("repository not initialized")
	}

	return nil
}
