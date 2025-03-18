package main

import (
	"log"
	"backend/graph"
	"backend/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// 连接 MongoDB
	db.ConnectDB()

	// 创建 Gin 服务器
	r := gin.Default()

	// 允许跨域请求（解决 OPTIONS 请求 404 问题）
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// **注册 GraphQL API 端点**
	r.POST("/graphql", gin.WrapH(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))))

	// 启动服务器
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
