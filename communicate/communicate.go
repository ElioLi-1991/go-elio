package communicate

import (
	"context"
	"net/url"
)

type EndPointer interface {
	Endpoint() (*url.URL, error)
}

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
