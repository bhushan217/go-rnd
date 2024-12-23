package main

import (
	"testing"
	"context"
	"log"
	"github.com/jackc/pgx/v5"
	"github.com/bhushan217/go-rnd/db/sqlc"
	// "github.com/bhushan217/go-rnd/store"
	"github.com/stretchr/testify/assert"
	// "github.com/jackc/pgtype"
)

func TestCanVote(t *testing.T) {
	testCases := []struct {
		name string
		age  int
		want string
	}{
		{
			name: "Can vote",
			age:  18,
			want: "Bhushan can Vote",
		},
		{
			name: "Cannot vote",
			age:  15,
			want: "Bhushan can not Vote",
		},
		{
			name: "Can vote - edge case",
			age:  17,
			want: "Bhushan can not Vote",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canVote("Bhushan", tc.age)
			if got != tc.want {
				t.Errorf("canVote(%q, %d) = %q; want %q", "Bhushan", tc.age, got, tc.want)
			}
		})
	}
}

//	func TestGetPermisions(t *testing.T){
//		permissions :=
//	}
func TestIf(t *testing.T) {
	t.Run("Condition is true", func(t *testing.T) {
		got := If(true, "true", "false")
		want := "true"
		if got != want {
			t.Errorf("If(true, \"true\", \"false\") = %q; want %q", got, want)
		}
	})

	t.Run("Condition is false", func(t *testing.T) {
		got := If(false, "true", "false")
		want := "false"
		if got != want {
			t.Errorf("If(false, \"true\", \"false\") = %q; want %q", got, want)
		}
	})

	t.Run("With integers", func(t *testing.T) {
		got := If(true, 1, 0)
		want := 1
		if got != want {
			t.Errorf("If(true, 1, 0) = %d; want %d", got, want)
		}
	})
}

func TestPermission(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	connStr := "postgresql://root:secret@localhost:5643/simple_blog?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)
	q := db.New(conn)
	t.Run("List permissions", func(t *testing.T){
		permissions, err := q.ListPermission(context.Background())
		assert.Nil(err)
		assert.NotEmpty(permissions)
		assert.Len(permissions, 2)
	})
	t.Run("Create and list permissions", func(t *testing.T){
		permission, err := q.CreatePermission(context.Background(), "DELETE_USER")
		assert.Nil(err)
		assert.NotEmpty(permission)
		assert.Equal(permission.Title, "DELETE_USER")
		permissions, err := q.ListPermission(context.Background())
		assert.Nil(err)
		assert.NotEmpty(permissions)
		assert.Len(permissions, 3)
	})
}
