package response

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type option func(*Success)

func DataOpts(data interface{}) option {
	return func(d *Success) {
		d.Data = data
	}
}

func NewSuccess(code int, opts ...option) *Success {
	s := &Success{}

	for _, opt := range opts {
		opt(s)
	}

	return &Success{
		Code: code,
		Data: s.Data,
	}
}
