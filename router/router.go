package router

import "fmt"

type Router interface {
}

func NewDefaultRouter() (Router, error) {
	fmt.Println("init router")
	return nil, nil
}
