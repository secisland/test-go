package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World! request url: "+c.Request.URL.Path)
	})
	// 3.监听端口，默认在8080
	// Run("")里面不指定端口号默认为8080
	r.Run(":8000")
}

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	kafka "github.com/segmentio/kafka-go"
// )

// func c() {
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
// 		GroupID:  "consumer-group-id",
// 		Topic:    "test22",
// 		MaxBytes: 10e6, // 10MB
// 	})

// 	for {
// 		m, err := r.ReadMessage(context.Background())
// 		if err != nil {
// 			break
// 		}
// 		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
// 	}

// 	if err := r.Close(); err != nil {
// 		log.Fatal("failed to close reader:", err)
// 	}
// }

// func p() {
// 	w := &kafka.Writer{
// 		Addr:         kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
// 		Topic:        "test22",
// 		Balancer:     &kafka.LeastBytes{},
// 		RequiredAcks: kafka.RequireAll, // ack模式
// 		Async:        true,             // 异步
// 	}

// 	err := w.WriteMessages(context.Background(),
// 		kafka.Message{
// 			Key:   []byte("Key-A"),
// 			Value: []byte("Hello World!"),
// 		},
// 		kafka.Message{
// 			Key:   []byte("Key-B"),
// 			Value: []byte("One!"),
// 		},
// 		kafka.Message{
// 			Key:   []byte("Key-C"),
// 			Value: []byte("Two!"),
// 		},
// 	)
// 	if err != nil {
// 		log.Fatal("failed to write messages:", err)
// 	}

// 	if err := w.Close(); err != nil {
// 		log.Fatal("failed to close writer:", err)
// 	}
// }
// func main() {
// 	go c()
// 	p()
// 	time.Sleep(300 * time.Second)
// 	// // to consume messages
// 	// topic := "test"
// 	// partition := 0

// 	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
// 	// if err != nil {
// 	// 	log.Fatal("failed to dial leader:", err)
// 	// }

// 	// conn.SetReadDeadline(time.Now().Add(300 * time.Second))
// 	// batch := conn.ReadBatch(1, 1e6) // fetch 10KB min, 1MB max

// 	// // b := make([]byte, 10e3) // 10KB max per message
// 	// b := make([]byte, 1) // 10KB max per message
// 	// for {
// 	// 	n, err := batch.Read(b)
// 	// 	if err != nil && err != io.EOF {
// 	// 		fmt.Println(err.Error())
// 	// 		break
// 	// 	}
// 	// 	if len(b[:n]) > 0 {
// 	// 		fmt.Println(string(b[:n]))
// 	// 	} else {
// 	// 		time.Sleep(1 * time.Second)
// 	// 	}
// 	// }

// 	// if err := batch.Close(); err != nil {
// 	// 	log.Fatal("failed to close batch:", err)
// 	// }

// 	// if err := conn.Close(); err != nil {
// 	// 	log.Fatal("failed to close connection:", err)
// 	// }
// }

// import (
// 	"fmt"
// 	"os"

// 	"github.com/spf13/cobra"
// )

// // 创建一个命令
// var rootCmd = &cobra.Command{
// 	Use:   "main",
// 	Short: "An main CLI application",
// 	Long:  "A detailed description of the CLI application",
// 	Args:  cobra.MinimumNArgs(0),
// 	PersistentPreRun: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
// 	},
// 	PreRun: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
// 	},
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Inside rootCmd Run with args: %v\n", args)
// 	},
// 	PostRun: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
// 	},
// 	PersistentPostRun: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
// 	},
// }

// // 创建一个子命令
// var versionCmd = &cobra.Command{
// 	Use:   "version",
// 	Short: "Print the version number of the CLI application",
// 	Long:  "All software has versions. This is CLI application's",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("1.0.0")
// 	},
// }

// var serveCmd = &cobra.Command{
// 	Use:   "server",
// 	Short: "start web server",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("hello world")
// 	},
// 	Args: cobra.MinimumNArgs(0),
// }

// // 添加子命令
// func init() {
// 	rootCmd.AddCommand(versionCmd) //go run .\cobra.go version
// 	rootCmd.AddCommand(serveCmd)   //go run .\cobra.go serve
// }

// func main() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
