// ReadObjectStorage defines the interface for storage operations.
package objects

import "io"

// ObjectListable is responsible for listing objects by storage location.
type ObjectListable interface {
	ListObjects(readerCloser io.ReadCloser, location []byte) (objects [][]byte, err error)
}

// ObjectReadable is responsible for reading files.
type ObjectReadable interface {
	OpenFile(filename string) (io.ReadCloser, error)
}

// ObjectWritable is responsible for writing files.
type ObjectWritable interface {
	CreateFile(filename string) (io.WriteCloser, error)
}

// ObjectRenamable is responsible for renaming files.
type ObjectRenamable interface {
	Rename(oldPath, newPath string) error
}

// ObjectDeletable is responsible for deleting files.
type ObjectDeletable interface {
	DeleteFile(filename string) error
}

// ObjectStorage combines all file operations into a single interface.
type ObjectStorage interface {
	ObjectListable
	ObjectReadable
	ObjectWritable
	ObjectRenamable
	ObjectDeletable
}
