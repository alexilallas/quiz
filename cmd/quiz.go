/*
Copyright © 2023 ALexi Lallas alexilallasengcomp@gmail.com
*/

package cmd

import (
	"bufio"
	"context"
	pb "github.com/alexilallas/quiz/internal/grpc"
	"github.com/alexilallas/quiz/internal/handler"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

// quizCmd represents the client command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start Quiz.",
	Long:  `Start the quiz where you can answer questions`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalln("failed to connect to server: ", err)
		}
		defer func(conn *grpc.ClientConn) {
			if err = conn.Close(); err != nil {
				log.Println("failed to close connection: ", err)
			}
		}(conn)

		if err = handler.ProvideHandler(bufio.NewScanner(bufio.NewReader(os.Stdin))).QuizHandler(context.Background(), pb.NewQuizClient(conn)); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
