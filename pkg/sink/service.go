package sink

import (
	"io"
	"log"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v2"
	"github.com/golang/protobuf/jsonpb"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v2.MetricsServiceServer = &server{}

// New ...
func New() v2.MetricsServiceServer {
	return &server{}
}

func (s *server) StreamMetrics(stream v2.MetricsService_StreamMetricsServer) error {
	log.Println("Started stream")
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		str, _ := s.marshaler.MarshalToString(in)
		log.Println(str)
	}
}
