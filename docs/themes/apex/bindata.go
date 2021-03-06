// Code generated by go-bindata.
// sources:
// files/css/index.css
// files/js/index.js
// files/js/smoothscroll.js
// files/views/index.html
// DO NOT EDIT!

package apex

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cssIndexCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x58\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x20\x12\x14\x48\x36\x4b\x90\x9c\x38\x75\x15\xa0\x58\xd7\x26\xe8\x1e\xba\x15\xcb\x8a\xed\x95\x92\x68\x9b\x0b\x25\x0a\x14\xe5\x24\x1d\xfa\xdd\x77\xe4\x91\x32\x69\xcb\x29\x8a\xa2\xb0\x74\xbc\xff\xf7\xe3\xdd\x29\x85\x92\x52\x93\xff\x4e\x08\x49\x12\x46\x7b\x56\x90\x6a\x28\x79\x95\x94\xec\x2b\x67\xea\x3c\x5d\xce\x67\x24\x9b\x91\x34\x87\xdf\xfc\xe2\xc6\xf2\x3d\xf2\x5a\x6f\x0a\xb2\xcc\xb2\xee\x09\x29\x1b\x46\x6b\xa6\xe0\x87\xaf\x37\xba\x20\x57\x78\x62\x8f\xb4\xa2\xd5\x03\x6f\xd7\x05\xc9\xd2\x6c\xa1\x58\x73\x13\x91\x93\x86\xd5\x7c\x68\xcc\xe9\xc4\xa1\xa0\x6a\xcd\xcc\xd9\xd2\x9e\xd9\xc3\x9a\xaa\x87\x82\x9c\x65\x59\x86\xcc\xa5\x18\x80\xe5\xec\xf2\xf2\xdd\x9b\xbb\x3b\x24\x09\xe3\x46\xb2\x56\xf4\x19\x0e\x56\xd4\xfc\x73\xc2\xe5\xda\x50\x56\x2b\xe4\x5b\x99\xb7\xe5\xf5\xf2\xf6\xcd\xb5\x27\x78\xf5\xf3\x7c\xbe\x98\xbf\x71\x52\x3d\x13\xac\xd2\x5c\xb6\x56\x7e\x4b\xd5\x39\x9a\x75\xf9\xd8\x1d\x1b\x85\x8f\x1b\xae\xd9\xcd\xc9\xb7\x93\x93\x9f\x6c\x5e\x4b\xf9\x94\xf4\xfc\xab\x4d\x41\x29\x95\xc9\x13\x90\x2c\xc3\x46\x37\x62\x06\xc4\xfa\xd9\x72\xae\x64\xab\x93\x15\x6d\xb8\x00\xbf\x13\xda\x75\x82\x25\xfd\x73\xaf\x59\x33\x23\xbf\x0a\xde\x3e\x7c\xa2\xd5\xbd\x7d\xbf\x03\xce\x19\x48\x10\x72\x7a\xcf\xd6\x92\x91\x2f\xbf\x9d\xce\xc8\xe9\x9f\xb2\x94\x5a\x9a\xa7\x3f\x9e\x9e\xd7\xac\x3d\x75\x3c\x5f\xca\xa1\xd5\x83\xa1\xbf\xa7\xad\xa6\x8a\x09\x61\x5e\xee\xb8\xa2\xe4\x9e\xb6\xbd\xe7\xfb\xa0\x24\xaf\x1d\x85\x9c\x7e\x64\x62\xcb\x34\xaf\x28\xf9\x9d\x0d\x0c\x28\x3d\x1c\x40\xb0\x8a\x63\xfa\x1e\x59\xf9\xc0\xc1\x61\xe3\x75\xdf\x00\x8c\x36\x36\x46\xb0\xc0\xa9\xe0\x80\xa5\xfa\xc6\x07\x05\xf1\x43\x8d\xf2\x6b\xc4\x8b\x25\x3d\x3a\xac\x5c\x62\x1d\x4b\xa8\xf8\x5a\xc9\xa1\xad\xc7\x04\xaf\x6d\x7a\x2b\x29\xa4\xf2\x34\x57\x1f\x7b\x00\x09\x61\x23\xe2\x72\x43\x69\x00\x2d\xbc\x05\xb8\x98\x97\x8e\xd6\x35\xc2\xce\xa6\xba\x28\xc6\x2a\x61\x55\x0e\xec\x85\x45\x3e\xb4\x1c\xd6\xf8\x02\x8b\x97\xcf\xc8\x06\x6e\xc5\xe6\x12\xfe\x5f\x59\xa5\xe8\x40\xa2\x65\x57\x90\xd7\x0b\x0c\xd6\xd1\xa0\x30\x5a\x36\xce\xb7\x30\x27\xe9\xdc\xc1\x3e\xca\xca\x35\x66\x25\x8e\x71\xbc\x21\x47\x72\x62\x9d\x3a\x70\xe4\xd0\x22\xaa\x41\xee\x9f\x49\xb7\xc3\x5e\xcc\x70\x60\xfe\xda\x0a\x51\xcb\xef\x5c\xe0\xed\x06\xf0\xa0\x0d\xb3\x66\x4f\x3a\xa9\x59\x25\x15\x35\x79\x2a\x48\x2b\x5b\xbc\x07\x5d\xe0\x53\x41\xe6\x90\x18\x32\x15\xdd\xf5\x54\x68\xf6\x1d\x2f\x34\x46\x28\x38\xe9\xb5\x92\xed\x7a\x06\x7a\xf1\x29\x74\xe8\x10\x27\x51\x5e\x17\x59\xe6\xb5\x50\xa3\x80\xfe\x80\xec\x55\x16\x02\x6b\xac\xe8\x25\xd6\x79\xbc\xda\x48\xcd\x21\xc8\x1a\x9e\x59\x4d\xce\xea\xba\x1e\x8d\x16\x1b\xb9\x65\xca\x9a\xc6\xc7\x43\x07\xc6\xce\xb2\xa7\x32\xc8\x67\x28\x1b\x71\x25\x91\x26\xfb\x82\x69\x1b\x44\x54\x84\x45\x66\x8a\x80\x3f\x97\xae\x8d\xef\xdd\x18\x90\xd8\x13\xca\x51\x28\xdf\x09\x21\x1b\xc4\x15\xe9\xf6\xf5\x7d\xb9\x96\x87\xe8\x76\xa8\x44\x8d\x2f\x94\xd6\x87\x75\xa4\xb6\x9d\x62\x47\xae\xf8\x6e\x2e\x1c\x5e\xf0\xb1\xe6\x63\x1a\x7c\x5e\x5c\x82\x15\x85\x41\xd5\x03\x7e\x91\x6a\xf2\xbf\x12\xf2\x31\x79\x2a\x48\x5f\x29\x29\x84\x77\xa8\x80\xa6\x2c\x07\x55\x31\xf2\x5e\xd6\x8c\x7c\x56\xa6\x23\x7f\x62\xad\x90\x33\xd2\xc8\x56\xf6\x1d\xad\xd8\xde\xa5\x4c\x97\x93\x37\xce\xa7\x04\xf2\xf1\x16\xdc\xad\x99\x01\x0e\x3e\x05\xb5\x47\xb4\xf5\x52\x40\xf3\x3e\xfb\x70\x7b\x3b\xbf\xbd\xde\x53\x9f\xa5\xaf\xfd\x9d\xde\x85\x07\x32\xf9\x74\x88\x0e\xd2\x76\x92\x25\xd6\x5d\x83\xbe\x47\x45\xbb\x63\x8d\x2a\x1a\x5d\x63\x53\x00\xcf\x6b\xa6\x29\x17\x3d\x38\xdd\x0f\x0d\x20\x04\x27\x5d\x35\xa8\xde\xa4\xbe\x93\xbc\xd5\x4c\xd9\x7c\x0e\xda\x04\xef\x61\x4e\xc8\x00\x93\xc6\xb5\xdd\x00\xfb\x3b\x7d\x5d\x88\x7e\xc1\x56\x1a\x43\xc2\x34\x4c\xd7\xd4\xb1\xe5\x0b\x87\xdd\xf4\x6f\x08\xa9\x73\xf7\xa8\xe6\x7d\x27\xcc\xc6\xb0\x12\xcc\x86\xff\xef\xd0\x6b\xbe\x7a\x06\xb8\x81\x8f\xa6\xac\x15\x43\x67\x8d\xe4\x7b\x20\x52\xf0\x17\x65\xdd\x46\x84\x56\xed\x8b\x35\x7b\xa0\xd2\xfc\x26\x35\x57\x38\x4a\x0a\x83\xc0\xa1\x69\x77\x1a\x99\xc9\xec\x71\x9f\x0c\xdb\x47\xbb\x6a\xfd\xa0\xc7\x84\xc0\x44\x5e\xb7\x09\xd4\xb3\xe9\x43\xb2\x87\x1a\x7a\x1e\xad\x71\xd8\x37\xd2\x7b\x5e\xb3\x92\xa2\x41\x63\x07\xd2\x07\xff\xe8\xa0\x65\xe4\x76\x98\x86\xd7\x8b\x57\x78\x06\xa0\x1f\xec\x41\x27\x7b\x8e\x11\xfb\xad\x01\x1c\xad\x1e\x9e\x6f\xa2\xb3\x1d\xcd\x0e\xae\x85\x03\xe7\xf1\x69\x87\x06\xde\x92\xd4\xc4\x85\x86\xf6\x9a\xf3\xae\xd4\x11\xeb\x5b\xd7\xf8\x77\xb6\x61\x27\x82\xb1\xb5\x3d\x86\xbc\xe9\x51\x00\x7b\x6a\xeb\x15\x58\x2f\xc9\x3c\xcb\x9a\x7e\xc2\xe9\x29\x7f\xc1\x89\x94\x56\xc6\xe8\x4b\x43\x68\x4a\xac\x28\xd9\x4a\x2a\x2f\xe6\x6a\x7d\x7a\x1a\x67\x93\x96\x70\x15\x06\x6d\x9d\x77\x85\xc9\xb3\xec\x55\x58\xf5\xdc\xdf\x7e\x4c\x56\xe2\x56\x16\xbc\x25\x7b\x6b\x59\x72\x74\x48\x6e\x79\xcf\x4b\x2e\xb8\x06\x24\x6e\x78\x5d\xb3\x76\xcc\x0d\x78\xd9\x98\xf6\x48\x05\xfb\xe7\x3c\xbb\x88\xe8\x89\x54\xdc\xce\x0b\x63\x2e\x80\x64\x98\x54\x2a\x04\xec\x0b\x90\x52\x67\xd6\x7c\xa1\x1c\xcb\xc9\x91\x69\xfa\x9d\x4c\x5a\xa9\x30\x9f\x61\x34\xf6\x59\xb0\xe9\x70\x72\xa7\xf2\x33\x5d\xb3\x83\xa5\x2b\xcf\xe2\xa1\x9a\x04\x88\xf6\x42\xc5\x8a\xab\x5e\x27\xd5\x86\x8b\x7a\x7a\x6b\x8b\x84\x71\xb8\xa5\x7f\x71\x2d\xd8\xf4\xbc\x8d\x86\xc7\x3c\x9d\xbb\xe1\x81\x22\x29\x66\x38\x34\x84\x75\xc6\xab\x3c\x12\x15\x8a\x7b\xaa\xdd\xea\x6c\xf3\x08\xdb\x46\x43\x9f\xfc\x57\xe0\x22\x1b\x83\x42\x3b\xa8\x27\xb4\xe3\xf1\xb5\x8c\x39\xcd\x4c\xe8\x68\x1b\xb7\xb2\x52\xc8\xea\x21\x64\x4a\xfb\xa1\x34\x4e\x1c\xdf\x04\x82\x8d\x22\x1a\x79\x4b\x37\xf1\x6c\x08\x41\xfd\x06\xd3\x5e\x2b\x40\x52\xd4\xa2\xc7\x09\xe3\xed\x8e\x46\x05\x83\x35\x4e\xd9\x39\x68\x07\x27\x5a\xf7\xdf\xa8\x17\xdf\x35\x11\x75\x8e\x52\x8a\xe9\xcf\xa2\x5d\x02\xfb\xc6\xc0\x7e\x67\x3f\x64\xbd\x42\x54\xed\xdb\xf3\x4d\x6a\xdf\xd5\x16\x4e\xa9\x38\xc4\x86\xff\x00\x48\xdf\xb5\xd5\x46\xaa\x17\x7a\x61\x04\x95\xc4\xdb\x97\xc6\x80\xb9\x21\x59\x9a\x2f\xa2\x34\xf2\xd6\x5a\x72\x65\xdc\x75\x1e\x27\x38\x76\x9e\x2b\x1f\xb3\xf3\xa0\xdf\xae\xf7\xbc\x08\xfb\x97\x03\xa5\xeb\x4f\xe1\x75\x08\xc4\x3b\xaa\x37\x4e\x87\xdd\x2a\x12\xb6\x05\xc0\xf6\xb6\x8b\x04\xbc\x41\xa3\x18\xc3\xc8\x51\xd9\x1d\x7c\xc4\xba\xa3\xd1\xd1\x11\xde\xbf\x98\xbf\x54\x50\xb3\xeb\x31\xd6\xc2\x57\x6e\x4d\xce\x83\x6b\xb0\x34\x77\xfb\x02\x45\xe3\xaf\xfa\x89\xbf\x8f\x5c\xfa\xee\x60\xce\x0e\xf6\xcd\xc9\x8a\x7f\x33\x7f\x89\xd8\xdb\x3a\x8e\xef\x13\xe6\x6c\xaf\xe7\xa3\x82\x71\x22\x1f\xab\xf7\x7e\x1b\xf3\x2e\xed\x3e\x0e\x1c\xfe\x1d\xe1\x22\x76\xce\x6f\x02\x07\xe6\xbf\xab\xe1\xdb\xc9\xff\x01\x00\x00\xff\xff\xaa\xb9\x74\x4f\x8b\x12\x00\x00")

func cssIndexCssBytes() ([]byte, error) {
	return bindataRead(
		_cssIndexCss,
		"css/index.css",
	)
}

func cssIndexCss() (*asset, error) {
	bytes, err := cssIndexCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "css/index.css", size: 4747, mode: os.FileMode(420), modTime: time.Unix(1505956668, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsIndexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\x4d\x6f\xda\x40\x10\xbd\xfb\x57\x4c\xd5\x83\x6d\x52\x2d\x9c\xdb\x12\x89\x44\x39\x44\x6a\x7b\x28\x52\x2f\x08\x89\x65\x3d\xe0\x15\xeb\x1d\xba\x5e\x83\x50\x95\xff\xde\xd9\x0f\x42\x83\xa8\x82\x64\xb4\xb6\xdf\xbc\x79\xef\xcd\x78\x3c\x1a\x15\x30\x82\x39\xfa\x61\x2f\xf8\x34\x2e\x0a\x45\xb6\xf7\xa0\x3d\x76\x3d\x4c\xa1\x21\x35\x74\x68\xbd\xf8\x3d\xa0\x3b\xcd\xd1\xa0\xf2\xe4\x66\xc6\x54\xe5\x42\x37\xcb\xb2\xce\x78\xa3\xed\xee\x1d\xbc\xf8\x8e\x76\x00\xc9\x25\x45\x6e\xfb\xd8\xa2\xda\x81\xde\xc0\x0a\xcd\x0a\x74\x0f\x34\xf8\x74\x6d\xe0\xa0\xf1\x98\x15\x6d\x06\xab\xbc\x26\xcb\x88\x07\x34\x74\x9c\x2b\x47\x4c\x88\xa6\x86\x3f\x05\x80\x63\xf1\xce\x02\x1a\xb1\x45\xff\x40\x83\x6d\xb4\xdd\x3e\x1a\xcd\x2a\x7e\x72\xf7\xaa\x16\x6b\xf2\x9e\x3a\xb8\x87\x49\xf1\xf2\xda\x7c\xc6\x9c\x07\xe9\x31\x5a\x85\x95\x5e\x5d\x77\x93\x19\xf0\xcc\xef\x2b\x9d\x7a\x45\x9b\x62\x43\xee\x49\xaa\xb6\x42\x98\xde\x03\x0a\x65\x64\xdf\x7f\xd3\xbd\x17\x0e\x3b\x3a\x60\x55\xc6\x52\x2c\xeb\xfa\x5c\xb2\xd0\xcb\x7f\x60\xb2\x69\x2e\x98\x5b\x92\x7c\x8b\xa0\xc8\x39\x96\x0f\x5d\x48\x2d\x6a\xe4\xb6\xe1\x4d\x40\x72\xe8\x9e\xfd\xf5\xa0\x6d\x04\x87\xb4\xf6\xe4\xfc\xff\x3c\x54\x59\x3e\xf2\x64\x79\x4a\x93\x82\x6f\x02\x5d\xf5\x85\xef\xbf\xa6\x69\x0b\x83\x76\xeb\x5b\x7e\x72\x77\x97\xe0\x10\x66\x53\xbd\x4d\x3d\x42\xd9\x4e\x7d\x86\x00\xac\x1d\xca\x5d\x3c\xbf\x14\xe1\xe2\xbf\xab\xec\x6e\x79\xec\x23\x1f\xf4\xfb\x13\x3b\xe0\x89\x9d\xb2\xf6\xa3\xb6\x0d\x1d\x43\x46\x4f\x07\xb6\x18\x02\x43\x8b\xae\x2a\x53\x41\xf9\x09\x62\xec\x17\x67\x97\x7d\x9a\x35\x0d\xf4\x1d\x91\x6f\x33\x3b\xd3\xbe\xc7\xaa\x8c\x56\xbb\x33\x69\x70\x94\xf6\x19\x0d\xc7\x84\xc2\x4b\xc7\x4b\x15\x1c\x8d\xc7\x79\xc7\xc9\x9a\x53\x91\x92\xe1\x95\xb3\xd4\xe0\x0f\xd9\x21\x7c\x98\x42\x39\x2b\xeb\xbc\x8d\xb9\x62\x70\xe6\x95\x91\xcf\x81\x32\x6e\xe9\xcc\x7b\xa7\xd7\x03\xab\x2f\x5b\x87\x9b\xf0\x49\x44\xbc\xb4\xaa\x25\xf7\xa6\x07\x97\x2d\x26\xcb\x48\xff\xf1\x9a\x3e\x99\xcc\x3f\x7e\x74\xfb\xe3\x0b\x14\xb5\x48\xd8\x67\xeb\xe9\x17\xef\x4a\x95\x66\xb7\xc6\x56\x1e\x34\xb9\xcf\x50\xa6\xdc\xca\x30\x40\x1e\x57\x5d\xfc\x0d\x00\x00\xff\xff\x79\x82\x0e\x63\x1a\x04\x00\x00")

func jsIndexJsBytes() ([]byte, error) {
	return bindataRead(
		_jsIndexJs,
		"js/index.js",
	)
}

func jsIndexJs() (*asset, error) {
	bytes, err := jsIndexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/index.js", size: 1050, mode: os.FileMode(420), modTime: time.Unix(1504199128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsSmoothscrollJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xd8\xbc\x9c\xa4\x9c\x42\xd9\x6d\xef\xa6\x23\x55\x6d\x4f\x89\x5a\xbb\xb5\xa3\x1b\x4b\x93\xda\xd3\xb9\xe9\xd0\x22\x24\xe1\x42\x11\x3a\x12\xd4\x9f\x89\x7d\x9f\xbd\x8b\xbf\x04\x40\xd2\x39\x4f\x33\x93\x97\xe6\xc5\x04\x76\xb1\xd8\x5d\xfc\x76\xb1\x58\x65\xf0\xfa\x0c\x5e\x43\xb1\x65\x8c\x6f\x8a\x65\xce\xd2\x14\x76\x2c\x3d\xad\x28\x7e\xbc\x81\xfd\x79\xf4\xfb\xe8\x3b\xc1\xb1\xe1\x7c\x57\x0c\x07\x03\x1a\x6f\x93\xb2\xe0\x71\x16\xad\x29\xdf\x94\x0f\x11\x65\x03\x77\xb5\xe0\xfd\xdd\xf9\xc5\xf7\xd0\x5d\xf6\xe0\x9d\xe4\x84\x7f\xc6\x05\x27\x59\x1f\xfe\x41\x72\xb2\xa5\x71\x01\x37\x24\xa3\xcb\x0d\x49\x53\x8a\x7b\xdc\x5c\x2d\xe0\x9a\x2e\x49\x56\x10\x5c\x3c\x38\x3b\xeb\xae\xca\x6c\xc9\x29\xcb\xba\x87\x3e\x24\x7d\x28\xb3\x84\xac\x68\x46\x92\x1e\x7c\x3a\x03\xe8\x94\x05\x81\x82\xe7\x74\xc9\x3b\xa3\x33\x9c\x10\x26\x00\x6e\x1b\xa7\x28\x9b\x14\x6a\x70\x18\xc2\x81\x66\x09\x3b\xc0\x3a\x65\x0f\x71\x0a\xec\xe1\x67\xb2\xe4\x8a\x98\x0c\x21\x61\xcb\x72\x4b\x32\x3d\x61\xb7\x18\x56\x9f\x92\x32\x90\x1b\x0c\xac\x4f\x70\x64\xb4\xb3\x73\x5d\xa5\x97\xe4\xcb\x09\x2f\xf3\x0c\x0e\x1b\x92\x81\x72\xc8\x84\x6c\xe2\x3d\x65\x39\xd0\x8c\x93\x7c\x15\x2f\x09\xd0\x02\x8a\x72\xb7\x63\x39\x57\xbb\x00\x5d\x41\xb7\xe3\xb3\x77\x90\x1f\x92\xc8\xa8\x39\x4d\x89\xf8\x13\x15\xfc\x94\x12\xb3\x1f\xe8\xed\x46\x72\xf4\x74\xa6\x94\x78\xad\x68\xaf\xb5\xe1\x85\x1e\x0e\xe4\xdf\x7d\x9c\x83\x96\x05\x63\x38\x44\x97\x8b\x9b\x6b\x33\x7e\x7c\xc4\x09\x3d\x18\x59\xee\xf9\xdb\xdb\xd9\xf5\xf5\x7f\x16\x57\x37\x53\x5c\xf1\x87\xef\xff\x38\x0a\xf7\x51\x9e\x85\x75\xcc\x37\x24\xa7\xd9\x1a\x58\x4e\xd7\x34\x43\x9f\x6b\x40\x6d\x09\xdf\xb0\xa4\xae\x88\xe5\x1b\x5b\x83\xd4\x0a\x3c\xbb\x48\xaf\x95\x4a\xa9\xef\x05\xeb\x7b\x5c\x93\x53\xc5\x37\x39\x19\x1a\x49\xe7\x5a\x86\xf1\xd9\x2e\x67\x9c\xf1\xd3\x8e\x38\x32\xd5\x97\xe6\xf0\xc5\x5e\x65\x9c\x7d\xa0\xe4\xd0\x2e\xc0\x70\x28\xb7\xd7\xfc\xa1\xf0\x03\x9c\x6e\x85\x33\x94\xf1\x35\xdb\x33\x44\xa6\x38\x80\x1d\x62\x82\xe5\xdb\x38\x43\x5c\x7c\xf3\x8d\x3f\x11\x21\x93\x56\xed\x2f\x75\x4a\xf4\x80\xf8\xee\x7a\xd3\x3d\x18\xc2\xbb\x98\x4b\xb2\x55\xcb\xea\xb5\xdc\xc4\xd9\x9a\x14\x60\xc3\xbc\xa0\x12\xc7\x34\x2b\x68\x42\x00\x03\x95\x28\x83\xcd\x82\xbf\x2a\xdd\x7d\x67\x59\xe2\x2e\xce\xe3\x2d\x7c\x7a\x5f\x6e\x1f\x48\xfe\x04\xc7\x36\xc2\xc9\xb3\xdd\x46\x8f\x27\xb4\x7b\xec\xc3\xa9\x82\x35\xdf\xd0\x42\xfb\xfa\x9a\xac\x04\x52\x8f\xa3\x3a\x69\xc1\x76\x48\x39\x05\xe8\xb7\xe6\xaa\xd8\x28\xf0\x6f\x51\xa6\x1c\xd8\x0a\xe2\xdd\x2e\x3d\x89\x33\x21\x98\x27\x60\x8b\x78\xad\xd4\xe1\x0c\x62\xc8\xa4\xca\xa1\xf9\x82\xbb\xcd\xb8\x8f\x96\x60\xb6\x33\xa4\x66\xab\x85\xac\xee\xc7\x30\x80\xe1\x3c\xfa\x0e\x65\x74\x2f\x44\x3a\x44\xb5\xa2\x25\x2b\xba\xf2\xe3\xc7\x2b\x9c\xff\xd8\xeb\xb5\x19\x89\x10\xa0\x4b\x3c\xf1\x42\xe4\x90\x58\xe7\x71\x78\x30\x39\xa7\xd8\xb0\x32\x4d\x70\x2c\x6d\xa7\x24\xa9\x1d\xad\x64\x98\xc4\x34\x9d\x95\x2d\x47\xfb\x38\x93\xf1\xed\x9e\xb0\xb5\x75\xc2\x58\x4a\xe2\xac\xc5\x58\x4f\x78\xf7\x58\x59\x2d\xf2\x9d\x88\x27\x3c\x93\x23\xbc\x1a\x8f\xa1\xa3\x72\x48\x47\xd3\xd5\x3f\x0c\xd3\x23\x8c\x91\x9a\x95\x32\xf1\xfa\x94\xc8\x1a\x29\x58\xbc\xac\xfd\x0c\x5f\x27\x2e\x39\xab\x6f\x13\x30\x61\x44\xe0\xbd\xc5\x3b\x95\xc6\x32\xb9\xaf\x68\x5e\x70\x88\xf3\x35\x46\x2f\x17\xf1\xa2\xb4\x1e\x78\xfa\x21\x1f\x8a\xb1\xf2\x30\xd7\x8b\x2d\xfb\xa0\x65\x0a\x62\x5d\x59\x0d\x03\x9e\x97\xc4\x00\x5d\x9f\x75\xe0\xac\x71\x9b\xb3\x30\x77\x84\x56\x28\x30\xb4\x1a\x21\xef\x15\xa9\x9f\x31\x04\xbf\x12\x4f\x73\x25\x21\xd4\x72\x85\x97\x4a\x5d\x4d\x14\xcd\x37\x39\x66\x35\x92\xe7\xb8\x5a\xde\x80\xae\x2c\xe1\x32\xff\xd6\x03\xbd\x20\x23\x07\x58\xa0\x81\x53\xb1\xb0\xdb\xb1\x8b\xc4\x8a\x3d\x5e\xea\x49\xa7\x15\xfe\xe8\xc5\xc4\x24\xb4\xf8\x21\x25\x80\xc8\x15\x56\x89\x60\x6f\x4d\x67\x62\xd1\xdc\x2e\xf9\x51\xae\xa8\x41\x9f\x25\xe4\x09\x05\x34\x84\xb7\x4f\x09\x20\xdf\x24\xbb\x4b\xd2\xea\x0c\x44\xee\xa7\xc5\x84\x25\xa7\x91\x33\xb3\x89\x8b\x6a\xd5\x7c\x87\x45\x42\x40\xfd\x40\x0b\x8a\xa4\xd9\x1e\x93\x7d\x6a\x93\x3b\x60\x19\xe3\x9c\x2e\x11\x37\x29\x49\x23\xe5\x04\xa1\xa8\xe5\x93\xe7\x53\x10\x0e\x4b\x86\x19\x43\xaa\x8a\x92\xa9\xd8\xae\xb0\x2c\x4a\x2f\x29\x43\x22\x28\x89\x1e\x1c\x3d\xa1\x41\x4b\x18\x3b\x20\xc4\x9d\x97\x98\x63\x32\x7e\x49\xe8\x7a\xc3\xe1\x4f\x62\x46\x9d\x8d\x9e\x79\x7c\x6c\x62\xff\x17\x4d\x30\x67\x39\xdc\x72\xc2\xdb\x36\x30\xdf\xdb\xf6\x10\xad\x09\x7f\xcb\xb6\xbb\x12\x81\x35\x17\xf5\x11\x3a\xbc\x2f\x93\x46\x2f\x62\x76\x85\x88\x88\xbd\x12\xd3\xb1\xe0\x45\x98\x52\x44\x4d\xf7\x95\x36\x1d\xa3\xe8\x55\xb7\xc1\x4c\x31\x5f\x57\x43\xa4\xe5\xb3\xc0\x75\x4d\x3e\x6a\xb4\x40\x2a\x68\xd7\xeb\xd0\x22\x69\x1b\xd2\x0b\x92\xae\x30\x87\xec\xd9\x47\x92\x38\xb7\xd6\x26\xe6\x7d\x58\xd3\x3d\x06\x5b\x2c\xce\x96\x93\x23\x4e\x60\xb9\xbd\x2b\x64\x74\x95\xeb\x8d\x0e\x0f\xbc\xf7\x6a\xc9\x1f\xd9\x42\xe0\x9b\x64\xaf\x65\xb5\x24\x76\x5c\xd8\xd5\x1c\x3e\xb4\xb1\xe8\x11\x06\x63\xfd\xd1\xed\xb9\x10\xc6\x38\x2e\x3d\x4c\x2f\xcb\x5c\x80\xf4\xae\x61\xee\xde\x9d\x23\x69\xbc\x2b\xd0\xe4\x31\x26\x41\x21\xfb\x8d\xd1\x0c\x4b\xe1\x38\xe7\x0b\x9c\xeb\xc1\xc0\xad\x52\x47\x4e\x46\x8a\xf7\x8c\x26\x56\x86\x10\x50\xc0\x06\xa1\x48\x72\xe1\x3a\xcc\x7b\x19\xb1\x65\xa3\xd9\xc7\x7c\xfd\x19\x2e\xb0\xec\xba\xc0\x92\x4a\xcf\x78\x82\x45\x29\x21\x6e\x73\x51\x4e\x60\xed\xe0\x6e\x61\xb5\x47\x93\x85\x3c\x71\xe5\x6b\x7a\x05\x18\x63\x3f\x32\x78\x06\xdd\xc1\xb7\x60\x5c\x1b\x1d\x43\x73\xef\x7a\x78\x54\x9e\x2f\x8d\xcf\x42\x39\xf7\xae\x9c\x53\x28\xe7\xde\x91\x63\x04\x69\xba\xc2\x46\xb4\x8c\xf1\x61\x63\xd7\x58\x44\xf7\xad\xe2\xf6\xeb\xbe\xe7\x3a\xc6\x94\xfc\x2c\x27\xe2\xe6\x3a\x10\x04\xff\x9e\xc8\x5c\x9e\x93\x18\x1f\x7d\x09\xb0\x32\xc7\x1a\xb9\xe0\x58\xfa\x0b\x34\x39\xd7\x9c\x75\x8a\xa8\x09\x2a\x27\xe0\x0d\x6d\xcd\x74\x29\x27\xf7\x6a\x3b\x44\x39\xf9\xa5\x44\xb1\x3f\x64\x74\x2b\x05\xff\x0d\x01\x4d\xba\x02\xab\xba\x54\xee\x9b\x95\xbd\x5e\x75\x7d\xb5\x84\x9b\x34\xa3\x30\x6f\xc8\x03\xbe\x75\xeb\xe5\x55\x2d\x9e\x24\x79\xae\xdf\xc1\x4d\x71\xf5\x58\xbb\x57\xfe\xd7\x2a\xda\xd9\x52\xa6\x3d\xbf\x90\x16\x11\x54\x1d\x9e\x1b\x57\x0a\x4e\xb5\x19\x2f\xf6\x94\x59\x35\x9e\x85\x17\xe3\xd5\xd1\xeb\x77\x8f\x46\x80\x97\x40\xe4\xe9\x7a\xb7\x8a\x7b\x74\xce\xfd\x8d\x4f\xa2\x2a\xf1\xeb\x80\x18\xdb\x57\xde\x9d\x7a\x0e\xee\xe2\x35\xb9\x9b\xad\x56\x78\x9f\x05\xcc\xf7\x0e\xf3\x7d\xc5\x7c\x1f\x32\xeb\x03\x1b\xdb\x27\xa8\x5e\x53\x5d\x0c\x04\x6b\x9c\x36\x15\x49\xda\xa0\xa3\xbd\xbd\xc4\x8b\xa5\x41\x2d\x4b\xc7\x67\x4b\x83\x22\xde\x73\xa8\xa9\xb8\xd2\x6e\x4d\x19\xdb\xc9\x37\x36\x5e\x24\x88\xc8\x95\xc0\xb8\x79\xbe\x8a\xac\xdc\xa4\xf2\xd0\xf9\xee\x07\x5b\x0f\xf5\xdf\xbe\xaf\xb1\x38\xe3\x61\xf5\x19\x50\xef\x34\xe9\x2e\x98\xbf\xd7\xf3\xf7\xd5\xfc\x71\x88\x90\xb4\x23\x7c\xb1\x9f\x8c\x6d\x61\x51\x67\x90\x3f\xbb\xbd\xfa\xfb\xd5\xfb\x1f\xae\xe1\x66\xba\xb8\x9c\xbd\x9b\xc3\xec\xc3\xf4\xf6\xf6\xea\xdd\x74\x6e\x43\xe0\x4c\xfb\xc4\x76\x09\x44\xe1\x5a\xb5\x09\x24\xd9\xd2\xc6\x0e\x05\x07\xb6\xb9\x54\x21\xd0\x5e\x14\xe1\xf3\x09\x51\xab\x12\xd7\x2f\x25\xcd\x6d\xd5\x2a\xb0\xec\x3f\x6d\x4c\x35\x5d\xfc\xfb\xfc\xa7\x9e\x8b\xec\x00\x5e\x2a\xab\xba\xb5\x4b\xdf\x19\xb8\x52\xa2\x54\xbc\x7b\x11\xc0\xee\x64\x2b\x33\xc7\x97\xb0\xc7\x7b\xf1\x93\x65\xed\x8d\x82\xfa\xbd\x09\x5c\xd7\xd3\x05\x2c\x2e\xa7\x30\xbf\x99\xcd\x16\x97\xef\xa7\xf3\x39\x4c\xa6\x78\x0a\xaf\x0c\xb4\x9c\x34\x13\x18\xe1\x98\xa0\x22\xbb\x1a\xff\xfa\x6b\xcd\xa4\x56\x22\x9a\x70\xe6\x29\x6c\xfb\x2a\x03\xa7\xc9\xe3\x9d\xec\xe4\xf4\x95\x8f\x73\x72\xfa\xff\x81\x8a\x32\xa3\x3d\x3d\xf7\x9e\x3b\x6f\x77\x69\x3d\x59\xf7\x5a\xe1\xd0\xda\xd7\x13\x59\xa0\x8d\xa8\xb3\x42\xeb\xda\xf1\x33\x2b\xbf\x1a\xcc\x4c\x2f\xb3\x06\x33\xd5\x07\xeb\x7b\x33\x2f\x42\xdb\x17\xc5\x9b\xa8\x0c\x52\xd5\xa7\xab\x29\xe1\x56\x0f\x5c\x36\xec\xc2\x8d\x47\x5f\x06\xb6\x75\x97\xd4\xc6\xaa\x89\xa2\x34\x15\xcf\x41\xd5\xee\xeb\x60\x95\x2f\xe7\x86\x61\xdf\xb1\x61\xb1\x34\xc1\x5f\x2b\xa6\x86\x41\x5f\xf2\xc5\xd8\xd5\xa9\xad\x9d\xdc\x8c\x41\xe1\x55\xf4\xe7\x79\xe0\xd6\x51\x43\xef\x48\xb1\x39\xed\x23\x17\x71\x8e\xf6\x4e\x05\x01\xd2\x2f\x43\xb9\xd4\x04\xfb\x73\x2e\x62\x3b\xcd\xab\x82\xdb\x77\x89\xcb\x69\x42\x45\xb3\xfb\xf5\x34\x54\xb5\x41\x43\x21\xe6\x6a\x1a\x9c\xb5\xdb\x31\xfe\x56\x4a\x0e\xd1\xe0\x34\x8e\xbf\xfd\x0c\xd8\x9f\x7e\xeb\xc9\x79\x3f\x06\x7c\x8e\xe9\x2b\x5f\x58\x46\x8d\xcf\x86\x8e\x2b\xd0\xef\xaf\x0a\xc0\xe7\xf8\xae\x1d\x7a\x3c\x5f\xf6\x72\xf2\x5f\x2d\xaa\x89\x26\x3c\xd7\xd4\x5b\x13\x9a\x7b\x9d\x06\xd5\xfc\xba\x45\x84\x17\xb6\xac\xae\xf8\x45\xaf\x68\xc2\xd0\x1a\xac\xa1\xdf\xca\xee\x93\xe0\xf4\x7b\x15\x4b\x3b\x2f\x24\x48\xdc\xb4\xae\x72\x8f\x24\xd4\xf8\x55\xe3\x63\x47\xfe\x48\xb8\x27\x71\x6a\xda\x93\xe6\xf7\x97\x9d\xd3\x88\x84\x17\xa6\xba\x70\xef\xe7\x68\x7e\x98\x38\xc6\xaa\x10\x7f\xe3\x7a\x30\xa8\xd9\x5a\x85\xa9\x88\x72\x65\x89\x14\xe0\x8b\xaa\x2a\x3c\x0f\x26\x95\x3f\x74\xef\x96\x66\xb0\x47\x8c\x8a\x46\x71\x55\x93\xd8\x3c\xd8\x90\x9e\x9e\xd3\x57\xe6\xa4\x40\x8b\xe6\x4c\x64\xda\xe5\xbf\x25\x09\x35\x9d\xe1\x0b\x95\x0e\xfd\x5e\x53\x3a\x70\xe6\xcb\x95\x36\x09\x4c\xc7\x9e\x73\x15\x90\xa3\xd0\xb3\x68\xbc\x0d\xd0\xb4\x25\xdb\x6e\x59\xf6\xb3\xea\x0b\x6f\x59\x52\xa6\x24\xb2\x4b\xe0\x93\xfd\xdd\x7b\x58\xfd\x4f\x01\xb5\x8d\xeb\x26\x94\xa3\x7e\x7f\x96\xa3\xea\xa7\x72\xa5\xcf\x53\xaf\xab\x3a\x2b\x7d\xfb\x43\x3c\x52\xfe\x1b\x00\x00\xff\xff\x13\x45\xe8\xd6\x8c\x20\x00\x00")

func jsSmoothscrollJsBytes() ([]byte, error) {
	return bindataRead(
		_jsSmoothscrollJs,
		"js/smoothscroll.js",
	)
}

func jsSmoothscrollJs() (*asset, error) {
	bytes, err := jsSmoothscrollJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/smoothscroll.js", size: 8332, mode: os.FileMode(420), modTime: time.Unix(1504199128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xc1\x6e\xd4\x30\x10\xbd\xf3\x15\xc6\x5c\xc9\x46\xbd\x71\xb0\x7b\x59\x40\x5c\x10\x95\x5a\x09\x71\xf4\xda\xd3\xb5\x8b\xe3\x84\x78\xb2\xbb\x55\xc9\xbf\xe3\xc4\x1b\xc9\x76\x02\xf4\xb4\x33\x7a\xef\x8d\xe7\x3d\xc7\xcb\xde\x7e\xfc\xb6\x7f\xf8\x71\xf7\x89\x68\x6c\xec\xed\x1b\x16\x7f\x08\x61\x1a\x84\x9a\x8a\x50\x36\x80\x82\x48\x2d\x7a\x0f\xc8\xe9\x80\x8f\xd5\x07\x9a\x42\x1a\xb1\xab\xe0\xd7\x60\x4e\x9c\x5e\xaa\x41\x54\xb2\x6d\x3a\x81\xe6\x60\x81\x12\xd9\x3a\x04\x17\x74\x06\x38\xa8\x23\x64\x4a\x27\x1a\xe0\xf4\x64\xe0\xdc\xb5\x3d\x26\xe4\xb3\x51\xa8\xb9\x82\x93\x91\x50\xcd\xcd\x7b\x62\x9c\x41\x23\x6c\xe5\xa5\xb0\xc0\x6f\x96\x41\x68\xd0\xc2\xed\xcb\xcb\xee\x61\x2a\xc6\x91\xfc\x26\xa1\xb9\x1f\x0e\x18\x7b\x56\x47\x46\x64\x5b\xe3\x7e\x92\x1e\x2c\xa7\x1e\x9f\x2d\x78\x0d\x10\xce\xd5\x3d\x3c\x72\x8a\x1a\x1a\xa8\x45\x07\x97\x5a\x7a\x5f\x1b\xa7\xe0\xb2\x0b\xd5\x7c\x12\xab\x97\x44\xd8\xa1\x55\xcf\xd7\x71\xca\x9c\x88\xb4\xc2\x7b\x4e\xbf\xf7\xa2\xeb\xa0\xbf\xae\x95\x63\xfb\xe0\x4b\x18\x97\xa0\x39\xfe\x25\x8c\xce\xc0\x1c\x9e\x9d\x11\x19\x92\x29\x48\x81\xe6\x3b\xe1\x16\x1e\xc2\x05\x69\x12\x05\xab\x27\xf4\x1f\x02\x1f\x52\x5a\x34\x69\x62\xa5\x8c\xd5\x61\x9b\x64\xf5\xd8\x6e\x5a\xd9\xc7\x2b\xac\xce\x45\x1c\x25\xef\xde\x28\x38\x88\x95\x9d\x84\xf1\x15\xdc\x50\xc0\x24\x5c\x6d\x2f\xdc\x11\xc8\xee\x4e\x1c\xc1\x8f\x63\x01\xe7\x13\x0c\x42\xb3\x9a\x30\x93\xc4\xf5\xca\xdf\x4d\xc6\xed\x70\x1c\xc7\x3c\x37\xb1\x56\x15\x19\x2c\xeb\x80\x53\xc5\x16\x2b\x62\x19\xd7\x66\x60\xc5\x9e\xff\xf1\x99\x0e\x98\x18\x94\x18\xc5\x69\x62\x66\xbd\xbe\xbe\xc9\x1c\x86\x76\xc5\x09\xf8\x75\x9b\xf5\x81\x6b\xf7\x9b\xde\x93\xbd\x3e\xb7\xed\xfc\xbd\x6e\x48\x99\x97\xbd\xe9\x90\xf8\x5e\x66\xcf\xee\xc9\xd7\xbe\x09\x32\x1d\xf0\xd6\xda\xdd\x93\x9f\xe4\x91\xfc\xea\x09\xf1\xdd\xfe\x4d\xba\xfd\x29\x67\x25\xab\xe3\x0b\x0f\x21\xcd\xff\x86\x7f\x02\x00\x00\xff\xff\x91\x39\x2f\x9b\x25\x05\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 1317, mode: os.FileMode(420), modTime: time.Unix(1504741248, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"css/index.css": cssIndexCss,
	"js/index.js": jsIndexJs,
	"js/smoothscroll.js": jsSmoothscrollJs,
	"views/index.html": viewsIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"css": &bintree{nil, map[string]*bintree{
		"index.css": &bintree{cssIndexCss, map[string]*bintree{}},
	}},
	"js": &bintree{nil, map[string]*bintree{
		"index.js": &bintree{jsIndexJs, map[string]*bintree{}},
		"smoothscroll.js": &bintree{jsSmoothscrollJs, map[string]*bintree{}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

