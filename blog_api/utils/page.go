package utils

import "math"

type Page struct {
	// 當前頁面
	Page int `json:"page`
	// 每頁記錄數
	Size int `json:"size"`
	// 總紀錄數
	Total int
}

func (page *Page) GetPage() int {
	// 最大頁數
	max := int(math.Ceil(float64(page.Total) / float64(page.Size)))
	// 當前頁數大於最大頁數 取最大值
	if page.Page > max {
		page.Page = max
	}
	return page.Page
}

// 讀取數據起始位置
func (page *Page) GetsStart() int {
	return (page.Page - 1) * page.Size
	// 40...49((5-1) * 10)
}
