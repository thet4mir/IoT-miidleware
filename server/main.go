package main

import (
	"context"
	"dashboard/server/devicepb"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type Device struct {
	VendorID  string `json:"vendor_id" redis:"vendor_id`
	ProductID string `json:"product_id" redis:"product_id"`
	SerialID  string `json:"serial" redis:"serial"`
	Action    bool   `json:"action" redis:"action"`
}

func (*server) DeviceUpdate(_ context.Context, req *devicepb.DeviceUpdateRequest) (*devicepb.DeviceUpdateResponse, error) {
	fmt.Println("Device Update Called")
	var ctx = context.Background()

	deviceId := idBuilder(req.Device.Serial, req.Device.IdVendor, req.Device.IdProduct)
	val, err := json.Marshal(req.Device)
	fmt.Println(val)
	if err != nil {
		log.Fatal(err)
	}
	setToRedis(ctx, deviceId, string(val))
	res := devicepb.DeviceUpdateResponse{}
	return &res, nil
}

func (*server) DeviceList(req *devicepb.DeviceListRequest, stream devicepb.DeviceService_DeviceListServer) error {

	var ctx = context.Background()

	fmt.Println("Start streaming to")
	var deviceList []*devicepb.Device

	for {
		deviceList = nil
		devices := getAllKeys(ctx, "*")

		for _, key := range devices {
			var d *devicepb.Device
			device := getFromRedis(ctx, key)
			json.Unmarshal([]byte(device), &d)
			deviceList = append(deviceList, d)
		}
		res := &devicepb.DeviceListResponse{
			Device: deviceList,
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("request canceled: %s", ctx.Err())
		default:
		}
		stream.Send(res)
		time.Sleep(time.Second * 1)
	}
}

func idBuilder(serial string, vendor string, product string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s.%s.%s", serial, vendor, product))
	id := builder.String()
	return id
}

type server struct{}

var redisClient *redis.Client

func main() {
	ctx := context.TODO()
	connectRedis(ctx)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}
	fmt.Println("server launched...")
	s := grpc.NewServer()
	devicepb.RegisterDeviceServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}

func connectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func setToRedis(ctx context.Context, key, val string) error {
	err := redisClient.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getFromRedis(ctx context.Context, key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func getAllKeys(ctx context.Context, key string) []string {
	keys := []string{}

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
