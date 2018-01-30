package ioutil

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// WriteUInt8 写入uint8
func WriteUInt8(w io.Writer, value uint8) error {

	wrote, err := w.Write([]byte{byte(value)})
	if err != nil {
		return err
	}

	if wrote != 1 {
		return fmt.Errorf("wrote %d != 1", wrote)
	}

	return nil
}

// WriteUint16 写入uint16
func WriteUint16(w io.Writer, value uint16) error {

	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return err
	}

	if wrote != 2 {
		return fmt.Errorf("wrote %d != 2", wrote)
	}

	return nil
}

// WriteUint32 写入uint32
func WriteUint32(w io.Writer, value uint32) error {

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return err
	}

	if wrote != 4 {
		return fmt.Errorf("wrote %d != 4", wrote)
	}

	return nil
}

// WriteUint64 写入uint64
func WriteUint64(w io.Writer, value uint64) error {

	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, value)

	wrote, err := w.Write(buffer)
	if err != nil {
		return err
	}

	if wrote != 8 {
		return fmt.Errorf("wrote %d != 8", wrote)
	}

	return nil
}

// WriteString 写入字符串
func WriteString(w io.Writer, text string) error {

	buffer := []byte(text)
	bufferLength := len(buffer)
	err := WriteUint64(w, uint64(bufferLength))
	if err != nil {
		return err
	}

	if bufferLength == 0 {
		return nil
	}

	wrote, err := w.Write(buffer)
	if err != nil {
		return err
	}

	if wrote != bufferLength {
		return fmt.Errorf("wrote %d != %d", wrote, bufferLength)
	}

	return nil
}

// WriteTime 写入时间
func WriteTime(w io.Writer, value time.Time) error {

	err := WriteUint64(w, uint64(value.Unix()))
	if err != nil {
		return err
	}

	return WriteString(w, value.Location().String())
}
