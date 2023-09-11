package parser

import (
	"net/url"
	"strconv"
)

const (
	page           = "page"
	size           = "size"
	defaultPage    = 1
	defaultSize    = 10
	pageLowerLimit = 1
	sizeLowerLimit = 1
	sizeUpperLimit = 20
)

func ParsePageAndSize(queryValues url.Values) (int, int) {
	pageStr := queryValues.Get(page)
	sizeStr := queryValues.Get(size)

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < pageLowerLimit {
		page = defaultPage
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < sizeLowerLimit || size > sizeUpperLimit {
		size = defaultSize
	}

	return page, size
}
