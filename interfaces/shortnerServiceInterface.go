package interfaces

type IShortnerService interface {
	ShortURL(url string) (string, error)
	GetRedirectURL(shortUrl string) (string, error)
	GetTopShortedDomains() (map[string]int, error)
}
