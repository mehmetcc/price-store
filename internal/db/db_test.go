// db_test.go
package db

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mehmetcc/price-store/internal/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestCreatePriceUpdate starts a Postgres container, connects using our Connect function,
// creates a PriceUpdate record, and verifies that it was inserted.
func TestCreatePriceUpdate(t *testing.T) {
	ctx := context.Background()

	// Define the container request for PostgreSQL.
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13", // Choose your preferred Postgres image and version.
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").
			WithStartupTimeout(60 * time.Second),
	}

	// Start the PostgreSQL container.
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("failed to start container: %v", err)
	}
	// Terminate the container when the test is finished.
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %v", err)
		}
	}()

	// Retrieve the container's host and mapped port.
	host, err := pgContainer.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get container host: %v", err)
	}
	mappedPort, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("failed to get mapped port: %v", err)
	}

	// Construct the DSN string.
	dsn := fmt.Sprintf("host=%s user=testuser password=testpass dbname=testdb port=%s sslmode=disable", host, mappedPort.Port())

	// Create a minimal configuration instance.
	cfg := &config.Config{
		Dsn: dsn,
	}

	// Connect to the database.
	Connect(cfg.Dsn)

	// Create a new PriceUpdate record.
	pu := &PriceUpdate{
		Symbol:    "AAPL",
		Price:     150.50,
		Timestamp: time.Now().UTC(),
	}
	if err := Create(pu); err != nil {
		t.Fatalf("failed to create PriceUpdate: %v", err)
	}

	// Verify that the record was inserted by querying it back.
	var fetched PriceUpdate
	if err := db.First(&fetched, pu.ID).Error; err != nil {
		t.Fatalf("failed to fetch PriceUpdate: %v", err)
	}

	// Check that the fetched record matches the inserted values.
	if fetched.Symbol != pu.Symbol {
		t.Errorf("expected Symbol %q, got %q", pu.Symbol, fetched.Symbol)
	}
	if fetched.Price != pu.Price {
		t.Errorf("expected Price %v, got %v", pu.Price, fetched.Price)
	}
}

// TestConnectInvalidDSN verifies that Connect fails when given an invalid DSN.
// Since Connect uses log.Fatalf (which calls os.Exit), we run the test in a subprocess.
func TestConnectInvalidDSN(t *testing.T) {
	// When BE_FATAL is set, we expect Connect to call log.Fatalf.
	if os.Getenv("BE_FATAL") == "1" {
		cfg := &config.Config{
			Dsn: "invalid_dsn",
		}
		Connect(cfg.Dsn)
		// If Connect returns normally, something is wrong.
		return
	}

	// Run this test in a subprocess.
	cmd := exec.Command(os.Args[0], "-test.run=TestConnectInvalidDSN")
	cmd.Env = append(os.Environ(), "BE_FATAL=1")
	err := cmd.Run()

	// Expect the subprocess to exit with a non-zero exit code.
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 0 {
		// Test passed.
		return
	}
	t.Fatalf("expected non-zero exit code when using an invalid DSN, got: %v", err)
}

// TestCreateWithoutConnect verifies that calling Create without first initializing
// the global database connection (i.e. without calling Connect) causes a panic.
func TestCreateWithoutConnect(t *testing.T) {
	// Reset the global db variable.
	db = nil

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when calling Create without connecting, but did not panic")
		}
	}()

	pu := &PriceUpdate{
		Symbol:    "AAPL",
		Price:     150.0,
		Timestamp: time.Now(),
	}
	// This call should panic because db is nil.
	Create(pu)
	t.Error("should not reach this line; Create() should have panicked")
}
