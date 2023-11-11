package resp

// 分页响应结果
type PageResult[T any] struct {
	PageSize int   `json:"pageSize"`
	PageNum  int   `json:"pageNum"`
	Total    int64 `json:"total"`
	List     T     `json:"list"` // ! 注意这里的别名
}
