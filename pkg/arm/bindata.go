// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/azuredeploy.json
// data/startup.sh
package arm

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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xdb\x38\x12\x7f\xef\x5f\x21\xe8\x0e\x70\xb3\x88\x6c\x39\x71\x7b\x45\xde\xd2\x64\x7b\x35\x6e\xd3\x1a\x51\xbb\xf7\xb0\x08\xba\x34\x35\xb6\x79\x91\x48\x81\xa4\x9c\xba\x81\xff\xf7\x03\xa9\x8f\xda\xfa\x76\x2c\x27\x4e\x6a\xf5\x21\xa9\x44\x0e\x67\x86\xbf\x19\xce\x07\x73\xff\xca\x30\x0c\xc3\xfc\xa7\xc0\x33\xf0\x91\x79\x66\x98\x33\x29\x03\x71\xd6\xeb\x45\x6f\xba\x3e\xa2\x68\x0a\x3e\x50\xd9\x45\x3f\x42\x0e\x5d\xcc\xfc\xf8\x9b\xe8\x9d\xd8\xfd\x37\x96\xdd\xb7\xec\x7e\xcf\x85\xc0\x63\x0b\x35\xee\x0b\xf8\x81\x87\x24\x74\xff\x27\x18\xfd\x87\x79\x1c\xad\x80\x19\x95\x40\xe5\x9f\xc0\x05\x61\x54\x2d\xd4\xef\xda\xea\x5f\x32\x20\x40\x1c\xf9\x20\x81\x0b\xf3\xcc\xb8\x5f\xc6\x6f\xe7\x88\x13\x34\xf6\x60\xed\x25\x07\xc1\x42\x8e\xf5\xcb\xbf\xf4\x2b\xf5\xdc\xa7\xbf\xe9\x41\x72\x11\x80\x5a\xe6\x8a\x60\xce\x04\x9b\xc8\xee\x27\x90\x77\x8c\xdf\xf6\x68\xf4\xd3\x01\x1c\x72\x22\x17\xff\xe6\x2c\x0c\x44\xcc\x46\x3a\x1d\x05\x64\x85\xd7\x13\xbb\xff\xd6\xb2\x4f\xad\x53\x3b\x3b\xce\x63\x18\xc9\x78\xd4\xfd\xbd\xd1\xbd\x42\x94\x4c\x40\xc8\xee\x1f\xf1\x07\x63\xb9\xcc\xce\xa1\xc8\xd7\xac\x51\x31\xb5\x30\xf3\x83\x50\x42\x76\x48\xc0\x59\x00\x5c\x92\x48\xf0\xb5\x6f\xfa\xbb\x88\x99\xbf\x0e\xbd\x8c\x1a\x56\x9f\xfc\xc4\x1c\x0f\xc8\xf3\xd8\xdd\x37\x21\x66\x19\x0e\x36\xe1\x66\x5d\x71\x18\x83\x50\xe3\xcc\x73\x45\xba\x82\xac\x1e\xee\x82\xc0\x9c\x04\x89\x0a\xf5\x1c\xc3\x71\x3e\x1a\x92\xa3\xc9\x84\xe0\x06\xf3\x25\xa1\x5a\xd3\xe7\xae\xcb\x41\x88\x11\x87\x09\xf9\xae\x88\xfd\xb6\xc1\xe4\x11\xe3\xf2\x1a\xd1\xa9\x56\xca\xc9\x89\x75\x72\x52\x3b\x99\x70\xc0\x09\xdf\x43\x3a\x66\x21\x75\xeb\xe6\x04\x9c\x30\xb5\x6d\xe6\x99\xd1\xb7\xfb\xb5\x83\x99\x64\x98\x79\x8a\xfe\x17\x1c\xd4\xd1\x8e\x6c\x62\x73\x25\x44\xf3\xd6\xe4\xff\xcd\x2c\x9d\xb2\x2c\xfc\x92\x7f\x7b\xf3\xaa\xf8\xfb\xf2\xf8\x99\x1b\x2c\xa1\x13\x8e\x0e\xe6\x7a\x30\xd7\xfc\xe0\xe7\x63\xae\xc5\x6b\x34\xc6\x9f\x8a\x11\xf6\x09\x80\x8a\x9f\x27\x45\xe0\xbb\xac\x9b\xc9\xcf\xdc\x16\x7e\x27\x07\xf8\xad\xc2\x2f\xeb\xff\xd7\x86\x3f\x05\xfe\xc4\x93\x02\x70\x30\x38\xdd\x39\x02\x4f\x5f\x10\x02\x77\x1a\xaf\xcc\x09\x97\x21\xf2\xe2\xff\x3e\x49\xa4\x32\xa7\x20\x37\x0d\x52\x50\xb4\x15\x4e\x80\x30\x94\xda\x4d\x32\x2a\xda\xb0\x8a\x60\x46\x0f\xee\xdb\x51\x92\xd7\x7b\x57\xbc\x43\x37\xb9\xb7\x05\xbe\xc1\x14\xe1\x98\x82\xdc\x2a\x6e\x8a\x48\xb4\xe6\x33\xb2\xa0\x4d\xe5\x3c\x19\xec\x17\x14\x83\x70\xec\x11\x3c\x1c\xc5\x66\x06\xf5\x60\xfc\x97\x65\xbf\xb3\xec\x7e\x9b\x60\x24\x81\xc5\x59\x28\x81\x6f\x8a\xc8\x94\x7b\x2f\x59\xfe\x0a\xe4\x8c\xb9\x8a\xa8\x23\x91\x24\x78\x5d\xdb\x19\xf0\x98\xe2\x36\x2c\x26\x9c\x70\xf6\x1e\x89\x1c\x8d\xad\xf4\xed\x31\xe4\xbe\x47\x1e\xa2\x18\x78\x13\x5d\xf7\xed\x02\x5d\xbb\x10\x00\x75\xc5\x67\x5a\x88\x78\xf3\xaf\xa4\xfe\x31\x74\x5f\x77\x1a\x6c\x79\xe7\xd8\xe8\xa4\x5b\xd0\x39\xba\x59\x97\xf7\xa6\xc5\x8d\xf6\xc6\x0f\xdc\xe8\x31\xc2\xb7\x40\xdd\xe4\x30\x60\xcc\xdb\xca\xda\x63\x72\xc5\xa6\x58\x60\x70\x05\x3e\x67\xc2\x75\xc1\xca\x1d\x8e\x2e\x18\x9d\x90\x69\xc8\xb5\xf0\x5b\xb1\x95\xd0\x6c\xcb\x0d\x05\x9c\xcc\x91\x84\x62\xfb\xb8\x5c\x50\xe4\xd7\x47\x23\x19\xb4\xd4\x2e\xaa\x27\x11\xbd\x42\xbb\x38\xcc\x3e\xc5\xde\xb2\xfc\x4b\xb3\x6d\x25\x51\xdc\xf3\x09\xc9\x14\x64\xd5\xc3\xd2\x8c\xbd\x68\xd8\x4f\x6b\x27\x74\xba\x75\x6e\x1f\x30\x2e\xad\xca\x4c\x62\x23\x74\xe4\x6d\x6a\xa3\xcd\xc5\x8c\x62\x24\x5f\x57\xef\xf1\x9a\xbb\x53\xfb\x9b\x7a\x80\xce\xd1\xb1\xd1\xe9\x15\xd8\x75\xf2\xae\x1e\x00\x35\xc8\x8d\xe9\xa8\x40\xd0\x3c\x33\xde\xd9\x35\xc3\x81\xa2\xb1\x07\x1f\x3c\x86\x24\xa1\xd3\xe1\xc8\x3c\x33\x26\xc8\x13\x50\x33\xad\xc4\x0f\x3c\xbe\x2a\xcb\x1c\x52\xfa\x61\x6b\x85\x26\x84\x9a\x6a\x94\xb8\x1e\x7c\x21\x3e\xb0\x50\x0e\xe9\x15\xa1\xa1\xd4\xb0\x1c\xd4\x4c\x53\x82\x5e\x12\x21\x39\x19\x87\xc9\x41\x73\x09\x13\x14\x7a\x55\xf1\x99\x11\xc3\x7f\x5c\x1e\x9a\x66\x98\x6b\x4f\xf5\x7a\x59\xd1\x8b\xed\x73\x6b\x3d\x67\xf3\xa4\x47\xcb\x9e\xb5\x00\xd5\x99\xe2\xc1\xc3\xe4\xe5\x8b\xed\x61\x30\xa8\x4b\x80\x0f\x2e\x66\x43\x17\x53\xaf\xd2\x5f\xd2\xc7\x0c\x06\xa7\x4f\xee\x64\x1a\x45\x53\x2c\x94\x8d\xe2\xa4\x48\xb8\x3d\x8a\x8d\x08\x95\xc0\xe7\xc8\x1b\x52\x07\x30\xa3\xae\x9a\xf2\xa6\x46\xa3\x34\xf4\xc7\xc0\x3f\x4f\x46\x89\x30\xb5\x15\xd9\x86\xe7\xe8\x8b\x39\x0f\x1e\x53\xab\xf5\xae\xa3\x7d\x0b\xd8\x93\x5a\x83\x23\x19\x47\x53\xe8\x89\xe8\xe7\x39\xc6\x2c\xa4\xb2\xbe\xda\xf0\xc6\xb2\xdf\x5a\xfd\x37\x6d\x56\x76\xd4\xf8\xe8\xac\xe8\x5e\xc3\x54\x39\xdc\x85\xb3\xc6\x55\xc1\xdc\xda\x32\x64\x34\xf3\x4b\x2c\xb7\x23\x11\x75\x11\x77\xbf\xfd\x71\xed\xb4\xa1\xbd\x8b\xe8\xc2\x45\x0f\xcd\x11\xf1\xd0\x98\x78\x44\x2e\x1c\x68\xa0\xbe\xb7\x96\x3d\xb0\x4e\x6d\x2b\xe0\x30\x27\x90\x6d\x14\x6c\xa5\x46\x24\x1e\x7a\x0f\x24\xba\x9e\xa3\xce\x21\xc9\xc3\x82\x40\xc3\x0c\x3c\x24\x27\x8c\xfb\x1f\xd4\xd9\x77\xc9\x7c\x44\xe8\x85\x52\x6f\xb1\xa5\xa5\xc3\xbf\x06\x2e\x92\xb0\x3e\xfe\x74\x1f\xeb\x92\x98\x05\x8b\x62\xd5\xe0\x98\xed\xb5\xcd\x88\x77\xff\x22\xc6\x66\x81\x06\xd2\x5a\x16\x63\x41\x75\x75\x71\x9b\x1d\x4f\x82\x86\x0e\x09\x92\xad\xb7\x3a\xc7\x86\x92\x66\x48\x5d\xf8\xfe\xfa\xe8\xe8\xa6\xcd\x6a\x69\x52\x0d\xda\x13\x17\x96\xb9\xd9\x31\x54\x07\xc7\x04\xe1\x06\x30\x28\xe9\x95\x3c\x21\x0c\xb6\xab\xd6\x66\x7a\x45\x2a\x32\x9c\x53\x90\x9d\xdc\xf6\x37\x21\x56\x78\x51\x46\x91\x5c\xb9\x69\xf6\x30\xca\x45\xc5\xbc\x7a\x08\xef\xb2\xd6\x9c\x5a\x10\x25\xb8\x35\x13\x8a\x32\xb8\x73\x8c\xc1\x03\x8e\x24\xb8\xb1\xfc\x84\x4e\x75\x01\x57\x65\x73\x45\xda\x8b\xe6\x0d\x47\x1f\x18\xbf\x43\xdc\x8d\x86\x97\x78\x64\x12\xb4\x57\x51\x26\x01\xd6\xb4\x5a\xac\x28\xfb\x88\x2f\x4a\x99\xcf\x0e\xde\xba\xfc\x1c\xf7\xe5\x5a\x4f\xb7\x4a\xcd\x4a\xe5\x5a\x71\x3f\x31\xfe\xb9\x7d\xaa\x15\xd9\x46\xb0\xb3\x12\x7a\xb0\x9d\xd5\xe5\xe4\xd9\x49\x2e\x58\xe4\x7b\xca\xbb\xc8\x8d\xc4\x6f\xea\xce\x72\x4b\x2c\x5b\x8c\x13\x63\x24\x5d\x21\x3c\x23\xb4\xd9\x01\x55\x15\x26\x3e\xdb\x83\x2a\x77\x54\xaf\x60\xb1\xc2\x05\x6f\x74\xda\x94\xc5\xe6\x6a\xcf\x7f\xc6\xc8\x3b\x6d\x61\xa6\xc7\xca\xdc\x6f\xed\x54\xc9\x88\xf3\x40\xab\x78\xa0\x72\x0a\x90\xa0\xd7\x9a\x21\xee\xde\x21\x0e\x23\xce\x26\xc4\xab\xb8\xf0\x31\xf7\x1d\xf2\x03\x72\xba\xfc\xf3\x4a\xbd\x56\x9a\x6c\xb4\x5c\x0c\x9f\xda\xd5\xf2\x11\x61\xd5\x05\x93\xba\x82\xc4\x06\x5e\xe6\x01\xc0\xde\xd0\x8f\x36\xbb\xec\xc2\x44\xad\x8e\x90\xeb\x13\xfa\x55\x00\x4f\x40\x8b\x3d\x16\xba\x56\x28\x72\x8d\xf7\x74\x4a\x2c\x04\xdf\x02\xe6\x29\x2d\x8f\xd0\xf0\x7b\xf3\x52\xb5\xe9\x12\xa1\xa2\xa3\x11\x12\xe2\x8e\x71\xf7\x3c\x94\x33\xa0\x92\xa4\x36\x5a\x1d\x6a\x98\x42\xcc\x1a\x04\x2e\xfa\xa4\xfc\x0f\x2c\xaa\xf1\x92\x3c\xf5\x07\xb4\xa6\x7a\x0b\x8b\x4b\x24\x51\x8c\x7d\xc7\xf9\x38\x4a\x96\x39\x17\x8e\xe4\x84\x4e\xd3\xf2\xc7\xea\xc7\xbc\x7b\x29\xe7\x1b\x49\x25\x9d\xd9\x9b\x31\x1f\x7a\x3f\xf7\xb1\xd7\x15\x62\xd6\x43\xa1\x9c\x31\x4e\x7e\x80\xfb\xed\x56\x89\x56\x4b\xb3\xfc\x68\x37\x0a\x01\x58\x3d\x2f\xff\xb6\xf0\x76\x56\x54\xf0\xa9\x45\x2d\xf1\xd1\x14\xae\x61\x02\x1c\x68\xc5\x05\x33\xa3\xc4\x70\xd7\x9c\xcf\x30\xa2\x15\x7d\xd5\x71\x81\xb1\x5c\x2a\x47\x98\x77\x97\x7a\x59\xed\x24\xcb\x29\x7c\x42\xbe\xf2\x64\xa5\x91\x60\x49\x04\x68\x32\x71\x49\xc4\x6d\xb5\x28\x58\x07\x0e\x3a\x7d\xb8\x06\xe4\xfe\x97\x93\x5c\x95\x67\x7d\x3c\x07\x24\xe1\x73\x7a\xc3\xf4\x03\x67\xbe\x66\xb6\xe9\x6d\x96\x1d\x06\x3f\x3d\xf8\x2e\x81\x0a\x9d\xb5\xfc\x2a\x89\x7a\x89\x2a\x56\x0e\x89\x72\x17\xfa\xb4\x81\x8a\x4a\x75\xb0\x28\x48\xf8\xeb\x03\x96\x50\xb2\xaf\xc1\x94\x23\x17\xae\x08\x65\xfc\xe7\xce\x96\x55\x19\x39\x93\x80\x25\xb8\x0e\x48\x49\xe8\xb4\x3c\x09\x32\xa3\xeb\xd3\xb1\xc4\xef\x91\x80\xb7\x83\xdf\x29\x66\x2e\x18\xaf\x1d\x89\xb8\x0c\x03\xe3\xef\x58\x9e\xbf\x8f\x1a\x87\x17\xda\xff\x8b\x19\xf0\x75\x1c\x9f\xeb\x3f\x59\xfc\xbd\x0c\xb5\x11\x43\x2b\x2c\x17\x91\x4e\xac\xe3\x22\x14\x92\xf9\x4e\xc4\x7e\xc9\xb8\x8f\x88\xba\x1e\xf0\x55\x4b\xe8\xda\x2f\xb8\x6a\xfd\xa0\x3f\x86\x3a\xd4\xac\x37\xf3\x81\x43\xa5\xe4\xbd\xaa\x58\xeb\x6d\x3f\xd4\xab\x1f\xaf\x5e\xdd\x1e\x04\x9e\x4b\xb5\x5a\x43\xec\x61\x74\xab\xae\x5a\xec\xa0\xf6\x5d\x68\x0c\x8f\x56\xf9\x6e\xc7\x14\x7f\xf5\xba\xf7\x2a\x64\xde\x6f\x70\x17\xbd\x19\xdb\x6b\x4b\xed\xcb\xe5\x37\xa3\x3e\x5d\xac\xaf\xff\x1f\x9a\x05\xd9\x67\x6f\x9b\x05\x8d\xdc\x54\x4e\x9a\xad\x2b\x06\xfb\xd0\x2a\x48\xcf\x92\xc7\xcc\x95\x1f\xb9\x51\xb0\x2f\x11\x42\x5d\x35\xb5\x0c\x86\xad\x35\x09\x8a\x37\x7b\x57\x99\x77\x3b\xc7\xef\xa3\x34\x08\x2a\xac\xe0\xd0\x1e\x68\x1f\xd0\x1b\x7a\xce\x7d\x6d\x0e\x34\x02\x78\x4a\xe9\xd0\x1a\x38\xb4\x06\x56\x44\x38\xb4\x06\x0e\xad\x81\x17\x5b\x13\x69\xd4\x18\x78\xba\xe2\x40\xc9\xe2\xcf\xb7\x29\xa0\xa5\x79\xe9\x2d\x01\xfd\x5b\x8c\x07\x93\x85\x32\x08\x65\xc4\xc5\xab\xe5\xab\xff\x07\x00\x00\xff\xff\x15\x9c\x92\x7e\xac\x51\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _startupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x7f\x4f\xe3\x46\x10\xfd\x7f\x3f\xc5\x34\x54\x4d\x7b\xaa\x6d\xb8\x9e\x50\xcb\x71\x48\x10\x02\x42\xfc\x08\x82\xf4\xa4\x53\x55\xa1\xf5\xee\x18\xaf\xb2\xde\x35\x3b\xe3\x08\x37\xca\x77\xaf\x6c\x43\x08\x21\x05\xae\xff\x25\x33\x6f\xe6\x3d\xbf\xb1\x67\x36\x7e\x48\x52\xe3\x92\x54\x52\x0e\x11\xde\x0b\xb1\x01\xe3\xd1\xe1\x68\x07\x12\x64\x95\x68\x47\x85\xa4\xbb\x58\x27\x3e\x98\x5b\xe3\xa2\xaa\x24\x0e\x28\x8b\x48\x3b\x8a\x95\x77\x19\x18\x02\x55\x85\x80\x8e\x6d\x0d\xb9\x0c\x5a\x79\x8d\xfa\x33\x18\x16\x1b\x50\x06\x9f\xca\xd4\xd6\x40\xb9\xaf\xac\x76\x7d\x86\x14\x85\xb8\x1e\x5e\x7d\x3d\x19\x0c\x6f\xc6\xdf\x2e\x87\x5f\xba\xce\xc2\x64\xf0\x17\x44\x19\xf4\x5a\x62\xaa\xa9\xe9\x6e\x6e\x13\xc9\xbe\x30\x2a\xf2\x25\x3a\xca\x4d\xc6\x91\xf3\x1a\x7b\xf0\xf7\x67\xe0\x1c\x9d\x00\x00\x78\xd6\x6e\x15\x2f\x32\x23\x04\xd5\xc4\x58\x28\xb6\x40\xec\x4b\xd0\x5e\x4d\x30\xc4\x84\x61\x6a\x14\x8a\x62\x92\x51\x7c\x9f\x51\x43\x9f\x68\x9c\x26\xa4\xd3\x2d\x81\x2a\xf7\xd0\x5f\xfc\x07\x48\xa6\x32\x24\xd6\xa4\x49\x57\x0e\xd0\x94\xc0\x6d\x28\xef\x2a\xcf\x12\x60\x13\x36\xfb\xb0\xb7\xd7\xea\xcf\x88\x65\x2a\x0a\x5f\x39\x5e\xad\x13\x01\x89\x7d\x40\xe5\x1d\x44\x57\x2f\xb2\xcb\x52\x65\xe0\x55\xad\x62\x36\x33\x19\xe0\x1d\xc4\xc3\x7b\x0e\x32\xbe\xf2\x16\xa1\x67\x5c\x16\x64\x6f\x3e\xef\x44\xf7\x0e\x46\xa3\xf1\xf5\xf8\x6a\xff\xf2\x66\x30\xba\x38\x3a\x39\xbe\xb9\xd8\x3f\x1f\x7e\x69\x8c\x8b\x3a\x57\xa3\xae\xe0\x51\xed\x93\xdb\x3f\xce\x96\xcd\x9c\xb7\x66\x8b\xd9\x0c\x2d\xe1\x77\x74\x57\xbe\x28\x2b\xc6\xef\xe8\xef\xf4\x7c\x2e\x04\xa1\x86\xc8\x40\x84\xd0\xa3\x8d\xc3\xe1\xc1\x9f\xc7\x37\x67\xa3\xe3\xb3\xe1\xd7\xe1\xd9\x97\x8f\xab\x81\x4f\x1b\x3d\x78\x4f\x77\x11\x0a\x88\x42\xd6\x61\x91\x95\x4e\x3e\x74\xbf\xbb\x17\x2f\x29\x24\x31\x86\xe4\x83\x10\xa9\x24\xdc\xfe\x04\x91\x86\xdd\xdd\x5d\x98\xcd\xe0\xa0\x0d\x0c\x5d\xf3\x4a\xc3\xcf\xdf\x64\x61\xcf\x65\xa0\x5c\x5a\x88\x07\x2d\x63\x7c\xe1\x35\x1e\x78\xcf\xc4\x41\x96\xa7\x55\x8a\x9d\x92\x5f\x60\x3e\x87\xbd\x65\x96\x46\x4a\x92\x3e\x22\xe3\xc9\x02\xfa\x16\xeb\x00\x03\xef\xd3\x41\xcd\x48\xeb\x59\x1b\xc0\x7f\xf0\xb5\x33\x59\x90\x96\x18\x62\x15\xf8\x2d\xc2\xcb\x60\xa6\x92\xf1\x14\xeb\x57\x69\x4f\xb1\x7e\x37\xeb\x04\x6b\xa1\xf2\xc2\x6b\xd8\xdc\xde\xdc\x84\xf7\x55\xbc\x84\xad\xb5\xef\xff\xf8\x37\x90\xaf\x98\xa6\x64\xeb\x92\x2a\x5f\x0a\xe8\x52\x5d\xbc\x9c\x98\x44\xc9\x88\x43\x45\x9c\x90\xaf\x82\xc2\x44\x3a\x95\xfb\x40\xc9\xd3\xb2\x7a\x68\x56\x95\x5a\x32\x46\x8f\x78\xf1\xb0\x5c\x9c\x2c\xb0\xf9\xb0\x31\xc0\xd6\xf6\xef\xf1\xf6\x6f\xf1\xd6\xc7\x3f\xe2\xad\xed\xfe\x1a\x59\x01\xc9\xdb\x69\xbb\x73\x45\x31\xd1\x26\x40\xf4\x5c\xa1\xb2\xbe\xd2\x65\xf0\x53\xa3\x31\x3c\xed\x71\xce\x0d\x35\x4b\x5a\x57\xa5\x35\x4a\x32\x6a\x30\xae\xd9\x9d\x90\xa3\x2d\x40\xe5\x32\x30\x41\xe6\x43\x1b\xeb\x3e\x06\x28\xbd\xa6\x5f\x41\x3a\x0d\x9c\x4b\xee\x13\x38\xcf\x60\x34\x4a\x2b\x94\xe4\xe7\xea\x9e\xf1\x26\xf2\x9f\x2a\x60\x77\x19\x76\x77\xfb\xc3\xd1\x51\x5f\x30\x3a\xe9\xf8\x44\xef\x34\xc3\x89\xcf\xa5\x33\x19\x12\xc7\xe3\x2e\x7c\x08\xf3\xb9\xa0\x2a\x25\x15\x4c\xc9\xc6\xbb\x17\xc0\xeb\xe5\x64\x0b\x97\x52\x0f\xac\xc1\x35\x4d\x1f\xc2\xcf\x51\xd7\xa8\x02\xf2\x5a\x64\x97\x7a\x40\x8f\xdf\x16\xda\x4c\xa1\x19\xf5\x71\xf0\x55\xb9\x82\xbb\x5a\xce\x35\x60\xeb\x95\x6c\x44\xaf\xe0\xce\x1e\xc2\xed\x83\xa3\xaa\x82\xe1\xba\xad\xb9\x90\x05\xee\x80\xa3\xc5\x02\x15\x65\x30\x85\x0c\xf5\xfe\x54\x1a\x2b\x53\x63\x0d\xd7\xd7\xc8\x1d\x4e\xd2\x02\x36\x1c\x1d\x35\x03\x77\x9e\x71\x07\xd6\x2d\x41\x50\xa1\x39\xf0\xd6\xfb\x92\xa0\x72\x6c\xec\xe3\xa4\x0d\x41\x55\x2e\x1d\x1d\x74\x32\xb5\xb8\xb6\xc9\xe2\x06\xad\x9e\xa8\xd7\xc0\xf0\x93\xf8\x37\x00\x00\xff\xff\xf7\x7f\xa9\x2e\x63\x08\x00\x00")

func startupShBytes() ([]byte, error) {
	return bindataRead(
		_startupSh,
		"startup.sh",
	)
}

func startupSh() (*asset, error) {
	bytes, err := startupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"azuredeploy.json": azuredeployJson,
	"startup.sh":       startupSh,
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
	"azuredeploy.json": {azuredeployJson, map[string]*bintree{}},
	"startup.sh":       {startupSh, map[string]*bintree{}},
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