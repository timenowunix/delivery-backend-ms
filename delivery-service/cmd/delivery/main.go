package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"delivery-service/internal/config"
	"delivery-service/internal/db"
	"delivery-service/internal/delivery/handler"
	"delivery-service/internal/delivery/repository"
	"delivery-service/internal/delivery/service"

	deliveryv1 "delivery-service/api/delivery/v1"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не удалось загрузить .env: ", err)
	}

	// Контекст с автоотменой при SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.NewConfig()

	// Подключение к PostgreSQL
	dbpool, err := db.ConnectPostgres(ctx, cfg)
	if err != nil {
		log.Println("Ошибка подключения к базе:", err)
		return
	}
	defer dbpool.Close()

	// gRPC-сервер и слушатель
	lis, err := net.Listen("tcp", ":50052") // Порт для delivery
	if err != nil {
		log.Println("Ошибка прослушивания порта:", err)
		return
	}

	grpcServer := grpc.NewServer()

	// Слои
	repo := repository.NewPgxRepository(dbpool)
	deliveryService := service.NewService(repo)
	deliveryHandler := handler.NewGRPCHandler(deliveryService)

	// Регистрируем gRPC-сервис
	deliveryv1.RegisterDeliveryServiceServer(grpcServer, deliveryHandler)

	// Запуск в горутине
	go func() {
		log.Println("🚚 delivery-service gRPC слушает на :50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("Ошибка при запуске gRPC: %v", err)
		}
	}()

	// Ждём завершения
	<-ctx.Done()
	log.Println("Получен сигнал завершения, останавливаем сервер...")
	grpcServer.GracefulStop()
}
