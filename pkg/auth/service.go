package auth

type Service interface {
	ValidUser(username, password string) bool
	Permission(username, userid string) bool
}
type Repository interface {
	UserExists(email string) bool
	GetUserIdByMail(email string) string
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}
func (s *service) ValidUser(email, password string) bool {
	return s.r.UserExists(email)
}

func (s *service) Permission(username, userid string) bool {
	uuid := s.r.GetUserIdByMail(username)
	return uuid == userid
}
