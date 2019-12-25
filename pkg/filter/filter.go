package filter

import (
	"fmt"

	pinyin "github.com/mozillazg/go-pinyin"
)

type filter struct{}

type IdTitile struct {
	Id    int64
	Title string
}

func NewFilter() *filter {
	return &filter{}
}

func (f *filter) Start(search string, allList []*IdTitile) {
	a := pinyin.NewArgs()
	a.Heteronym = true
	pinyin.Pinyin(search, a)

	fmt.Println()
}

func init() {
	fmt.Println("filter")

}
