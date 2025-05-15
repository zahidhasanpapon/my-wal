package wal

import (
	"fmt"
	"io"
	"os"
	"sync"
	//pb "github.com/zahidhasanpapon/my-wal/proto" // Alias for your proto package
	//"google.golang.org/protobuf/proto"
)

// WAL represents a Write-Ahead Log.
type WAL struct {
	file   *os.File
	mu     sync.Mutex
	reader io.Reader // Internal reader for ReadNextEntry, initialized with the file
}

// NewWAL creates or opens a WAL file.
// If the file exists, it's opened for appending and reading.
// If it doesn't exist, it's created.
func NewWAL(filePath string) (*WAL, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open/create WAL file %s: %w", filePath, err)
	}
	return &WAL{
		file:   file,
		reader: file, // Initialize reader
	}, nil
}

//// AppendEntry appends a new LogEntry to the WAL.
//// The entry is first marshaled into bytes, then prefixed with its length (uint32),
//// and then written to the file.
//func (w *WAL) AppendEntry(entry *pb.LogEntry) error {
//	w.mu.Lock()
//	defer w.mu.Unlock()
//
//	data, err := proto.Marshal(entry)
//	if err != nil {
//		return fmt.Errorf("failed to marshal log entry: %w", err)
//	}
//
//	lenBuf := make([]byte, 4) // Use 4 bytes for uint32 length
//	binary.LittleEndian.PutUint32(lenBuf, uint32(len(data)))
//
//	if _, err := w.file.Write(lenBuf); err != nil {
//		return fmt.Errorf("failed to write entry length: %w", err)
//	}
//
//	if _, err := w.file.Write(data); err != nil {
//		return fmt.Errorf("failed to write entry data: %w", err)
//	}
//	return nil
//}
//
//// ReadNextEntry reads the next LogEntry from the WAL.
//// It reads the length prefix, then the actual entry data, and unmarshals it.
//// Returns io.EOF if no more entries can be read.
//func (w *WAL) ReadNextEntry() (*pb.LogEntry, error) {
//	// Note: For simplicity, this ReadNextEntry is not thread-safe if the same WAL instance
//	// is used for both appending and reading concurrently without external synchronization
//	// for the read operations, or if multiple goroutines call ReadNextEntry on the same WAL.
//	// The file offset is shared.
//
//	lenBuf := make([]byte, 4)
//	if _, err := io.ReadFull(w.reader, lenBuf); err != nil {
//		if err == io.EOF || err == io.ErrUnexpectedEOF {
//			return nil, io.EOF // Clean EOF
//		}
//		return nil, fmt.Errorf("failed to read entry length: %w", err)
//	}
//
//	entryLen := binary.LittleEndian.Uint32(lenBuf)
//	if entryLen == 0 { // Should not happen with valid entries
//	    return nil, fmt.Errorf("read zero length for entry, possible corruption or EOF")
//    }
//
//	dataBuf := make([]byte, entryLen)
//	if _, err := io.ReadFull(w.reader, dataBuf); err != nil {
//		if err == io.EOF || err == io.ErrUnexpectedEOF { // Should indicate corruption if length was read
//		    return nil, fmt.Errorf("failed to read entry data after reading length (expected %d bytes): %w", entryLen, io.ErrUnexpectedEOF)
//		}
//		return nil, fmt.Errorf("failed to read entry data: %w", err)
//	}
//
//	entry := &pb.LogEntry{}
//	if err := proto.Unmarshal(dataBuf, entry); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal log entry: %w", err)
//	}
//	return entry, nil
//}

// SeekToStart resets the WAL's internal reader to the beginning of the file.
// This is necessary if you want to read entries after appending some, or re-read.
func (w *WAL) SeekToStart() error {
	w.mu.Lock() // Potentially lock if file operations need it, though Seek itself is often atomic
	defer w.mu.Unlock()

	_, err := w.file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to start of WAL file: %w", err)
	}
	// Re-initialize the reader to the file, as the internal state of any previous buffered reader might be stale.
	w.reader = w.file
	return nil
}

// Close closes the WAL file.
func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.file != nil {
		err := w.file.Close()
		w.file = nil // Mark as closed
		return err
	}
	return nil
}
