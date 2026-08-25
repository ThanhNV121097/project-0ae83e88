package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-0ae83e88/backend/migrations"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type app struct {
	db *pgxpool.Pool
}

type errorBody struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type greetingBody struct {
	Text string `json:"text"`
}

func main() {
	ctx := context.Background()
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer db.Close()

	if err := applyMigrations(ctx, db); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}

	srv := &http.Server{
		Addr:              ":" + port(),
		Handler:           routes(&app{db: db}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("serve: %v", err)
	}
}

func routes(a *app) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", a.healthz)
	mux.HandleFunc("GET /v1/greeting", a.greeting)
	return mux
}

func port() string {
	if v := os.Getenv("PORT"); v != "" {
		return v
	}
	if v := os.Getenv("APP_PORT"); v != "" {
		return v
	}
	return "8080"
}

func (a *app) healthz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := a.db.Ping(ctx); err != nil {
		http.Error(w, "database unavailable", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok\n"))
}

func (a *app) greeting(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var text string
	err := a.db.QueryRow(ctx, "select text from greetings where id = 1").Scan(&text)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "not_found")
		return
	}
	if err != nil {
		log.Printf("query greeting: %v", err)
		writeError(w, http.StatusInternalServerError, "internal_error")
		return
	}

	writeJSON(w, http.StatusOK, greetingBody{Text: text})
}

func applyMigrations(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, `create table if not exists schema_migrations (
		filename text primary key,
		applied_at timestamptz not null default now()
	)`)
	if err != nil {
		return fmt.Errorf("ensure schema_migrations: %w", err)
	}

	files, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}

	names := make([]string, 0, len(files))
	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, ".up.sql") {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		if err := applyMigration(ctx, db, name); err != nil {
			return err
		}
	}
	return nil
}

func applyMigration(ctx context.Context, db *pgxpool.Pool, name string) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin migration %s: %w", name, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var exists bool
	if err := tx.QueryRow(ctx, "select exists(select 1 from schema_migrations where filename = $1)", name).Scan(&exists); err != nil {
		return fmt.Errorf("check migration %s: %w", name, err)
	}
	if exists {
		return tx.Commit(ctx)
	}

	sql, err := migrations.Files.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, string(sql)); err != nil {
		return fmt.Errorf("execute migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, "insert into schema_migrations (filename) values ($1)", name); err != nil {
		return fmt.Errorf("record migration %s: %w", name, err)
	}
	return tx.Commit(ctx)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, code string) {
	writeJSON(w, status, errorBody{Error: errorPayload{Code: code, Message: "Request failed"}})
}
