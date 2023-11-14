package todos

type TodoService interface {
	CreateTodo(text string, userId uint) (*Todo, error)
	GetTodoByID(id uint) (*Todo, error)
}

type todoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) CreateTodo(text string, userId uint) (*Todo, error) {
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

func (s *todoService) GetTodoByID(id uint) (*Todo, error) {
	return s.repo.FindByID(id)
}
