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
	"google.golang.org/grpc/metadata"
)

const (
	REDIS_CHANEL_NAME string = "world"
	REDIS_ADDR        string = "localhost:6379"
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

func (s *ServerChat) ClientChat(stream pb.ChatService_ClientChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't auth connection")
	}
	clients, ok := md["client"]
	if !ok {
		return fmt.Errorf("Can't auth client func client chat")
	}

	client_name := clients[0]
	fmt.Println(client_name)

	if r := redis_client.SIsMember(context.Background(), REDIS_CLIENT, client_name); r.Val() == false || r.Err() != nil {
		return fmt.Errorf("Client didt regist")
	}

	defer redis_client.SRem(context.Background(), REDIS_CLIENT, client_name)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		pub := redis_client.Publish(context.Background(), REDIS_CHANEL_NAME, fmt.Sprintf("%s:%s", res.Name, res.Message))
		if pub.Val() == 0 || pub.Err() != nil {
			return fmt.Errorf("Error send message")
		}
	}
	return nil
}

func (s *ServerChat) ServerChat(client *pb.Client, stream pb.ChatService_ServerChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't auth connection")
	}
	clients, ok := md["client"]
	if !ok {
		return fmt.Errorf("Can't auth client func server chat")
	}

	client_name := clients[0]

	if r := redis_client.SIsMember(context.Background(), REDIS_CLIENT, client_name); r.Val() == false || r.Err() != nil {
		return fmt.Errorf("Client didt regist")
	}
	defer redis_client.SRem(context.Background(), REDIS_CLIENT, client_name)
	sub := redis_client.Subscribe(context.Background(), REDIS_CHANEL_NAME)

	for {
		mess, err := sub.ReceiveMessage(context.Background())
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		payload := strings.Split(mess.Payload, ":")
		name := payload[0]
		text := payload[1]
		if name != client_name {
			err := stream.Send(&pb.ClientMessage{Name: name, Message: text})
			if err != nil {
				return err
			}
		}
	}
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
	defer redis_client.ShutdownSave(context.Background())

	defer lis.Close()
	log.Println("Init server ", SERVER_ADDR)
	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterChatServiceServer(s, &ServerChat{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Cant init grpc server\n", err)
	}
}
