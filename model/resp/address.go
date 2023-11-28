package resp

type AddressListVO struct {
	ID        int    ` json:"id" mapstructure:"id"`
	MemberID  int    ` json:"member_id"`
	Name      string ` json:"name"`
	Phone     string ` json:"phone"`
	Address   string ` json:"address"`
	IsDefault string ` json:"is_default"`
}
