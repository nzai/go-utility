package ioutil

import (
	"encoding/binary"
	"io"
	"time"
)

// WriteBytes 写入bytes
func WriteBytes(w io.Writer, buffer []byte) error {

	bufferLength := len(buffer)
	err := WriteUint64(w, uint64(bufferLength))
	if err != nil {
		return err
	}

	wrote, err := w.Write(buffer)
	if err != nil {
		return err
	}

	if wrote != bufferLength {
		return io.ErrShortWrite
	}

	return nil
}

// WriteUInt8 写入uint8
func WriteUInt8(w io.Writer, value uint8) error {

	wrote, err := w.Write([]byte{byte(value)})
	if err != nil {
		return err
	}

	if wrote != 1 {
		return io.ErrShortWrite
	}

	return nil
}

// WriteUint16 写入uint16
func WriteUint16(w io.Writer, value uint16) error {

	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, value)

	wrote, err := w.Write(lengthBytes)
	if err != nil {
		return err
	}

	if wrote != 2 {
		return io.ErrShortWrite
	}

	return nil
}

// WriteUint32 写入uint32
func WriteUint32(w io.Writer, value uint32) error {

	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, value)

	wrote, err := w.Write(lengthBytes)
	if err != nil {
		return err
	}

	if wrote != 4 {
		return io.ErrShortWrite
	}

	return nil
}

// WriteUint64 写入uint64
func WriteUint64(w io.Writer, value uint64) error {

	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, value)

	wrote, err := w.Write(lengthBytes)
	if err != nil {
		return err
	}

	if wrote != 8 {
		return io.ErrShortWrite
	}

	return nil
}

// WriteString 写入字符串
func WriteString(w io.Writer, text string) error {
	return WriteBytes(w, []byte(text))
}

// WriteTime 写入时间
func WriteTime(w io.Writer, value time.Time) error {

	err := WriteUint64(w, uint64(value.Unix()))
	if err != nil {
		return err
	}

	return WriteString(w, value.Location().String())
}
