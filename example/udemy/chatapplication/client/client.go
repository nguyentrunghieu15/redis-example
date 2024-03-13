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

func ClientChat(cc chat_proto.ChatServiceClient, id string, c chan (bool)) {
	md := metadata.Pairs("id", id)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := cc.ClientChat(ctx)
	if err != nil {
		log.Fatalln("Can't chat with server")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		mess := strings.TrimFunc(text, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '.' || r == '-' || r == '_')
		})
		err := stream.Send(&chat_proto.ClientMessage{Id: id, Message: mess})
		if err != nil {
			break
		}
	}
	close(c)
}

func ServerChat(cc chat_proto.ChatServiceClient, id string, c chan (bool)) {
	md := metadata.Pairs("id", id)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := cc.ServerChat(ctx, &chat_proto.Client{Id: id})
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
		fmt.Printf("%s:%s\n", res.GetId(), res.GetMessage())
	}
	close(c)
}

func main() {
	addr := "localhost:5000"
	// Client chat
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Can't connect to server with error", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()
	client_chat := chat_proto.NewChatServiceClient(conn)

	//Server chat
	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Can't connect to server with error", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()
	server_chat := chat_proto.NewChatServiceClient(cc)

	// Request join world chat
	res, err := client_chat.ReqJoinChat(context.Background(), &chat_proto.Client{Id: ""})
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Your ID in world chat
	id := res.GetId()
	fmt.Println("Your id:", id)

	c := make(chan (bool), 0)
	go ClientChat(client_chat, id, c)
	go ServerChat(server_chat, id, c)
	<-c
}
