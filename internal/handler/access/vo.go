package access

type PolicyReq struct {
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}

type UpdatePolicyReq struct {
	OldPolicy []string `json:"oldPolicy"`
	NewPolicy []string `json:"newPolicy"`
}
