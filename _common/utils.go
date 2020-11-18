package _common

type ErrorFields struct {
	Errors map[string]interface{} `json:"errorFields"`
}

func NewErrorField(key string, err error) ErrorFields {
	res := ErrorFields{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

type error interface {
	Error() string
}

func ErrorToString(err error) string {
	return err.Error()
}
