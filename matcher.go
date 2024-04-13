package btcvanity

import "strings"

// isMatch checks if the btc wallet matches the prefix of pattern
func isMatch(pattern string, addr string) bool {
	addr = strings.ToUpper(addr[1 : len(addr)-1])
	pattern = strings.ToUpper(pattern)
	return strings.HasSuffix(addr, pattern)
}
