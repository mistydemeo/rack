// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3c\xfd\x73\xdb\x38\x76\xbf\xe7\xaf\xc0\xa0\x69\x65\xef\xca\xb2\xec\xed\x5e\x7b\xdc\xa6\x33\x8e\xec\x24\xbe\xb3\x13\x55\x72\x92\xe9\x25\x9e\x1b\x9a\x84\x24\x9e\x29\x80\x0b\x40\x8e\xbd\x1a\xfd\xef\x1d\x7c\x90\xc4\x27\x25\xdb\xc9\x5c\x73\x73\x9b\x88\x7c\x78\x78\x78\x78\xdf\x78\xe0\x7a\x0d\x72\x34\x2b\x30\x02\x30\xad\x2a\x08\x36\x9b\x17\x00\xac\xd7\xe0\x65\x5a\x55\x20\x79\x05\x06\x27\x55\xd5\x3e\x5c\xa6\xb8\x98\x21\xc6\xe5\x9b\xcb\xfa\x87\x7a\xfd\x02\x00\x00\xe0\xc9\xe7\xe9\x15\x5a\x56\x65\xca\xd1\x1b\x42\x97\x29\xff\x84\x28\x2b\x08\x86\x20\x01\xf0\x78\x78\x34\x3c\x18\xfe\xf9\x60\xf8\x67\xd8\x57\xe0\x23\x82\xf3\x82\x17\x04\x33\x98\x68\x14\x72\x26\xae\x71\x00\x78\x93\x96\x29\xce\x10\x3d\xc8\x5a\x50\x77\x6e\x6f\x50\x45\x49\x86\x18\xdb\x36\x06\x9e\x63\x8e\x28\x4e\x4b\x31\x39\x80\x6f\x70\x92\x9c\xfd\xbe\x4a\x4b\x41\xcc\x17\xf1\x64\x82\x66\x30\x31\xc0\xc0\xa6\x0f\xe0\xff\x22\x06\xc1\x35\xd8\xf4\x6b\x2c\x63\x5a\xdc\xa5\x1c\x6d\x41\x52\x43\x85\x71\xbc\x2e\x53\x7c\x3b\x45\xd9\x8a\x16\xfc\xe1\x2d\x25\xab\x4a\x70\x6c\x6d\xa2\x03\x09\xf8\xb2\x96\xd8\x04\x2f\x6d\x58\x81\x13\x5e\xab\x75\x69\xa4\x70\x9c\xd2\x74\x89\x38\xa2\x72\x68\x37\x73\x2b\x01\xfb\x08\xc6\x06\xe1\xeb\xb5\x8c\xca\x15\xe3\x88\x1a\x3b\x0a\x00\xbc\x7a\xa8\x90\x22\x9c\xd3\x02\xcf\x61\xbf\x7d\x75\x8a\x66\xe9\xaa\xe4\xf2\xad\xfd\x9c\x65\xb4\xa8\x78\x2d\x3e\x50\xbf\x6a\xb9\x36\x5a\x31\x4e\x96\x57\xa4\x2a\xb2\xc9\x0a\xf3\x62\x89\x02\x93\x76\xce\x99\x00\x88\x49\x8e\xfe\xc1\x62\x13\x07\xe7\x3d\x45\x55\x49\x1e\x96\x08\xf3\xcb\xf4\xbe\x58\xae\x96\xe1\x69\xdf\xaf\x96\x37\x88\x46\xa6\x3d\x1e\x0e\x3b\xe6\xd4\x78\x41\x85\x68\x86\x30\x4f\xe7\x08\x90\x19\xd0\xec\x47\x0c\x70\x02\x6e\x11\xaa\x00\x5d\x61\x5c\xe0\x39\xf8\xb6\x28\x4a\x04\x72\x49\x97\x58\x6a\x17\xc9\x05\x7e\x22\xc9\xbf\x76\x52\xac\xd0\x7e\x3f\x8a\xcf\xf0\x5d\x41\x09\x16\x24\x3f\x61\x57\x1f\xb5\x9f\xa6\x1d\xd8\x6d\x1e\x0b\xe1\x07\x5c\x3e\x80\xb4\x2c\xc9\x37\x90\x66\x62\xb9\x62\xb1\x7c\x51\x30\x20\xac\xe8\x8c\x92\x25\x28\x30\x2b\x72\x04\xf8\x02\x81\x4f\xe3\x51\x84\xe6\xf7\xc4\x7c\x71\x22\x10\xa2\xfc\x53\x5a\xae\x90\x32\x26\xd2\x6c\xf4\x25\x1c\xb8\xf6\x16\xf1\x57\xf4\xf0\xa3\xf9\x64\x58\xba\x27\xb0\xe9\x23\x43\x60\xba\xba\xc1\x88\x33\x8d\x48\xf0\x89\x55\x28\x2b\x66\x0f\x82\x2d\x07\x92\x47\x25\x49\x73\x50\x5b\x26\x80\x70\x5e\x91\x02\x73\xf6\x43\x78\x36\x41\x25\x4a\x59\x68\x41\xdf\xdb\x54\x4d\x50\x45\x58\xc1\x09\x0d\x6d\xd2\xf3\x26\x9b\x92\x15\xcd\x10\xc8\x48\x8e\x00\x6d\xa7\xf1\x48\xb0\x5d\xc6\xf7\xa6\xe2\x6a\x81\xc0\x85\xb5\x75\x4c\xcf\x07\xe6\x62\x42\x30\x23\xb4\x51\x8a\x00\x71\x4a\x30\x22\x64\x5d\x14\x8c\xff\xd7\xc9\xe7\x69\x92\x9c\x8d\x8e\x93\x44\x01\x27\xc9\x79\xfe\xdf\x4f\x21\xf5\xd3\x78\x04\x98\x9a\x6f\x37\xaa\xe2\x72\xff\x63\x88\xab\xb4\x7a\xec\x46\x64\x1d\x62\x59\xd4\x39\xba\xb7\x37\x39\xfb\x9f\x8f\xe7\x93\xb3\xd3\x7d\x70\x91\x2e\x6f\xf2\x14\x18\xbe\x13\xbc\x4b\x71\x5e\x22\x0a\xb4\x3a\x80\x1a\xa3\x41\xe6\x65\x81\x2f\x10\x9e\xf3\x85\x24\xf2\xc8\x7c\xe5\x18\x00\x9f\xbe\xf1\x28\xc2\xb9\x96\x69\x9f\xc6\x23\xc1\xb1\xa7\x32\x6c\x0b\x83\xc6\xa3\xd1\xf9\xe9\xe4\xbb\x8b\xbc\x98\x59\x20\x0e\x4f\x6f\x05\x63\x97\x69\x55\x15\x78\x6e\xca\x37\x1c\x13\xca\xc7\x94\x70\x92\x11\xc7\xf3\x2c\x38\xaf\x54\x38\x29\x64\x0b\x61\x44\x0d\x38\xf8\xee\xea\x6a\x2c\x4c\xda\x39\x66\x5c\x68\x5a\xe8\x9d\xd4\x75\x14\x83\x98\xc2\x96\x3b\x7a\x3a\xd6\x3d\xdf\xf4\xd9\x13\x5a\x33\xf2\xac\x63\x7d\x57\xa3\xe8\xf2\xf4\xab\xf8\x64\xd3\xe9\x85\x3b\x55\xd9\xb1\x34\x01\xfe\xbc\xa9\xc0\x26\xb8\xdf\x13\xc4\xa4\x55\xb6\x36\xdc\x0c\x57\x49\x19\x71\xa3\x52\x27\xce\x4f\x2e\x93\x44\xc2\x18\x2b\x19\x53\x52\x21\xca\x0b\x64\x5b\x49\xe1\xf6\x18\x5b\x2d\x91\x80\x1f\x93\xb2\xc8\x1e\x4e\x49\xb6\xf2\xe2\x26\xc7\x56\x88\x6c\xec\xf8\xe0\x68\x78\x70\xf4\x1f\xc6\x24\x12\x68\xca\x53\x8e\xf4\xf8\x2f\xd6\x2b\xe0\xe0\x93\xe0\x67\xb3\x19\xca\xa4\x33\x96\xee\xd7\xc1\xa6\x49\x2f\x70\x56\x54\x75\xa6\x35\x45\xf4\xae\xc8\x90\x72\xd0\xa5\xb4\x47\x83\x74\x99\xfe\x41\x70\xfa\x8d\x0d\x32\xb2\xb4\x92\x23\x73\xa1\x99\x36\x68\x5f\x00\x64\x9c\x25\xed\xc2\x5b\xef\x5e\xff\xd9\x58\xbf\xcd\xb7\x16\x66\x38\x4e\xf9\x42\x10\x7f\x98\x11\x7c\x47\xee\x0f\xa1\xfd\x56\x30\x54\xb1\xdc\x66\x85\xcb\x08\x05\xf9\xf0\x3e\x5d\xaa\x6d\xcc\x97\x05\x2e\x18\xa7\x29\x27\xd4\x63\x09\xdc\xb2\x4f\x60\xd7\xbd\x02\xde\x7e\x09\xfe\x7a\x3b\x62\x70\x0e\xfe\x24\x7e\xd6\xf2\xa9\x1e\x80\xcd\x16\xee\x99\xbf\x5a\xc8\x4d\x57\x42\xd6\x21\xdd\xca\x03\x25\xc9\x9b\x15\x56\x54\xed\x24\xe4\x23\x92\x23\x5f\xa0\xa7\xbf\xbc\x5e\x65\xb7\x88\xb7\xd9\xf7\x5f\x48\xa1\x25\xe4\x00\xf6\xc5\x5f\x6a\x5f\x61\xdf\x48\xc6\x25\x19\x13\x34\x97\x96\x7c\x03\xae\x7d\x71\x83\xd3\x5f\x74\x40\xed\x62\x55\x48\xa9\x72\x95\x87\x16\xda\xa6\xd8\x21\xf2\xf1\x43\x25\xd8\x87\x33\x59\x07\x29\x08\x1e\xfc\x51\x54\x50\xcd\x15\x15\x46\xed\x89\x05\xb2\x02\xe7\xe8\x7e\x80\xee\x75\x6a\x62\x81\x5d\xa2\x25\xa1\x0f\xd3\xe2\x0f\xc9\xd4\xa3\xe3\xff\xb4\x5f\xd7\xd6\x45\x91\xfe\x16\xf1\x13\xae\x64\xc3\x33\x41\x42\x32\x28\xf6\xd4\x0d\x1a\xe9\x74\xb3\xba\x40\xb6\xed\x8c\xba\x2a\x96\x88\xac\xa4\xe0\xfd\x32\x1c\xc2\xb8\xa0\x84\xab\x10\xb4\x31\x9a\x60\x10\x29\x40\x64\x94\xe0\x7f\x90\x9b\x5d\x40\xeb\x5a\x85\x09\xba\x63\x79\x83\x29\xfb\xd4\x81\x9c\xa2\xb9\x50\xee\x87\x28\xf6\xd0\xa0\x3a\x20\x86\x11\xa4\x8c\xab\x02\x91\xed\x4a\x3e\xac\x78\xb5\xe2\xdb\x0b\x64\x44\xc3\x81\x41\xf7\xe2\x5a\xb8\x2d\xdc\x68\xd6\x18\x1e\xd1\xa6\x15\x9c\x3b\xa1\x8d\x30\x5e\x22\x05\xb3\xc4\xa7\x81\x73\x5d\xe6\x0b\xf1\xff\xf5\x5a\xa4\x7a\x12\xaf\x51\x93\x0c\x15\xf2\xea\x6a\x24\x4d\xf1\x1c\x81\x97\xb7\xb2\x18\x79\x86\x39\x95\xb6\x97\xd5\x8b\x81\x67\x38\xbd\x29\x51\xbe\x5e\x83\x55\x55\x21\x2a\x20\x37\x9b\x56\x2b\xde\x13\xa9\x12\xc1\x92\x9d\x78\x32\x45\xa5\xb2\xa1\x5f\xc0\xd0\xd4\x71\x1b\xdf\x9b\x5a\xb9\x95\x19\x11\x7a\x7f\x70\x24\xd5\x49\x6b\x54\xbb\xae\xee\x15\xd6\x15\x35\x67\x75\x28\xb6\xba\x96\x0c\x64\x91\x61\x84\x1b\xb5\xcd\x1d\x91\xe5\x32\x3d\x45\x65\xb1\x2c\x38\xca\x45\x18\x04\x8d\xb2\x50\x93\x49\x1f\xf5\x87\xfd\xe3\x5f\xff\x64\xbe\xb3\x52\x08\x55\x1a\xf2\x8a\x3a\x74\x85\xfb\x60\x34\xfe\x08\x56\xb8\xe0\xea\x09\x12\xfa\x83\xfa\x20\xc5\x39\xb8\x7c\x2d\x46\x4c\x4e\x2e\x8d\x37\xb0\x95\xef\x5d\xd9\xd3\x88\xa0\x5c\x3f\xbc\x20\x73\x3b\x8b\x0d\xc8\x5b\x03\xa3\x24\xac\xbf\x65\x06\x43\x91\x63\x73\xd8\x4e\x8c\xcc\x99\xfc\xaf\x02\xda\x65\x8a\xd6\xac\xec\x54\x51\x8f\x54\xe1\x8b\x59\x3b\x6c\xf0\x2e\x65\xe3\x66\x37\xb4\x6c\x38\xd2\xd3\x02\xeb\xb0\x8b\x19\x15\x70\x43\x8c\x06\x42\xc0\xc0\x66\x73\x36\x9a\x5e\xa5\xec\xf6\x54\x10\x5f\xf0\x40\x62\x59\x21\x9c\xb3\x0f\xd2\x1b\x5a\x0e\xbf\xdf\x04\x76\xd2\xb5\x5c\x07\x52\x44\x05\x2e\x72\x3e\x77\x0e\x03\xd8\x88\x7b\x8e\x06\xc3\xdd\x82\x03\x3d\xf1\x15\xb9\x45\x78\xab\xe7\x8b\x7a\x3d\x1d\xbc\x45\x02\x09\x27\x7c\x98\xf2\x34\xbb\x95\x23\xa4\xda\x8b\xed\x6a\x78\x08\xfd\x90\xc2\xac\x35\x35\x88\xea\x67\x0e\xa8\x53\xfa\x6c\xc0\xcd\xe7\xce\x90\x26\x58\xd1\xa0\xe2\xb7\xeb\x9d\x53\x76\xbb\x43\x1c\x5b\x47\xb0\xf6\x82\xbc\x08\xf6\x7c\x99\xce\x0d\x38\xf9\x33\x04\xb8\x5e\x0b\x81\x45\x03\x69\x85\x70\x3e\x38\xa1\x34\x7d\xd8\x6c\xfc\x28\x56\x03\x04\x72\x0e\x60\x09\xb5\x8c\x8b\xfa\xe0\x25\x2a\x65\xcc\x2b\x45\x7c\x3b\x7a\x93\x18\x89\x61\xb3\xe9\xaf\xd7\xa8\x64\x68\xb3\x59\xaf\x11\xce\xa3\x63\xe0\x7a\x5d\xcf\xb5\xd9\xc0\x20\x69\xe1\xe1\xd7\x3e\x2b\xc4\x7c\x42\x81\x31\x32\x69\x56\x15\x08\x00\x61\x37\x5b\xd6\x6b\x70\x27\xac\x5c\x60\xe8\xc6\x4b\x96\xc2\x44\xc1\x51\xb5\x6a\x05\xdc\x70\x71\x47\x61\x17\xd7\xec\xbf\xe7\xe7\x5c\xc4\x2a\x22\x0d\xe2\x3e\x7e\x2e\xee\xd8\x49\x40\x03\x70\x32\x1e\xd7\x92\x28\x4c\x65\x54\x68\x85\x16\x9e\x8c\xfe\xaa\x61\x11\xbe\xd3\xbf\x23\xb0\x27\x9f\xa7\x7f\x9f\x9c\xbd\x3d\xff\xf0\xde\x1c\x61\x3c\x0d\x8f\x33\x62\x13\xf4\xd0\x07\x2f\xd5\xa6\x29\x31\x35\x96\x02\x02\xbb\x2d\xe5\x53\x08\x87\x1a\x03\x61\x08\x48\xdb\x6d\x81\x5d\x47\x34\x8d\x60\xa8\xbf\x7c\x69\x88\x0b\x69\xeb\xb1\x76\x5e\xc6\xe0\xa2\xc0\xb7\x9f\x52\xca\xc2\xc4\x79\xb4\x75\x52\x15\x9b\x1d\x5e\x7c\x78\xfb\xf7\xb7\x93\x0f\x1f\xc7\x31\xa7\x1e\x2a\x33\x4c\x3e\x8c\xce\xa6\x53\xdf\x7a\xb9\xc9\xad\x27\x62\x9f\x48\xb9\x5a\x06\xb2\x7c\x9b\x11\x68\x70\x49\x56\x98\x8b\xb8\x52\x0f\x08\xb3\x40\x79\xe9\xc1\x39\x9b\x3e\x30\x8e\x96\x91\x4d\x14\x44\x0e\xde\x11\xc6\x37\x9b\x64\xbd\x1e\x8c\x08\xe6\x69\x81\x11\x0d\x0a\x95\xe2\x95\x30\x1f\x11\x64\x91\x3c\xf5\xf0\x4e\x11\x7a\xe8\xe7\xbf\x8e\x03\x3b\x14\x76\x4e\x72\x4c\x58\xc4\x08\x61\xa1\x54\xb9\x25\x2f\x2a\x48\xb1\x37\xa0\x39\xcf\x96\x14\xbd\x27\x2a\x8c\x03\x2e\xa8\x67\x4a\xe1\xd9\x3d\xa7\xa9\xa0\x71\xdb\x9e\x05\x74\xb0\x19\x7a\x99\x56\x91\x0d\x0c\xef\x97\x18\x64\xba\x47\x2d\xe5\x21\x76\x08\x0f\x59\x9d\xe4\x39\x45\x8c\xd5\xe0\xb5\x1e\x84\x9c\xc8\xa3\x94\xe3\x19\x7c\xab\x63\xc0\x30\xd7\x9e\x8e\x77\x4c\x28\x37\x8a\xdc\x1d\x3b\x32\x10\xa0\x9d\x8a\x23\xf2\x86\x3d\xf4\x3b\x18\xd4\xe5\x56\x55\x30\xde\x07\x7b\x2f\x91\x88\x78\x5f\xeb\x94\x77\x7f\x67\x5d\x48\x84\x32\xc4\x62\x82\x2e\xdf\x24\x68\x5d\xaf\xc1\xa0\x9e\x12\x6c\x36\x42\x08\x82\xe6\x47\x73\x42\x80\x37\x2a\x03\x36\x9b\x43\xf1\xa0\x59\x49\x78\xf7\x41\xb7\x5e\x75\xa8\x3d\x74\xa8\x4b\xb6\x4e\xff\xff\x40\x79\xc7\xb4\xb8\x2b\x4a\x34\x47\x79\x6b\xaa\xdb\x67\x1e\x8f\x76\xad\x3e\x6a\xd9\x0e\x6c\xa3\x9d\xc2\x34\x5d\x47\x2a\xa6\x76\xca\x04\xa1\x98\xd7\xce\x7a\x5e\x58\xec\x51\xe1\xad\x21\x96\x2e\xb3\xa0\x99\xe8\x35\x7b\x55\x57\x60\xe5\x64\x91\x40\x3b\xc4\x7c\x3b\x91\x09\xe4\x40\x32\xf9\x7a\x11\xe2\xbe\x9d\xc0\x9e\x8d\x84\x0f\xd0\x95\xf8\xdd\x2a\xb0\x6d\x63\x4f\x5b\x14\xd4\xcf\x9c\x5c\xa3\x6d\x37\x19\x11\x3c\x2b\xe6\x2b\xea\x96\x27\x34\xa0\x6e\x1b\x79\x87\xd2\x92\x2f\x1e\xc6\xaa\x79\xa4\x95\x0a\xaf\x6d\xc5\x97\xe0\xba\x57\xa6\x6b\xac\xee\xa6\xb1\x05\xcb\xa5\x98\x15\x14\xe5\x23\xe1\xe0\x83\x61\x6c\xa4\x0a\xb4\x53\x18\xbb\x93\x98\x5c\x90\x34\xaf\x5f\x86\x6c\x68\x20\xf0\x6d\x14\x7d\xb7\xa4\xcd\x1c\x21\x2c\x9b\x1e\xb1\x27\x13\x22\xa9\x84\xd2\x34\x0f\xf7\x2d\x0b\x12\x42\x63\xd2\xda\x66\xcd\x2d\x73\x76\x97\x77\xcf\x5c\x38\xe7\x36\xce\x66\xc7\x4b\xdb\xa6\xf8\x47\x12\xfc\xa0\x3e\xf9\xc5\x8e\xae\x5d\xf6\x2b\x17\x06\xc1\x8e\x49\x32\xa7\xdb\x56\xe8\x0a\xf6\x39\xda\xc5\xc0\x86\x95\x66\xa5\xe7\xa5\x2e\x2e\x49\xf2\x92\x57\x9a\xde\xc1\xd8\x78\x6a\x00\xd7\xb3\x8c\x29\x9a\x15\xf7\x02\xbe\xa2\x05\xe6\x33\x00\x6b\xdc\xff\xca\xa0\x8d\xd3\x2d\x2a\x0d\x4c\x4f\x6f\x54\x92\x64\x07\x63\x60\x8e\xa0\x0f\x1d\x09\x03\x33\x2b\x32\xaf\xa9\x22\xda\x3e\xe9\x2e\x75\x2b\x5a\x19\xd3\x7a\x3d\x3f\x4f\xda\x92\x70\x6d\x36\xbc\x1d\x4d\xf7\x8b\x48\x94\x76\x66\x5e\x2b\x68\xf5\x78\x67\x07\x1f\xc3\xc3\x1f\xd2\xbf\xf4\x14\x0a\x65\xa4\xf4\x14\xd2\x84\xb9\x54\x26\xa9\x99\x6c\x92\xe2\x9c\x2c\x19\xd8\x2b\x38\x49\xdb\x59\xf6\x3d\x3f\xdd\xb9\x90\x27\x6d\xbf\x5d\x7b\x8e\x95\x65\xf5\x06\x5f\xba\x76\x6f\xbb\x74\x34\xba\xd7\xf0\xd8\x61\xad\xc3\xc7\xee\xf8\xc5\x19\xdb\x96\xf3\x8d\x0a\xb9\x6b\x3a\xc5\xbe\x59\xf6\x59\x8c\x03\xf0\xf4\xfd\x54\xa5\x87\xd7\x76\x77\xc3\x0f\x11\xe7\xfa\x9f\x8f\x09\xd5\x22\xd8\xad\x62\xb2\x5e\xb5\x1b\x27\x7f\x1f\x09\x77\x5d\xe0\x0f\x20\xdc\x14\x9b\x81\xeb\x76\x01\xa7\x2b\x24\xe5\x71\x60\x1a\xeb\xe7\xc9\xbb\x7b\x12\xf2\x03\x24\x3e\x20\x70\xb1\xee\xc4\x67\x72\xd2\x8d\x79\x8f\x45\x4c\x67\xce\x64\x34\xb7\x06\xe3\x5e\x28\xc1\xec\xb3\x30\x2f\xbf\x04\x3b\x1c\x10\x1c\xd4\xa4\x7a\x45\x14\xbb\x33\xf3\x1c\xcf\x75\xd1\xc0\x49\x34\x3a\x75\x4e\x43\xb9\x21\xa3\x2a\x44\x8d\x57\x37\x65\x91\xf9\x49\x1c\x1c\x15\x39\x3d\x17\xdc\x86\xc3\x81\xfc\xdf\xe1\x30\x50\xc0\x8f\x24\xa0\xed\x68\xa3\x03\x42\xb7\xda\xf9\xb9\x6c\x2c\x91\x84\xe7\x95\xd9\x55\xb5\x2d\x5b\x85\x6f\x28\x59\x1a\xb1\xab\xa5\xd3\x1e\xf0\x15\x89\x81\xda\x09\xe6\xb6\x20\xd1\xd9\xd9\x40\xaa\x6b\xa6\x59\x9f\xaa\xec\x3c\x77\xd9\xe2\x1d\x7b\xf7\xa3\xaa\x10\x3a\xc4\x55\xe2\x5b\xa6\x8c\x17\x59\x6b\x05\x0a\x3c\x4f\x12\xd3\x28\xb4\xe2\xfc\x34\x27\x61\xe5\xb9\x3b\xe8\x69\xbb\xee\x98\xfe\x28\x11\x44\xbf\x83\xc1\x34\x5b\xa0\x25\x02\xb0\x68\xaf\xd3\x58\x81\xb8\x7a\xaf\x5a\x5f\x42\x4d\x2f\x46\x8b\xb0\xd2\xbf\xf3\x99\xa2\xb2\x6e\xcf\xb5\xb7\xdf\xe8\x3c\xb0\xbb\x78\x5d\xd9\xf4\x00\xed\x9c\xc4\x52\xd5\xa0\x32\xb4\x94\x3b\x84\x35\xf7\x0a\xfa\xe6\x9a\xe2\xd2\xe4\x9d\x19\x46\x97\x7c\x1e\xc2\xe6\xaf\x33\xb8\x36\x7f\x45\xb6\xb8\x0b\xd1\xc1\x48\x76\x66\x9d\xd2\xb4\xc0\x05\x9e\xab\x76\x35\x45\x86\x96\x25\x98\x48\xe7\xd3\x37\xbb\x7e\xfe\x34\xb4\xcc\x5a\x8b\xc7\xec\x11\x01\xf0\x3c\x2f\x91\xd1\x2a\x24\x84\xcc\x78\xa4\x92\x42\x13\x0d\x25\x8c\xfd\x8d\x60\x54\x4f\xd9\xbe\x52\x65\x83\xd1\x02\x65\xb7\x6e\xb1\x42\x57\x14\xae\x16\x14\xb1\x05\x29\x65\xa5\xe9\xd8\x16\x28\xc9\xc4\xbb\xb4\xb1\x3b\x6a\x48\xfd\xd4\x35\x28\xf0\x2a\xa5\xf3\x70\xa3\x99\x57\x5d\x34\xd0\x19\xc6\x2c\x89\x4a\x68\x4c\x31\xeb\x90\x43\xa3\x22\x94\xc7\x0a\x90\xe6\x8c\x29\x5f\x38\x26\xce\x3f\x8d\x76\xf8\xaf\x46\x1a\x3b\x60\x01\x7f\xc4\x8b\x20\x37\xdb\xc4\xd7\xd8\x93\xba\xcf\xf6\x7b\x7a\x30\xcb\xcd\x2b\x76\x0e\x82\xa7\x46\x8e\x2b\x31\x22\x27\xa7\xfb\x57\x8e\xdf\xdd\xd7\xd9\xa8\x1d\x65\x94\xa9\xaf\x17\xc4\x3f\x31\x95\xeb\xb7\xad\xc7\xd3\x8b\x60\x1f\x6c\xd4\x93\x9a\x8e\x60\x67\x17\x19\x6a\x6d\xb6\x38\xe7\x02\x84\x39\xd7\xe2\x51\x13\x87\x8a\x28\x8f\xcc\x1c\x03\xc7\x74\xd3\xe9\x85\xc1\xab\xda\xc9\xfe\xb8\xbd\xf0\xa4\x20\x6a\xba\xbb\x40\x9f\x4b\x86\x5f\x51\x77\xdb\x73\xbf\x6f\xfc\x12\xe9\x79\xde\x51\x81\x7d\x85\xbd\x7f\xe8\xd2\xda\x40\x59\xd3\x6e\xa5\x56\x0e\xc7\xc2\x13\xec\x31\x97\x83\xea\x78\xc9\x02\x37\x5e\x85\x8e\xf9\x39\xa7\xc5\xcd\x8a\xab\x05\x47\x8e\xfe\x6a\x62\xb6\x91\x01\xac\xa4\x53\xb8\x2b\xff\xbc\x67\xe3\x9d\x88\x38\xfa\xc3\x74\x67\xe2\xf3\x35\xc8\x6b\xeb\xee\xbb\x9b\xe5\xcb\xca\xb3\xe5\xe7\xe2\xf5\x88\x90\xdb\x02\x4d\x79\x91\xdd\x16\x18\x31\xd6\xc4\x0f\x62\x55\xf6\xee\xa6\x33\x59\x49\x7d\x80\x16\x5b\x82\x05\xe6\x35\xd8\x21\x01\x8e\xa5\x55\xfa\x02\x75\x63\x2d\x40\x2b\xdc\xa1\xdb\xd7\x4d\xa7\x72\x73\xc8\xb5\x35\x16\xde\xf8\x63\x1c\x80\x96\x5b\xcd\xc6\x18\xc9\xc0\xb6\x04\x3d\xd0\xfb\x6c\xf4\x04\xca\xc6\x97\x11\x25\xf8\x2f\xe4\x86\xf9\xbd\xbd\x22\x8a\xc2\xce\xad\x93\x6d\x77\x4e\xa2\x29\xf1\x8e\xf7\x4d\x76\xb8\xc1\xd0\x71\xd7\xc4\x6b\x4c\xdb\x76\xcf\xe4\xfb\xdc\x32\x79\xc4\x1d\x93\xc8\xb9\xa4\x69\x49\xe3\x77\x4b\xa2\x56\xd6\x0e\xeb\x6c\x65\xd1\xfb\xeb\x9e\x80\x6d\xbd\x4d\xb2\xe3\x5d\x92\xce\x9b\x3f\xe1\xae\x88\x1d\x6e\xff\x98\x3c\x85\x28\x63\xc9\x64\x85\xaf\x52\x76\x1b\x06\xb5\x6f\xa6\x04\x41\xcc\xd4\x36\x62\xae\x4f\x28\x6e\x8e\x11\xc2\x20\x40\xd1\x92\x99\x67\x9a\x5b\xa2\x7a\x6b\x70\x4a\x71\x92\x7e\x63\x89\x40\x12\xf1\x03\xc0\xb7\x9c\xcd\x9d\x93\xf8\x08\xf8\x08\x74\x27\x59\x46\x56\x98\x9f\xe7\x5b\x30\xea\x55\x1e\x76\x60\x6e\xfa\xdb\x46\x17\x1f\xa7\x57\x67\x13\x18\x6d\x55\x00\xde\x1d\x96\xf6\x4f\xe8\xa9\xff\x2c\x14\x2d\x3d\x5f\xb6\xc2\x9b\x05\x4b\x32\x67\xc9\x88\xa2\x94\xa3\xa6\x7b\x2c\xe2\xb7\x6d\xd0\x29\xa7\x28\x5d\x76\xc2\x8e\x57\xfc\x82\xcc\xcf\xee\x10\xe6\x2c\xc4\xac\x80\xa3\xb7\x45\x3c\xd2\xb5\x55\x0b\x97\x9c\xa4\xe3\xe6\x52\x5f\xb5\xb6\x74\x49\x05\x80\x02\xcb\x81\xbc\x83\x9d\x1c\xa6\xdf\x58\x7d\x2f\xe9\x27\xff\x2e\x92\xfa\xf3\x4f\xdc\x9d\x7f\x1e\xcb\x03\xa7\x22\xad\xb4\x18\xa7\xc9\x00\x26\x71\xc6\xb9\x61\x7a\xd4\x5b\x18\x91\x80\x1f\x07\x68\x67\xdd\xdc\x8e\x8b\x39\xec\xe8\x35\x3a\xb7\x8a\xd7\xb8\xfe\xed\xd5\x3a\xff\x02\xda\x42\x3f\x30\x3c\x57\xc7\xf5\xb2\x7a\xa6\xe0\x01\xbc\x71\xab\x4c\x7f\x75\xe5\xdf\x07\xbf\x58\x37\xc9\x9b\xa2\xc4\xaf\x43\xab\x8e\xe4\xdd\xfc\x83\x7f\x2b\xaa\x37\x45\x19\xd8\x4f\xf8\x15\xfb\xe5\x98\xde\x8a\x21\xc0\x38\x2d\x32\xde\xfb\xcd\xf5\x9e\x77\x29\x05\xe9\x37\x06\x5e\x01\x8a\x7e\x5f\x15\x14\xed\xf5\xd2\x6f\xec\x80\xe5\xb7\xbd\xfd\x20\x30\xca\x04\x30\x46\xdf\xc4\xb0\xc1\xd9\x68\xba\xb7\x5e\xa6\xf7\x13\xc4\x69\x81\x58\x72\x34\xdc\x84\x87\x09\xf1\x35\xc6\x8d\x4a\xb2\xca\x3f\xa7\x3c\x5b\x5c\x90\x39\xdb\x0b\x8f\xd1\x86\x1b\xbc\x02\xbd\x80\x7d\xf6\xd6\x12\x31\x27\x7a\x76\x29\xcd\x02\x95\x65\x32\xcc\x96\x5a\x00\x7b\xbf\x05\x7b\x3c\x3b\x10\xeb\x9b\x90\x1e\x5e\xe3\x6e\x43\x14\xad\x44\xc0\xad\x7e\x0a\xc1\xa2\xb5\xbb\x2c\xff\xb0\xc9\xbf\xcb\xb2\x85\x54\xf1\x4a\xf0\xb0\xce\x13\x7a\x89\x43\x6f\x5b\x83\xeb\x68\xf5\x10\x4b\xe9\x87\x39\x14\xac\xca\xa8\x69\x7b\x49\xaf\xe7\xee\xae\xd7\xc6\x84\xee\x2b\x91\xfa\xd5\x0a\x07\x5e\x81\x99\x56\xec\x3d\x24\xac\x5d\x1f\x64\x04\x73\x74\xcf\xf7\x3d\xfe\xc8\x59\xa4\xb8\xa8\x6b\x01\xe0\x15\x90\x43\x06\xfa\xf7\x80\xa2\xaa\x4c\x33\xb4\x77\xf8\x6f\xff\xb2\xf7\xf5\x6b\xfe\xf3\xfe\x6f\x87\xf3\x7e\x8b\x7f\x29\xa4\xb0\x0f\x72\x94\x45\x70\x03\x40\x11\x5f\x51\x0c\xd4\xb9\xfd\x60\x46\xc9\x72\xb4\x48\xa9\xd0\xcc\x3d\x31\xcc\x13\x5e\xf1\x9f\x80\x1e\xd4\x84\xaa\x96\x8a\xc0\x56\xc3\xfa\x1f\x8c\xa7\x94\xa3\xfc\xf5\x43\x02\x7a\x22\xf3\xe9\xf5\x63\x90\xb6\xfc\x24\xae\x3c\x7d\x51\xac\xd0\xcd\x23\xd7\x51\x34\x5a\xd5\x92\xfa\x1f\x71\x40\xe1\x5c\x13\x70\x14\x05\x20\x77\x88\xd2\x22\x47\x2c\x89\x2f\x4f\x21\xd2\x4d\x56\x1f\xda\x01\x5f\xba\x06\x00\x29\xde\x38\x5d\xa2\xc4\x5a\x54\xbf\xde\xf8\xe4\x0b\xe8\xb1\x45\xaf\x0f\x7a\x07\x59\xaf\x79\x2a\x84\xb5\x0b\xed\x75\xec\x65\x70\xd4\x26\xba\xa9\xec\x16\x7d\x03\xaf\xc0\x65\xca\x17\x83\x59\x49\x08\xdd\x93\xff\xa4\xb2\x75\x63\x6f\xff\xa7\xa3\xe1\x70\x38\x0c\xcb\x44\x49\xe6\x7b\xd6\x92\xc0\xcf\xa0\x97\xf4\xc0\xcf\x8d\x79\xf9\x19\xf4\x0e\x85\x1c\xc8\x59\x5e\x89\x37\x72\xba\x9f\x41\x6f\xc9\xea\x85\xca\xc7\x96\xe4\x1b\x42\x8e\x28\x8d\x4a\xb7\xdc\x0b\x46\x4a\x34\x10\x84\xf4\x10\xa5\xc7\xbd\x3e\x10\x23\x82\xd4\x8a\x3f\x0c\x71\xed\xae\xf6\x9a\x29\xf6\xc1\x5a\x38\x87\x01\x55\x09\xce\x9e\x92\xf2\x46\x71\x07\x39\xc1\x68\x5f\x18\x11\x41\xfa\xce\x3a\xe3\x33\xbc\x9e\x50\xb2\x6d\x89\x18\x4b\xe7\xa8\x0f\xb2\x9b\x0e\xcb\xc0\x64\x64\x25\x8c\xb4\x60\xe2\x81\x60\xd4\x9e\xf0\x44\xa7\x29\x47\x7b\xfb\xfb\x83\xb9\x5a\x4e\xc0\x0d\x81\x9d\x55\xb6\x76\x31\xc2\x7e\x26\xcd\xaf\xa8\x9a\x94\x75\xbc\xa7\xe0\x59\x28\xf6\x53\x3c\x89\x48\x0c\x1b\x64\x76\xe0\xd8\x30\xfc\x89\x9b\xbe\x6d\xcf\x77\x63\x83\xa6\x4e\x85\xa8\x3b\x69\xb4\xde\xc2\x04\x34\x7b\x29\xc2\x24\xc6\xd3\x65\x95\x44\xb6\x69\x8b\x46\x47\x99\x0e\x1e\xbf\x4f\xe0\x31\x7b\x05\x62\xfb\x05\x1c\x7e\xeb\xa5\xc6\x99\x2d\x37\xb8\x32\xa2\xfd\x56\x9d\x6e\x62\xba\xe3\x3b\x59\xb3\x44\x63\x05\xed\x81\xd3\x78\xa7\x5b\x46\x87\x1c\x5b\xea\x6a\x22\xa4\x98\x2e\x08\xe5\x3a\x6a\x98\xac\x3a\x6a\x6c\x5a\x26\x12\x09\xb4\x35\x16\x37\x9a\x7a\x07\x17\x04\xcf\xf5\x0c\x07\x2c\x5b\xa0\x7c\x65\x7f\x1c\x66\xaa\x9f\x9d\xdd\x57\x14\xb1\xba\xd6\x23\x89\xd3\x6f\x9c\xfe\x22\x75\x9a\xe9\x95\xbc\x65\xd8\x1e\x8d\xed\xdb\x54\x23\x72\x83\xf6\x3c\x0f\x10\xac\x0f\x4e\x9d\xb3\xd7\x4a\x9f\x3e\x7e\xad\xef\x47\x7f\x85\x09\xf8\x5a\x37\x80\x48\x1f\xb0\xd9\x7c\x85\x7d\xf0\x15\x6a\x63\xde\x02\xe8\xeb\x8f\x12\xc0\xd8\xe3\x50\x55\x35\xb0\x45\x2a\x71\x1a\x23\xba\x2c\x18\x0b\x65\x58\xc0\x4d\xb1\x0c\xd8\xd0\xae\x01\xbb\x3a\x9a\x35\xed\x87\x2a\xd5\x4e\xce\xf1\x1d\xb9\x45\xa1\xcf\x9d\xd4\xcf\x54\x9b\xd0\x13\xf9\x6e\x14\x3e\xc5\xa4\xd2\x01\x32\xa7\xd4\x69\x8a\x8a\xcc\x84\x25\x9a\x68\x37\xa0\x27\xd1\xc6\xc4\x5d\x8a\x13\xae\x5c\x87\xbe\x96\xa1\x2e\xd9\x4f\xf4\x9b\xc0\xe7\xf1\x02\x97\xc6\x27\x06\x98\xde\xdb\xe0\x65\xf1\xb0\x46\x3d\xff\x92\xb8\xf1\x15\x3f\xaf\xe9\xdd\x6b\xf9\x7a\x51\x33\x29\xfa\x5d\x00\xff\x03\x1b\x36\x53\xd4\xf9\xa6\x5e\xa7\xf7\x7d\x03\xaf\xd4\xf3\x42\xef\x46\x27\x53\x3b\xba\x40\x03\xc3\xfa\xde\x92\xb5\x00\x44\xd7\xe4\x7c\xe3\xc4\xea\xca\x0f\x7f\x48\xc1\x3d\x77\x88\x6c\xde\x4e\x67\x0e\xd1\x02\xb6\x53\x55\x8f\x56\xa5\xdc\x8a\xbf\xf5\xda\xbd\x8b\xd0\x5d\x00\xb3\x8f\x23\xdc\x79\x8c\xc3\x09\xaf\x8e\x04\x45\x0c\x69\xab\xef\x53\x2a\x4b\x86\xca\x87\xcb\xff\xb5\x22\xb7\x26\x2d\x76\x5c\x11\x3e\xac\xb0\xfc\x86\x7d\x50\x61\x5d\xc6\xf0\xc1\xe2\x1f\x28\xfb\xfe\xdf\x1e\xeb\xac\x40\x42\xa4\x5a\xea\x4a\x92\xe6\x37\x4d\x4b\x9d\xea\xf6\xbc\x41\x91\xf3\x87\xc8\x18\xa5\xce\x88\xd6\x27\xb9\xec\x0d\x25\xcb\x60\x73\xde\x76\x6c\x13\x17\xd7\xe7\x82\x2f\x76\xc0\x95\x1d\x6f\x25\x3e\x3b\x4e\x4e\x56\x7c\x41\x68\xf1\x07\x0a\x36\x9e\x7a\xa3\x42\xc7\xd5\x46\x41\x35\xc8\xd7\x9f\x02\x68\x9c\x27\xce\x45\x9f\xa0\x10\xd7\xff\xba\xde\x6e\x4d\xcd\xaf\x27\xf9\x1f\x25\xb2\x6d\xce\xf4\x97\x24\xd1\xdf\x0d\xd3\x46\xe7\x14\x95\x48\xc8\x49\x73\x52\x0d\x27\x48\x24\xe7\x5b\x8c\x92\xfc\xdc\xef\x88\x60\x4e\x55\xe7\x8c\xdb\x8a\x08\xaf\x52\xe7\x3e\xec\xba\xfe\x56\x07\x64\xf2\x72\xb8\xb0\xb1\x4d\xbf\x80\xfe\x54\x99\xd5\x9f\xd7\xc0\xa7\x55\x65\x02\x77\xb8\x9e\x10\xdb\x0c\xae\xfd\x5f\x00\x00\x00\xff\xff\x0b\x04\x17\x63\xf5\x5d\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
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

