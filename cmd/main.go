package main

import (
	"flag"

	"github.com/smallnest/blockchain"
	"github.com/smallnest/log"
)

var (
	addr     = flag.String("addr", ":8972", "listened address")
	dataFile = flag.String("data", "./data", "data file")
)

func main() {
	flag.Parse()

	store, err := blockchain.NewLevelDBStore(*dataFile)
	if err != nil {
		log.Fatalf("failed to create leveldb store: %v", err)
	}
	defer store.Close()

	// 创建一个区块链
	var bc = &blockchain.Blockchain{
		Store: store,
	}
	bc.GenerateGenesisBlock()

	// 创建 rpc server
	var server = blockchain.NewServer(*addr, bc)

	// 启动服务
	if err := server.Serve(); err != nil {
		log.Errorf("failed to serve: %v", err)
		return
	}

	log.Info("exit mormally")
}
