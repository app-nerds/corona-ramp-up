// +build dev

/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package assets

import (
	"net/http"
	"os"
	pathpkg "path"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/union"
)

/*
 * Skip files we don't need to include
 */
var skipFunc = func(path string, fi os.FileInfo) bool {
	fname := fi.Name()
	extension := pathpkg.Ext(fname)

	return extension == ".go" ||
		extension == ".DS_Store" ||
		extension == ".md" ||
		extension == ".svg" ||
		fname == "LICENSE"
}

/*
 * Create a union of all paths that you want to include as
 * static assets
 */
var Assets = union.New(map[string]http.FileSystem{
	"/app": filter.Skip(http.Dir("./app"), skipFunc),
	// "/app": filter.Skip(minifyfs.Dir("./app"), skipFunc),
})
