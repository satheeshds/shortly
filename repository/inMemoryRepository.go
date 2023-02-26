package repository

import (
	"fmt"
	"math"
	"sort"
)

type InMemoryRepository struct {
	urlDictionary map[string]string
	reverseLookUp map[string]string
	domainTracker map[string]int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		urlDictionary: make(map[string]string),
		reverseLookUp: make(map[string]string),
		domainTracker: make(map[string]int),
	}
}

func (r *InMemoryRepository) Store(shortUrl, original string) error {
	if err := r.isRepoInitialized(); err != nil {
		return err
	}

	if existing, ok := r.urlDictionary[shortUrl]; ok {
		return fmt.Errorf("%s already mapped with %s value", shortUrl, existing)
	}

	r.urlDictionary[shortUrl] = original
	r.reverseLookUp[original] = shortUrl
	return nil
}

func (r *InMemoryRepository) Get(shortUrl string) (string, error) {
	if err := r.isRepoInitialized(); err != nil {
		return "", err
	}

	if result, ok := r.urlDictionary[shortUrl]; ok {
		return result, nil
	}

	return "", fmt.Errorf("no url mapped for the given short url : %s", shortUrl)
}

func (r *InMemoryRepository) GetPreviousShortenedIfExist(original string) (string, error) {
	if err := r.isRepoInitialized(); err != nil {
		return "", err
	}

	if result, ok := r.reverseLookUp[original]; ok {
		return result, nil
	}

	return "", fmt.Errorf("no url mapped for the given url : %s", original)
}

func (r *InMemoryRepository) GetTopShortedDomains() (map[string]int, error) {

	if err := r.isRepoInitialized(); err != nil {
		return nil, err
	}

	// Convert the map to a slice of key-value pairs
	pairs := make([]struct {
		Key   string
		Value int
	}, len(r.domainTracker))
	i := 0
	for k, v := range r.domainTracker {
		pairs[i] = struct {
			Key   string
			Value int
		}{k, v}
		i++
	}

	// Sort the slice based on the values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	//return top 3 values
	result := make(map[string]int, 3)

	for j := 0; j < int(math.Min(3, float64(len(pairs)))); j++ {
		result[pairs[j].Key] = pairs[j].Value
	}
	return result, nil
}

func (r *InMemoryRepository) AddDomain(domain string) error {
	if err := r.isRepoInitialized(); err != nil {
		return err
	}

	count := r.domainTracker[domain]
	r.domainTracker[domain] = count + 1

	return nil
}

func (r *InMemoryRepository) isRepoInitialized() error {
	if r == nil || r.domainTracker == nil || r.urlDictionary == nil || r.reverseLookUp == nil {
		return fmt.Errorf("repository not initialized")
	}

	return nil
}
