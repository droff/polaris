package libs

import (
	"net/http"
	"os"
)

// Assets uses for loading assets
type Assets struct {
	Fs http.FileSystem
}

// Open method
func (fs Assets) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)

	if err != nil {
		return nil, err
	}

	return readDirFile{f}, nil
}

type readDirFile struct {
	http.File
}

func (f readDirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
