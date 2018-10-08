// Copyright © 2018 J. Strobus White.
// This file is part of the blocktop blockchain development kit.
//
// Blocktop is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Blocktop is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with blocktop. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"net/http"
	"strconv"

	consensus "github.com/blocktop/go-consensus"
	kernel "github.com/blocktop/go-kernel"
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
	s.RegisterService(new(kernel.RPC), "Kernel")
}
