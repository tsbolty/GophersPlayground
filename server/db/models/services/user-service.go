package services

// func (s *UserService) DeleteUser(userID uint) error {
// 	// Check if the user has any todos
// 	todos, err := s.todoRepo.FindByUserID(userID)
// 	if err != nil {
// 		return err
// 	}

// 	if len(todos) > 0 {
// 		// Handle the situation, maybe you want to prevent deletion or delete all todos first
// 		return fmt.Errorf("user has todos and cannot be deleted")
// 	}

// 	// If no todos are associated with the user, proceed to delete the user
// 	return s.userRepo.Delete(userID)
// }
