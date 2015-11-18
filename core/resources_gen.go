// Code generated by go-bindata.
// sources:
// res/Readme.md
// res/actions/Readme.md
// res/default/actions/f.sh
// res/default/actions/goed.rc
// res/default/actions/goed.sh
// res/default/actions/goed_helper.ank
// res/default/actions/goimports.sh
// res/default/actions/s.sh
// res/default/actions/search.ank
// res/default/actions/stats.sh
// res/default/config.toml
// res/default/themes/acme.toml
// res/default/themes/default.toml
// res/resources_version.txt
// res/themes/Readme.md
// DO NOT EDIT!

package core

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _resReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\xb1\x6e\x1b\x31\x0c\xdd\xef\x2b\x5e\xb7\x04\x70\xed\x3d\x5b\x51\x67\x08\x50\xc4\x05\x9c\xec\xa6\x25\xea\x2c\x54\x77\x3a\x48\x54\x93\x76\xe8\xb7\x57\x92\x75\x71\xce\x68\x33\x8a\x7c\xef\xf1\xf1\x51\x5f\xfd\x28\x3c\x4a\x84\x37\xf8\xb3\x59\xf7\x9e\x35\xee\x3a\xe0\x33\x94\x1f\x8d\xed\xd7\xe2\x07\x87\x3b\x3c\x47\x0e\xb8\x51\x29\xe6\xb7\xfd\xcd\xfa\xb6\xf5\x53\x20\xb1\x7e\x84\xb1\x8e\x2b\x4d\x4e\x3c\x70\xdc\xfc\x93\x72\xee\xad\xa0\x02\x93\x30\x4c\xf0\x03\xa2\xca\x0a\xea\x04\x1f\xb2\xe2\x64\xf3\xf8\x73\x59\x68\xd4\x14\xf4\xa6\xe9\x55\x6d\x52\x65\xd6\x7f\xc4\x5b\xb3\xab\x48\xcd\x86\x92\x93\x82\xdc\x05\xdb\xdb\x91\x1c\xea\x6e\xc5\x67\x76\xa0\x3d\x46\x2f\x60\x6d\x05\xda\x06\x56\xe2\x7e\x81\x22\x5e\xac\x73\x38\x32\x02\x4f\x8e\x54\xc6\xa7\x29\x2f\x97\xa6\x3e\x90\xe6\xb8\x5e\x68\x2f\x03\xda\x37\xc3\x73\x6e\xd8\x5e\x46\x2c\x79\x97\x84\xe6\x25\x5b\x30\x1f\x70\xde\x6d\xfe\x46\x6a\xb5\x2b\x56\xa5\x1d\x93\x31\x1c\x32\x5c\x9d\xcf\x8b\x40\x2f\x73\x11\x9a\x84\x56\xd8\xee\xf0\xb8\x7b\xc2\xfd\xf6\xe1\x09\x9f\x2a\xc9\x8e\x45\x5a\x71\xa3\x51\x7e\x17\x63\x50\x29\x84\xa2\x51\xf3\x9b\x41\x37\xf1\x16\x5f\xbe\x3f\x20\x7a\xf5\x83\xa5\x24\xd3\xbd\x25\x50\x33\x06\x05\x46\x8a\x99\xf2\xfc\xf8\xed\x7e\xbf\x07\xe1\x72\x2d\xfc\xcc\x46\xca\xb7\xe1\x57\x1b\x25\xae\xba\xe8\x61\xf2\x0f\xe0\x57\x1a\x26\xc7\xb0\x06\x47\x2f\x27\x1c\xe6\xbd\x7b\x6f\x06\xa9\x51\x1f\x90\xa7\xe0\x70\x1d\xcc\x7b\x40\x13\x9d\xd9\x8b\xde\x7c\xe0\xe2\x6c\xdd\xfd\x0d\x00\x00\xff\xff\x4d\x01\x5a\xb0\xfe\x02\x00\x00")

func resReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resReadmeMd,
		"res/Readme.md",
	)
}

func resReadmeMd() (*asset, error) {
	bytes, err := resReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/Readme.md", size: 766, mode: os.FileMode(420), modTime: time.Unix(1440537445, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resActionsReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x2e\x2d\x2e\xc9\xcf\x55\x48\x4c\x2e\xc9\xcc\xcf\x2b\x56\x48\xcf\x57\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\xc4\x35\xc4\xb2\x17\x00\x00\x00")

func resActionsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resActionsReadmeMd,
		"res/actions/Readme.md",
	)
}

func resActionsReadmeMd() (*asset, error) {
	bytes, err := resActionsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/actions/Readme.md", size: 23, mode: os.FileMode(420), modTime: time.Unix(1435782197, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsFSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8d\x3d\x6e\x84\x30\x10\x46\xeb\xcc\x29\xbe\x0c\xae\x90\x08\x22\x65\xfe\x2e\x41\x89\x10\xb2\x89\x8d\x2d\x11\x13\xe1\x41\x82\xdb\x2f\x5e\x9a\xdd\x29\x9f\xe6\x7d\xaf\x78\xad\x4d\x88\xb5\xd1\xc9\x13\x15\x68\xfd\xb2\xca\xb8\x09\x64\x41\xb2\x7a\x1d\x3d\x5c\x98\x6d\x82\x39\x10\xf5\x9f\x25\x0a\x0e\x1d\x58\x15\x8c\x6a\x16\x34\xe8\x3f\x21\xde\x46\xc2\x79\x76\xf4\x0b\xb8\x3d\xa2\xe8\xfd\x03\x0e\x5f\xd9\x1d\xb2\x37\xa4\xcd\x24\x59\x43\x9c\x7e\xd0\xfd\x6b\xf1\x3d\x5f\xc6\x1e\xce\x11\x72\x81\x28\xd3\x6f\x7e\xe3\xa7\xc4\xf4\x98\x78\xb9\x5e\xd4\x3b\xdf\x05\x17\xe2\x2f\x54\x66\xa8\x72\x03\x5c\xaa\xa6\xe4\x5b\x00\x00\x00\xff\xff\x04\xb7\x77\x9f\xd2\x00\x00\x00")

func resDefaultActionsFShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsFSh,
		"res/default/actions/f.sh",
	)
}

func resDefaultActionsFSh() (*asset, error) {
	bytes, err := resDefaultActionsFShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/f.sh", size: 210, mode: os.FileMode(493), modTime: time.Unix(1442593713, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsGoedRc = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x41\x4f\xc2\x40\x10\x85\xcf\xec\xaf\x98\x84\x3d\x08\x21\x34\x7a\x34\xe9\x81\x48\x03\x1c\x04\x23\x44\x8e\x58\xbb\xbb\x74\x92\xba\xdb\x74\x07\x2b\x69\xea\x6f\x77\x77\x0b\x0a\x7a\xda\x66\xe6\x9b\x37\xef\x4d\xfb\x30\x33\x52\x40\x95\x01\x6a\x24\xc6\x66\xab\x64\xba\x5b\x2c\xd7\x9b\xc9\xf2\x21\x89\xf9\x6d\x57\x78\x59\x24\xdb\x98\xdf\x31\xd6\x87\xb5\x4c\xab\x2c\x87\xbd\x9f\xca\x0e\x96\xcc\x7b\xf4\x76\xc0\x82\x50\x03\x19\x53\x58\x50\x58\x59\x02\xf6\x34\xd9\xcc\x63\x3e\x5f\x3d\x26\xd1\xd8\xc3\x51\x9a\x11\x1a\x6d\xa3\xfb\xcb\xa2\x90\x2a\x3d\x14\x74\xd1\xf4\x73\x8c\x29\x1d\x36\xec\x4c\x29\x35\x34\x6e\x6d\xf8\x48\x9d\x78\x21\x9d\xd3\xd0\x64\xbd\x80\xa4\x25\x76\x5d\x7e\xe5\x1d\x5e\x9b\xb2\x16\x2d\xb8\x08\x6d\xd0\xcb\x44\x10\x72\x0f\x6a\x32\x4e\x4b\x60\x25\x33\x32\xd5\x71\x04\xa9\x16\xa0\x0d\xa1\x3a\x76\xc1\x8c\x02\xca\x25\x68\x59\x7b\x8a\xf5\xce\x09\xdd\x30\x1f\x5e\xec\xfd\x40\x59\xef\xb2\x5a\xfc\xdd\xcd\x7f\xae\x76\xb2\xc1\x5a\x60\x2c\x99\x2e\x36\xab\xe7\xf8\x37\x58\x1f\xb6\xb9\x7b\x6a\xa4\xfc\x94\x69\x04\xbc\xa3\x00\x6d\x17\xd2\x5b\x37\xd0\x9c\xcb\x7c\xd8\xc2\xd5\x35\x22\x6f\xd0\x31\x16\x1a\x3b\xb6\xf9\x09\xb0\xdd\x5f\x22\xf9\x49\x70\x13\x7f\xc1\xbe\x92\xe5\x00\x3c\xa7\xa0\x51\xff\x38\x2f\x64\x03\xa8\x50\x8b\x01\xfb\x0e\x00\x00\xff\xff\x8c\x87\xcf\x61\x16\x02\x00\x00")

func resDefaultActionsGoedRcBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoedRc,
		"res/default/actions/goed.rc",
	)
}

func resDefaultActionsGoedRc() (*asset, error) {
	bytes, err := resDefaultActionsGoedRcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed.rc", size: 534, mode: os.FileMode(420), modTime: time.Unix(1447443819, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsGoedSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\x4d\x8b\x22\x31\x14\x3c\x9b\x5f\xf1\x88\x39\x28\x88\x8d\x7b\x5c\x68\x58\x59\x1b\xf5\xb0\xba\x8c\x32\x1e\x35\xe6\xc3\x0e\xf4\x24\x4d\x27\x3d\xad\x0c\xce\x6f\x9f\xa4\x3f\xa4\xf5\x16\xea\x55\x55\xaa\xde\x1b\xc2\xd2\x08\x0e\x67\x6a\x53\x50\x5a\x39\x84\xc4\x35\x37\x85\x83\xe5\x36\x59\x1c\xd7\x9b\xdd\x7e\xbe\xf9\x9b\xc4\x64\xf6\x84\xbf\xaf\x93\x43\x4c\x7e\x21\x34\x84\x9d\xa0\x05\x4b\xe1\x12\x5c\x58\x69\x9d\xf9\x88\xce\xa5\xca\x9c\xd2\xe0\x8c\xc9\x2c\x48\x55\x58\x07\x9d\xfc\xff\x7c\xbf\x8a\xc9\x6a\xfb\x2f\x89\xa6\x41\x13\x51\xe6\x94\xd1\x36\xfa\xdd\x07\xb9\x90\xb4\xcc\x5c\x6f\x18\x74\x08\xc9\x52\xd7\x48\xfd\xdd\xd1\xe4\x42\x8f\xc6\xf0\xe5\x53\x84\x27\x50\xff\x57\x26\x7c\x8d\x7a\x8c\x06\x35\x89\xe6\xaa\x99\x92\xa7\x46\x80\x4f\x79\xc5\x4f\x18\x30\x99\x61\x74\x7f\xb5\x66\xbc\x35\x66\xdc\xfb\x39\xe3\xbd\xb9\x2a\x04\x73\xa6\xb8\x4d\x80\x6a\x0e\xda\x38\x25\x6f\x4d\x6f\x23\xc1\xa5\x02\xb4\xa8\x02\x0b\x0d\xba\x05\x78\x31\xf9\xd3\xcb\xf1\xa9\x44\x75\x64\x15\x7f\xcd\x42\x1e\x4b\xed\x62\xa1\x3b\x3c\x2e\x91\x2c\xd6\xfb\xed\x5b\x8c\x1f\x9d\x31\x0c\xe1\x90\xfa\x4a\x95\x72\x69\xdb\x76\x02\xa4\xe1\x81\xb2\x4d\x7d\x44\x33\x45\xad\xcf\xd0\x2a\x19\xc7\x1d\x66\x5e\xcc\x7a\xdb\x8b\x42\x81\x86\x65\x63\x6c\xa7\x36\x0d\x04\xdb\x1c\xd9\x89\xab\x83\x51\xfc\x0d\x97\x42\xe4\x63\x68\x79\x32\xc6\xf2\x99\x17\x8c\x6c\x4d\x94\x4a\xf3\x31\x42\x3f\x01\x00\x00\xff\xff\x16\xca\x6c\xe0\x66\x02\x00\x00")

func resDefaultActionsGoedShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoedSh,
		"res/default/actions/goed.sh",
	)
}

func resDefaultActionsGoedSh() (*asset, error) {
	bytes, err := resDefaultActionsGoedShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed.sh", size: 614, mode: os.FileMode(420), modTime: time.Unix(1447444742, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsGoed_helperAnk = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\xc1\x4e\xfb\x30\x0c\xc6\xcf\xcd\x53\xf8\xdf\xff\x25\x95\x50\x81\x07\xd8\x61\xda\xaa\xa9\x97\x21\x31\x04\x07\x84\x50\xd4\xba\x10\xd1\x3a\x95\x93\x0e\x24\xb4\x77\xc7\x69\xc3\xe0\xc8\x2d\xfe\x3e\xfb\x67\x7f\xf9\x0f\x3b\x87\x2d\xbc\x62\x3f\x22\xc3\xe0\xda\xa9\x47\xe8\x1c\x83\xa1\x37\x07\xbe\x61\x3b\x06\x5f\xaa\xa3\x61\xc0\x0f\x6c\x60\x05\x76\x18\x1d\x07\x9d\x3b\x7f\x19\x95\xbc\x50\x2a\x8d\xcd\xa4\x4f\x95\x25\xe6\xd1\xe2\x3b\xd4\x5b\x50\x59\x37\x51\x33\x97\xba\x88\x7e\xc6\x18\x26\x26\x08\xae\xa6\xa0\x9d\x2f\x77\x18\x90\x8e\x3a\xdf\xdd\x54\xdb\xe7\xfb\xba\x7a\xc8\x8b\x42\x65\x27\x95\x9d\x59\x96\x7c\x30\xd4\xa0\xf0\x12\xee\x5b\xf9\x0b\xb2\xde\x1f\xee\xd6\xfb\x4d\xf5\x1b\x7b\x3b\x91\x07\x03\x8d\x1b\x06\x43\xed\x85\xa4\xb3\x41\x5f\x17\x60\x3b\xe8\x8c\xed\x7d\x5a\xc3\x13\x69\xc3\x2f\x7e\xd9\xd2\x0c\xad\xfc\x40\x8c\x5d\x6e\x96\xc1\xd9\x7c\xbc\x7a\x2a\x16\xb7\x5c\x4b\x29\x2d\x51\x15\x65\x64\x4b\xa1\x4f\x84\xd4\x71\x08\xad\x9b\x82\xf4\xc8\x95\xcb\xfb\xc7\x40\xe6\xb3\x21\x6f\x31\x16\x25\xda\x72\xb0\x8e\x0c\x39\x30\x8a\xff\x56\x40\xb6\x9f\xaf\x3a\xaf\x11\x3d\x76\x64\x02\xa8\x96\x38\x52\x9d\x62\xe4\x93\xfa\x0a\x00\x00\xff\xff\x98\x29\x53\x99\xe9\x01\x00\x00")

func resDefaultActionsGoed_helperAnkBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoed_helperAnk,
		"res/default/actions/goed_helper.ank",
	)
}

func resDefaultActionsGoed_helperAnk() (*asset, error) {
	bytes, err := resDefaultActionsGoed_helperAnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goed_helper.ank", size: 489, mode: os.FileMode(420), modTime: time.Unix(1440536590, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsGoimportsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x41\x4b\x03\x31\x10\x85\xef\xf9\x15\x4f\xba\x47\xeb\x82\xf7\x0a\xa2\x45\x7a\xa9\xa0\xa2\xc7\x6d\x36\x3b\xe9\x06\x93\x4c\x49\xb2\xad\x42\x7f\xbc\x13\x5b\xa4\x1e\x27\xf3\xe5\xbd\x6f\x66\x57\x6d\xef\x62\xdb\xeb\x3c\xaa\x4c\x05\x73\xfa\xc2\xcc\x6a\xe7\x41\x3a\xf9\x6f\xa5\x66\x78\x99\x22\xb6\xec\xc2\x8e\x53\xc9\xe0\x08\x2d\x23\x32\x4f\xc9\x10\xac\xf3\x74\x2d\x90\xd5\xde\x67\xf4\xda\x7c\xa2\xb0\xec\x6d\x28\x70\xf6\xe2\x5f\xe4\x02\xbd\x97\x60\xdd\x7b\xba\x51\x6a\xef\xe8\xb0\x68\x9e\x9e\x97\x8f\xdd\xfb\x6a\xf9\xa1\x5c\xcc\xe5\x3c\xaf\xd6\xaf\x6f\xf7\xeb\x87\x65\x2d\xdf\x8a\x53\xed\x80\x67\xa3\x8b\xe3\xa8\x72\x32\x8b\xcd\x96\x69\xe8\xf4\xce\xa1\xc6\x74\xf2\xd4\xc9\x1e\x4d\x0d\x41\x53\xdf\x36\xca\x84\x61\xf1\x57\xaf\x0e\xa3\x33\xe3\x85\xce\xed\x1d\xda\x81\xf6\x6d\x9c\xbc\xc7\xf1\x88\x13\x2d\xd6\xd2\x99\xe4\xe0\x32\x12\x0c\x87\xa0\xe3\x50\x4f\xae\x63\xd5\x50\x8d\x80\x98\x1f\xd0\x48\x67\x45\xc9\xb3\x1e\x7e\x2d\xd0\x4f\xd6\x52\x82\x4d\x1c\x4e\xec\x7f\xc9\x33\x7a\xe1\x08\xf5\x13\x00\x00\xff\xff\x61\x60\xee\x9a\x7e\x01\x00\x00")

func resDefaultActionsGoimportsShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsGoimportsSh,
		"res/default/actions/goimports.sh",
	)
}

func resDefaultActionsGoimportsSh() (*asset, error) {
	bytes, err := resDefaultActionsGoimportsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/goimports.sh", size: 382, mode: os.FileMode(493), modTime: time.Unix(1446851503, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsSSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xb1\x8e\x83\x30\x0c\x86\xe7\xf3\x53\xfc\x67\x58\x73\x88\x1b\xef\xda\xbe\x04\x23\x62\x00\x1a\x88\x25\x94\xa0\xc4\x95\xe8\xdb\x37\xc0\xd2\x7a\xb2\xac\xef\xfb\xe4\xe2\xbb\x1a\xc4\x57\x43\x9f\x1c\x51\x81\xc6\x85\xa8\xe3\x43\xa1\x01\x73\xb4\x2b\x26\x59\x6c\x42\x85\xbb\xc4\x44\x24\x13\x5a\x70\x59\x30\xcc\xa2\xa8\xd1\xfd\x43\x9d\xf5\x84\x3c\x76\x74\x01\xdc\x3c\xbd\xf6\xdb\x1f\x12\x2e\x6b\xaf\x36\xfa\x1b\xda\xbc\xb8\x8e\x4f\x68\x93\xec\xd1\x24\x44\xfb\xf5\xca\x3f\xfc\x51\x9d\xdf\xab\x5f\x27\x52\xfe\xf2\x21\x1c\xff\x98\xe8\x05\xc6\x8c\x61\x09\x31\x3b\x35\xa3\xdc\x29\xbc\x02\x00\x00\xff\xff\x00\x70\x5a\x41\xc8\x00\x00\x00")

func resDefaultActionsSShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsSSh,
		"res/default/actions/s.sh",
	)
}

func resDefaultActionsSSh() (*asset, error) {
	bytes, err := resDefaultActionsSShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/s.sh", size: 200, mode: os.FileMode(493), modTime: time.Unix(1446160181, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsSearchAnk = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xcf\x4a\x03\x31\x10\x87\xcf\xce\x53\x8c\xe9\x25\xc1\xed\x9f\x17\x58\x41\x45\xbc\x8a\x3d\x96\x22\x61\x3b\xdd\x06\xd3\x64\x99\x24\x52\x11\xdf\xdd\x49\x44\xad\x5e\x16\x66\xe6\xfb\x7d\x33\x9b\xd9\xa5\x0d\x2f\x11\x60\x86\xf7\x27\x7b\x9c\x3c\x61\xad\x31\x0d\xec\xa6\xdc\xe1\x50\x98\x29\x64\xff\x86\x25\x94\x44\x3b\x00\x1f\xed\x4e\xab\x31\xd2\xee\xf9\x40\x7e\x22\x5e\x48\x40\x19\x80\x57\xcb\x18\x53\x87\x7b\x6f\xc7\x0e\xe9\x44\x03\xf6\xe8\x8e\x53\xe4\xac\x55\x4c\xca\x74\x3f\x55\x45\xce\xeb\x98\x96\x95\xaf\x96\x19\x3e\x5a\x4e\x72\x05\x8f\x09\x97\x4d\x96\xc0\x8d\x21\x32\xdd\x59\xe9\xf7\xad\xb5\xb8\x8d\xd1\x6b\xe5\x94\xac\xb3\x3e\x51\x87\xea\x97\x11\x4d\x63\x9a\x48\x1b\xb9\x78\xf8\x8e\xdd\x88\x55\x9b\xcd\x6a\x0b\xcd\xdf\xe3\x46\x8d\x4c\x93\x68\xd4\x3c\xb4\xaf\x57\x5b\x70\x7b\x3c\xdf\xd8\x63\xe6\x42\xf8\x0e\x17\x2d\x74\x55\x53\x73\x27\xdc\x47\x25\x3d\x05\x2d\x1b\x0c\x5e\xe3\xea\x1f\xc3\x62\x94\x91\x80\x48\x72\xe4\x9f\xe9\x57\xbf\xfe\xef\x53\x09\x98\xc8\xf2\x70\x80\x07\x79\xd5\x05\x97\xa0\x73\x5c\x67\x76\x61\x5c\x7b\x37\x90\xae\x21\x63\xe0\x33\x00\x00\xff\xff\x7e\x24\xb8\x7b\xab\x01\x00\x00")

func resDefaultActionsSearchAnkBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsSearchAnk,
		"res/default/actions/search.ank",
	)
}

func resDefaultActionsSearchAnk() (*asset, error) {
	bytes, err := resDefaultActionsSearchAnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/search.ank", size: 427, mode: os.FileMode(493), modTime: time.Unix(1442596451, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultActionsStatsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\x4a\x4d\xce\xc8\x57\xd0\xd5\x55\x70\x0b\x72\x75\x05\xd2\x5c\x69\x45\xa9\xa9\x70\x41\x17\x37\x90\x50\x4a\x9a\x82\x6e\x85\x42\x49\x6e\x41\x5a\x31\x5c\x26\xc4\x3f\x40\x21\x20\xc8\xdf\x39\x18\xa4\xa0\xa0\x58\x21\xb1\xb4\x42\xa1\x46\xa1\x38\xbf\xa8\x44\x41\xb7\x28\x5b\xc1\x58\xc7\x18\xc8\xcd\x48\x4d\x4c\x51\xd0\xcd\x53\x30\x03\x04\x00\x00\xff\xff\x85\x91\x8c\xbb\x71\x00\x00\x00")

func resDefaultActionsStatsShBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultActionsStatsSh,
		"res/default/actions/stats.sh",
	)
}

func resDefaultActionsStatsSh() (*asset, error) {
	bytes, err := resDefaultActionsStatsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/actions/stats.sh", size: 113, mode: os.FileMode(493), modTime: time.Unix(1446839550, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultConfigToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\x70\x49\x4d\x4b\x2c\xcd\x29\x51\x48\xce\xcf\x4b\xcb\x4c\xe7\x0a\xae\xcc\x2b\x49\xac\xf0\xc8\x4c\xcf\xc8\x01\xe2\x92\xcc\xbc\x74\xdb\x92\xa2\xd2\x54\xae\x90\x8c\xd4\xdc\x54\x5b\xa5\x14\x88\x6a\xbd\x92\xfc\xdc\x1c\x25\x2e\xdf\xc4\x0a\xe7\xdc\x14\xa7\xd2\xb4\xb4\xd4\x22\x9f\xcc\xbc\xd4\x62\x5b\x63\x03\x03\x03\x40\x00\x00\x00\xff\xff\x52\xe0\x53\xf9\x54\x00\x00\x00")

func resDefaultConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultConfigToml,
		"res/default/config.toml",
	)
}

func resDefaultConfigToml() (*asset, error) {
	bytes, err := resDefaultConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/config.toml", size: 84, mode: os.FileMode(420), modTime: time.Unix(1446832834, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultThemesAcmeToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xbd\x6e\xc2\x30\x10\xc7\xf7\x3e\x05\xca\x9c\xc1\x81\x7c\x90\x21\x0b\x4e\xb2\x54\x55\x87\xd0\xee\x01\x2c\x1a\x29\xc1\xc8\x18\x51\xa4\x0e\xdd\x90\xba\xf5\x7b\xa1\x43\x55\xa9\xea\x73\xf1\x24\x3d\xc7\xb9\x24\x94\xa8\xb0\x38\xff\xfb\xdd\xcf\x17\x3b\xa3\x79\x60\x44\x2e\x29\x7f\xc6\x59\xac\x9e\x86\xf8\x34\x9a\x27\x2c\x67\x53\x19\x18\x23\x9f\x0c\x90\xc0\xac\xcd\xd1\xb5\x58\x71\x71\xcc\x61\xd6\x70\x94\x17\x05\x5b\x40\x6b\xec\x11\x8f\xc4\x2a\x4a\xa4\xc8\x16\xb0\xa9\xed\x13\x57\x27\xe7\x6c\xbb\xe1\x62\x66\x05\x46\xdf\x02\x59\x4c\xac\x3a\xeb\x07\x86\x1b\xff\xc9\x06\x90\xf9\x98\x25\x6c\x99\x8a\x54\x72\x01\xdd\x36\x21\x4e\xb5\x07\xa6\xd0\xef\x3b\x27\x29\x18\xa2\x7e\x9d\x6e\x8b\x09\xcf\xa1\x9d\x52\x62\x57\xd2\x32\x82\xde\x61\x78\x1c\x41\x63\x68\xd5\x91\x4c\xe5\x7a\x35\x49\x45\x2f\xe8\x19\x87\x8f\x07\x93\x0e\xf4\x51\xd4\x8b\x16\x33\x66\xb7\x52\x71\xcd\xd1\x1c\x95\x22\x51\x5a\x88\x05\x7f\xa2\xe4\xb4\x98\xb5\xcc\x6e\x54\x99\x7b\xb8\x42\xe2\xd4\xdb\xe4\x97\x8b\x56\x05\x9c\xd7\x19\xdb\xa0\xf4\xfe\xb9\x6b\xdc\x8a\x38\x95\xc6\x59\xce\x68\xce\xd2\xd2\x78\xd8\xbf\x98\x58\x32\xfd\xe6\xfe\x81\x09\x33\x21\xb7\x9a\x79\x37\xa9\xad\x77\x6e\x31\xc9\x54\xf0\x3c\xc7\x21\xde\x1e\x4d\xd7\x86\x7b\x50\x1e\x5c\x20\x33\x4e\x27\x9a\x79\x32\xf1\x63\x55\x0b\x4f\x33\x17\x5c\x30\x35\x65\x92\xcd\x98\x7e\x9f\x6f\x33\x76\x74\xb5\x03\xbb\x5a\x96\xd0\xee\xf3\x3f\x28\xe4\x1b\xfd\x7a\xbb\xaf\x2e\x0c\x06\xa2\x37\xd5\xe0\xbb\x9f\x4e\x51\x2a\xe6\x59\xa9\xb8\xeb\x2a\xd3\x9c\xaf\xf4\xb0\xfb\xd7\xe6\xfc\x1c\x0f\xcf\xe6\x37\x00\x00\xff\xff\xd8\xe8\x2f\xf0\x98\x03\x00\x00")

func resDefaultThemesAcmeTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultThemesAcmeToml,
		"res/default/themes/acme.toml",
	)
}

func resDefaultThemesAcmeToml() (*asset, error) {
	bytes, err := resDefaultThemesAcmeTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/themes/acme.toml", size: 920, mode: os.FileMode(420), modTime: time.Unix(1444676960, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resDefaultThemesDefaultToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xbb\xee\xda\x30\x14\xc6\xf7\x3e\xc5\x5f\x99\x3d\xd8\x49\x08\x64\xc8\x40\x7c\x59\xaa\xaa\x43\x68\xf7\x00\x16\x8d\x94\x60\x64\x82\x28\x52\x87\x6e\x48\xdd\x7a\x5f\xe8\x50\x55\xaa\xfa\x5c\x3c\x49\x8f\x6f\x81\x52\x44\xa6\xf8\x3b\xbf\xef\x3b\xc7\x47\x2e\x57\x45\xc4\xa7\xd8\x7e\xd1\x33\x01\x27\x26\x70\x82\x85\x39\x95\xab\x4a\xb6\x72\xd1\x5f\x6b\x62\xd0\xbc\x8b\x18\x8e\xee\xf4\x56\xe9\x22\x8a\x09\xce\x02\x17\xb4\x0b\x47\x55\xd7\xc9\x35\x58\xc5\x18\x8f\x1d\x56\xf5\xba\x59\x43\xd3\x34\x0f\xc6\xe7\xf2\xb0\x57\x7a\x49\x6c\x98\x69\x4a\x06\x2d\x2e\xa2\x4c\xdc\x68\x09\x68\x79\xd0\x2a\xb9\xa9\x75\xdd\x2b\x0d\xee\x14\xe3\x91\xef\x11\x54\xf0\xe7\xa3\xff\x54\x48\xe0\xf1\xa0\x1e\xba\xb9\x6a\xc1\x4e\x29\x4e\x7d\xa8\x95\xc0\x3b\x61\xff\x4a\x60\x64\x64\x90\xfa\xba\xdf\x6d\xe7\xb5\x7e\x2a\x9e\xa2\xf3\x8f\x0f\x88\x97\x70\x47\xf8\xcc\x8f\xdf\xee\xc0\xcc\xe4\xdb\xde\x70\x25\xcb\xb2\x3b\x25\xae\x6d\x0a\x26\x74\xe4\x17\xd7\x2d\x1f\x27\x3b\xe0\x12\x0b\x0b\xb9\xd1\x5f\xae\x6d\x45\xd8\x86\x10\xf9\xba\x91\xfb\x90\xf9\xfe\x33\xe2\x2c\x64\xb2\x90\xe9\x89\x10\xca\xc7\x38\xf6\x8f\xa4\x69\x25\x6d\x65\x6d\x13\xcf\xa7\x2f\x88\x50\x57\x42\x3c\x1f\x1e\x12\x30\xac\xd1\xfd\xc1\x31\xdf\x11\x3c\x02\x62\x3a\x5f\x31\xd5\x42\xab\xb6\x0d\x43\x7c\xfb\x88\x60\x36\xbb\x0e\xe4\xb6\x72\x61\x66\xf5\xdc\x31\x9f\x50\x96\x06\x26\x0d\xcc\x0b\xa5\xa5\x99\xb2\x6a\x96\xd2\xdd\xe7\x37\x22\x02\xa7\x6e\xa4\xe9\x2d\xf6\x6a\x63\xa1\xe3\xcf\x47\x10\x53\x7b\x77\xbd\xe3\xaf\x7b\x18\x0c\x44\xdf\xf8\xc1\x8f\x7f\x10\x17\xae\x60\x88\x89\x0f\xaa\xf5\xaa\xb1\x11\xef\xee\x95\x69\xab\xb6\x6e\xd8\xd3\x57\x94\x24\xfe\x4e\x9c\x86\xfa\xdf\x00\x00\x00\xff\xff\x58\x21\xa3\x13\x97\x03\x00\x00")

func resDefaultThemesDefaultTomlBytes() ([]byte, error) {
	return bindataRead(
		_resDefaultThemesDefaultToml,
		"res/default/themes/default.toml",
	)
}

func resDefaultThemesDefaultToml() (*asset, error) {
	bytes, err := resDefaultThemesDefaultTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/default/themes/default.toml", size: 919, mode: os.FileMode(420), modTime: time.Unix(1445963239, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resResources_versionTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\x34\x31\xb1\x30\x30\x33\x31\x36\x31\xe3\x02\x04\x00\x00\xff\xff\xa8\xe3\x5a\x40\x0b\x00\x00\x00")

func resResources_versionTxtBytes() ([]byte, error) {
	return bindataRead(
		_resResources_versionTxt,
		"res/resources_version.txt",
	)
}

func resResources_versionTxt() (*asset, error) {
	bytes, err := resResources_versionTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/resources_version.txt", size: 11, mode: os.FileMode(420), modTime: time.Unix(1448064346, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _resThemesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x2e\x2d\x2e\xc9\xcf\x55\x28\xc9\x48\xcd\x4d\x2d\x56\x48\xcf\x57\xc8\x48\x2d\x4a\xe5\x02\x04\x00\x00\xff\xff\x26\x56\x89\xea\x16\x00\x00\x00")

func resThemesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_resThemesReadmeMd,
		"res/themes/Readme.md",
	)
}

func resThemesReadmeMd() (*asset, error) {
	bytes, err := resThemesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/themes/Readme.md", size: 22, mode: os.FileMode(420), modTime: time.Unix(1435782185, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"res/Readme.md": resReadmeMd,
	"res/actions/Readme.md": resActionsReadmeMd,
	"res/default/actions/f.sh": resDefaultActionsFSh,
	"res/default/actions/goed.rc": resDefaultActionsGoedRc,
	"res/default/actions/goed.sh": resDefaultActionsGoedSh,
	"res/default/actions/goed_helper.ank": resDefaultActionsGoed_helperAnk,
	"res/default/actions/goimports.sh": resDefaultActionsGoimportsSh,
	"res/default/actions/s.sh": resDefaultActionsSSh,
	"res/default/actions/search.ank": resDefaultActionsSearchAnk,
	"res/default/actions/stats.sh": resDefaultActionsStatsSh,
	"res/default/config.toml": resDefaultConfigToml,
	"res/default/themes/acme.toml": resDefaultThemesAcmeToml,
	"res/default/themes/default.toml": resDefaultThemesDefaultToml,
	"res/resources_version.txt": resResources_versionTxt,
	"res/themes/Readme.md": resThemesReadmeMd,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"res": &bintree{nil, map[string]*bintree{
		"Readme.md": &bintree{resReadmeMd, map[string]*bintree{
		}},
		"actions": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resActionsReadmeMd, map[string]*bintree{
			}},
		}},
		"default": &bintree{nil, map[string]*bintree{
			"actions": &bintree{nil, map[string]*bintree{
				"f.sh": &bintree{resDefaultActionsFSh, map[string]*bintree{
				}},
				"goed.rc": &bintree{resDefaultActionsGoedRc, map[string]*bintree{
				}},
				"goed.sh": &bintree{resDefaultActionsGoedSh, map[string]*bintree{
				}},
				"goed_helper.ank": &bintree{resDefaultActionsGoed_helperAnk, map[string]*bintree{
				}},
				"goimports.sh": &bintree{resDefaultActionsGoimportsSh, map[string]*bintree{
				}},
				"s.sh": &bintree{resDefaultActionsSSh, map[string]*bintree{
				}},
				"search.ank": &bintree{resDefaultActionsSearchAnk, map[string]*bintree{
				}},
				"stats.sh": &bintree{resDefaultActionsStatsSh, map[string]*bintree{
				}},
			}},
			"config.toml": &bintree{resDefaultConfigToml, map[string]*bintree{
			}},
			"themes": &bintree{nil, map[string]*bintree{
				"acme.toml": &bintree{resDefaultThemesAcmeToml, map[string]*bintree{
				}},
				"default.toml": &bintree{resDefaultThemesDefaultToml, map[string]*bintree{
				}},
			}},
		}},
		"resources_version.txt": &bintree{resResources_versionTxt, map[string]*bintree{
		}},
		"themes": &bintree{nil, map[string]*bintree{
			"Readme.md": &bintree{resThemesReadmeMd, map[string]*bintree{
			}},
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

