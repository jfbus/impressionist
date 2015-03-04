package storage

import (
	"errors"
	"image"
	"os"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

var ErrFileNotFound = errors.New("file not found")
var ErrAccessDenied = errors.New("access denied")

func Read(storage, file string) (image.Image, error) {
	s, err := Get(storage)
	if err != nil {
		return nil, err
	}
	fd, err := s.Open(file)
	if err != nil {
		log.Warn(err.(*os.PathError).Error())
		switch err.(*os.PathError).Err {
		case syscall.ENOENT:
			return nil, ErrFileNotFound
		case syscall.EPERM:
			return nil, ErrAccessDenied
		}
		return nil, err.(*os.PathError).Err
	}
	i, _, err := image.Decode(fd)
	fd.Close()
	return i, err
}
