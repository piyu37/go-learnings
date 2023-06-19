package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"practice/grpc"

	gGrpc "google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := gGrpc.Dial(":8888", gGrpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend: %v", err)
	}

	client := grpc.NewTasksClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(context.Background(), client)
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, client grpc.TasksClient, text string) error {
	t, err := client.Add(ctx, &grpc.Text{
		Text: text,
	})

	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully:", t)

	return nil
}

func list(ctx context.Context, client grpc.TasksClient) error {
	l, err := client.List(ctx, &grpc.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}

	for _, t := range l.Tasks {
		if t.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s\n", t.Text)
	}

	return nil
}
