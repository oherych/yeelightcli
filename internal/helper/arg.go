package helper

type Arg interface {
	Name() string
	Example() string
	Description() string
}
