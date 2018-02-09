package ioutil

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// WriteUInt8 写入uint8
func WriteUInt8(w io.Writer, value uint8) (int, error) {

	wrote, err := w.Write([]byte{byte(value)})
	if err != nil {
		return 0, err
	}

	if wrote != 1 {
		return 0, fmt.Errorf("wrote %d != 1", wrote)
	}

	return wrote, nil
}

// WriteUint16 写入uint16
func WriteUint16(w io.Writer, value uint16) (int, error) {

	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return 0, err
	}

	if wrote != 2 {
		return 0, fmt.Errorf("wrote %d != 2", wrote)
	}

	return wrote, nil
}

// WriteUint32 写入uint32
func WriteUint32(w io.Writer, value uint32) (int, error) {

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return 0, err
	}

	if wrote != 4 {
		return 0, fmt.Errorf("wrote %d != 4", wrote)
	}

	return wrote, nil
}

// WriteUint64 写入uint64
func WriteUint64(w io.Writer, value uint64) (int, error) {

	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return 0, err
	}

	if wrote != 8 {
		return 0, fmt.Errorf("wrote %d != 8", wrote)
	}

	return wrote, nil
}

// WriteString 写入字符串
func WriteString(w io.Writer, text string) (int, error) {

	buffer := []byte(text)
	bufferLength := len(buffer)
	size, err := WriteUint32(w, uint32(bufferLength))
	if err != nil {
		return 0, err
	}

	if bufferLength == 0 {
		return 0, nil
	}

	wrote, err := w.Write(buffer)
	if err != nil {
		return 0, err
	}

	if wrote != bufferLength {
		return 0, fmt.Errorf("wrote %d != %d", wrote, bufferLength)
	}

	return size + wrote, nil
}

// WriteTime 写入时间
func WriteTime(w io.Writer, value time.Time) (int, error) {

	size1, err := WriteUint64(w, uint64(value.Unix()))
	if err != nil {
		return 0, err
	}

	size2, err := WriteString(w, value.Location().String())
	if err != nil {
		return 0, err
	}

	return size1 + size2, nil
}
