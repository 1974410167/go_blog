package e

type ParamsError struct {
}

func (p *ParamsError) Error() string {
	return "ID is must greater than 0"
}

var ParamsInsError = &ParamsError{}
