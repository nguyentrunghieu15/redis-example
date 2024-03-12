package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/nguyentrunghieu15/redis-example/example/udemy/chatapplication/chat_grpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	REDIS_CHANEL_NAME string = "world"
	REDIS_ADDR        string = "192.168.88.129:6379"
	REDIS_PASSWORD    string = ""
	REDIS_DB          int    = 0
	REDIS_CLIENT      string = "client"

	SERVER_ADDR string = "0.0.0.0:5000"
)

type ServerChat struct {
	pb.ChatServiceServer
}

var redis_client = redis.NewClient(&redis.Options{
	Addr:     REDIS_ADDR,
	Password: REDIS_PASSWORD, // no password set
	DB:       REDIS_DB,       // use default DB
})

func (s *ServerChat) ReqJoinChat(ctx context.Context, client *pb.Client) (*pb.ResponseJoinReq, error) {
	log.Println("Invoke a client : ", client.Name)
	res := redis_client.SAdd(context.Background(), REDIS_CLIENT, client.Name)
	if res.Err() != nil {
		log.Println("Server error: ", res.Err().Error())
		return &pb.ResponseJoinReq{IsAccept: false}, res.Err()
	}

	if res.Val() == 0 {
		return &pb.ResponseJoinReq{IsAccept: false}, nil
	}
	return &pb.ResponseJoinReq{IsAccept: true}, nil
}

func (s *ServerChat) Chat(stream pb.ChatService_ChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "BidirectionalStreamingEcho: failed to get metadata")
	}
	var client_name string
	if t, ok := md["client"]; ok {
		client_name = t[0]
	}

	isMember := redis_client.SIsMember(context.Background(), REDIS_CLIENT, client_name)
	if !isMember.Val() {
		return fmt.Errorf("Server rufuse your chat")
	}

	defer func() {
		redis_client.SRem(context.Background(), REDIS_CLIENT, client_name)
	}()

	var client = redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: REDIS_PASSWORD, // no password set
		DB:       REDIS_DB,       // use default DB
	})

	sub := client.Subscribe(context.Background(), REDIS_CHANEL_NAME)
	defer sub.Close()
	c := make(chan byte, 0)
	go func() {
		for {
			// Revice message from another
			subres, suberr := sub.ReceiveMessage(context.Background())
			if suberr != nil {
				err := stream.Send(&pb.ClientMessage{Name: "Server", Message: suberr.Error()})
				if err != nil {
					log.Println("Error:", err)
				}
				break
			} else {
				mess := strings.Split(subres.Payload, ":")
				if mess[0] == client_name {
					continue
				}
				fmt.Printf("Server send for %s mess from %s: %s\n", client_name, mess[0], mess[1])
				err := stream.Send(&pb.ClientMessage{Name: mess[0], Message: mess[1]})
				if err != nil {
					log.Println("Error:", err)
					break
				}
			}
		}
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(c)
				break
			}
			if err != nil {
				log.Println("Error: ", err)
				close(c)
				break
			}
			pub := redis_client.Publish(context.Background(), REDIS_CHANEL_NAME, fmt.Sprintf("%s:%s", res.Name, res.Message))

			if pub.Val() == 0 {
				stream.Send(&pb.ClientMessage{Name: "Server", Message: "Can't send message"})
				break
			}
		}
	}()
	<-c
	return nil
}

func main() {
	lis, err := net.Listen("tcp", SERVER_ADDR)
	if err != nil {
		log.Fatalln("Cant init server listen on addr", SERVER_ADDR)
	}
	status := redis_client.Ping(context.Background())
	if status.Err() != nil {
		log.Fatalln("Cant connect Redis", status.Err(), status.Val(), redis_client)
	}

	redis_client.FlushDB(context.Background())

	defer lis.Close()
	log.Println("Init server ", SERVER_ADDR)
	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterChatServiceServer(s, &ServerChat{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Cant init grpc server\n", err)
	}
}
