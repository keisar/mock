//go:generate mockgen -package source -destination ../output/source_mock.go -source=source.go
package source

type X struct{}

type S interface {
	F(X)
}
