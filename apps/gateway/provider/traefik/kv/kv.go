package kv

import (
	"path"
	"reflect"

	"github.com/kvtools/valkeyrie/store"
	"github.com/traefik/paerser/parser"
)

// Decode decodes the given KV pairs into the given element.
// The operation goes through three stages roughly summarized as:
// KV pairs -> tree of untyped nodes
// untyped nodes -> nodes augmented with metadata such as kind (inferred from element)
// "typed" nodes -> typed element.
func Decode(pairs []*store.KVPair, element interface{}, rootName string) error {
	if element == nil {
		return nil
	}

	filters := GetElementKvs(rootName, element)
	node, err := DecodeToNode(pairs, rootName, filters.Keys()...)
	if err != nil {
		return err
	}

	metaOpts := parser.MetadataOpts{TagName: "kv", AllowSliceAsStruct: false}
	err = parser.AddMetadata(element, node, metaOpts)
	if err != nil {
		return err
	}

	return parser.Fill(element, node, parser.FillerOpts{AllowSliceAsStruct: false})
}

func NewKVSet() *KVSet {
	return &KVSet{
		Items: []*store.KVPair{},
	}
}

type KVSet struct {
	Items []*store.KVPair
}

func (s *KVSet) Add(items ...*store.KVPair) {
	s.Items = append(s.Items, items...)
}

func (s *KVSet) Keys() []string {
	keys := []string{}
	for i := range s.Items {
		keys = append(keys, s.Items[i].Key)
	}
	return keys
}

func GetElementKvs(rootName string, element any) *KVSet {
	if element == nil {
		return nil
	}

	rootType := reflect.TypeOf(element)
	kvs := GetFieldKvs(rootName, rootType)
	return kvs
}

func GetFieldKvs(rootName string, rootType reflect.Type) *KVSet {
	set := NewKVSet()

	if rootType.Kind() == reflect.Ptr {
		rootType = rootType.Elem()
	}

	if rootType.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < rootType.NumField(); i++ {
		field := rootType.Field(i)

		if !parser.IsExported(field) {
			continue
		}

		if field.Anonymous &&
			(field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct || field.Type.Kind() == reflect.Struct) {
			kvs := GetFieldKvs(rootName, field.Type)
			set.Add(kvs.Items...)
			continue
		}

		set.Add(&store.KVPair{Key: path.Join(rootName, field.Name)})
	}

	return set
}
