package services

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/satheeshds/shortly/interfaces"
)

type ShortnerService struct {
	Repo interfaces.IShortnerRepository
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (s *ShortnerService) ShortURL(original string) (string, error) {

	domain, err := extractDomain(original)
	if err != nil {
		return "", err
	}

	prev, err := s.Repo.GetPreviousShortenedIfExist(original)
	if err == nil {
		// do we need to increment the counter if the same url shortened several times.
		return prev, nil
	}

	shortId := generateRandomId()
	shortUrl := fmt.Sprintf("https://short.ly/%s", shortId)

	err = s.Repo.Store(shortUrl, original)
	if err != nil {
		return "", err
	}

	err = s.Repo.AddDomain(domain)
	if err != nil {
		return "", err
	}

	return shortUrl, err
}

func (s *ShortnerService) GetRedirectURL(shortUrl string) (string, error) {
	return s.Repo.Get(shortUrl)
}

func (s *ShortnerService) GetTopShortedDomains() (map[string]int, error) {
	return s.Repo.GetTopShortedDomains()
}

func generateRandomId() string {
	// Generate a random short ID
	rand.Seed(time.Now().UnixNano())
	shortId := make([]rune, 6)
	for i := range shortId {
		shortId[i] = letters[rand.Intn(len(letters))]
	}

	return string(shortId)
}

func extractDomain(original string) (string, error) {
	u, err := url.Parse(original)

	return u.Hostname(), err
}
