package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\r\n<body>\r\n  <p>Can you see this? → {{.Bar}}</p>\r\n</body>\r\n"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n</body>\r\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1541167388, 1541167388719312300),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1541331111, 1541331111950580600),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1541330534, 1541330534467116500),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1541167388, 1541167388718307800),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}}, "")
