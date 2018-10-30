package yaconapi

type Api struct {
	XOrgID string
	Token  string
}

func New(token string) (*Api, error) {
	a := Api{}
	return &a, nil
}
