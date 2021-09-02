package auth

type Service struct {
	repoService *repoStruct
}

func NewAuthService(repo *repoStruct) *Service {
	return &Service{
		repoService: repo,
	}
}
func (authService *Service) loginUser(Login Login) (User, error) {
	userf, err := authService.repoService.FindByName(Login.Name)
	if err != nil {
		return User{}, err
	}
	return userf, nil

}
