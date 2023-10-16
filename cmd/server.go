/*
Copyright © 2023 ALexi Lallas alexilallasengcomp@gmail.com
*/

package cmd

import (
	"github.com/alexilallas/quiz/internal/core/usecase"
	pb "github.com/alexilallas/quiz/internal/grpc"
	"github.com/alexilallas/quiz/internal/server"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const serverPort = ":8080"

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the api for Quiz",
	Long:  `Start the api for Quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", serverPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var grpcServer = grpc.NewServer()
		server := server.ProvideServer(server.Validator{}, usecase.ProvideQuizUseCase())
		pb.RegisterQuizServer(grpcServer, server)

		log.Printf("gRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
