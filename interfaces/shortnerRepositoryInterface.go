package interfaces

type IShortnerRepository interface {
	Store(shortUrl, original string) error
	Get(shortUrl string) (string, error)
	GetTopShortedDomains() (map[string]int, error)
	AddDomain(domain string) error
}
