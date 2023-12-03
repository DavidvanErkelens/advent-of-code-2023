package challenges

type Challenge interface {
	RunPartOne(string) string
	RunPartTwo(string) string
	DataFolder() string
}
