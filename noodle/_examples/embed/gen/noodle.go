//Code generated by gommon from config/gommon.yml DO NOT EDIT.
package gen

import (
	"time"

	"github.com/dyweb/gommon/noodle"
)

func init() {

	dirs := map[string]noodle.EmbedDir{"": {
		FileInfo: noodle.FileInfo{
			FileName:    "assets",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1519202702, 0),
			FileIsDir:   true,
		},
		Entries: []noodle.FileInfo{{
			FileName:    "404",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1518678160, 0),
			FileIsDir:   true,
		}, {
			FileName:    "idx",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1519187199, 0),
			FileIsDir:   true,
		}, {
			FileName:    "noidx",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1518677629, 0),
			FileIsDir:   true,
		}, {
			FileName:    ".noodleignore",
			FileSize:    147,
			FileMode:    436,
			FileModTime: time.Unix(1518819613, 0),
			FileIsDir:   false,
		}, {
			FileName:    "index.html",
			FileSize:    105,
			FileMode:    436,
			FileModTime: time.Unix(1519202702, 0),
			FileIsDir:   false,
		}},
	}, "404": {
		FileInfo: noodle.FileInfo{
			FileName:    "404",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1518678160, 0),
			FileIsDir:   true,
		},
		Entries: []noodle.FileInfo{},
	}, "idx": {
		FileInfo: noodle.FileInfo{
			FileName:    "idx",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1519187199, 0),
			FileIsDir:   true,
		},
		Entries: []noodle.FileInfo{{
			FileName:    "sub",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1519187227, 0),
			FileIsDir:   true,
		}, {
			FileName:    "index.html",
			FileSize:    180,
			FileMode:    436,
			FileModTime: time.Unix(1518677561, 0),
			FileIsDir:   false,
		}, {
			FileName:    "main.css",
			FileSize:    38,
			FileMode:    436,
			FileModTime: time.Unix(1518677599, 0),
			FileIsDir:   false,
		}, {
			FileName:    "main.js",
			FileSize:    51,
			FileMode:    436,
			FileModTime: time.Unix(1518677599, 0),
			FileIsDir:   false,
		}},
	}, "idx/sub": {
		FileInfo: noodle.FileInfo{
			FileName:    "sub",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1519187227, 0),
			FileIsDir:   true,
		},
		Entries: []noodle.FileInfo{{
			FileName:    "index.html",
			FileSize:    115,
			FileMode:    436,
			FileModTime: time.Unix(1519187227, 0),
			FileIsDir:   false,
		}},
	}, "noidx": {
		FileInfo: noodle.FileInfo{
			FileName:    "noidx",
			FileSize:    4096,
			FileMode:    2147484157,
			FileModTime: time.Unix(1518677629, 0),
			FileIsDir:   true,
		},
		Entries: []noodle.FileInfo{{
			FileName:    "main.css",
			FileSize:    37,
			FileMode:    436,
			FileModTime: time.Unix(1518677629, 0),
			FileIsDir:   false,
		}, {
			FileName:    "main.js",
			FileSize:    50,
			FileMode:    436,
			FileModTime: time.Unix(1518677629, 0),
			FileIsDir:   false,
		}, {
			FileName:    "noindex.html",
			FileSize:    192,
			FileMode:    436,
			FileModTime: time.Unix(1518677561, 0),
			FileIsDir:   false,
		}},
	}}

	box := noodle.EmbedBox{
		Dirs: dirs,
		Data: []byte{0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x86, 0xb2, 0x50, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd, 0x0, 0x0, 0x0, 0x2e, 0x6e, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x54, 0xcb, 0xb1, 0xaa, 0x85, 0x30, 0x10, 0x84, 0xe1, 0x7e, 0x9f, 0x62, 0xc0, 0x2e, 0x85, 0xde, 0xc2, 0xf7, 0xb9, 0x4, 0x77, 0x34, 0x81, 0x4d, 0x2, 0xc9, 0x86, 0xe3, 0xe3, 0x1f, 0xac, 0xe4, 0x94, 0xc3, 0x37, 0xff, 0x2, 0x4f, 0x79, 0x20, 0xe2, 0x68, 0xa5, 0xb0, 0xba, 0xd4, 0xa6, 0xfc, 0x2f, 0x4d, 0xa7, 0x71, 0x60, 0x41, 0xbe, 0x6a, 0xeb, 0x44, 0x8d, 0x85, 0x22, 0x61, 0x4d, 0xf3, 0xe2, 0x99, 0x8d, 0xbf, 0x82, 0x4f, 0x36, 0x3d, 0x62, 0x57, 0x91, 0xfd, 0x6f, 0xdf, 0xc2, 0xab, 0xd1, 0xc, 0x9e, 0x88, 0xa7, 0x19, 0x98, 0x55, 0xd9, 0x11, 0x71, 0x36, 0x53, 0x76, 0xc9, 0x7a, 0x6f, 0x61, 0xf5, 0xdb, 0xdf, 0xff, 0xb3, 0xbe, 0x1, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x75, 0x72, 0x72, 0x7f, 0x6d, 0x0, 0x0, 0x0, 0x93, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa1, 0x45, 0x55, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0xb2, 0xc9, 0x28, 0xc9, 0xcd, 0xb1, 0xe3, 0xb2, 0xc9, 0x48, 0x4d, 0x4c, 0xb1, 0xe3, 0x52, 0x50, 0x50, 0x50, 0xb0, 0x29, 0xc9, 0x2c, 0xc9, 0x49, 0xb5, 0xf3, 0x54, 0x48, 0xcc, 0x55, 0x28, 0xca, 0xcf, 0x2f, 0xb1, 0xd1, 0x87, 0x8, 0x70, 0xd9, 0xe8, 0x43, 0x14, 0xd9, 0x24, 0xe5, 0xa7, 0x54, 0x82, 0xb4, 0x18, 0x22, 0x14, 0xa9, 0x17, 0x2b, 0x64, 0xe6, 0xa5, 0xa4, 0x56, 0xe8, 0x81, 0x8c, 0xb3, 0xd1, 0xcf, 0x30, 0x4, 0x29, 0x87, 0xaa, 0xd3, 0x7, 0x5b, 0x1, 0x8, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x3e, 0xe1, 0xdb, 0x8, 0x51, 0x0, 0x0, 0x0, 0x69, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x94, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x64, 0xce, 0x31, 0xae, 0x3, 0x21, 0xc, 0x4, 0xd0, 0x9e, 0x53, 0x58, 0x1c, 0x60, 0xd1, 0xf6, 0x5e, 0xf7, 0xff, 0x18, 0x7c, 0x70, 0x4, 0x89, 0x21, 0x11, 0x76, 0x91, 0xbd, 0x7d, 0x44, 0xd8, 0x2e, 0xed, 0x68, 0xfc, 0x3c, 0x58, 0xac, 0x9, 0x39, 0x2c, 0x1c, 0x33, 0x39, 0x0, 0x0, 0xb4, 0x6a, 0xc2, 0xf4, 0x7, 0xb1, 0x41, 0xed, 0x99, 0xdf, 0xdb, 0xac, 0x60, 0x58, 0xf1, 0xaa, 0x48, 0xed, 0xf, 0x18, 0x2c, 0x87, 0x57, 0x3b, 0x85, 0xb5, 0x30, 0x9b, 0x87, 0x32, 0xf8, 0x76, 0xf8, 0x16, 0x6b, 0xdf, 0x92, 0xaa, 0x27, 0x87, 0x61, 0xb1, 0xf8, 0xff, 0xcc, 0xe7, 0x7c, 0xb2, 0xff, 0xb2, 0x65, 0x27, 0x87, 0x9a, 0x46, 0x7d, 0x19, 0xe8, 0x48, 0xd7, 0xfd, 0x5d, 0x3d, 0x61, 0x58, 0xf1, 0x74, 0x2e, 0x20, 0x7c, 0xd7, 0x7e, 0x2, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0xfe, 0x33, 0x1f, 0xd9, 0x7c, 0x0, 0x0, 0x0, 0xb4, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa9, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x4a, 0xca, 0x4f, 0xa9, 0x54, 0xa8, 0xe6, 0x52, 0x50, 0x50, 0x50, 0x48, 0x4a, 0x4c, 0xce, 0x4e, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4d, 0xce, 0xcf, 0xc9, 0x2f, 0xb2, 0x52, 0xa8, 0x4c, 0xcd, 0xc9, 0xc9, 0x2f, 0xb7, 0xe6, 0xaa, 0x5, 0x4, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x48, 0x46, 0x3, 0xbf, 0x2c, 0x0, 0x0, 0x0, 0x26, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa9, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x52, 0x2f, 0x2d, 0x4e, 0x55, 0x28, 0x2e, 0x29, 0xca, 0x4c, 0x2e, 0x51, 0xb7, 0xe6, 0xe2, 0x4a, 0xce, 0xcf, 0x2b, 0xce, 0xcf, 0x49, 0xd5, 0xcb, 0xc9, 0x4f, 0xd7, 0x50, 0x2f, 0xc9, 0xc8, 0x2c, 0x56, 0xc8, 0x2c, 0x56, 0xa8, 0x4c, 0xcd, 0xc9, 0xc9, 0x2f, 0x57, 0x28, 0x48, 0x4c, 0x4f, 0x55, 0x54, 0xd7, 0xb4, 0x6, 0x4, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x39, 0xd, 0x75, 0x6d, 0x39, 0x0, 0x0, 0x0, 0x33, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x63, 0x23, 0x55, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x12, 0x0, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0xb2, 0xc9, 0x28, 0xc9, 0xcd, 0xb1, 0xe3, 0xb2, 0xc9, 0x48, 0x4d, 0x4c, 0xb1, 0xe3, 0x52, 0x50, 0x50, 0x50, 0xb0, 0x29, 0xc9, 0x2c, 0xc9, 0x49, 0xb5, 0xf3, 0x54, 0x48, 0xcc, 0x55, 0xc8, 0xcc, 0x4b, 0x49, 0xad, 0xd0, 0x3, 0x29, 0xb1, 0xd1, 0x87, 0x8, 0x73, 0xd9, 0xe8, 0x43, 0x94, 0xda, 0x24, 0xe5, 0xa7, 0x54, 0x82, 0x34, 0x1a, 0xa2, 0x2b, 0x55, 0xc8, 0xcc, 0x2b, 0xce, 0x4c, 0x49, 0x55, 0x28, 0x2e, 0x4d, 0xb2, 0xd1, 0xcf, 0x30, 0x4, 0x69, 0x81, 0xaa, 0xd5, 0x7, 0x5b, 0x6, 0x8, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0xdc, 0xc9, 0x85, 0x5d, 0x55, 0x0, 0x0, 0x0, 0x73, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xb8, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x4a, 0xca, 0x4f, 0xa9, 0x54, 0xa8, 0xe6, 0x52, 0x50, 0x50, 0x50, 0x48, 0x4a, 0x4c, 0xce, 0x4e, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4d, 0xce, 0xcf, 0xc9, 0x2f, 0xb2, 0x52, 0x48, 0x2f, 0x4a, 0x4d, 0xcd, 0xb3, 0xe6, 0xaa, 0x5, 0x4, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x64, 0x6c, 0x1a, 0xf4, 0x2b, 0x0, 0x0, 0x0, 0x25, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xb8, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd, 0x0, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x52, 0x2f, 0x2d, 0x4e, 0x55, 0x28, 0x2e, 0x29, 0xca, 0x4c, 0x2e, 0x51, 0xb7, 0xe6, 0xe2, 0x4a, 0xce, 0xcf, 0x2b, 0xce, 0xcf, 0x49, 0xd5, 0xcb, 0xc9, 0x4f, 0xd7, 0x50, 0x2f, 0xc9, 0xc8, 0x2c, 0x56, 0xc8, 0x2c, 0x56, 0x48, 0x2f, 0x4a, 0x4d, 0xcd, 0x53, 0x28, 0x48, 0x4c, 0x4f, 0x55, 0x54, 0xd7, 0xb4, 0x6, 0x4, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x2c, 0x13, 0x1c, 0x99, 0x38, 0x0, 0x0, 0x0, 0x32, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x3, 0x4, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x94, 0x36, 0x4f, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x12, 0x0, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x6c, 0xce, 0x31, 0xe, 0xc3, 0x20, 0xc, 0x40, 0xd1, 0x9d, 0x53, 0x58, 0x1c, 0x20, 0x28, 0xbb, 0xe3, 0xbd, 0xc7, 0xa0, 0xe0, 0xca, 0xb4, 0x40, 0x2a, 0xec, 0xa1, 0xb9, 0x7d, 0x95, 0x92, 0xb1, 0xab, 0xfd, 0xfd, 0x64, 0x14, 0x6b, 0x95, 0x1c, 0xa, 0xc7, 0x4c, 0xe, 0x0, 0x0, 0xad, 0x58, 0x65, 0xba, 0x41, 0x6c, 0xd0, 0x77, 0x83, 0xd2, 0x33, 0x7f, 0x96, 0x33, 0xc3, 0x30, 0x57, 0x33, 0xab, 0xa5, 0xbf, 0x60, 0x70, 0xdd, 0xbc, 0xda, 0x51, 0x59, 0x85, 0xd9, 0x3c, 0xc8, 0xe0, 0xc7, 0xe6, 0x5b, 0x2c, 0x7d, 0x49, 0xaa, 0x9e, 0x1c, 0x86, 0x49, 0xe3, 0x7d, 0xcf, 0xc7, 0x75, 0x2a, 0xeb, 0x7f, 0x5e, 0x56, 0x72, 0xa8, 0x69, 0x94, 0xb7, 0x81, 0x8e, 0x74, 0x39, 0x4f, 0xf5, 0x84, 0x61, 0x8e, 0x4f, 0x6f, 0x42, 0x18, 0x7e, 0x9f, 0x7f, 0x3, 0x0, 0x0, 0xff, 0xff, 0x50, 0x4b, 0x7, 0x8, 0x8c, 0x98, 0x6f, 0x99, 0x7e, 0x0, 0x0, 0x0, 0xc0, 0x0, 0x0, 0x0, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x86, 0xb2, 0x50, 0x4c, 0x75, 0x72, 0x72, 0x7f, 0x6d, 0x0, 0x0, 0x0, 0x93, 0x0, 0x0, 0x0, 0xd, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0x0, 0x0, 0x0, 0x0, 0x2e, 0x6e, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa1, 0x45, 0x55, 0x4c, 0x3e, 0xe1, 0xdb, 0x8, 0x51, 0x0, 0x0, 0x0, 0x69, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0xa8, 0x0, 0x0, 0x0, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x94, 0x36, 0x4f, 0x4c, 0xfe, 0x33, 0x1f, 0xd9, 0x7c, 0x0, 0x0, 0x0, 0xb4, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0x31, 0x1, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa9, 0x36, 0x4f, 0x4c, 0x48, 0x46, 0x3, 0xbf, 0x2c, 0x0, 0x0, 0x0, 0x26, 0x0, 0x0, 0x0, 0xc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0xe9, 0x1, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xa9, 0x36, 0x4f, 0x4c, 0x39, 0xd, 0x75, 0x6d, 0x39, 0x0, 0x0, 0x0, 0x33, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0x4f, 0x2, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x63, 0x23, 0x55, 0x4c, 0xdc, 0xc9, 0x85, 0x5d, 0x55, 0x0, 0x0, 0x0, 0x73, 0x0, 0x0, 0x0, 0x12, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0xc1, 0x2, 0x0, 0x0, 0x69, 0x64, 0x78, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xb8, 0x36, 0x4f, 0x4c, 0x64, 0x6c, 0x1a, 0xf4, 0x2b, 0x0, 0x0, 0x0, 0x25, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0x56, 0x3, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0xb8, 0x36, 0x4f, 0x4c, 0x2c, 0x13, 0x1c, 0x99, 0x38, 0x0, 0x0, 0x0, 0x32, 0x0, 0x0, 0x0, 0xd, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0xbd, 0x3, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x50, 0x4b, 0x1, 0x2, 0x14, 0x3, 0x14, 0x0, 0x8, 0x0, 0x8, 0x0, 0x94, 0x36, 0x4f, 0x4c, 0x8c, 0x98, 0x6f, 0x99, 0x7e, 0x0, 0x0, 0x0, 0xc0, 0x0, 0x0, 0x0, 0x12, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x81, 0x30, 0x4, 0x0, 0x0, 0x6e, 0x6f, 0x69, 0x64, 0x78, 0x2f, 0x6e, 0x6f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x50, 0x4b, 0x5, 0x6, 0x0, 0x0, 0x0, 0x0, 0x9, 0x0, 0x9, 0x0, 0x19, 0x2, 0x0, 0x0, 0xee, 0x4, 0x0, 0x0, 0x0, 0x0},
	}

	noodle.RegisterEmbedBox("test", box)

}