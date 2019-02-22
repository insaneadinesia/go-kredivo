package kredivo

type Env int8

const (
	_ Env = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)

var typeString = map[Env]string{
	Sandbox:    "https://sandbox.kredivo.com/kredivo/v2",
	Production: "https://api.kredivo.com/kredivo/v2",
}

// implement stringer
func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}
