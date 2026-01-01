package registry

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
)

type Registerable interface {
	Register(registry *Registry)
}

type Registry struct {
	Items    map[string]any
	initList *list.List
}

func (r *Registry) Init() {
	r.Items = make(map[string]any)
	r.initList = list.New()
}

func (r *Registry) Add(item Registerable) {
	r.initList.PushBack(item)
}

func (r *Registry) Build() error {
	// start processing the init list until all items are registered
	for e := r.initList.Front(); e != nil; e = e.Next() {
		item, ok := e.Value.(Registerable)
		if ok {
			item.Register(r)
		} else {
			return fmt.Errorf("item in init list does not implement Registerable")
		}
	}
	return nil
}

func (r *Registry) Get(item any) Registerable {
	// BuildIndex the qualified type name
	name := getQualifiedType(item)

	if strings.Index(name, "*") != 0 {
		name = "*" + name
	}

	if obj, ok := r.Items[name]; !ok {
		panic(fmt.Sprintf("Dependency with key " + name + " not found in map"))
	} else {
		return obj.(Registerable)
	}
}

func (r *Registry) Set(item Registerable) {
	name := getQualifiedType(item)
	(r.Items)[name] = item
}

func getQualifiedType(i interface{}) string {
	t := reflect.TypeOf(i)
	name := t.String()
	return name
}
