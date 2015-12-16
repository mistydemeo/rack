// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// models/templates/service/mysql.tmpl
// models/templates/service/papertrail.tmpl
// models/templates/service/postgres.tmpl
// models/templates/service/redis.tmpl
// models/templates/service/webhook.tmpl
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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5b\x6b\x6f\x1b\xbb\xd1\xfe\xee\x5f\x41\x2c\xf2\x29\x90\x65\x3b\xc1\xc9\xfb\x56\x68\x0b\x38\xb2\x73\xa2\xc6\x76\x0c\x49\xc9\x41\x1b\x18\xc5\x7a\x45\x59\x5b\x49\xcb\x3d\x24\xd7\x27\x8e\xa1\xff\xde\xe1\x6d\x97\xb7\x5d\xc9\x97\x18\x6d\x91\x73\xb1\x97\x43\xce\x70\x66\x38\xf3\xcc\x90\xb9\xbf\x47\x33\x3c\xcf\x0b\x8c\x92\xb4\x2c\x13\xb4\xd9\xec\x21\x74\x0f\xff\x22\x94\x1c\xff\x36\x99\xe2\x75\xb9\x4a\x39\xfe\x40\xe8\x3a\xe5\x5f\x31\x65\x39\x29\x12\x34\x40\xc9\x9b\xc3\xa3\xc3\xfd\xc3\x3f\xc1\x3f\x49\x4f\x91\x0f\x49\x31\xcb\x39\x8c\xb3\x64\xa0\x97\x80\xa5\xee\x11\xd7\x6b\xa0\xe4\x3a\x5d\xa5\x45\x86\xe9\x7e\xd6\x90\xa2\xbe\xe2\x19\x10\x97\x94\x64\x98\xb1\x36\xda\xe4\x3d\xac\xb5\x1c\xae\x2a\xc6\x31\x15\x0c\x51\xf2\xa1\x18\x0c\x4e\x7f\xaf\xd2\x95\x10\xe0\x9b\xf8\x32\xc6\x73\xf8\x31\x31\x54\x68\xd3\x43\x49\x82\xae\x90\x5a\x64\xa3\x05\xbf\x4c\x69\xba\xc6\x40\xc0\xc4\xce\xba\x25\x2f\x05\xed\x0e\x52\x3b\x74\x46\x64\x4b\x5a\xfd\x09\x3e\x4e\xef\x4a\x2c\x35\x3a\xe1\x34\x2f\x6e\xb4\x36\xe5\xd0\x09\x9e\xa7\xd5\x8a\xcb\x51\xf7\x3b\xcb\x68\x5e\x72\x63\x8b\x44\x0f\x6d\x7a\x35\xa7\xb2\x8a\x70\x01\xd2\x8b\x6a\x7d\x0d\x12\x44\x98\x48\x9b\x1e\xb6\xb1\x11\x5a\xbc\xfc\x82\xd8\x22\xa5\x98\x21\x32\x47\x38\xcd\x16\x48\xef\x36\xe4\x7f\x5a\xdc\xe6\x94\x14\x6b\x5c\xf0\xb8\x1c\xed\x9b\xed\xd8\x6b\x74\xab\x9f\xf0\xdd\xcf\x66\x31\xc6\x2b\x9c\x32\xfc\x02\x76\x1b\xe3\x92\xb0\x9c\x13\x1a\xdb\xd3\xd3\x98\x4d\x48\x45\x33\x8c\x32\x32\xc3\x88\x36\x6c\x02\x11\x26\xd5\x75\x81\x39\x6b\xe1\x7f\x96\x33\xfe\x67\x08\x0c\x70\xd2\x86\x6f\x06\x03\x45\x3c\x18\x8c\x66\x7f\x7d\x8c\x4c\x5f\x2f\x87\x88\x29\x7e\x68\x4e\x28\xe2\x8b\x9c\x21\x11\x87\x02\xa9\x4c\xe8\x71\xa4\xb2\xec\x29\x4e\x1f\xe3\x5d\xde\x4b\x8a\x5b\xf2\x1d\x76\x2e\x4d\x89\x6e\xf5\x7a\xbd\x56\xbf\x09\x45\xb8\x1c\xb6\x28\xa5\xd1\x07\xd0\x08\x65\x3c\x56\x17\x51\x1d\x38\xb1\x6a\x8c\x99\xb4\xa3\x6d\x9f\x64\x08\x61\x85\xac\xa7\xa4\xcc\xb3\x31\x59\xc5\xfc\xd4\x08\x39\x3a\x3e\x1f\x0c\x24\x8d\x25\xc9\x25\x25\x25\xa6\x3c\xc7\xae\xd1\x45\x06\x60\xac\x5a\x63\x41\x7f\x49\x56\x79\x76\x77\x42\xb2\x2a\x38\xd3\x9e\x7d\x44\x66\x78\xb3\x0f\xc9\xe1\xe8\xff\x2c\x26\xca\xb5\x38\x58\x49\xcf\xff\xe6\x0c\x21\x6f\x3d\x49\x7e\x3a\x9f\xe3\x4c\x5a\xf7\x78\xb5\x22\x7f\x78\xab\x69\xd1\xf3\x22\xcb\xcb\x74\xa5\x32\xc0\x04\xd3\xdb\x3c\xc3\x32\xfc\x83\x4b\xac\xaf\x67\x69\x3f\x5d\xa7\x3f\x48\x91\xfe\xc1\xfa\x19\x59\xcb\xe0\x1f\x59\xe7\x38\xd3\x7e\x02\xf3\x18\x67\x83\x66\xe3\x30\xc3\x23\xdf\x38\xbf\xdb\xa3\xce\xca\x90\x56\xf8\x42\x08\x7f\x90\xb8\x9f\x85\x26\x95\xae\x5d\x1d\xf8\x1a\x50\x94\x77\x17\x90\x9b\xa4\x0e\x66\xeb\xbc\x80\xd3\x47\x53\x38\xb7\x81\x2e\x92\x2d\x06\x92\x34\xbb\x18\x49\x12\x3a\x86\x12\x8a\x0d\x4c\x61\xa9\x2c\x79\x2d\x7e\x35\x8e\xa9\x3e\xa0\xcd\x16\xb5\xd9\xbf\x35\x94\x9b\x30\x91\x35\xae\xdd\xe1\xd6\x67\xd2\xd4\x83\xc1\x87\xaa\x50\x52\xed\xe4\xdd\x43\x08\x85\xa1\x27\x4f\xde\xbe\xaf\xb2\x25\xe6\x0d\xa6\xf8\x1b\xc9\xb5\x6b\xec\xc3\x4e\xe1\x7f\x99\x8c\x25\xf0\x73\x03\x31\xa4\x18\x63\x7c\x23\x4f\x33\x6c\x3e\xf4\x33\x58\x58\xa7\x2a\x7f\x55\xb5\xa8\x8e\x4c\x07\xce\xb2\x35\xe2\x12\xc8\xe5\x60\x2e\x51\x18\xfc\xde\xff\x91\x97\x89\x62\xd2\xea\x7e\x1f\xd3\x62\xb6\x92\x60\xc3\x9c\x04\xfc\x1d\xc0\x47\x01\x67\xc5\xa1\x3b\xc7\x6b\xc8\x03\x93\xfc\x87\x54\xe7\xd1\x9b\xff\x77\x87\x4d\x40\x51\x42\xff\x8a\xf9\x31\x57\x5e\x11\x44\x1d\xe1\x13\xb4\x08\x4e\x58\x32\xae\x0a\x9e\x2b\x1f\x2e\x40\xe3\xff\x62\x2e\x83\x29\x8c\x91\x4a\xfa\xd6\xdb\xc3\xa4\xdd\x15\xe2\x78\x8c\xd6\xf1\x70\x2b\x24\x7b\x00\x29\x53\x71\xc4\xc3\x6f\x0e\x29\xc3\x59\x45\x73\x7e\x97\xb4\x2c\xc5\xc4\x11\xaa\x07\x4d\xf8\xfe\x5c\xf1\xb2\xe2\xdb\x01\x32\xd1\x74\x5b\x25\x75\x09\xeb\x1c\x8e\x39\x87\x1c\xe6\x25\xf1\xaf\xe9\xaa\xd2\xb6\xd4\xee\x55\xd3\x35\xea\xde\x33\xff\xdd\xec\x01\x43\x5c\xcc\xe4\xba\x56\x99\x10\xc3\xe5\xaa\x6a\xb8\x47\x34\x2d\x6e\x30\x7a\xb5\x44\x83\xbf\xa0\xfe\x69\xc1\xa9\x8c\x5e\xcc\xec\x41\x61\x76\xa0\xab\x4a\x38\x92\x82\x6e\xb3\x69\x42\x76\x17\x82\x8f\xcf\x69\xf0\x7c\x4f\xf1\xd7\xe2\x76\x0b\x6e\xa0\xb9\x27\x34\x96\x42\xd7\xa2\x36\x1c\x71\x5f\x6c\x02\x06\x86\x64\xbd\x86\x43\x65\x67\xde\x36\x5c\xe6\x64\x7d\x58\x0a\x16\xd1\xb3\x15\x21\xac\x66\xd3\xc6\x60\xa1\x71\x99\x88\x20\x30\x21\xa7\x78\x36\x24\x95\x13\xe9\x1b\x69\x3c\x94\xef\x48\x73\xd4\xce\x78\xba\xc0\xa8\x90\x53\x05\xc2\xcf\x0b\xf0\x61\x38\xba\x32\xe2\x48\xcc\xcf\x61\x5c\x6b\x11\x71\x82\xc0\x15\x01\x50\xc2\x9e\xd0\x12\xe3\x12\xd1\xaa\x28\xc4\xd6\x48\x81\xee\xe0\xa0\xa1\x4c\x57\x3b\xdb\x76\x33\x5a\xa7\x37\xf8\xc1\x4a\x7d\x82\xfa\x54\xc4\x0b\x38\x76\xe8\x4d\x24\xcc\x5f\xde\xc5\x59\xc2\xd8\xf9\x7b\xa1\x9d\xf1\xf1\xb9\xd0\x0a\x84\x19\x70\x4f\xbc\x55\x0a\xcb\xf1\x9f\x7f\xe3\xbb\x9e\x85\x3a\x7c\x48\xb7\x4f\x3e\xc1\x20\xcb\x9b\xb0\x11\x09\x19\x86\x44\xc5\x88\xde\x96\xf5\xad\x98\xdb\xc2\xc1\xc9\xe3\x7a\x14\xea\x0a\x4e\x71\xba\x36\x9d\x85\x68\x0a\x4f\x26\x50\x8b\xd6\x67\xe0\xa8\x89\x5b\x7a\xff\xf9\x1c\xf5\x3f\xa6\xec\x52\x49\x62\x05\xa1\x33\x72\xc3\xbe\x30\xa7\x0c\x8f\x80\x64\x49\x51\xab\xba\x05\x44\x34\x18\x4f\x41\x82\x03\x07\x78\xc4\x81\x9e\x07\x37\x5c\x90\x27\x64\xf3\xe0\xb9\x45\xd5\x81\xee\x76\xc4\x76\x9d\x10\x3c\x06\xc2\x77\x82\xe1\x1e\x80\x5e\x6a\x2b\x5e\x56\x7c\x8c\x33\x42\x67\x60\xfd\xab\xe8\x2c\x0b\x34\x7e\x6b\xc3\x45\x29\x2d\x06\x80\xde\x07\x66\xd5\xd7\xf0\x87\x49\xf7\x38\x08\x11\x18\x6c\x2f\x5b\x4a\x65\xca\xdc\xb0\xff\x5a\xa3\xa4\x00\x8d\xfa\x78\x14\x79\x14\x0e\xae\xda\xf3\x69\xdc\xa4\x2e\x8c\x76\x9c\xc9\x46\x48\xa7\x4b\x29\x1a\x01\x02\xb7\xfa\x15\x04\x87\x5c\x16\x35\x56\xb8\xd6\xd6\xab\x98\x5c\x15\xf4\x7d\xeb\x54\x71\xc2\x61\xb5\x1b\x35\x4a\xa9\x7d\xdd\x4f\xf0\x0f\x09\x12\x0d\x1a\x32\x29\xf3\x15\x64\xb2\x7c\x0e\x05\xb7\x95\x33\x3b\x8e\x5c\x5b\x92\x8d\x07\xc6\xd3\xe1\x64\x9a\xb2\xe5\x89\x90\x22\xe7\x91\xa2\xbf\x04\x59\xd9\x67\xe9\x24\x4e\x7d\xd0\xab\x0b\x40\x79\x84\xae\x22\xb5\xbd\x22\x17\xc5\xba\xcf\xc3\x22\xb6\x8e\xd2\x51\xff\x70\xb7\x5a\x42\x33\x9e\x92\x25\x2e\xb6\xc2\xe5\x56\xa8\xdc\xd8\x2f\x56\x77\x74\xfb\xba\x42\x19\x5a\x87\x49\x58\x81\xd8\x5d\xac\x7a\x21\xf3\xcd\x23\xf5\xda\x77\x35\xb9\xfd\xdd\x9b\x52\xd7\x36\x26\x4f\xe0\x3b\x9f\x44\x68\x5c\x03\x3b\x1b\xca\x62\x81\x13\xff\xc9\x61\x50\xa1\xd8\xce\x7a\x30\xee\x2f\x61\x2e\xf5\x1d\xc5\x0e\xc0\xd2\x51\x61\x3e\xb8\xea\x7b\x0d\xb8\x1b\x8f\x6c\xf8\xd4\x4e\xde\x37\x54\x26\x5c\x49\xc6\x96\xb2\x7b\xee\xea\xfa\x28\x59\xab\xd9\xb6\x6f\xbe\x6f\x71\x50\xb3\xa9\xff\x08\xcf\x74\x1a\xed\x61\x5b\xdd\x26\xf5\xd1\x69\x04\xc4\xc7\x91\xec\x4b\x1e\x86\x30\xc2\x74\x89\x19\x86\x0b\xff\x6c\xd5\x55\x72\x5d\x59\x35\x81\xc8\xa3\x3d\x23\xe9\xcc\x78\x54\xa4\x0b\xd4\xc4\xca\xfe\x25\xa1\xfc\x3c\x2d\x4b\x51\xa2\xf9\x1e\x8a\x42\xc5\x0c\x3c\xc5\x3c\xc8\x87\x23\x6a\x53\x1f\xa0\x70\x29\x78\x0a\x09\x81\xc6\x55\x19\x77\x79\xe4\x1b\xe8\x82\x28\x2c\xb9\x5b\xe7\xc7\x5e\x72\x5b\x9a\x8a\xde\x2a\xb9\xc5\x5d\xbd\x7b\x3b\x25\xbd\x32\x13\x9d\x9c\xd4\xcc\x89\x6b\xdf\xaf\x61\xcd\x22\x7d\x9d\xf5\xb4\xf2\xc4\x64\xa1\x3c\x2b\xc0\x0c\xc5\xf9\x9c\xe7\x99\xe8\x0b\xec\x5a\xed\x3e\x6e\x75\xb7\x2e\xd6\x51\x20\xab\x28\x7e\xcc\xc2\x3f\x41\x56\x25\x8c\x12\xf3\xef\x02\x57\x5c\x3d\xc9\xee\xf1\x82\xfe\x39\x6d\xfe\x98\x4d\x9a\x1f\x5b\xee\x0c\xba\x2f\x72\x7c\x1b\x3c\xe8\x12\xe9\x19\x9c\xf3\xc1\x02\xff\x7c\x09\x3f\x12\x16\xbb\x4b\xdc\x49\x97\x79\x31\xc3\xdf\x2d\x66\xe3\xb4\x98\x91\x35\x7b\x61\x1d\x6b\xb7\x7f\xcc\x1e\x2e\x48\xa7\x78\xd6\x98\xac\xd1\xf0\x4c\x06\x5b\x75\x56\xe5\x11\xeb\xc9\x35\x9a\x02\xe6\x09\xc7\xcd\xed\x19\x3c\xe6\xbc\x25\x86\xb4\xd6\xa2\xa7\x3c\xcf\xd6\x76\x1b\xc2\x07\x2f\x42\xc7\x4e\x4e\x53\x89\xeb\xe4\x62\xa2\xd0\x80\x77\xd9\xff\x82\x47\x3d\xec\x9d\x3c\x69\xed\x67\x49\x91\x7e\x3f\xe6\x51\xd6\x8b\xa8\x7c\xa2\xbb\xe1\xbf\x52\x52\x95\xad\x55\xb0\xba\x32\x76\x48\xb7\x56\xc2\x92\xcc\xed\xb7\x05\xc0\x07\xed\x80\x08\xf7\xaf\x6b\x45\x7a\x10\x26\x71\x24\x1a\x15\x37\x54\x55\xf2\x5e\xcd\xd0\xe9\x36\x9a\xca\xef\xb6\x0c\xf3\x19\x1d\x09\x7d\x24\x87\x7d\xf9\xe7\xe0\x30\x6c\xc9\x8c\x4a\xd8\x3b\x27\x19\x91\xf5\x3e\xcf\xca\x90\xe4\x03\x25\x6b\xc1\xf8\x39\xbd\x29\x60\x32\x25\xcf\xcd\xc2\xed\xa8\x6c\xab\x92\xb6\x03\x46\xbb\x60\xfa\x5a\x66\xa3\x99\x23\xac\xb8\x9f\x8f\xb5\x39\xe2\x1e\xdb\xee\xa4\xab\x94\xf1\x3c\x6b\x50\x3a\x58\x7a\x30\xb0\x41\xfb\x0e\xed\x9b\xe6\x0d\x45\x53\x11\xe8\x6f\xb6\x22\x54\x29\x8a\x7f\x47\xfd\x49\xb6\xc0\x20\x57\x92\x17\xfa\x8e\xce\xab\x20\xd5\xb8\xb6\x89\xa1\x76\x92\x56\xa8\x52\xf1\xe6\xa1\xc0\xb2\x31\x77\x42\x01\xc1\xc3\x46\x54\x1f\x51\xc9\x75\x5a\xa4\xd7\x2b\x2c\x74\xc8\x69\x85\x7b\xf6\x4d\xdc\xbb\x43\xe7\x80\x34\xeb\xd8\x17\x4b\xe0\xb9\xb3\x15\x6e\x26\xbd\x7d\x77\xe8\x4d\xa3\x84\xb1\x7f\x90\x02\x1b\x16\xcd\xd0\x47\x9c\xae\xf8\x62\xb8\xc0\xd9\xd2\xaf\x57\xd5\xd0\xdd\x74\x01\xc7\x70\x41\x56\x33\xd9\xd1\x74\xaf\x0c\x47\x42\x49\xb7\xb2\x3f\xf6\x8b\x57\xd5\xd1\x9b\xf8\xcd\xad\xaa\x8d\x92\xe9\xf0\xd2\x69\x1a\xb6\x25\x20\xe3\xd7\x1f\x72\xca\xb8\xf8\xc5\x24\xa5\xd8\xdd\xae\xa5\xb8\xb7\xce\xf7\x2f\xc5\x22\xba\x99\xa6\xea\xb1\x54\x22\x5e\xd6\xe0\x22\x2c\x09\x77\x0c\x3d\x6a\xc3\xa3\xb9\xda\xee\xb3\x14\x29\x7e\x84\x88\xbc\x28\xb0\x23\x97\xd4\x6d\x48\x62\x9f\x9c\x9f\x1d\xc3\xa4\x6f\x88\x5b\xab\x0c\xef\x20\x59\x4d\xfa\x1c\x52\x69\xef\xf0\xb8\x04\x12\x7a\x76\x7a\x42\x59\x16\x6e\x28\xda\xbd\xb7\xf5\x30\x99\x9c\xc5\xbb\xf7\x2f\x6b\xa3\x16\x2b\xb5\x4a\xf7\x33\xec\x14\x65\x04\x12\x58\xfe\xef\xa5\x97\xa7\x17\xe5\x01\xcb\x88\x14\xff\xe5\x36\x8c\x9f\xb4\xff\x69\x1b\x7a\x5f\xae\xfc\xf7\x61\xcf\x0d\x7c\xce\xde\x0f\x09\x59\xe6\x78\x02\x20\x45\xde\x89\xb1\x3a\xa9\x7f\xbb\xf7\xef\x13\xd3\xb9\xec\x58\x8a\x56\xbc\xb3\x86\xe5\x2b\xa6\xcd\x0a\xbb\xf6\x3f\x83\x98\x6d\x30\x39\xd2\x06\xea\xae\x09\xac\x3b\xb8\x07\x5c\x41\x79\xaf\x6c\x9c\xb6\x6a\xfc\x0a\xd9\x7f\xe1\xd8\x72\x79\xbc\xd3\xcb\xc6\xd6\xdb\x52\xef\x0e\xb7\xb9\xfc\xf4\x31\xb5\xf7\x96\xd0\xf5\x0c\xef\x8e\xb7\xfb\x9e\xd5\x7d\xea\xe8\xf3\xb1\x1e\x3e\x86\x79\x0e\x67\xcc\x7b\x05\xe9\xfb\xac\x0b\xd4\xf7\x62\x3f\x5f\xc5\xaf\x21\xac\xbb\xe6\xa0\x47\x12\x3e\x80\x8c\xdf\x89\x3b\xf0\xcf\xf5\x5f\xdb\xde\xe1\x2b\xca\x8e\xf7\xa8\xcf\xff\xd4\xb4\xd5\xc8\x72\x14\xab\xaa\x61\x05\x27\xe8\xba\xae\x1a\x54\xd9\x7a\x8d\x5f\xc7\x63\x62\xcb\x1c\x8a\x6f\x04\x14\xa4\x26\x64\x32\x51\xfd\x45\xeb\x8f\xed\xab\x8d\xfd\xb5\x7e\xcb\xf9\x62\x87\xb5\xb2\x37\x5b\x85\x07\x92\xe3\x8a\x2f\x08\xcd\x7f\xe0\x68\x05\x1d\xcc\x8a\xbc\x01\x70\x5e\x00\xc4\xd8\xbc\x8e\x2c\xe3\x63\xac\xbd\xb6\xd1\x66\xc4\xfc\xa4\x46\x3b\x9f\xae\xd8\xd7\x91\x22\x2e\x75\x01\x6b\xff\x2d\x8f\x0c\x69\x2a\xec\x18\x0f\x16\x81\xd1\xbd\x27\x4c\xea\x07\x4f\x91\x32\xc4\x79\x13\x55\xd7\xaf\xd6\xcb\xb3\xf6\x39\x86\xa8\x9e\xd5\xbc\x73\x6a\x9f\xa4\x69\xea\x39\xf1\xbf\xb8\xd1\x5c\x4a\x2d\xf1\x5d\x0f\xbd\xba\x15\x09\x4a\x3f\xf2\xbb\x85\x92\xc4\xa9\x3a\x45\xe3\x08\xc8\x54\x8d\x2d\x7f\x53\xe4\x16\x66\xf5\xb3\x60\xf2\x69\x74\x71\x3a\x19\x4d\x92\xf8\x4b\xa3\x26\x7c\x8c\x3f\x0f\x4f\x27\x13\x5f\xaf\x4e\xad\xff\x95\xac\x20\x26\xd8\x11\xa6\xa9\xa0\xce\xc5\x3d\xa3\xa8\x7a\x35\x91\x2f\x77\xbf\x4b\x48\x2d\x19\x6a\xc9\xd3\xda\xbb\x4d\xc8\x8a\x0b\x70\x96\x17\x4b\x87\x6b\xb7\x7b\x79\xae\x65\x4d\x6a\x91\xc4\xa3\xea\xb8\x1f\x44\x8e\x37\x20\xef\x51\xa5\xd1\x85\xf7\xc6\xf8\x6a\xf3\x44\xe5\xd8\x25\xac\xab\xa0\xe0\x69\xc8\xf6\xba\x37\xbe\xc1\x28\xb0\x8a\xdc\xe9\xee\x86\x2b\xc3\x1b\xd0\x46\x19\x36\xb6\x7b\xa8\x3e\xf6\x0c\x00\x6a\xa5\x03\x9a\x56\x54\x64\xbd\x2d\x0e\xdf\xfa\xba\x78\x68\xf2\x76\x30\xd0\x0f\xda\xb5\x0d\x4e\xf0\x0a\x8b\x54\x56\xc3\x46\x90\x40\x6c\x6f\x0b\x60\x92\x8f\x97\x84\x26\xa8\x82\xfa\x80\x47\x6e\xed\x2e\x41\x32\x4d\x6f\xbc\xbc\x6e\x5e\x85\x24\xec\x0e\x92\xd0\x5a\xf4\xe5\x4d\x53\xdc\xbc\xa1\x77\xba\x61\x35\xbd\xf8\xeb\x2f\xbd\x58\x07\x3d\xe8\xec\xc6\x22\xbb\xa5\xb5\x7f\x07\x00\x00\xff\xff\xe1\xab\xfd\x1f\xdf\x38\x00\x00")

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

	info := bindataFileInfo{name: "templates/app.tmpl", size: 14559, mode: os.FileMode(420), modTime: time.Unix(1450153741, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServiceMysqlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5b\x6f\xda\x30\x14\x7e\xe7\x57\x58\x7e\xda\x24\xc6\x28\x95\x2a\x2d\x9a\x26\x51\xa0\x55\xa4\xb5\x43\x85\x75\x0f\x53\x1f\x8c\x7d\x40\xd6\x82\xed\xd9\x4e\x27\x56\xf1\xdf\x67\x87\x24\xd8\x84\xd0\x8b\xb4\x5e\x50\xc8\x39\xdf\xf9\xce\xdd\x7e\x7a\x42\x0c\x96\x5c\x00\xc2\x06\xf4\x23\xa7\x80\xd1\x76\xdb\x41\xe8\xc9\xfd\x23\x84\x87\x3f\x66\x73\x58\xab\x8c\x58\xb8\x92\x7a\x4d\xec\x3d\x68\xc3\xa5\xc0\x28\x41\x78\xd0\x3f\xeb\x7f\xe8\x7f\x72\x7f\xb8\xbb\x53\x9f\x12\x4d\xd6\x60\x9d\x0e\x4e\x4a\x13\xde\x48\x96\x49\xea\x2c\xb0\x99\x95\x9a\xac\x20\x90\x39\xe9\x7c\xa3\xa0\x30\x77\x9b\xaf\x17\xa0\x4b\x53\x85\x68\x0c\x4b\x92\x67\xb6\x90\x9e\xf5\x63\x89\xa1\x9a\x2b\x5b\xb9\x52\x53\x20\xb3\xe3\x40\x86\xff\x05\xf4\xee\xfa\xf2\x3d\x2e\x51\xdb\x0a\x8e\xc7\xc4\x92\x05\x31\x6d\x7e\xcc\xac\xe6\x62\xd5\xe6\x07\x51\xea\x94\x23\xa5\x2a\x62\x25\x07\x12\x2e\x21\x4d\x17\x52\x61\x2c\x11\x14\x0a\xd2\xb7\xb8\xc1\x16\x3d\x3b\xe8\xad\x39\xd5\xf2\x94\x3b\x15\x0f\xa2\x19\x31\x06\x2d\xa5\x0e\x3c\x93\x0c\x4c\xd3\xb5\xa9\x53\xfc\x23\x35\x7b\x85\x5b\x31\xe7\xcc\x35\x12\x68\xa4\x2a\x3b\x0d\x86\x59\xbe\x10\x60\xcd\x11\x02\x87\xfe\xca\x8d\xfd\xec\xda\x2e\x49\x26\xa3\x41\x92\xec\x74\x93\x24\x65\x5f\xda\x38\x1d\xe8\x7e\x3a\x42\xa6\xb4\xda\xa0\xfb\xee\x3a\xbb\xa8\xc2\x7f\x28\x77\x19\x6b\x5e\x51\x34\xc8\xef\x15\x3d\x1e\xe7\x3e\x44\xe7\xbc\x8f\xef\x74\x78\xb5\xe5\x4e\x60\x1f\x7f\xcb\xad\xca\xa3\x4c\xe2\xa9\xd4\xf6\xfc\xbc\x7f\x31\xa7\x6a\xc8\x98\xf6\x22\x67\x80\x64\x39\xec\x1e\xaf\x44\x92\x5c\x83\x1d\x5a\xeb\xbe\xff\xdc\x77\x08\xee\x22\x3c\x11\x4c\x49\x2e\x6c\xcf\x23\xc1\x18\x8c\x1e\xd0\x36\x6c\x8d\xbd\x6d\xff\xf8\x36\xdb\x05\xf2\xc0\xf0\x44\x3c\xde\x6c\xcc\xef\x2c\x9c\xcc\xc8\xf2\x1d\x2c\x7d\x22\x6a\xf9\x51\x74\xd8\xb9\xc7\xd0\xb5\xfc\x28\x3a\x6c\x93\x63\xe8\x5a\xee\xd1\x51\x15\xee\xc0\xc8\x5c\x53\x88\xea\x30\x03\x9a\x6b\x6e\x37\xd7\x5a\xe6\xea\xb9\x16\x88\x95\x83\x46\x98\x6a\xa9\x40\x5b\x0e\xf1\xb4\x38\x49\xa1\x7a\xd0\x27\x6b\x1f\x07\xaa\x16\x79\x37\x54\x8f\x18\x52\xb1\x2a\xca\xeb\x8a\x14\xe8\x20\x1f\x6c\xaa\x1c\xa5\x95\x54\x66\xde\xa0\xa5\xca\xd7\xee\x4a\xcb\x75\x59\x70\xec\xeb\xef\xdf\xcd\xe5\xe1\x9b\x11\x67\x3a\xf5\xa1\xba\x5d\xdd\x2b\x7e\x3f\x9e\x5d\xe0\x32\x57\xbb\x9f\x87\xc8\x27\x37\x1b\x29\x8b\x72\xec\xa7\x25\x00\x6c\x5b\x56\xc7\x73\x39\xbd\x1b\xbb\x8f\xf1\x65\xa8\xfc\xa2\x9c\x46\x90\x57\xe4\xb6\x00\xa5\xcc\x44\xb1\x54\x5b\xee\x64\x3c\xf5\x84\x3c\x1b\xcc\x7e\x96\x5e\x12\xc9\xb1\x23\xb7\xf6\xac\x21\xdc\x3b\x54\xa6\xa1\x22\x1b\xf9\x53\x23\xc2\x46\x07\x57\x2b\x2e\x65\x20\x2c\x5f\x72\xd0\x31\xb1\x8f\x67\x66\x09\xfd\x75\xbb\x1b\xa4\x03\xf8\x6d\x3d\x7e\xcd\x79\xef\xb6\x16\xaa\x81\x0a\xeb\x7e\x00\x9c\x88\x95\xbb\xec\xd4\xf5\x8c\xeb\x78\x43\x8c\xbb\xba\xc4\x7b\xa0\x39\xfc\x2d\x90\x78\xf9\x34\x37\x4e\x04\x8b\x47\x27\x92\xe4\x8b\x8c\xd3\x6c\x33\xa4\x6e\x9f\x18\xbe\xc8\x0a\x67\x97\x24\x33\x87\x4d\xb7\xab\x5d\xd5\x2a\x2b\x35\x88\xe5\xee\xcc\x88\x66\xbe\x98\xf6\x30\x49\xd1\xca\x71\x3b\xed\xa1\xd9\xa6\x9d\xea\x73\xdb\x71\x17\x45\x10\xcc\xdf\x0d\xff\x05\x00\x00\xff\xff\xc3\x37\xc1\xe6\x33\x0a\x00\x00")

func templatesServiceMysqlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceMysqlTmpl,
		"templates/service/mysql.tmpl",
	)
}

func templatesServiceMysqlTmpl() (*asset, error) {
	bytes, err := templatesServiceMysqlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/mysql.tmpl", size: 2611, mode: os.FileMode(420), modTime: time.Unix(1446238118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServicePapertrailTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x7f\x6b\x1b\x39\x10\xfd\xbf\x9f\x42\x88\x42\xa1\x38\x4e\x9c\x50\x8e\x13\xdc\x1f\xbe\x34\x69\x7d\x75\x12\xb3\x4e\x7b\x70\x25\x1c\xca\xee\xd8\xd1\x79\x2d\x2d\x92\x36\x4d\x1a\xfc\xdd\x6f\x24\xed\x7a\x7f\xc6\x0e\x81\xe3\xa0\x6e\x58\xe9\x69\x34\x7a\x33\xf3\x66\x9e\x9e\x48\x02\x0b\x21\x81\x50\x03\xfa\x5e\xc4\x40\xc9\x66\xf3\x86\x90\x27\xfc\x11\x42\xc7\x7f\xce\xaf\x61\x9d\xa5\xdc\xc2\xb9\xd2\x6b\x6e\xbf\x81\x36\x42\x49\x4a\x18\xa1\xc7\x47\xa3\xa3\x83\xa3\x5f\xf1\x1f\x1d\x04\xf8\x8c\x6b\xbe\x06\x8b\x18\xca\x0a\x13\xb8\xfa\x55\xa7\xb5\x4f\x5c\xb8\x7e\xcc\xc0\x5b\x98\x5b\x2d\xe4\xb2\x38\xed\xb7\x3e\x82\x89\xb5\xc8\x6c\x79\xc7\x8c\x67\xa0\xad\xe6\x22\x25\x5f\xa3\xe9\x80\xc0\x70\x39\x24\xef\x52\xb5\x34\xa3\x61\xb6\xdd\xe3\x59\x36\x8c\xd5\x9a\x8d\x46\xc7\x27\x1f\xde\xd1\xc2\xdc\xc6\xff\xdd\x14\xbe\x45\x60\x54\xae\x63\x68\xb8\x36\xe5\xeb\xdb\x84\x9f\x3d\x40\x9c\xbb\x2b\x23\x95\x42\x8f\xab\xcc\x13\xc1\xd8\x64\x7c\xc1\x98\xc7\xd4\x3c\x9e\x69\xe5\xdc\x10\x0d\xc3\x81\x3c\x63\xf2\x35\x38\xfc\x4c\xa5\x22\x7e\xfc\xa8\x62\xfc\x96\xb6\x85\x43\x64\xc9\x6a\x20\xf5\xf8\x00\x79\x1d\xfd\x52\xbb\xc4\x83\xe6\x16\xa3\x50\x9c\xff\xde\xd8\x22\x2d\x7b\x1e\x7e\xb6\x58\x40\x6c\xbd\xef\x69\xaa\x7e\xb4\xac\x15\xae\x0b\x19\x8b\x8c\xfb\xf0\xe0\x05\x45\x06\xa0\x79\x42\x53\xcf\xcc\x90\xaf\xf9\x4f\x25\xf9\x0f\xe3\xf8\xa5\xe4\xa6\xa4\xb3\x61\x67\x1c\xdb\xe0\x3d\x9e\x33\xd6\xb0\xea\xe1\x78\xa2\x05\xdf\x34\xbe\xeb\xbb\x0d\xcb\x18\x78\x7b\xe7\x9c\x3f\xa4\xcd\x65\xc7\x64\xe0\xba\xc9\x41\x9b\x81\x80\x7c\xbc\xc4\x7c\x74\x66\x42\xa0\x4f\x53\x95\x27\x21\x91\xd1\xe1\x2f\x98\xf7\x46\x98\x0e\x33\x74\x4f\xb8\x3c\xe6\x25\x21\xf3\xc0\x5d\x61\xeb\x73\xbc\x38\xb6\x3f\x7c\x1e\x56\x51\xdf\xbb\x4f\xca\x40\xb2\x89\xbc\x57\x2b\x38\xcf\x65\x38\xd0\x8b\xbe\x79\xe6\x92\xb2\x74\x76\x5d\xf3\xfe\x19\x93\x3d\xab\x3d\x29\xf4\xdf\xd3\xb0\x0a\xc1\x66\x9f\xc0\x46\x10\x2b\x9d\x74\xe3\xde\x87\x9d\xdf\x71\x9d\x4c\x50\xd1\xb8\x55\x7a\xff\x89\x20\x5f\xb7\x80\xc2\x06\x7c\xbd\x1f\x3f\x15\xc6\x06\xec\x0e\x77\x9c\xda\xb1\x53\x04\x59\x98\xaa\xe5\x27\xad\xf2\xec\xa5\xe0\x7d\x7e\x78\xf4\x2c\xb7\x08\x3d\xbb\xc7\x1c\x35\xaf\x4e\x8c\xfe\x04\xf8\x5f\x42\x1d\xbb\x2a\x5f\x94\x55\x5e\x8b\x09\x8f\x57\xaf\x7f\x20\x0a\xe4\xb9\x64\xec\x0f\x25\x0a\xa5\xa3\x03\xf7\x3f\xd7\x92\xa1\x3c\xb2\xd6\xa5\xb8\xf9\xe4\x8e\x2f\xb6\xcd\x23\x82\xa5\x6f\x6a\x9b\x01\xa1\x3d\xdb\xe3\x38\x56\xb9\xb4\x93\xa4\x40\x18\xe7\xed\xa1\xc3\x79\x18\x29\x71\xfe\x15\x5e\xd5\x1c\xec\xf0\xbd\xd3\xe4\x9b\x96\xaa\x16\xdc\x77\xd6\x76\x6b\x71\xfd\xab\x42\x96\xab\xdb\x48\xe2\xc8\xa0\xb9\x5c\x02\x79\xbb\x1a\x90\xb7\xf7\x84\xfd\x46\x86\xe3\xe8\xd2\x84\xb9\xa1\xe0\x0d\x41\x79\x86\x6d\x11\x41\xb8\x7e\x81\xfd\xd9\xf5\xf9\x56\x7f\xdc\x8e\x01\xfe\x61\x41\xa1\x19\xf3\x79\x38\xf7\xac\x97\xe7\x9a\x1d\xa0\xea\xb7\xa4\xd3\x48\xcf\x24\xbf\x4d\x21\x71\x3b\x56\xe7\xd0\xea\xa0\x35\xd3\x63\x1d\xe6\x0b\x74\x14\xdf\xb0\xd9\xb4\x9b\x6d\x29\x94\x9e\x69\xb2\x8d\x3e\x2a\xc2\xd8\x5a\xb7\xf0\xbd\x36\x9a\x60\xf5\x2c\x01\xd5\x01\x5f\x82\x66\x3b\x3d\xd2\xb5\x00\x74\x58\x2e\x67\xca\x88\xed\x64\x73\x1d\x4d\x2e\xfe\xfe\x7c\x15\x4d\xfe\xba\xba\xac\x27\x65\xc5\x62\x9d\x72\x90\x49\x45\x70\xf7\xea\xe7\x67\x96\x92\xd6\xad\xf2\xbf\x68\x74\xf9\xcc\x65\x92\x7a\xbb\x54\xc8\x04\x1e\x86\x77\xc5\x42\x23\x14\xe5\xb8\xd4\xe5\xa6\x6f\xae\xea\xa7\x87\x9e\xaa\x04\xba\x13\xd1\xfc\xe4\xf7\x3c\x5e\x81\xed\x2b\xbc\x83\x50\x79\xb1\xc2\x96\xf6\xb0\xa3\xd0\x42\x65\xb4\x63\x71\xf2\x05\x1e\x1d\x36\xf4\xc5\xc3\x6a\x88\x1c\xfe\x14\x19\x7d\x76\x1e\x89\xb0\x38\x45\x98\x25\x24\x7a\xfc\x4f\x53\xaf\xe9\x35\xee\xa9\xdc\xab\xd7\xf1\x07\xda\xad\x9e\xc6\x28\x7a\x95\xdb\x2c\xb7\x75\xd2\x5f\x51\x54\x53\x21\x57\xed\xb0\x7d\xe3\x69\xee\x5d\xdc\xa6\xf5\x0b\x92\xa9\x33\x9e\x97\x56\x2a\x5e\x1d\xa4\xfd\x96\x37\xee\x57\x19\xfb\x37\x00\x00\xff\xff\xcd\x2b\x5b\x59\x4a\x0c\x00\x00")

func templatesServicePapertrailTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePapertrailTmpl,
		"templates/service/papertrail.tmpl",
	)
}

func templatesServicePapertrailTmpl() (*asset, error) {
	bytes, err := templatesServicePapertrailTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/papertrail.tmpl", size: 3146, mode: os.FileMode(420), modTime: time.Unix(1447373411, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServicePostgresTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\x5b\x6f\xe2\x38\x14\x7e\xe7\x57\x58\x7e\xda\x95\x58\x96\xb2\x17\x69\xa3\xd5\x48\x14\x68\x15\x69\xa6\x83\x0a\xd3\x79\x18\xf5\xc1\xd8\x07\x64\x4d\xb0\x2d\xdb\xe9\xa8\x53\xf1\xdf\xe7\x38\x24\x21\x26\x04\x7a\x47\x21\xe7\x7c\xdf\xb9\x1f\xfb\xe5\x85\x08\x58\x4b\x05\x84\x3a\xb0\x4f\x92\x03\x25\xbb\x5d\x8f\x90\x17\xfc\x27\x84\x8e\xbf\x2e\x96\xb0\x35\x19\xf3\x70\xa3\xed\x96\xf9\x07\xb0\x4e\x6a\x45\x49\x42\xe8\x68\x78\x35\xfc\x63\xf8\x1f\xfe\xd1\xfe\x5e\x7d\xce\x2c\xdb\x82\x47\x1d\x9a\x94\x14\x81\x24\xcb\x34\x47\x06\xb1\xf0\xda\xb2\x0d\x34\x64\x28\x5d\x3e\x1b\x28\xe8\xee\xf2\xed\x0a\x6c\x49\x55\x88\xa6\xb0\x66\x79\xe6\x0b\xe9\xd5\x30\x96\x38\x6e\xa5\xf1\x95\x2b\xb5\x09\xe2\xf6\x36\x88\x93\x3f\x81\xfc\x76\x7b\xfd\x3b\x2d\x51\xbb\x0a\x4e\xa7\xcc\xb3\x15\x73\x5d\x7e\x2c\xbc\x95\x6a\xd3\xe5\x07\x33\xe6\x9c\x23\xa5\x2a\x11\xa5\x0d\xa2\x30\x21\x6d\x17\x52\xe5\x3c\x53\x1c\x0a\xa3\xef\x71\x43\xac\x06\x7e\x34\xd8\x4a\x6e\xf5\x39\x77\x2a\x3b\x84\x67\xcc\x39\xb2\xd6\xb6\xe1\x99\x16\xe0\xda\xae\xcd\x51\xf1\x87\xb6\xe2\x0d\x6e\xc5\x36\x17\xd8\x48\x60\x89\xa9\x78\x5a\x16\x16\xf9\x4a\x81\x77\x27\x0c\x20\xfa\xa3\x74\xfe\x7f\x6c\xbb\x24\x99\x4d\x46\x49\xb2\xd7\x4d\x92\x54\x7c\xe8\xb2\x89\xa0\x87\xf9\x84\xb8\x92\xb5\x65\xee\x0b\x76\x76\x51\x85\xf7\xe4\xd9\x68\xe7\x37\x16\xf3\x74\x39\xe0\xbc\xb2\xd3\xf2\xe0\xc1\xf0\xd3\xc1\x1e\xe2\xc4\x08\x42\x90\xe7\x63\xac\x99\x7b\x0d\x7e\xfa\x39\xf7\x26\x8f\xd2\x49\xe7\xda\xfa\x7f\xfe\xfe\x6b\xb4\xe4\x66\x2c\x84\x0d\x22\x24\x60\x59\x0e\xfb\xc7\x1b\x95\x24\xb7\xe0\xc7\xde\xe3\xf7\x6f\x87\x36\xa1\x7d\x42\x67\x4a\x18\x2d\x95\x1f\x04\x24\x38\x47\xc9\x23\xd9\x35\xfb\xe3\xc0\x1d\x1e\xdf\xc7\x5d\x20\x8f\x88\x67\xea\x69\x5e\x66\xbb\x39\xa1\x11\xf9\x3d\xac\x43\x2e\x6a\x79\x17\x41\xb3\x89\x4f\x11\xd4\xf2\x2e\x82\x66\xd3\x9c\x22\xa8\xe5\x81\x20\x2a\xc7\x3d\x38\x9d\x5b\x0e\x51\x41\x16\xc0\x73\x2b\xfd\xf3\xad\xd5\xb9\xb9\xd4\x0b\xb1\x72\xa3\x23\xe6\x56\x1b\xb0\x5e\x42\x3c\x3b\x28\x29\x54\x8f\x1a\xa6\x6a\x5d\x52\x6d\xf6\x7e\x13\x11\x19\x49\xd5\xa6\x28\x35\x16\xac\xa1\x43\x42\xbc\xa9\x41\xab\x5e\x73\x9d\x05\x4e\xcf\x4d\xa8\xe3\x8d\xd5\xdb\xb2\xf8\x34\xf4\x42\x78\xb7\xd4\xc7\x6f\x26\x52\xd8\x34\x44\x8b\xcb\x7b\x50\xfc\xfe\x79\xf5\x2f\x2d\xd3\xb5\xff\x79\x8c\x7c\xc2\x39\x49\x45\x94\xe6\x30\x39\x0d\xc0\xae\x63\x97\x5c\x4a\xeb\xfd\x14\x3f\xa6\xd7\x4d\xe5\x57\xa5\x35\x82\xbc\x2d\xbd\x05\x2e\x15\x2e\x0a\xa7\xda\x7c\x67\x43\xaa\x07\xe6\x62\x3c\x87\xd1\x7a\x4d\x30\xa7\x8e\xe1\xda\xb3\x96\xf0\xe0\x50\x99\x89\xca\xd8\x24\x9c\x24\x11\x36\x3a\xcc\x3a\x71\xa9\x00\xe5\xe5\x5a\x82\x8d\x0d\x87\x78\x16\x9e\xf1\xef\x77\xfb\x71\x3a\x82\xdf\xd5\x43\xd8\x9e\xfd\x7e\x67\xad\x5a\xa8\x66\xe9\x8f\x80\x33\xb5\xc1\x0b\x10\x3d\xbd\xec\x51\xfe\x89\x39\xbc\xd1\xc4\x0b\xa1\xbd\x05\x3a\x20\xf1\x22\x6a\x6f\x9f\x08\x16\x0f\x50\x24\xc9\x57\x99\xe4\xd9\xf3\x98\xe3\x62\x71\x72\x95\x15\xfe\xae\x59\xe6\x8e\xfb\x6e\x5f\xbe\xaa\x5b\x36\xe6\x88\x08\x4f\x91\x68\xf2\x8b\x99\x6f\xe6\x29\xda\x3d\xb8\xdc\x1e\xdb\x9d\xda\xab\x3e\x77\x3d\xbc\x3f\x82\x12\xe1\xca\xf8\x2b\x00\x00\xff\xff\xa1\x5a\xfe\x19\x4a\x0a\x00\x00")

func templatesServicePostgresTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePostgresTmpl,
		"templates/service/postgres.tmpl",
	)
}

func templatesServicePostgresTmpl() (*asset, error) {
	bytes, err := templatesServicePostgresTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/postgres.tmpl", size: 2634, mode: os.FileMode(420), modTime: time.Unix(1446238118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServiceRedisTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xd1\x6f\xda\x3e\x10\x7e\xe7\xaf\xb0\xfc\xf4\xfb\x49\x8c\x41\x27\x75\x6a\x34\x4d\x42\x0c\xaa\x48\x5b\x87\x80\x76\x0f\x55\x1f\x8c\x7d\x50\x6b\x89\x1d\xd9\x4e\xb7\xaa\xe2\x7f\xdf\xd9\x09\x10\x27\x1d\x6d\xa5\xb5\x80\x22\xdf\xdd\xf7\xdd\x7d\x77\xb6\xf3\xf4\x44\x04\x6c\xa4\x02\x42\x2d\x98\x07\xc9\x81\x92\xdd\xae\x47\xc8\x13\x7e\x09\xa1\xe3\x1f\xcb\x15\xe4\x45\xc6\x1c\xcc\xb4\xc9\x99\xbb\x01\x63\xa5\x56\x94\x24\x84\x9e\x0d\x47\xc3\x77\xc3\x0b\xfc\xd0\x7e\xe5\x3e\x67\x86\xe5\xe0\xd0\x87\x26\x35\x04\xae\x7e\x61\x8e\xad\x99\x85\xc6\x1a\xae\xae\x1e\x0b\x08\x30\x4b\x67\xa4\xda\xd6\x10\x55\x00\x6c\x58\x99\xb9\x60\x1d\xc6\x06\xcb\x8d\x2c\xdc\x3e\x83\xda\x91\x88\x9a\x81\x48\x25\xe0\x37\xad\x03\x76\xfb\x48\x9a\x2a\xeb\x98\xe2\x10\x38\xbb\x59\x9c\x4c\x02\x8d\x9c\xf1\x7b\x18\xb8\xb3\x41\x2e\xb9\xd1\x7f\x4b\x08\x1d\x57\xf7\x40\x1c\x22\x12\xbd\xc1\x54\x2a\x4e\xe2\x34\x29\xb1\xf8\x4e\x52\x73\x66\xed\x2f\x6d\xc4\x1b\x64\x89\xab\xbf\x56\x88\x2b\xc8\x7f\x48\xb0\x06\x62\x20\xd7\x0f\x20\xfe\xef\x12\x2d\xcb\xb5\x02\x67\x9f\x2f\xfc\xab\xb4\xee\x13\xb6\x39\x49\xa6\x93\xb3\x24\xa9\x7c\x93\x24\x15\x9f\x4f\xd4\x79\x33\x9f\x10\x5b\xa3\x76\xe8\x6e\x0a\xfe\x3c\xd5\x91\x05\xe3\x3d\xc5\x69\x86\x03\x72\xaf\x81\x4f\xbf\x97\xae\x28\xa3\x62\xe8\x5c\x1b\x77\xfe\xe1\xe3\xc5\x8a\x17\x63\x21\x8c\x37\x21\x00\xcb\x4a\xa8\x1e\x67\x2a\x49\x2e\xc1\x8d\x9d\x6f\xe6\x2d\xa1\x0b\x28\x32\xc9\x99\xa7\xba\x34\xba\x2c\x68\x1f\x31\x8c\xcc\x99\x79\x9c\x2a\x31\xd7\x52\xb9\x81\x07\x02\x6b\x29\xb9\x23\xbb\x66\xcf\x8e\x54\xfe\xf1\x9f\x50\x05\xa0\x16\xcf\x54\x3d\x2c\x40\x48\xdb\xdc\x3a\x11\xd1\x02\x36\x5e\xa6\x83\xdd\x47\x47\x3a\x2d\xc0\xea\xd2\x70\x88\x94\x5a\x02\x2f\x8d\x74\x8f\x55\x2e\x2f\x34\x29\x76\x6e\xb4\x6a\x6e\x74\x01\xc6\x49\x88\x47\x0a\x2d\xc1\xb5\xd5\x49\xe3\xeb\x20\xfb\xd3\xa5\xdf\x74\x8f\x18\x52\xb5\x0d\x8a\xa3\x6e\x0d\x1f\xe2\x8b\x4d\x0b\xa4\x74\x9a\xeb\xcc\x03\x3a\x1e\x64\x9c\x19\x9d\xd7\x3d\xa0\xbe\x25\x7e\x6d\xa5\xdb\x2b\x13\x29\x4c\xea\x4b\xa5\xa3\xe1\x20\xfc\xbf\x1f\x9d\xd3\x5a\xab\xea\xef\x2e\xca\x09\xa7\x37\x15\x91\xc6\x7e\x9e\x1b\x01\xbb\xce\xc0\x4f\xfc\xf1\x50\x6d\x9c\x17\x85\xcd\x98\x75\x32\x04\xec\xf7\xda\x1b\xe4\x6d\x29\xbb\xa8\x94\x0d\x28\x64\xdb\x82\x39\xec\xfb\x54\xd8\xa8\x9c\xfd\x69\x70\xb2\xa4\xce\xdc\xbe\xb2\xa4\xee\xbc\xbf\xa6\xae\x71\xe9\x34\xde\x2d\x92\xcf\x98\xcc\xf0\x10\x33\x53\xc5\xd6\x19\xf8\x36\x6c\x58\x66\xa1\xdf\x76\xfe\x26\x95\x36\xf5\x55\x74\x5d\x6c\x0d\x13\x3e\x1f\x67\xca\xd8\x35\x24\x75\xa5\xc5\xe1\xdc\x3f\x88\x10\x5d\x08\xc7\xba\x9f\x6d\xe7\x15\x5e\x69\x51\x70\xa7\xdf\x2d\x80\xa9\xda\xe2\x8d\x7a\x98\xfd\xb8\x29\x57\x65\x1e\xe2\x27\x59\x69\xeb\x8b\x92\x8e\x62\x97\x78\x86\x9b\x96\xb6\xbc\xf1\x40\x1c\x33\x0c\x8d\x59\x3a\xc6\x7f\x86\xe4\x5b\xf9\xc5\xdb\x2e\x8c\xc7\x6d\x73\x40\xa2\x7d\x8f\x07\xcb\x5d\x77\x50\x7a\xfb\xdf\x5d\x0f\x5f\x21\x40\x09\xff\xd6\xf0\x27\x00\x00\xff\xff\x64\xe5\x91\xd9\x4d\x08\x00\x00")

func templatesServiceRedisTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceRedisTmpl,
		"templates/service/redis.tmpl",
	)
}

func templatesServiceRedisTmpl() (*asset, error) {
	bytes, err := templatesServiceRedisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/redis.tmpl", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1446238118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServiceWebhookTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x92\x4d\x6f\xb2\x40\x10\xc7\xef\x7c\x8a\xcd\x5e\xbc\x28\xa0\xcf\x73\x29\x37\xd3\x97\x53\x63\x8d\xa0\x9e\x71\x19\x75\x23\x30\x9b\x65\x30\x6d\x88\xdf\xbd\xbb\x4b\x51\x48\xeb\xa5\x4d\x48\xc8\xce\xcb\x6f\x66\xfe\x33\x4d\xc3\x32\xd8\xcb\x12\x18\xaf\x40\x9f\xa5\x00\xce\x2e\x17\xaf\xf1\x18\xe3\xf3\x6d\x9c\x40\xa1\xf2\x94\xe0\x05\x75\x91\xd2\x06\x74\x25\xb1\xe4\x2c\x62\x7c\x16\x4e\xc3\x49\xf8\x60\x3e\x3e\xb6\xc1\xcb\x54\xa7\x05\x90\x89\xe0\x11\xb3\xe9\xc6\xb6\xd6\xf9\xf5\x61\x9e\xc9\x87\x02\x97\x1b\x93\x96\xe5\xc1\xe5\x39\xc7\x13\x54\x42\x4b\x45\x1d\x7b\x0b\xbb\x23\xe2\x89\xad\x57\xaf\x63\x06\xfe\xc1\x67\xa3\x23\x91\xaa\xa2\x20\x38\x68\x99\xf9\x02\xcb\x33\xbe\x9b\x5f\x11\xe8\x54\x9c\x26\x36\x38\x98\xce\xfe\xfd\x1f\x71\x87\xbc\xb4\x64\xfe\x58\x57\x84\x45\x82\x4a\x8a\x5f\xf5\x31\xa4\x2d\x90\xe4\x5e\x8a\xd4\xba\xff\xcc\xf4\xbe\xb8\x7c\x05\x15\xd6\x5a\x40\x4f\xb6\x7e\xa1\xea\xc7\x22\xed\x60\x51\x14\x2f\xe2\xb8\xde\xdd\x0a\x5c\xab\x76\x9b\x32\xb1\x53\x3f\xbc\xd9\x97\x1a\x15\x68\x92\xd0\xe7\x1a\x7b\xdc\xee\x3e\xc1\x13\xd8\xa4\xc6\xf6\xb5\xe7\xd1\x50\xc2\x4e\x89\xb6\x17\x6b\x9a\x6b\x37\xd4\x2d\xfc\xbb\x46\x83\x24\x53\x9e\x50\x60\xee\x86\xb0\x2b\xe5\x3d\xe7\x73\x99\x29\x94\x25\x0d\x89\xf6\x86\x9c\x5c\x9d\x6c\x7d\xf1\xde\x6a\x52\x35\xdd\xbf\xb8\x4d\x9a\xd7\xc0\xef\xe0\x1c\xc6\x33\xb7\xde\x30\x28\x33\x7b\xf5\x9f\x01\x00\x00\xff\xff\x54\xf9\xe1\x63\x0d\x03\x00\x00")

func templatesServiceWebhookTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceWebhookTmpl,
		"templates/service/webhook.tmpl",
	)
}

func templatesServiceWebhookTmpl() (*asset, error) {
	bytes, err := templatesServiceWebhookTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/webhook.tmpl", size: 781, mode: os.FileMode(420), modTime: time.Unix(1446238118, 0)}
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
	"templates/service/mysql.tmpl": templatesServiceMysqlTmpl,
	"templates/service/papertrail.tmpl": templatesServicePapertrailTmpl,
	"templates/service/postgres.tmpl": templatesServicePostgresTmpl,
	"templates/service/redis.tmpl": templatesServiceRedisTmpl,
	"templates/service/webhook.tmpl": templatesServiceWebhookTmpl,
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
		"service": &bintree{nil, map[string]*bintree{
			"mysql.tmpl": &bintree{templatesServiceMysqlTmpl, map[string]*bintree{}},
			"papertrail.tmpl": &bintree{templatesServicePapertrailTmpl, map[string]*bintree{}},
			"postgres.tmpl": &bintree{templatesServicePostgresTmpl, map[string]*bintree{}},
			"redis.tmpl": &bintree{templatesServiceRedisTmpl, map[string]*bintree{}},
			"webhook.tmpl": &bintree{templatesServiceWebhookTmpl, map[string]*bintree{}},
		}},
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

