package ioutil

import (
	"encoding/binary"
	"io"
	"time"
)

// WriteBytes 写入bytes
func WriteBytes(w io.Writer, buffer []byte) (int, error) {

	bufferLength := len(buffer)
	lengthSize, err := WriteUint64(w, uint64(bufferLength))
	if err != nil {
		return 0, err
	}

	wrote, err := w.Write(buffer)
	if err != nil {
		return 0, err
	}

	if wrote != bufferLength {
		return 0, io.ErrShortWrite
	}

	return lengthSize + bufferLength, nil
}

// WriteUInt8 写入uint8
func WriteUInt8(w io.Writer, value uint8) (int, error) {

	return w.Write([]byte{byte(value)})
}

// WriteUint16 写入uint16
func WriteUint16(w io.Writer, value uint16) (int, error) {

	lengthBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBytes, value)

	return w.Write(lengthBytes)
}

// WriteUint32 写入uint32
func WriteUint32(w io.Writer, value uint32) (int, error) {

	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, value)

	return w.Write(lengthBytes)
}

// WriteUint64 写入uint64
func WriteUint64(w io.Writer, value uint64) (int, error) {

	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, value)

	return w.Write(lengthBytes)
}

// WriteString 写入字符串
func WriteString(w io.Writer, text string) (int, error) {
	return WriteBytes(w, []byte(text))
}

// WriteTime 写入时间
func WriteTime(w io.Writer, value time.Time) (int, error) {

	unixSize, err := WriteUint64(w, uint64(value.Unix()))
	if err != nil {
		return 0, err
	}

	locationSize, err := WriteString(w, value.Location().String())
	if err != nil {
		return 0, err
	}

	return unixSize + locationSize, nil
}
