package todo

type Service struct {
	repoService *repoStruct
}

func NewTodohService(repo *repoStruct) *Service {
	return &Service{
		repoService: repo,
	}
}
func (todoService *Service) insertTodo(S ITodo) error {
	err := todoService.repoService.InsertByName(S)
	return err

}
func (todoService *Service) users() []ITodo {
	return  todoService.repoService.FindUserByName()
}
func (todoService *Service) user(name string) ITodo {
	return  todoService.repoService.singleUser(name)
}
func (todoService *Service) dlete(req ITodo) error {
	return  todoService.repoService.deleteTodoMessage(req)
}
func (todoService *Service) update(req ITodo) error {
	return  todoService.repoService.updateTodoMessage(req)
}