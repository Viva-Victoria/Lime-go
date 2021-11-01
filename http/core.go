package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Viva-Victoria/go-pool"
)

type RequestBuilder func(body []byte) (*http.Request, error)

func DefaultRequestBuilder(ctx context.Context, url string) RequestBuilder {
	return func(body []byte) (*http.Request, error) {
		return http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	}
}

type Writer struct {
	buildRequest RequestBuilder
	pool         pool.Pool
	client       http.Client
}

func DefaultCore(url string) *Writer {
	p, _ := pool.NewFixedPool(1)
	return &Writer{
		pool: p,
		client: http.Client{
			Timeout: time.Second,
		},
		buildRequest: DefaultRequestBuilder(context.Background(), url),
	}
}

var (
	ErrPoolRequired           = errors.New("pool.Pool instance required")
	ErrRequestBuilderRequired = errors.New("HttpRequestBuilder required")
)

func NewCore(pool pool.Pool, client http.Client, builder RequestBuilder) (*Writer, error) {
	if pool == nil {
		return nil, ErrPoolRequired
	}
	if builder == nil {
		return nil, ErrRequestBuilderRequired
	}

	return &Writer{
		pool:         pool,
		client:       client,
		buildRequest: builder,
	}, nil
}

func (c *Writer) Sync() error {
	c.pool.Wait()
	return nil
}

func (c *Writer) Write(data []byte) error {
	request, err := c.buildRequest(data)
	if err != nil {
		log.Printf("cannot create request to write log: %v", err)
		return err
	}

	response, err := c.client.Do(request)
	if err != nil {
		log.Printf("cannot send log through http: %v", err)
		return err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		log.Printf("invalid status code received: %d", response.StatusCode)
		return fmt.Errorf("ivalid status received: %d %s", response.StatusCode, response.Status)
	}

	return nil
}

func (c *Writer) PostWrite(data []byte) {
	c.pool.Add(func(_ int) {
		_ = c.Write(data)
	})
}
