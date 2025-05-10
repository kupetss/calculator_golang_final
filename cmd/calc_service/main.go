package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"calculator/api"
	"calculator/internal/grpcserver"
	"calculator/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Инициализация БД
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal("Failed to init DB:", err)
	}
	defer db.Close()

	// Запуск gRPC сервера вычислений
	go grpcserver.StartGRPCServer()
	time.Sleep(1 * time.Second) // Даем время на запуск сервера

	r := gin.Default()

	// Public routes
	r.POST("/register", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required,min=3"`
			Password string `json:"password" binding:"required,min=6"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := storage.RegisterUser(db, req.Username, req.Password); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "User created"})
	})

	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, err := storage.AuthenticateUser(db, req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Генерация JWT токена
		expiresAt := time.Now().Add(storage.TokenExpire)
		claims := &storage.Claims{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiresAt),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				Issuer:    "calculator-app",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(storage.JWTSecretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":   tokenString,
			"expires": expiresAt.Unix(),
		})
	})

	// Protected routes
	auth := r.Group("/", AuthMiddleware(db))
	{
		auth.POST("/calculate", func(c *gin.Context) {
			userID := c.GetInt("userID")
			var req struct {
				Expression string `json:"expression" binding:"required"`
			}

			if err := c.BindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Сохраняем задачу
			taskID, err := storage.SaveTask(db, userID, req.Expression, "pending")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
				return
			}

			// Вычисление через gRPC
			conn, err := grpc.Dial(
				"localhost:50051",
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithBlock(),
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Calculation service unavailable"})
				return
			}
			defer conn.Close()

			client := api.NewCalculatorClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			res, err := client.Calculate(ctx, &api.CalculationRequest{
				Expression: req.Expression,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if res.Error != "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
				return
			}

			// Обновляем результат
			if err := storage.CompleteTask(db, taskID, res.Result); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"id":     taskID,
				"status": "completed",
				"result": res.Result,
			})
		})

		auth.GET("/tasks", func(c *gin.Context) {
			userID := c.GetInt("userID")
			tasks, err := storage.GetUserTasks(db, userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, tasks)
		})
	}

	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			return
		}

		claims := &storage.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(storage.JWTSecretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
