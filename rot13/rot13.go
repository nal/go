package rot13

// Interface
type rot13 interface {
	Encode(string) string
	Decode(string) string
}
