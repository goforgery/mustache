package mustache

import (
	"github.com/hoisie/mustache"
)

type Mustache struct {
	cache    map[string]*mustache.Template
	useCache bool
}

func Create(useCache ...bool) *Mustache {
	this := &Mustache{}
	this.cache = map[string]*mustache.Template{}
	if len(useCache) == 1 {
		this.useCache = useCache[0]
	}
	return this
}

func (this *Mustache) Render(f string, o ...interface{}) (string, error) {
	t, ok := this.cache[f]
	if ok == false {
		tmpl, err := mustache.ParseFile(f)
		if err != nil {
			return "", err
		}
		if this.useCache {
			this.cache[f] = tmpl
		}
		t = tmpl
	}
	return t.Render(o...), nil
}

// Clean the cache.
func (this *Mustache) Clean() {
	this.cache = map[string]*mustache.Template{}
}
