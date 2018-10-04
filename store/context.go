package store

import "github.com/gin-gonic/gin"

const key = "store"

type Setter interface {
	Set(string, interface{})
}

func FromContext(c *gin.Context) *Store {
	return c.Value(key).(*Store)
}

func ToContext(c *gin.Context, store *Store) {
	c.Set(key, store)
}
