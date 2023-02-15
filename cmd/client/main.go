package main

import (
	"context"
	"fmt"
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
	"github.com/pablogolobaro/dockertool-legend/pkg/console"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	flags, err := getFlags()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := chooseContext(flags.Timeout)
	defer cancelFunc()

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%v", flags.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	getStatsStream(ctx, conn)
}

func getStatsStream(ctx context.Context, conn grpc.ClientConnInterface) {
	client := pb.NewContainerStatsServiceClient(conn)

	stream, err := client.GetStatsStream(ctx, &pb.GetStatsMessage{})
	if err != nil {
		log.Fatal(err)
	}
	for true {
		stats, err := stream.Recv()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		console.WriteToConsoleClient(stats)
	}

}

func chooseContext(timeout int) (context.Context, context.CancelFunc) {
	if timeout > 0 {
		return context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	} else {
		return context.WithCancel(context.Background())
	}
}
