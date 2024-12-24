package future

type Future interface {
	Resolve(interface{}) error
}
