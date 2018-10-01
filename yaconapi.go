package yaconapi

type Api struct {
	XOrgID string
	Token  string
}

func NewApi(token string) (*Api, error) {
	a := Api{}
	return &a, nil
}
