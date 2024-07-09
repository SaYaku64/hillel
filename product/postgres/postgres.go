package main

import (
	"database/sql"
	"fmt"
	"os"

	secrets "github.com/ijustfool/docker-secrets"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	godotenv.Load()
	dockerSecrets, _ := secrets.NewDockerSecrets("")
	password, err := dockerSecrets.Get("postgres-password")
	if err != nil {
		panic(err)
	}

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Підключення до PostgreSQL встановлено!")

	if err := createTable(db); err != nil {
		fmt.Println("createTable error: err =", err)

		return
	}

	if err := insertUser(db, "Alex", 30); err != nil {
		fmt.Println("insertUser error: err =", err)

		return
	}

	// userId := 1

	// userFromDB, err := getUser(db, userId)
	// switch err {
	// case sql.ErrNoRows:
	// 	fmt.Println("Користувач не знайдений")
	// case nil:
	// default:
	// 	fmt.Println("getUser error: err =", err)
	// }
	// fmt.Printf("Get user: %+v\n", userFromDB)

	// if err := updateUser(db, userId, 12); err != nil {
	// 	fmt.Println("updateUser error: err =", err)

	// 	return
	// }
	// fmt.Printf("User with id %d updated\n", userId)

	// if err := deleteUser(db, userId); err != nil {
	// 	fmt.Println("deleteUser error: err =", err)

	// 	return
	// }
	// fmt.Printf("User with id %d deleted\n", userId)

	fmt.Println("All actions done!")
}

func createTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INT
    )`
	_, err := db.Exec(query)

	return err
}

func insertUser(db *sql.DB, name string, age int) error {
	sqlStatement := `
    INSERT INTO users (name, age)
    VALUES ($1, $2)
    RETURNING id`
	id := 0
	return db.QueryRow(sqlStatement, name, age).Scan(&id)
}

func getUser(db *sql.DB, id int) (user User, err error) {
	sqlStatement := `
    SELECT id, name, age
    FROM users
    WHERE id = $1`
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age)

	return
}

func updateUser(db *sql.DB, id int, age int) error {
	sqlStatement := `
    UPDATE users
    SET age = $2
    WHERE id = $1`
	_, err := db.Exec(sqlStatement, id, age)

	return err

	// check num of affected rows
	// count, err := res.RowsAffected()
	// if err != nil {
	//     panic(err)
	// }
}

func deleteUser(db *sql.DB, id int) error {
	sqlStatement := `
    DELETE FROM users
    WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)

	return err
}
