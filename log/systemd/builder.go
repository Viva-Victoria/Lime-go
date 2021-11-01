package systemd

import (
	"bytes"
	"io"
)

type StringBuilder struct {
	length int
	buffer bytes.Buffer
}

func (s *StringBuilder) AppendString(str string) *StringBuilder {
	s.buffer.WriteString(str)
	return s
}

func (s *StringBuilder) AppendStrings(all ...string) *StringBuilder {
	for _, str := range all {
		s.AppendString(str)
	}
	return s
}

func (s *StringBuilder) WriteTo(w io.Writer) (int64, error) {
	return s.buffer.WriteTo(w)
}

func (s *StringBuilder) Bytes() []byte {
	return s.buffer.Bytes()
}

func (s *StringBuilder) String() string {
	return s.buffer.String()
}
