package user

import (
	"github.com/RafaelFernando12/api-golang/domain"
	"github.com/RafaelFernando12/api-golang/pkg/log"
)

type userService struct {
	userRepository domain.UserRepository
	log            *log.MultiLogger
}

func NewUserService(userRepository domain.UserRepository, log log.MultiLogger) *userService {
	service := &userService{
		userRepository: userRepository,
		log:            &log,
	}

	return service
}

func (u *userService) Create(user *domain.User) error {
	return nil
}
