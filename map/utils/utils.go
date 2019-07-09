package utils

var PRIME_LIST = [...]uint32{
	5,
	7,
	1009,
	2003,
	4001,
	8009,
	16001,
	// for testing
	100003,
	4248551,
	10000019,
}

func GetClosestPrime(val uint32) uint32 {
	min := uint32(1 << 31)
	for _, prime := range PRIME_LIST {
		if (prime-val)*(prime-val) < (min-val)*(min-val) {
			min = prime
		}
	}
	return min
}
