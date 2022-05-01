package utils

import "github.com/graphql-go/graphql"

var _cachedListType map[string]*graphql.List
var _cachedObjectType map[string]*graphql.Object

// Creates a cached type, useful for wrapper graphql functions that return type
func CreateList(key string, interfaceConfig *graphql.List) *graphql.List {
	if _cachedListType == nil {
		_cachedListType = make(map[string]*graphql.List)
	}

	if config, ok := _cachedListType[key]; ok {
		return config
	} else {
		_cachedListType[key] = interfaceConfig
		return _cachedListType[key]
	}
}

func CreateObject(key string, interfaceConfig *graphql.Object) *graphql.Object {
	if _cachedObjectType == nil {
		_cachedObjectType = make(map[string]*graphql.Object)
	}

	if config, ok := _cachedObjectType[key]; ok {
		return config
	} else {
		_cachedObjectType[key] = interfaceConfig
		return _cachedObjectType[key]
	}
}
