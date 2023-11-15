package todos

// type TodoService interface {
// 	CreateTodo(text string, userId uint) (*Todo, error)
// 	GetTodoByID(id uint) (*Todo, error)
// 	GetAllTodos() ([]*Todo, error)
// }

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}
func (s *TodoService) CreateTodo(text string, userId uint) (*Todo, error) {
	newTodo := &Todo{
		Text:   text,
		Done:   false,
		UserID: userId,
	}
	createdTodo, err := s.repo.Create(newTodo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}

func (s *TodoService) GetTodoByID(id uint) (*Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) GetAllTodos() ([]*Todo, error) {
	return s.repo.FindAll()
}
