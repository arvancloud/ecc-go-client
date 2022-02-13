package resource

type SSHKey struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}
