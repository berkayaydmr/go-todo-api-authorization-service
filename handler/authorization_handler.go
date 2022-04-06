package handler

import (
	pb "authorization-service/authorizationproto"
	"authorization-service/repository"
)

type AuthorizationService struct {
	RedisRepository repository.RedisClientInterface
}

func NewAuthorizationHandler(redisRepository repository.RedisClientInterface) *AuthorizationService {
	return &AuthorizationService{RedisRepository: redisRepository}
}

func (handler *AuthorizationService) AuthorizationUser(request *pb.AuthorizationRequest) *pb.AuthorizationResponse {
	response := &pb.AuthorizationResponse{
		ResponseStatusCode: 401,
	}

	redisToken := handler.RedisRepository.GetData("userAuth_" + request.GetUserId())
	sentToken := request.GetAuthorizationToken()

	if redisToken == sentToken {
		response.ResponseStatusCode = 200
		return response
	}

	return response
}
