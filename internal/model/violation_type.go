package model

type ViolationType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo,omitempty"`
}
