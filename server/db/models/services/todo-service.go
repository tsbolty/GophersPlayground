package services

// func (s *ComplexService) CreateTodoForUser(todoText string, userID uint) (*models.Todo, error) {
// 	// Check if the user exists using the UserRepository
// 	_, err := s.userRepo.FindByID(userID)
// 	if err != nil {
// 		// The user doesn't exist, return an error
// 		return nil, err
// 	}

// 	// The user exists, create the Todo using the TodoRepository
// 	todo := &models.Todo{
// 		Text:   todoText,
// 		Done:   false,
// 		UserID: userID,
// 	}
// 	err = s.todoRepo.Create(todo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return todo, nil
// }
