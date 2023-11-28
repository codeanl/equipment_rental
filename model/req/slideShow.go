package req

type SaveOrUpdateSlideshow struct {
	ID     int    `json:"id"`
	Name   string `json:"name" mapstructure:"name"`
	Url    string `json:"url" mapstructure:"url"`
	Status string `json:"status" mapstructure:"status"`
}

type SlideshowList struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Name     string `form:"name"`
	Status   string `form:"status"`
}
