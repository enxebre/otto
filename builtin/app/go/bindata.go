// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/build/template.json
// data/common/dev/Vagrantfile.tpl
// data/common/dev-dep/build/Vagrantfile.fragment.tpl
// data/common/dev-dep/build/Vagrantfile.tpl
// data/common/dev-dep/build/build.sh.tpl
// data/common/dev-dep/build/upstart.conf.tpl
// DO NOT EDIT!

package goapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
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

var _dataAwsVpcPublicPrivateBuildTemplateJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xdc\xcc\xbc\xf8\x82\xc4\xe4\xec\xd4\xa2\xf8\xb2\xd4\xa2\xe2\xcc\xfc\x3c\x25\x2b\x05\x25\x03\x3d\x0b\x3d\x03\x25\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x1b\x82\xd0\x08\x26\x00\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJson,
		"data/aws-vpc-public-private/build/template.json",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJson() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json", size: 38, mode: os.FileMode(420), modTime: time.Unix(1435862031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xef\x6b\xe4\x36\x10\xfd\xee\xbf\x62\xea\x5c\x7a\x2d\xc4\x32\x17\xc2\x7d\x08\x77\x81\x23\x85\x34\xd0\x92\x6b\x93\xde\x97\x52\xb2\x5a\x7b\x2c\x8b\xc8\x1a\x55\x92\xf7\x47\x36\xfb\xbf\x77\x64\x7b\xb3\xbb\x07\x57\x0e\x16\xd6\x1e\xcd\x7b\x6f\x66\xde\xc8\x27\x70\x83\x16\xbd\x8c\x58\xc3\x7c\x0d\x77\x31\xd2\x19\xd4\x04\x96\x22\x60\xad\xe3\x0f\xd9\x49\x76\x02\x0f\xad\x0e\xc0\xbf\xd8\x22\x7c\x91\xca\x4b\x1b\x1b\x6d\x10\xd4\xd7\x58\x68\xc8\x0f\x59\x35\x2e\xd0\x90\xeb\xd0\x46\xa0\x86\x29\x62\xa2\x90\xce\x19\x5d\xc9\xa8\xc9\x96\x01\xfd\x42\x57\x28\xe0\x36\x42\x68\xa9\x37\xf5\x20\x3a\x47\x68\xa5\xad\x8b\x24\x8e\xb5\x80\x07\x82\x8e\x6a\xdd\xac\x13\x2d\xf3\x1c\xc8\x9f\x41\x1f\x70\x50\xfb\xe4\x5c\x0a\x88\x2c\x9b\x8e\x45\x45\xb6\xd1\xaa\xf7\xf8\x53\x7e\x9e\xff\x9c\x3a\x7a\x19\x43\x2f\x19\xc0\xf8\x24\x16\x9d\x98\xd3\x0a\x3e\x42\xde\xca\xd0\xea\x8a\xbc\x2b\x9d\xc7\x4a\x07\x7c\x7f\x91\x67\x9c\x78\x02\xf7\x18\x7b\x07\x12\xc2\xda\x56\xdc\x66\x43\xa6\x46\x0f\x8d\xa7\x0e\xa8\xf7\xb0\x24\xff\xa4\xad\x82\x5a\x33\x2e\x92\xe7\x2a\x09\xca\xc5\x58\xc4\x91\xd2\x48\xf0\x38\x11\xe4\x9b\x0d\x38\x19\x5b\xb1\x23\xd8\x6e\xf3\x33\xc8\x77\xc8\x49\xfc\x37\x92\x35\x48\x63\x06\xa9\xc6\x4b\x95\xc6\x19\xa0\x45\x8f\xc3\xa0\xa5\x5d\xf3\xa0\x1d\xda\x1a\x6d\xa5\x31\x08\x46\x6d\x4e\x87\xa3\x5d\x36\x68\x9b\xbc\x78\xdc\xa3\x4f\xb7\x29\x6b\xf3\x9a\xf1\xe2\x91\x55\xb6\xdb\x11\xcb\x5c\x09\xce\x49\x43\x05\xb7\x36\xc4\x54\xc0\x0d\xc1\xbc\xd7\x6c\x11\xda\x85\xf6\x64\x13\xf0\xa8\x3d\xe7\x69\xa1\x03\xfb\x0a\x79\x68\xd1\x18\xee\x46\x5b\xa3\x2d\x5e\xc2\x9b\x50\x79\xed\xe2\xa3\x22\x23\xad\x1a\x79\x7f\x97\x4f\x08\x9a\x8d\x27\xf6\x4f\x46\x98\x4d\x8d\x43\x08\xed\x0c\x14\x61\x98\x46\x6a\x86\x89\x26\x8f\xd9\x9e\x14\x48\xf1\xef\x54\xe6\x34\x80\xd3\x3f\xfe\xc6\xaa\x25\xc8\xab\xfa\xd5\x98\x1c\xae\xae\xa0\x6c\xa9\xc3\x5d\xa4\x14\x73\x5e\x01\x5f\xfd\x93\xf1\x00\xb2\xec\xb8\x64\xde\x90\x0f\x1f\xee\xaf\xff\xbc\xfd\xfc\x90\x05\x8c\x50\x60\x96\x8d\x9c\xbf\xd0\xd2\x1a\x36\x29\x39\x78\x43\x42\x88\x3c\x5b\xaa\x94\xf1\x2f\x14\x77\x5f\x29\x28\x12\x51\x7a\xa1\x9e\xa1\x8d\xd1\x85\xcb\xb2\x0c\xbc\x30\x52\xa1\x50\x44\xca\xa0\x74\x3a\xf0\xda\x76\xe5\x28\xca\x7f\xef\xc4\x85\x38\x17\xdc\x4a\xbf\x2a\x64\x57\xbf\xbf\x98\x08\x76\xea\x7f\x59\x7e\xf7\x07\xda\xa1\xe7\x3d\xe7\x18\x14\xd7\x50\xf6\xc1\x97\x86\x2a\x69\xa0\x58\x3d\x37\xdf\x2a\x66\xc7\xc5\x86\x0c\x44\x77\x9f\x3f\x3d\xfc\xba\x27\xeb\x9e\x78\xdc\x50\x38\x28\xc9\x25\x54\xda\xd9\xf1\x84\x51\x4b\x0b\xb3\x65\x4b\xb2\xd3\xb3\xcb\xdd\xc3\x51\xe2\xc4\xcd\x77\x28\x26\x72\xbe\x48\xaf\xec\xc3\xc9\x5b\x5c\x39\xf2\x71\x88\x7e\x3c\x00\x96\x73\x6d\x2f\xf7\x0d\x70\x74\x88\xbc\x49\x79\x6f\xbf\xe9\xdd\x31\xe7\xd8\xc9\x21\xeb\xff\x20\xa7\x42\xa7\x6d\x4f\xb5\x7e\xb9\xbe\x0f\xc3\x45\x52\xc4\x5f\xb9\xb8\x9f\x88\x74\xb1\x48\x16\xf7\xae\xe6\x0f\x1f\x14\x6b\xb8\x2a\xf9\x82\x95\xb6\xe7\x6b\x72\x7e\xf5\xe3\xbb\xe3\x34\x3d\x5d\x20\xce\x53\xbc\xef\xf3\x67\x0f\x1d\xfa\xaa\xf7\x5a\x9a\x6c\xda\xa9\xff\x02\x00\x00\xff\xff\x55\x6e\x3e\x31\x84\x05\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 1412, mode: os.FileMode(420), modTime: time.Unix(1440180769, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildVagrantfileFragmentTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\x3f\x6b\xf3\x40\x0c\xc6\xf7\xfb\x14\x0f\xce\x3b\xbe\xf1\xed\x21\x9d\x32\x74\x6c\x87\x2e\x9d\x8a\x63\x2b\x8d\xe0\xfe\x71\xa7\x73\x31\x21\xdf\xbd\x72\x08\xc4\xa5\x0d\x74\xe8\x78\x48\xfa\x3d\x3f\x89\x5b\xe1\x91\x02\xe5\x4e\x68\xc0\x7e\xc2\x40\x89\xc2\x40\xa1\x9f\x70\x88\x59\x9f\x23\xb9\x98\x3c\x05\xd9\xe0\x74\x42\xe8\x3c\xe1\x7c\x36\xe6\xdf\xed\xf1\x56\x48\x6a\xc2\x03\xb6\xdb\xdd\xd3\xf3\xab\x59\x61\x17\xd3\x84\x58\x33\xf6\x1c\xba\x3c\x19\x3f\xc2\x8a\x4f\x56\xe1\xeb\xdb\x1c\x6c\x2d\xd9\xba\xd8\x77\xce\x6a\xa3\x5d\xe2\xaf\x0c\x39\x12\x6a\x2a\xd2\x65\xc1\x81\x1d\xdd\x21\xb5\xd7\x9e\xb6\x8f\xe1\x00\x4b\xd2\x5b\x0e\x2c\x0b\xe2\xa5\x62\xcc\xd2\x6f\xe6\x15\x70\x90\x88\x0e\x42\x3e\x61\xe0\x4c\xbd\xc4\x3c\xb5\x78\xd1\xe0\xd2\x67\x4e\x82\x0f\x76\x0e\x3e\x8e\x34\xdb\xf8\xd6\xcc\x24\x7e\x6f\x47\xdf\xa6\x1c\x47\x2e\x1c\x03\x9a\x19\xd6\xfc\x37\x40\xd1\xb5\x7b\xda\xa0\xd1\xec\xd4\xc9\x51\x83\x7d\xd2\xe2\xa0\x12\x6a\x3d\xae\xd5\xdc\xee\x2b\xbb\xc1\xc6\x2a\xa9\xca\x65\x6a\xa0\x22\x7a\x2a\x51\x96\x8e\xfe\xb0\x61\x63\xfe\x2e\x77\x79\xac\xdf\xa5\x7f\xb9\xef\x3d\x95\x72\x24\xe7\x2e\x3c\x0e\x8e\x83\xba\x7c\xfb\x23\xe6\x33\x00\x00\xff\xff\xf9\x1f\xa6\x3e\x6d\x02\x00\x00"

func dataCommonDevDepBuildVagrantfileFragmentTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildVagrantfileFragmentTpl,
		"data/common/dev-dep/build/Vagrantfile.fragment.tpl",
	)
}

func dataCommonDevDepBuildVagrantfileFragmentTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildVagrantfileFragmentTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/Vagrantfile.fragment.tpl", size: 621, mode: os.FileMode(420), modTime: time.Unix(1440181528, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x4f\xfb\x46\x10\x7c\xf7\xa7\xd8\x3a\xa8\xb4\x12\x3e\x0b\x84\x78\x88\x20\x12\x02\x29\xe5\xa1\x02\x95\x94\x3e\xc2\xc5\xde\xd8\x27\xec\xdb\xeb\xdd\x39\x7f\x80\x7c\xf7\xee\x9d\x1d\x20\x48\xf4\xf7\x93\x22\xc5\x5e\xef\xce\xcc\xde\x8c\x3d\x82\x29\x6a\xb4\xd2\x63\x09\xf3\x0d\xdc\x7a\x4f\x47\x50\x12\x68\xf2\x80\xa5\xf2\xbf\x24\xa3\x64\x04\xb3\x5a\x39\xe0\x9f\xaf\x11\x1e\x64\x65\xa5\xf6\x0b\xd5\x20\x54\x5f\x67\x61\x41\x16\xe6\x9d\x6a\x4a\xa5\xab\xd0\xce\xc3\x73\xa5\xa5\xdd\xf0\x8d\xf4\x01\xa3\x73\xdc\x2e\x1d\x48\x28\xd1\xa0\x2e\x51\x17\x9b\x38\x56\xe2\x12\x1b\x32\x2d\x6a\x2f\x22\xeb\x75\x2f\xa3\x96\xba\xcc\x82\x16\x86\xe0\xf9\x40\x2c\x60\x46\xd0\x52\xa9\x16\x9b\x58\x3c\x0a\xa8\x51\xdd\xa5\x31\xb1\x21\x49\x06\x9d\xa2\x20\xbd\x50\x55\x67\xf1\xb7\xf4\x24\xfd\x3d\xec\xf6\xd6\x97\xde\x12\x80\xfe\x4a\x2c\x5b\x31\xa7\x35\x5c\x40\x5a\x4b\x57\xab\x82\xac\xc9\x8d\xc5\x42\x39\x3c\x3b\x4d\x13\x6e\x1c\xc1\x3d\xfa\xce\xb0\x6a\xb7\xd1\x05\x6f\xb0\xa0\xa6\x44\x0b\x0b\x4b\x2d\x50\x67\x61\x45\xf6\x39\xec\x5c\x2a\x9e\xf3\x14\x16\x26\xc8\x97\xbd\x88\x3d\xa6\x1e\xe0\x71\x00\x48\x5f\x5f\xc1\x48\x5f\x8b\x1d\xc0\x76\x9b\x1e\x41\xba\x9b\xfc\x31\xf9\xaa\x46\x8b\x51\x42\x41\xad\xe1\xdd\x4b\x28\xa5\x97\xd1\x2e\x8a\xc3\x39\xb1\x33\x02\xfe\xc1\xb0\x7c\x3c\x43\x96\x26\x8b\x02\x5d\xef\x68\xf4\x0b\x5c\x61\x95\xe1\x93\xff\x09\xa9\xef\x44\xdb\x6d\xce\xae\x65\x6c\x64\x1e\x41\xa2\xf2\xc0\x36\xc8\xbe\xd1\xce\xcb\xa6\x81\x29\x0d\x24\xa8\x97\xca\x92\x0e\x1e\xef\x11\x19\x4b\x4b\xe5\x14\x69\x48\x5d\x8d\x4d\xc3\x40\x4a\x37\x4a\xe3\x18\x0e\x7a\x61\x8f\x15\x35\x52\x57\x09\x27\x26\x49\xf6\x6b\xec\xdb\xf9\xf9\xfd\xd5\x5f\x37\x77\xb3\xc4\xa1\x87\x0c\x93\x04\x8b\x9a\x20\xbd\xa6\x95\x6e\x48\xc6\x30\x4e\x49\x08\x91\x26\xab\x2a\x74\xfc\x0b\xd9\x2d\xe4\x35\xb5\xb8\x3b\xe8\xbc\x22\xe1\xa5\x15\xd5\x0b\xd4\xde\x1b\x37\xce\x73\xc7\x36\xca\x0a\x45\x45\x54\x35\x28\x8d\x72\x61\xf3\xbc\x27\xe5\xbf\x63\x71\x2a\x4e\x04\xab\xec\xd6\x99\x6c\xcb\xb3\xd3\x01\x60\xc7\xfe\xb7\xe6\x7b\xfb\x89\xdb\x75\xc1\x00\x69\x21\xbb\x82\xbc\x73\x36\x6f\xa8\x90\x0d\x64\xeb\x97\xc5\x77\x62\x12\xdb\x7e\xfb\x68\xa0\xf9\x53\xc6\xdc\x4c\x6f\xef\x2e\x67\x7f\x7c\xf0\xb4\xcf\x1c\x45\xc8\x0c\xdb\x6f\xc2\x54\x70\xae\x7f\xc2\x53\x2b\x0d\x4f\xab\x9a\x64\xab\x9e\xc6\xbb\x8b\xbd\xc6\x01\x9b\x73\xe7\x03\x38\x87\xef\x1d\x3d\x3e\x39\xc4\xb5\x21\xeb\x63\xf5\xe2\xd3\x60\xce\x6f\xfb\xf8\x63\x37\xae\xc6\xca\x41\xe8\x3b\x84\xc9\xe4\xcb\x32\x62\xce\xef\x9b\x2d\xf6\x31\xfb\x4d\x3e\xa3\xfe\xcf\xe4\x20\x74\x48\x5a\xd0\xfa\x70\x75\xef\xe2\x07\xa5\x22\xfe\x40\xf9\x8f\x13\x91\xc6\x67\xc1\xfd\xce\xf0\x0b\x82\x90\x6d\x60\x12\xe2\x9b\xeb\x8e\x23\x7a\x32\xf9\xf5\x78\xbf\x4d\x0d\xe1\xe5\xbe\x8a\x3f\x3e\xf3\x17\x0b\x2d\xda\xa2\xb3\x4a\x36\xc9\x10\xb7\xff\x02\x00\x00\xff\xff\x17\x58\x96\x17\x3f\x05\x00\x00"

func dataCommonDevDepBuildVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildVagrantfileTpl,
		"data/common/dev-dep/build/Vagrantfile.tpl",
	)
}

func dataCommonDevDepBuildVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/Vagrantfile.tpl", size: 1343, mode: os.FileMode(420), modTime: time.Unix(1440118643, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildBuildShTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x91\xc1\x6e\xdb\x30\x10\x44\xef\xfa\x8a\xa9\x7d\xc8\xc9\x12\xfa\x01\x2d\xd0\x5e\x72\x6c\x81\xb4\x1f\xb0\xa2\x56\x12\x1b\x99\x4b\xac\x96\x51\xfc\xf7\x5d\x4a\x0e\x62\xc0\x30\x04\x2e\x67\xe6\xed\xf0\xfc\xa5\xeb\x63\xea\x7a\x5a\xe7\xe6\xdc\x9c\xf1\xa3\x98\x5c\x26\x4e\xac\x64\x3c\xa0\xbf\xe1\x97\x99\xb4\xfb\xec\xcf\x1c\x57\xf8\xcf\x66\x46\x5f\xe2\x32\x60\x0d\x1a\xb3\x61\x14\x05\xe1\x59\x2e\x6e\xe3\xa2\xac\xf2\x8f\x83\xb5\xcd\xca\x86\x0b\x37\x2e\x7d\xf1\x2f\xab\x72\x13\x8c\xf4\xca\x87\x87\x87\x6a\x80\x24\xac\x72\x65\xe4\x85\xcc\x9d\xae\x35\x80\x0c\x1b\x3f\x29\x23\x26\x73\x94\x60\xf1\x8d\x1d\x02\x7f\xfb\x92\xac\xf8\x29\x32\xa9\xc5\x50\x16\xd2\x0a\x39\xf0\x48\x65\x71\x91\xa4\x27\xc3\x22\x34\x3c\x26\xc4\xf1\x08\x8f\xab\x4f\xdd\xc5\xb9\xda\x86\xdf\xb3\xa8\xe1\xf7\xcb\xd7\x6f\xa7\xef\x38\xed\x94\x52\x34\x30\xfc\xbf\xee\x30\xc6\x85\x9d\xcc\x41\x30\x39\xfe\x7e\x4a\x36\xd7\x51\x66\x5d\x6e\xd5\xa6\xe4\xa6\x45\x37\x3b\x7e\xf7\x46\x93\x52\xb2\xae\x3d\x42\xab\xdf\xb3\x54\x7e\xd9\xa5\x9b\xe8\x6b\x4c\x13\x86\xa8\xde\x8d\xe8\xad\x09\x03\x3e\x44\xf5\xf2\xcf\xbd\xd1\x4a\x7d\xef\x0f\x94\x06\x6c\x1a\xed\x68\x4b\x8a\xe5\x62\x9f\x86\xeb\x4c\xea\x65\x7f\xfa\x9d\xb1\x45\xc7\xab\x77\x83\x5c\xb3\xd3\x3f\x4c\xeb\x22\xf7\x5a\x11\x28\x81\x69\x8d\xbe\x02\xbf\x5b\x6d\x17\xd1\x0b\x99\xe4\xfe\xaa\x17\xc1\xa9\x13\x7f\xf6\xee\xc8\x3c\x35\xff\x03\x00\x00\xff\xff\xf2\xf1\x39\xcc\x26\x02\x00\x00"

func dataCommonDevDepBuildBuildShTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildBuildShTpl,
		"data/common/dev-dep/build/build.sh.tpl",
	)
}

func dataCommonDevDepBuildBuildShTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildBuildShTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/build.sh.tpl", size: 550, mode: os.FileMode(493), modTime: time.Unix(1440178707, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevDepBuildUpstartConfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcf\x31\x4f\x03\x31\x0c\x05\xe0\xdd\xbf\xe2\xa9\x48\x30\x95\xd0\xc2\x31\xde\xca\xc8\xc8\x80\x3a\xa4\x17\x17\x45\x4a\xed\x28\x76\x5b\x50\xd5\xff\xce\x11\x81\x54\xa6\x28\xef\x49\x9f\xed\xc4\x36\xb5\x5c\x3d\xab\x60\x71\x3e\x43\xe2\x9e\x71\xb9\x60\x89\x17\x16\x6e\xd1\x39\x61\xfb\x85\x57\x77\x5d\x10\x35\xb6\x1a\x4f\xf2\xf7\xa2\xe4\x7d\x76\xac\x06\x0c\x44\xe6\xb1\x39\x66\xa6\x1d\xa4\xf0\x91\x0b\xde\xd7\x8f\x4f\xc3\x66\x2e\xb4\xfe\xcf\x1f\x9e\x37\x44\x37\x78\x63\x24\x95\x3b\xc7\x29\x8a\xc3\x15\x33\xdb\x11\x57\xc5\x2e\x9a\x63\xa7\x0d\x89\xab\x51\x55\xf3\x65\x87\xf8\x93\x27\x58\x61\xae\x7d\x68\x5f\x9e\x80\x70\xb0\x16\x8a\x4e\xb1\x84\x6d\x96\x70\x75\xc9\x38\x86\x63\xfc\xe9\x3e\xae\xd2\xfb\xf9\x8b\xf5\x78\xbb\x22\x96\x84\x5f\xe5\x3b\x00\x00\xff\xff\x9a\xf6\x9d\x7c\x0c\x01\x00\x00"

func dataCommonDevDepBuildUpstartConfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevDepBuildUpstartConfTpl,
		"data/common/dev-dep/build/upstart.conf.tpl",
	)
}

func dataCommonDevDepBuildUpstartConfTpl() (*asset, error) {
	bytes, err := dataCommonDevDepBuildUpstartConfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev-dep/build/upstart.conf.tpl", size: 268, mode: os.FileMode(420), modTime: time.Unix(1440179369, 0)}
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
	"data/aws-vpc-public-private/build/template.json": dataAwsVpcPublicPrivateBuildTemplateJson,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
	"data/common/dev-dep/build/Vagrantfile.fragment.tpl": dataCommonDevDepBuildVagrantfileFragmentTpl,
	"data/common/dev-dep/build/Vagrantfile.tpl": dataCommonDevDepBuildVagrantfileTpl,
	"data/common/dev-dep/build/build.sh.tpl": dataCommonDevDepBuildBuildShTpl,
	"data/common/dev-dep/build/upstart.conf.tpl": dataCommonDevDepBuildUpstartConfTpl,
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
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"template.json": &bintree{dataAwsVpcPublicPrivateBuildTemplateJson, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
			"dev-dep": &bintree{nil, map[string]*bintree{
				"build": &bintree{nil, map[string]*bintree{
					"Vagrantfile.fragment.tpl": &bintree{dataCommonDevDepBuildVagrantfileFragmentTpl, map[string]*bintree{
					}},
					"Vagrantfile.tpl": &bintree{dataCommonDevDepBuildVagrantfileTpl, map[string]*bintree{
					}},
					"build.sh.tpl": &bintree{dataCommonDevDepBuildBuildShTpl, map[string]*bintree{
					}},
					"upstart.conf.tpl": &bintree{dataCommonDevDepBuildUpstartConfTpl, map[string]*bintree{
					}},
				}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

