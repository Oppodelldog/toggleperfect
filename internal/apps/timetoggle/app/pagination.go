package app

import (
	"math"
	"reflect"
)

type Pagination struct {
	Page     int
	PerPage  int
	NumItems int
}

func (p Pagination) Index(itemNo int) int {
	return itemNo + (p.PerPage * (p.Page - 1))
}
func (p Pagination) NextPage() int {
	if p.HasPage(p.Page + 1) {
		return p.Page + 1
	}

	return 1
}

func (p Pagination) HasPage(i int) bool {
	if p.NumItems == 0 || i <= 0 || p.PerPage <= 0 {
		return false
	}

	return p.GetLastPage() >= i
}

func (p Pagination) GetLastPage() int {
	return int(math.Ceil(float64(p.NumItems) / float64(p.PerPage)))
}

func (p Pagination) GetCurrentPageItems(items interface{}) []interface{} {
	v := reflect.ValueOf(items)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("items must be slice or array")
	}

	var pageItems []interface{}
	for _, i := range p.GetPageIndices() {
		pageItems = append(pageItems, v.Index(i).Interface())
	}

	return pageItems
}

func (p Pagination) GetPageIndices() []int {
	var indices []int
	for i := 0; i < p.PerPage; i++ {
		index := p.Index(i)
		if index >= p.NumItems {
			break
		}
		indices = append(indices, index)
	}

	return indices
}
