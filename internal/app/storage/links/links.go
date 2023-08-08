package links

type LinksStorage struct {
	Links *map[string]string
}

func (l *LinksStorage) AddLink(key, value string) {
	(*l.Links)[key] = value
}

func (l *LinksStorage) GetLink(key string) (string, bool) {
	v, ok := (*l.Links)[key]
	return v, ok
}

var linksData = make(map[string]string)

func GetLinksStorage() LinksStorage {
	return LinksStorage{
		Links: &linksData,
	}
}
