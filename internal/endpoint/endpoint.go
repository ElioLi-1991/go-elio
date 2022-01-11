package endpoint

import "net/url"


func NewEndpoint (scheme string,host string,isSecure bool)*url.URL {
	var query string
	if isSecure {
		query = "isSecure=true"
	}
	return &url.URL{Scheme: scheme, Host: host, RawQuery: query}
}
