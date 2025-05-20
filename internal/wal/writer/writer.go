package writer

import (
	"encoding/binary"
	"fmt"
	pb "github.com/zahidhasanpapon/my-wal/proto"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

// Writer handles writing entries to the WAL file
type Writer struct {
	file *os.File
}

// NewWriter creates a new writer instance
func NewWriter(file *os.File) *Writer {
	return &Writer{
		file: file,
	}
}

// Write serializes and writes an entry to the WAL
// Format: [entry size (4 bytes) [entry bytes])
func (writer *Writer) Write(entry *pb.LogEntry) error {
	// Serialize the entry to bytes
	data, err := proto.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal entry: %w", err)
	}

	// Get the current position of appending
	_, err = writer.file.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("failed to seek to end: %w", err)
	}

	// Write entry size (uint32) - 4 bytes
	sizeBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuf, uint32(len(data)))
	if _, err = writer.file.Write(sizeBuf); err != nil {
		return fmt.Errorf("failed to write entry size: %w", err)
	}

	// Write entry data
	if _, err = writer.file.Write(data); err != nil {
		return fmt.Errorf("failed to write entry data: %w", err)
	}

	// Sync to ensure durability
	if err := writer.file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	return nil
}
