package defanginganipaddress

import "strings"

func defangIPaddr(address string) string {
	adds := strings.Split(address, ".")
	return strings.Join(adds, "[.]")
}
