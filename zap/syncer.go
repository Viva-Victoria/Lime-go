package zap

import (
	"github.com/Viva-Victoria/Lime-go/writer"
	"go.uber.org/zap/zapcore"
)

type CoreWriteSyncer struct {
	sync   bool
	writer writer.Writer
}

func NewWriteSyncer(sync bool, writer writer.Writer) zapcore.WriteSyncer {
	return &CoreWriteSyncer{
		sync:   sync,
		writer: writer,
	}
}

func (c *CoreWriteSyncer) Write(data []byte) (int, error) {
	if c.sync {
		err := c.writer.Write(data)
		return len(data), err
	}

	c.writer.PostWrite(data)
	return len(data), nil
}

func (c *CoreWriteSyncer) Sync() error {
	return c.writer.Sync()
}
