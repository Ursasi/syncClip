package util

type GetRequest struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
	MAC  string `json:"mac"`
}
type ConnectRequest struct {
	SID string `json:"sid"`
	DID string `json:"did"`
}

type DisconnectRequest struct {
	SID string `json:"sid"`
	DID string `json:"did"`
}

type RegisterRequest struct {
	Host string `json:"host"`
	Port string `json:"port"`
	MAC  string `json:"mac"`
}
type ReceiveRequest struct {
	Msg string
}
