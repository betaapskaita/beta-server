package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/betaapskaita/beta-server/libs/rpcs"
	"github.com/sirupsen/logrus"
)

func Start(httpPort string, rpcServer *rpcs.RpcServer) {
	grpc := rpcServer.WrappedGrpc

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()

		allowCors(resp, req)
		if grpc.IsGrpcWebRequest(req) || grpc.IsAcceptableGrpcCorsRequest(req) {
			grpc.ServeHTTP(resp, req)
		}

		go logEnd(start, req)
	})

	fmt.Println("HTTP server listening on", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		log.Fatal("Failed to start a HTTP server:", err)
	}
}

func logEnd(tm time.Time, req *http.Request) {
	logrus.WithFields(logrus.Fields{
		"host":     req.Host,
		"method":   req.Method,
		"url":      req.URL,
		"proto":    req.Proto,
		"duration": time.Since(tm),
	}).Info("Request")
}

func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}
