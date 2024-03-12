package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	chat_proto "github.com/nguyentrunghieu15/redis-example/example/udemy/chatapplication/chat_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func toValidASCII(s string) string {
	var result = []rune{}
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || v == '_' || v == '.' || v == '-' {
			result = append(result, v)
		}
	}
	return string(result)
}

func main() {
	addr := "localhost:5000"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Can't connect to server with error", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()
	client := chat_proto.NewChatServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)
	var name string
	for {
		fmt.Print("Enter your username: ")
		name, _ = reader.ReadString('\n')
		name = toValidASCII(strings.Trim(name, "\n"))
		res, err := client.ReqJoinChat(context.Background(), &chat_proto.Client{Name: name})
		if err != nil {
			log.Fatal("Error:", err)
		}
		if !res.IsAccept {
			fmt.Print("Your username is used\n")
			continue
		}
		break
	}

	md := metadata.Pairs("client", name)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("failed to call server: %v\n", err)
	}
	c := make(chan byte, 0)

	go func() {

		greet := []string{"Hi", "Im student", "Bye"}
		for _, v := range greet {
			err := stream.Send(&chat_proto.ClientMessage{Name: name, Message: v})
			if err != nil {
				stream.CloseSend()
				break
			}
			time.Sleep(time.Second * 5)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				log.Fatalln("Server closed stream")
				close(c)
				break
			}
			if err != nil {
				log.Fatalln("Error: ", err)
				close(c)
				break
			}
			fmt.Printf("%s:%s\n", r.Name, r.Message)
		}
	}()
	<-c
}
