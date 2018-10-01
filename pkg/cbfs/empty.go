package cbfs

import (
	"fmt"
	"io"
	"log"
)

func init() {
	if err := RegisterFileReader(&SegReader{Type: TypeDeleted, Name: "CBFSEmpty", New: NewEmpty}); err != nil {
		log.Fatal(err)
	}
	if err := RegisterFileReader(&SegReader{Type: TypeDeleted2, Name: "CBFSEmpty", New: NewEmpty}); err != nil {
		log.Fatal(err)
	}
}

func NewEmpty(f *File) (ReadWriter, error) {
	r := &EmptyRecord{File: *f}
	Debug("Got header %v", *r)
	r.Data = make([]byte, r.Size)
	return r, nil
}

func (r *EmptyRecord) Read(in io.ReadSeeker) error {
	_, err := in.Read(r.Data)
	if err != nil {
		return fmt.Errorf("empty read: %v", err)
	}
	Debug("Empty data read OK")
	return nil
}

func (r *EmptyRecord) String() string {
	return recString("(empty)", r.RecordStart, r.Type.String(), r.Size, "none")
}

func (r *EmptyRecord) Write(w io.Writer) error {
	if err := Write(w, r.FileHeader); err != nil {
		return err
	}
	return Write(w, r.Data)
}

func (r *EmptyRecord) Header() *File {
	return &r.File
}
