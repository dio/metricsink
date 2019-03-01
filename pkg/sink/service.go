package sink

import (
	"io"
	"log"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2"
)

type server struct{}

var _ v2.MetricsServiceServer = &server{}

// New ...
func New() v2.MetricsServiceServer {
	return &server{}
}

func (s *server) StreamMetrics(stream v2.MetricsService_StreamMetricsServer) error {
	log.Println("Started stream")
	for {
		_, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
