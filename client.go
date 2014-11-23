package firebase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type client struct {
	baseUrl string
}

type Client interface {
	BaseUrl() string
	// SetBaseUrl(url string)

	GetEntity(key string) (*firebaseEntity, error)
	// GetEntity2(key string, v interface{}) error
}

func NewClient(url string) client {
	return client{baseUrl: url}
}

func (c *client) BaseUrl() string {
	return c.baseUrl
}

// func (c *client) SetBaseUrl(url string) {
// 	c.baseUrl = url
// }

// func (c *client) get(path string) (resp *http.Response, err error) {
// 	return http.Get(c.baseUrl + path)
// }

func appendExtension(s string, ext string) string {
	parts := []string{s, ext}
	return strings.Join(parts, ".")
}

func appendJsonExtension(s string) string {
	return appendExtension(s, "json")
}

func (c *client) getEntityUrl(key string) string {
	return appendJsonExtension(c.BaseUrl() + key)
}

func (c *client) GetEntity(key string) (*firebaseEntity, error) {
	var url = c.getEntityUrl(key)
	if response, err := http.Get(url); err != nil {
		log.Println("could not GET "+url, err)
		return nil, err
	} else {
		defer response.Body.Close()
		if body, err := ioutil.ReadAll(response.Body); err != nil {
			log.Println("could not response from GET "+url, err)
			return nil, err
		} else {
			var data = make(map[string]interface{})
			json.Unmarshal(body, &data)

			var entity = NewFirebaseEntity()
			entity.Data = data

			return entity, nil
		}
	}
}
