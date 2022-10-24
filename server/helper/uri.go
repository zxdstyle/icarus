package helper

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"strings"
)

func GetResourceUriIndex(name, base string) string {
	return GetResourceUri(name, base)
}

func GetResourceUriCreate(name, base string) string {
	return fmt.Sprintf("%s/create", GetResourceUri(name, base))
}

func GetResourceUriStore(name, base string) string {
	return GetResourceUri(name, base)
}

func GetResourceUriShow(name, base string) string {
	return fmt.Sprintf("%s/:%s", GetResourceUri(name, base), base)
}

func GetResourceUriEdit(name, base string) string {
	return fmt.Sprintf("%s/edit", GetResourceUri(name, base))
}

func GetResourceUriUpdate(name, base string) string {
	return fmt.Sprintf("%s/:%s", GetResourceUri(name, base), base)
}

func GetResourceUriDestroy(name, base string) string {
	return fmt.Sprintf("%s/:%s", GetResourceUri(name, base), base)
}

func GetResourceUri(resource, base string) string {
	if !strings.Contains(resource, ".") {
		return resource
	}
	segments := strings.Split(resource, ".")
	uri := GetNestedResourceUri(segments)

	current := fmt.Sprintf("/:%s", base)
	return strings.Replace(uri, current, "", -1)
}

func GetNestedResourceUri(segments []string) string {
	var uri []string
	for _, segment := range segments {
		resource := fmt.Sprintf("%s/:%s", segment, GetResourceWildcard(segment))
		uri = append(uri, resource)
	}
	return strings.Join(uri, "/")
}

func GetResourceWildcard(value string) string {
	val := inflection.Singular(value)
	return strings.Replace(val, "_", "_", -1)
}

func GetBaseName(resource string) string {
	uri := strings.Split(resource, ".")
	return GetResourceWildcard(uri[len(uri)-1])
}
