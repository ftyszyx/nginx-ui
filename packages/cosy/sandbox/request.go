package sandbox

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddHeader add header
func (c *Client) AddHeader(key, value string) {
	c.Header[key] = value
}

// attachHeader attach header to the given request
func (c *Client) attachHeader(req *http.Request) {
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}
}

// Get send a get request
func (c *Client) Get(uri string) (r *Response, err error) {
	return c.Request(http.MethodGet, uri, nil)
}

// Post send a post request
func (c *Client) Post(uri string, body gin.H) (r *Response, err error) {
	return c.Request(http.MethodPost, uri, body)
}

// Put send a put request
func (c *Client) Put(uri string, body gin.H) (r *Response, err error) {
	return c.Request(http.MethodPut, uri, body)
}

// Patch send a patch request
func (c *Client) Patch(uri string, body gin.H) (r *Response, err error) {
	return c.Request(http.MethodPatch, uri, body)
}

// Delete send a delete request
func (c *Client) Delete(uri string, body gin.H) (r *Response, err error) {
	return c.Request(http.MethodDelete, uri, body)
}

// Options send a option request
func (c *Client) Options(uri string, body gin.H) (r *Response, err error) {
	return c.Request(http.MethodOptions, uri, body)
}
