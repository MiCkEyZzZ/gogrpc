package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/MiCkEyZzZ/protoapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = ":8080"

func AskingDateTime(ctx context.Context, m protoapi.RandomClient) (*protoapi.DateTime, error) {
	request := &protoapi.RequestDateTime{
		Value: "Пожалуйста отправьте мне дату и время",
	}
	return m.GetDate(ctx, request)
}

func AskPass(ctx context.Context, m protoapi.RandomClient, seed int64, place int64) (*protoapi.RandomInt, error) {
	request := &protoapi.RandomParams{
		Seed:  seed,
		Place: place,
	}
	return m.GetRandom(ctx, request)
}

func AskRandom(ctx context.Context, m protoapi.RandomClient, seed int64, place int64) (*protoapi.RandomInt, error) {
	request := &protoapi.RandomParams{
		Seed:  seed,
		Place: place,
	}
	return m.GetRandom(ctx, request)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Использование порта по умолчанию:", port)
	} else {
		port = os.Args[1]
	}
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.NewSource(time.Now().Unix())
	seed := int64(rand.Intn(100))
	client := protoapi.NewRandomClient(conn)
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Дата и время работы сервера:", r.Value)
	length := int64(rand.Intn(20))
	p, err := AskPass(context.Background(), client, 100, length+1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Случайный пароль:", p.Value)
	place := int64(rand.Intn(100))
	i, err := AskRandom(context.Background(), client, seed, place)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Случайное число 1:", i.Value)
	k, err := AskRandom(context.Background(), client, seed, place-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Случайное число 2:", k.Value)
}
