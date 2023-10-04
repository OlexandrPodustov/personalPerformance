package bandwidththrottle

import (
	"net"
	"time"

	"golang.org/x/time/rate"
)

// BandwidthThrottle is a struct that manages bandwidth throttling for TCP connections.
type BandwidthThrottle struct {
	globalRate *rate.Limiter // Global bandwidth rate limiter
	connRates  map[net.Conn]*rate.Limiter
}

// NewBandwidthThrottle creates a new BandwidthThrottle with a global bandwidth rate.
func NewBandwidthThrottle(globalRate rate.Limit) *BandwidthThrottle {
	return &BandwidthThrottle{
		globalRate: rate.NewLimiter(globalRate, int(globalRate)),
		connRates:  make(map[net.Conn]*rate.Limiter),
	}
}

// SetGlobalRate sets the global bandwidth rate for all connections.
func (bt *BandwidthThrottle) SetGlobalRate(rate rate.Limit) {
	bt.globalRate.SetLimit(rate)
}

// SetConnectionRate sets the bandwidth rate for an individual connection.
func (bt *BandwidthThrottle) SetConnectionRate(conn net.Conn, rt rate.Limit) {
	bt.connRates[conn] = rate.NewLimiter(rt, int(rt))
}

// RemoveConnectionRate removes the bandwidth rate limit for an individual connection.
func (bt *BandwidthThrottle) RemoveConnectionRate(conn net.Conn) {
	delete(bt.connRates, conn)
}

// ThrottleConn wraps a net.Conn with bandwidth throttling capabilities.
type ThrottleConn struct {
	net.Conn
	limiter *rate.Limiter
}

// WrapConnection wraps a net.Conn with bandwidth throttling capabilities.
func (bt *BandwidthThrottle) WrapConnection(conn net.Conn) net.Conn {
	limiter, exists := bt.connRates[conn]
	if !exists {
		limiter = bt.globalRate
	}
	return &ThrottleConn{
		Conn:    conn,
		limiter: limiter,
	}
}

// Write throttles the write operation to match the specified bandwidth rate.
func (tc *ThrottleConn) Write(p []byte) (n int, err error) {
	if tc.limiter.Allow() {
		return tc.Conn.Write(p)
	}
	time.Sleep(time.Second)

	return 0, nil
}
