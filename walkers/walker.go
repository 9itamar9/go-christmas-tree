package walkers

type Walker interface {
	Walk(path string, level int) string
}
