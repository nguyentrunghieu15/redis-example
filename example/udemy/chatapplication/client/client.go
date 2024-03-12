package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	chat_proto "github.com/nguyentrunghieu15/redis-example/example/udemy/chatapplication/chat_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func ClientChat(cc chat_proto.ChatServiceClient, name string, c chan (bool)) {
	md := metadata.Pairs("client", name)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	stream, err := cc.ClientChat(ctx)
	if err != nil {
		log.Fatalln("Can't chat with server")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		mess := strings.Trim(text, "\n")
		fmt.Println("Send:", mess)
		err := stream.Send(&chat_proto.ClientMessage{Name: name, Message: mess})
		if err != nil {
			log.Fatalln("Can't send message with error:", err)
		}
	}
}

func ServerChat(cc chat_proto.ChatServiceClient, name string, c chan (bool)) {
	md := metadata.Pairs("client", name)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := cc.ServerChat(ctx, &chat_proto.Client{Name: name})
	if err != nil {
		log.Fatalln("Error:", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error:", err)
		}
		fmt.Printf("%s:%s\n", res.Name, res.Message)
	}
	close(c)
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
		name = strings.Trim(name, "\n")
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
	c := make(chan (bool), 0)
	go ClientChat(client, name, c)
	go ServerChat(client, name, c)
	<-c
}
