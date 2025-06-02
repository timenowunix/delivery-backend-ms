package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"order-service/internal/config"
	"order-service/internal/db"
	"order-service/internal/kafka"
	"order-service/internal/order/handler"
	"order-service/internal/order/repository"
	"order-service/internal/order/service"

	orderv1 "order-service/api/order/v1"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не удалось загрузить .env: ", err)
	}

	// Создаём context, который автоматически завершится при Ctrl+C или SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop() // Освобождает ресурсы после срабатывания сигнала

	cfg := config.NewConfig()

	// Подключаем Kafka producer
	brokers := []string{os.Getenv("KAFKA_BROKER")}
	topic := os.Getenv("KAFKA_TOPIC")
	producer := kafka.NewProducer(brokers, topic)
	defer producer.Close()

	// Подключение к базе Postgres
	dbpool, err := db.ConnectPostgres(ctx, cfg)
	if err != nil {
		log.Println("Ошибка подключения к базе:", err)
		return
	}
	defer dbpool.Close() // Закроется при завершении программы

	// Настройка gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Ошибка прослушивания порта:", err)
		return
	}

	grpcServer := grpc.NewServer()

	// Инициализация слоёв
	repo := repository.NewOrderRepository(dbpool)
	orderService := service.NewOrderService(repo)
	orderHandler := handler.NewOrderHandler(orderService, producer)

	// Регистрируем gRPC сервис
	orderv1.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Println("gRPC сервер слушает на :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("Ошибка при запуске gRPC: %v", err)
		}
	}()

	// Ждём сигнал завершения (например, Ctrl+C)
	<-ctx.Done()
	log.Println("Получен сигнал завершения, останавливаем сервер...")

	// Завершаем сервер без обрыва соединений
	grpcServer.GracefulStop()
}
