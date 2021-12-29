package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	query := "SELECT " +
	"posts.title FROM posts LEFT JOIN  views " +
		"ON posts.id = views.post_id ORDER BY views.counter;"
	ctx := context.Background()
	url := "postgres://postgres:password@localhost:5432/blog"
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unablew to connect to the database %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	var result string
	err = conn.QueryRow(ctx, query).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}
	cfg.MaxConns = 8;cfg.MinConns = 4
	cfg.HealthCheckPeriod = 1 * time.Second
	cfg.MaxConnLifetime = 2 * time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		Timeout: cfg.ConnConfig.ConnectTimeout,
	}).DialContext
	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	var result2 string
	err = dbpool.QueryRow(ctx, query).Scan(&result2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result2)
}
