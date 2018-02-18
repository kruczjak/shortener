package main

import (
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
	"net"
	"shortener/config"
	"log"
	pb "shortener/grpc_server/proto"
	"google.golang.org/grpc/reflection"
	"context"
	"shortener/redis_client"
	"io"
	"fmt"
)

type server struct{}

func (s *server) ShortenUrl(ctx context.Context, in *pb.ShortenUrlRequest) (*pb.ShortenUrlResponse, error) {
	url, err := redis_client.ShortenUrl(in.Url, uint(in.OwnerId), in.Options)
	if err != nil {
		fmt.Println("test	")
		return nil, err
	}

	fmt.Println("test	")
	return &pb.ShortenUrlResponse{ShortenedUrl: url}, nil
}

func (s *server) ShortenUrls(stream pb.Shortener_ShortenUrlsServer) error {
	for {
		_, err := stream.Recv() // TODO
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}


	}
	return nil
}

func (s *server) Info(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	shortenedUrl, err := redis_client.FindShortenedUrl(in.ShortenedUrl)
	if err != nil {
		return nil, err
	}

	return &pb.InfoResponse{
		ShortenedUrl: &pb.ShortenedUrl{
			Url: shortenedUrl.Url,
			Clicks: uint32(shortenedUrl.Clicks),
			Options: string(shortenedUrl.Options),
			OwnerId: uint32(shortenedUrl.OwnerId),
		},
	}, nil
}

func (s *server) RemoveShortenedUrl(ctx context.Context, in *pb.RemoveShortenedUrlRequest) (*pb.RemoveShortenedUrlResponse, error) {
	success, err := redis_client.RemoveShortenedUrl(in.ShortenedUrl)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveShortenedUrlResponse{Success: success}, nil
}

func main() {
	godotenv.Load()

	lis, err := net.Listen("tcp", config.Getenv("GRPC_PORT", ":5001"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterShortenerServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
