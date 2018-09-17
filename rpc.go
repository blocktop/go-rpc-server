package rpc

import (
	"net/http"
	"strconv"

	consensus "github.com/blckit/go-consensus"
	grpc "github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/spf13/viper"
)


func Start() {
	s := grpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	registerServices(s)

	port := strconv.FormatInt(int64(viper.GetInt("rpc.port")), 10)
	http.Handle("/rpc", s)
	go http.ListenAndServe("localhost:"+port, nil)
}

func registerServices(s *grpc.Server) {
	s.RegisterService(new(consensus.RPC), "Consensus")
}