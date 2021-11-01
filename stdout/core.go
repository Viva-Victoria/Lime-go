package stdout

import (
	"fmt"

	pool "github.com/Viva-Victoria/go-pool"
)

type Writer struct {
	pool pool.Pool
}

func NewPooledCore(pool pool.Pool) *Writer {
	return &Writer{
		pool: pool,
	}
}

func NewSyncedCore() *Writer {
	return &Writer{
		pool: nil,
	}
}

func (c *Writer) Sync() error {
	if c.pool != nil {
		c.pool.Wait()
	}
	return nil
}

func (c *Writer) Write(data []byte) error {
	fmt.Println(string(data))
	return nil
}

func (c *Writer) PostWrite(data []byte) {
	if c.pool == nil {
		_ = c.Write(data)
	}

	c.pool.Add(func(_ int) {
		_ = c.Write(data)
	})
}
