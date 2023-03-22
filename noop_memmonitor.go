package gomemwatermark

// ############ NOOP Mem Monitor ################
type noopMemMonitor struct{}

// UseNoopMemMonitor creates a noop implementation of mem-monitor if we want to fully disable it
// with minimal penalty
func UseNoopMemMonitor() {
	memMonitorInstance = &noopMemMonitor{}
}

func (n *noopMemMonitor) Increment(a Sizeable) {}
func (n *noopMemMonitor) Decrement(a Sizeable) {}
func (n *noopMemMonitor) Fence()               {}
func (n *noopMemMonitor) Breached() bool       { return false }
