package email

type service struct {
	Repository
}

func NewEmailService(r Repository) Service {
	return service{Repository: r}
}
