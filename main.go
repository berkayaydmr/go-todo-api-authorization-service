package main

import (
	pb "authorization-service/authorizationproto"
	"authorization-service/common"
	server "authorization-service/server"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	env := common.GetEnvironment()

	logger := common.NewLogger(env.Debug)
	zap.ReplaceGlobals(logger)

	redisClient := redis.NewClient(&redis.Options{
		Addr: common.GetEnvironment().RedisUrl,
	})

	newServer := server.NewServer(redisClient)

	list, err := net.Listen(env.Network, env.RouterUrl)
	if err != nil {
		zap.S().Error("Error: ", err)
		return
	}
	zap.S().Info(env.Network, env.RouterUrl, " listening")

	s := grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(s, newServer)

	if err := s.Serve(list); err != nil {
		zap.S().Error("Error: ", err)
		return
	}

}
