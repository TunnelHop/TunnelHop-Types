package messages

type TunnelMessage uint64

const (
	NewControlTunnel TunnelMessage = iota
	NewProxyTunnel
	ProxyTunnelEstablished
	DomainActivated
	ControlTunnelHealthCheck
	ControlTunnelHealthCheckAck
)

type Message struct {
	MessageType TunnelMessage
	Payload     []byte
}

// ApplicationDomain is the key to the registry. This represents
// the protocol being requested, the host and the port. This allows
// us to differentiate requests that might be requesting HTTPS vs HTTP, etc.
type ApplicationDomain struct {
	Protocol  string `json:"protocol"`
	Host      string `json:"host"`
	AuthToken string `json:"auth_token"`
}

type ControlMessage struct {
	ApplicationName string `json:"application_name"`
	AuthToken       string `json:"auth_token"`
	ApplicationDomain
}

type HealthCheckAckMessage struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

type ProxyMessage struct {
}
