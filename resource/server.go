package resource

type Server struct {
	Name           string          `json:"name,omitempty"`
	DiskSize       int32           `json:"disk_size,omitempty"`
	NetworkIds     []string        `json:"network_ids,omitempty"`
	FlavorId       string          `json:"flavor_id,omitempty"`
	ImageId        string          `json:"image_id,omitempty"`
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
	SSHKey         bool            `json:"ssh_key,omitempty"`
	KeyName        string          `json:"key_name,omitempty"`
	Count          int32           `json:"count,omitempty"`
	CreateType     string          `json:"crate_type,omitempty"`
}
