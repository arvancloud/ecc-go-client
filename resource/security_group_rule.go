package resource

type SecurityGroupRule struct {
	// ID of security group rule
	Id string `json:"id,omitempty"`
	// group id of security group rule
	GroupId string `json:"group_id,omitempty"`
	// description about security group rule
	Description string `json:"description,omitempty"`
	// The remote IP prefix to be associated with this security group rule.
	Ip string `json:"ip,omitempty"`
	// The maximum port number in the range that is matched by the securitygroup rule.
	PortEnd int32 `json:"port_end,omitempty"`
	// The minimum port number in the range that is matched by the security group rule.
	PortStart int32 `json:"port_start,omitempty"`
	// The protocol that is matched by the security group rule
	Protocol string `json:"protocol,omitempty"`
	// The direction in which the security group rule is applied(ingress or egress)
	Direction string `json:"direction,omitempty"`
	// IPV4 or IPV6
	EtherType string `json:"ether_type,omitempty"`
	// Rule creation time
	CreatedAt string `json:"created_at,omitempty"`
	// Rule update time
	UpdatedAt string `json:"updated_at,omitempty"`
}
