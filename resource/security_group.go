package resource

type SecurityGroup struct {
	// ID of security group
	Id string `json:"id,omitempty"`
	// name of security group
	Name string `json:"name,omitempty"`
	// description about security group
	Description string `json:"description,omitempty"`
	// Real name of security group
	RealName string `json:"real_name,omitempty"`
	// Security group is read only
	ReadOnly bool `json:"read_only,omitempty"`
	// Instances with this security group
	Abraks []interface{} `json:"abraks,omitempty"`
	// rules of security group
	Rules []SecurityGroupRule `json:"rules,omitempty"`
	// list of security group tags
	// Tags []Tag `json:"tags,omitempty"`
}
