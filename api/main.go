package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/marioarranzr/kubernetes-go-grpc/pb"
	"google.golang.org/grpc"
)

var gcdClient pb.GCDServiceClient

type Result struct {
	GCD string `json:"gcd"`
}

func main() {
	var target = flag.String("target", "gcd-service", "help message for flagname")
	var port = flag.String("port", "3001", "help message for flagname")
	flag.Parse()

	// Connect to GCD service
	conn, err := grpc.Dial(*target+":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcdClient = pb.NewGCDServiceClient(conn)

	// Set up HTTP server
	router := httprouter.New()
	router.GET("/gcd/:a/:b", gcdHandler)

	// Run HTTP server
	log.Fatal(http.ListenAndServe(":"+*port, router))
}

func gcdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a, err := strconv.ParseUint(ps.ByName("a"), 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid parameter A :: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	b, err := strconv.ParseUint(ps.ByName("b"), 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid parameter B :: %v", err.Error()), http.StatusInternalServerError)
		return
	}
	// Call GCD service
	req := &pb.GCDRequest{A: a, B: b}
	res, err := gcdClient.Compute(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := Result{
		GCD: fmt.Sprint(res.Result),
	}
	jsonObj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonObj)
}
