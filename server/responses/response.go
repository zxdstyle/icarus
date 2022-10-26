package responses

type (
	Response interface {
		StatusCode() int
		Content() any
	}
)
