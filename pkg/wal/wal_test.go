package wal

import (
	"os"
	"path/filepath"
	"testing"
	//pb "github.com/zahidhasanpapon/my-wal/proto"
	//"google.golang.org/protobuf/proto"
)

func helperCreateTestWAL(t *testing.T) (string, *WAL) {
	t.Helper()
	tempDir := t.TempDir()
	walFile := filepath.Join(tempDir, "test.wal")
	w, err := NewWAL(walFile)
	if err != nil {
		t.Fatalf("NewWAL() error = %v", err)
	}
	return walFile, w
}

func TestNewWAL(t *testing.T) {
	walFile, w := helperCreateTestWAL(t)
	defer w.Close()

	if _, err := os.Stat(walFile); os.IsNotExist(err) {
		t.Errorf("WAL file %s was not created", walFile)
	}
}

//func TestAppendAndReadEntry(t *testing.T) {
//	_, w := helperCreateTestWAL(t)
//
//	entry := &pb.LogEntry{
//		SequenceNumber: 1,
//		Timestamp:      time.Now().UnixNano(),
//		OperationType:  pb.OperationType_OPERATION_TYPE_PUT,
//		Key:            []byte("test_key"),
//		Value:          []byte("test_value"),
//	}
//
//	if err := w.AppendEntry(entry); err != nil {
//		t.Fatalf("AppendEntry() error = %v", err)
//	}
//
//	if err := w.Close(); err != nil {
//		t.Fatalf("Close() error = %v", err)
//	}
//
//	// Reopen to read
//	w2, err := NewWAL(w.file.Name()) // w.file is nil after close, use its stored name
//	if err != nil {
//		t.Fatalf("NewWAL() for reopen error = %v", err)
//	}
//	defer w2.Close()
//
//	// Must seek to start to read from beginning if same file instance was used for write
//    // or if it's a fresh instance opened from existing file.
//    if err := w2.SeekToStart(); err != nil {
//        t.Fatalf("SeekToStart() error = %v", err)
//    }
//
//	readEntry, err := w2.ReadNextEntry()
//	if err != nil {
//		t.Fatalf("ReadNextEntry() error = %v", err)
//	}
//
//	if !proto.Equal(entry, readEntry) {
//		t.Errorf("Read entry %+v, does not match appended entry %+v", readEntry, entry)
//	}
//
//    // Ensure no more entries
//    _, err = w2.ReadNextEntry()
//    if err != io.EOF {
//        t.Errorf("Expected io.EOF after reading all entries, got %v", err)
//    }
//}
//
//func TestAppendAndReadMultipleEntries(t *testing.T) {
//	_, w := helperCreateTestWAL(t)
//
//	entries := []*pb.LogEntry{
//		{SequenceNumber: 1, OperationType: pb.OperationType_OPERATION_TYPE_PUT, Key: []byte("key1"), Value: []byte("value1")},
//		{SequenceNumber: 2, OperationType: pb.OperationType_OPERATION_TYPE_DELETE, Key: []byte("key2")},
//		{SequenceNumber: 3, OperationType: pb.OperationType_OPERATION_TYPE_PUT, Key: []byte("key3"), Value: []byte("value3")},
//	}
//
//	for _, entry := range entries {
//		if err := w.AppendEntry(entry); err != nil {
//			t.Fatalf("AppendEntry() error for entry %d = %v", entry.SequenceNumber, err)
//		}
//	}
//
//	if err := w.Close(); err != nil {
//		t.Fatalf("Close() error = %v", err)
//	}
//
//	w2, err := NewWAL(w.file.Name())
//	if err != nil {
//		t.Fatalf("NewWAL() for reopen error = %v", err)
//	}
//	defer w2.Close()
//    if err := w2.SeekToStart(); err != nil {
//        t.Fatalf("SeekToStart() error = %v", err)
//    }
//
//	for i, expectedEntry := range entries {
//		readEntry, err := w2.ReadNextEntry()
//		if err != nil {
//			t.Fatalf("ReadNextEntry() error for entry %d = %v", i+1, err)
//		}
//		if !proto.Equal(expectedEntry, readEntry) {
//			t.Errorf("Read entry %d: %+v, does not match appended entry %+v", i+1, readEntry, expectedEntry)
//		}
//	}
//
//    _, err = w2.ReadNextEntry()
//    if err != io.EOF {
//        t.Errorf("Expected io.EOF after reading all entries, got %v", err)
//    }
//}
//
//func TestReopenWAL(t *testing.T) {
//	walFile, w := helperCreateTestWAL(t)
//
//	entry1 := &pb.LogEntry{SequenceNumber: 1, Key: []byte("initial_key")}
//	if err := w.AppendEntry(entry1); err != nil {
//		t.Fatalf("AppendEntry() error = %v", err)
//	}
//	if err := w.Close(); err != nil {
//		t.Fatalf("Close() error = %v", err)
//	}
//
//	// Reopen WAL
//	w2, err := NewWAL(walFile)
//	if err != nil {
//		t.Fatalf("NewWAL() for reopen error = %v", err)
//	}
//
//	// Append another entry
//	entry2 := &pb.LogEntry{SequenceNumber: 2, Key: []byte("second_key")}
//	if err := w2.AppendEntry(entry2); err != nil {
//		t.Fatalf("AppendEntry() on reopened WAL error = %v", err)
//	}
//	if err := w2.Close(); err != nil {
//		t.Fatalf("Close() reopened WAL error = %v", err)
//	}
//
//	// Reopen again to read all entries
//	w3, err := NewWAL(walFile)
//	if err != nil {
//		t.Fatalf("NewWAL() for final read error = %v", err)
//	}
//	defer w3.Close()
//    if err := w3.SeekToStart(); err != nil {
//        t.Fatalf("SeekToStart() error = %v", err)
//    }
//
//	read1, err := w3.ReadNextEntry()
//	if err != nil {
//		t.Fatalf("ReadNextEntry() for entry1 error = %v", err)
//	}
//	if !proto.Equal(entry1, read1) {
//		t.Errorf("Read entry1 does not match: expected %+v, got %+v", entry1, read1)
//	}
//
//	read2, err := w3.ReadNextEntry()
//	if err != nil {
//		t.Fatalf("ReadNextEntry() for entry2 error = %v", err)
//	}
//	if !proto.Equal(entry2, read2) {
//		t.Errorf("Read entry2 does not match: expected %+v, got %+v", entry2, read2)
//	}
//
//    _, err = w3.ReadNextEntry()
//    if err != io.EOF {
//        t.Errorf("Expected io.EOF after reading all entries, got %v", err)
//    }
//}
