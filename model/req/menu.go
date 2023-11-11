package req

type SaveOrUpdateMenu struct {
	ID        int    `json:"id"`
	Name      string `json:"name" mapstructure:"name"`
	Path      string `json:"path" mapstructure:"path"`
	Component string `json:"component" mapstructure:"component"`
	Icon      string `json:"icon" mapstructure:"icon"`
	Sort      int    `json:"sort" validate:"required,min=1" mapstructure:"sort"`
	ParentId  int    `json:"parent_id" mapstructure:"parent_id"`
	Status    string `json:"status" mapstructure:"status"`
}
