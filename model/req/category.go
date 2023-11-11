package req

type SaveOrUpdateCategory struct {
	ID       int    `json:"id"`
	Name     string `json:"name" mapstructure:"name"`
	Pic      string `json:"pic" mapstructure:"pic"`
	Sort     int    `json:"sort" validate:"required,min=1" mapstructure:"sort"`
	ParentId int    `json:"parent_id" mapstructure:"parent_id"`
	Status   string `json:"status" mapstructure:"status"`
}
