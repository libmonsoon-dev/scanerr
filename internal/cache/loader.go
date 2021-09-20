package cache

import (
	"strings"
	"sync"

	"github.com/libmonsoon-dev/scanerr/config"

	"golang.org/x/tools/go/packages"
)

type UpstreamLoader interface {
	Load(conf config.LoaderConfig) ([]*packages.Package, error)
}

func NewPackageLoader(upstreamLoader UpstreamLoader) *Loader {
	l := &Loader{
		upstream: upstreamLoader,
		cache:    make(map[string]*loaderCacheItem),
	}

	return l
}

type loaderCacheItem struct {
	sync.Once

	result []*packages.Package
	error  error
}

type Loader struct {
	upstream UpstreamLoader

	cacheMutex sync.Mutex
	cache      map[string]*loaderCacheItem
}

func (l *Loader) Load(conf config.LoaderConfig) ([]*packages.Package, error) {
	item := l.getCacheItem(conf)

	item.Do(func() {
		item.result, item.error = l.upstream.Load(conf)
	})
	return item.result, item.error
}

func (l *Loader) getCacheItem(conf config.LoaderConfig) (item *loaderCacheItem) {
	l.cacheMutex.Lock()
	defer l.cacheMutex.Unlock()

	key := getLoaderKey(conf)
	item = l.cache[key]
	if item == nil {
		item = new(loaderCacheItem)
		l.cache[key] = item
	}

	return
}

func getLoaderKey(conf config.LoaderConfig) string {
	return strings.Join(conf.Patterns, ",")
}

var _ UpstreamLoader = (*Loader)(nil)
