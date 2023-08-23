package links

type LinksStorage struct {
	links map[string]string
}

func (l LinksStorage) AddLink(key, value string) {
	l.links[key] = value
}

func (l LinksStorage) GetLink(key string) (string, bool) {
	v, ok := l.links[key]
	return v, ok
}

var linksData = make(map[string]string)

func New() *LinksStorage {
	return &LinksStorage{
		links: linksData,
	}
}

var LinksStorageInstance = New()
