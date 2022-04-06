package server

import (
	pb "authorization-service/authorizationproto"
	"authorization-service/repository"
	"context"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

type server struct {
	pb.UnimplementedAuthorizationServiceServer
	redisClient *redis.Client
}

func NewServer(redis *redis.Client) *server {
	return &server{redisClient: redis}
}

func (server *server) AuthorizationUser(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {

	redisRepository := repository.NewRedisRepository(server.redisClient)
	zap.S().Info("grpc ready to response time: ", time.Now())

	response := &pb.AuthorizationResponse{
		ResponseStatusCode: 401,
	}

	redisToken := redisRepository.GetData("userAuth_" + in.GetUserId())
	sentToken := in.GetAuthorizationToken()

	if redisToken == sentToken {
		response.ResponseStatusCode = 200
		return response, nil
	}

	return response, nil
}
