package writer

type Writer interface {
	Sync() error
	Write(data []byte) error
	PostWrite(data []byte)
}
