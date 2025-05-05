package main

// Array represents an MD array
type Array int

const (
	MD3200 Array = 3200
	MD3220 Array = 3220
	MD3400 Array = 3400
	MD3460 Array = 3460
)

// Map arrays to strings for the CLI
var arrayIds = map[Array][]string{
	MD3200: {"MD3200"},
	MD3220: {"MD3220"},
	MD3400: {"MD3400"},
	MD3460: {"MD3460"},
}
