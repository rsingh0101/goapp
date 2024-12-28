package mariadb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB struct holds the database connection
type DB struct {
	Conn *sql.DB
}

// NewDB creates a new database connection
func NewDB(user, password, host, dbName string) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error verifying connection: %w", err)
	}

	fmt.Println("Connected to the database successfully!")
	return &DB{Conn: conn}, nil
}

// CreateTable creates the users table
func (db *DB) CreateTable() error {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		age INT NOT NULL,
		PRIMARY KEY (id)
	);`

	_, err := db.Conn.Exec(createTable)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	fmt.Println("Table created successfully!")
	return nil
}

// InsertUser inserts a user into the users table
func (db *DB) InsertUser(name string, age int) error {
	insertUser := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := db.Conn.Exec(insertUser, name, age)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	fmt.Println("User inserted successfully!")
	return nil
}

// QueryUsers queries all users from the users table
func (db *DB) QueryUsers() ([]User, error) {
	rows, err := db.Conn.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
func (db *DB) DeleteUser(name string, age int) error {
	deleteUser := `DELETE FROM users WHERE name =? AND age =? `
	_, err := db.Conn.Exec(deleteUser, name, age)
	if err != nil {
		fmt.Printf("Error deleting user: %v", err)
	}
	fmt.Println("User deleted successfully!")
	return nil
}

// User struct represents a user in the users table
type User struct {
	ID   int
	Name string
	Age  int
}
