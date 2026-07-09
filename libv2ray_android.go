package libv2ray

// ProcessFinder is an interface for Android process finding functionality.
// Apps using AndroidLibXrayLite should implement FindProcessByConnection()
// and pass the implementation to RegisterProcessFinder() before starting the core.
type ProcessFinder interface {
	// FindProcessByConnection finds the UID of the process that owns the given connection.
	//
	// network: Protocol type: "tcp" or "udp"
	// srcIP: Source IP address
	// srcPort: Source port
	// destIP: Destination IP address
	// destPort: Destination port
	// Returns the UID of the owning process, or -1 if not found.
	FindProcessByConnection(network, srcIP string, srcPort int, destIP string, destPort int) int
}

// RegisterProcessFinder keeps the Android API surface compatible with current v2rayNG.
// Xray-core v1.260327.0 does not expose common/net.RegisterAndroidProcessFinder,
// so process-based routing is unavailable for the exact 26.3.27 AAR and this is a no-op.
func (x *CoreController) RegisterProcessFinder(finder ProcessFinder) {
}
