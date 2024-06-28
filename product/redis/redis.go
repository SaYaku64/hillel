package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type DBClient struct {
	cache *redis.Client
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{ // Створює новий клієнт для підключення до Redis
		Addr:     "localhost:6379",
		Password: "", // пароль, якщо необхідний
		DB:       0,  // номер бази даних
	})

	pong, err := rdb.Ping(ctx).Result() // Перевірка підключення до сервера Redis
	if err != nil {
		fmt.Println("Помилка підключення:", err)
		return
	}
	fmt.Println("Підключено до Redis:", pong)

	dbClient := DBClient{cache: rdb}

	if err := dbClient.Set(ctx, "key", "value"); err != nil {
		fmt.Println("Помилка збереження:", err)
	}

	val, err := dbClient.Get(ctx, "key")
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Ключ не знайдено")
		} else {
			fmt.Println("Помилка отримання:", err)
		}
	}
	fmt.Println("Значення:", val)

	dbClient.Del(ctx, "key")

	// List (slice)
	dbClient.PushSlice(ctx, "slice", []string{"1", "2", "3"})

	vals, err := dbClient.RangeSlice(ctx, "slice")
	if err != nil {
		fmt.Println("Помилка отримання списку:", err)
	}
	fmt.Println("Список:", vals)

	// Hash (map)
	dbClient.SetMap(ctx, "keyForMap", map[string]any{"key1": "value1", "key2": 22})

	val, err = dbClient.GetMap(ctx, "keyForMap", "key1")
	if err != nil {
		fmt.Println("Помилка отримання геша:", err)
	}
	fmt.Println("Геш значення:", val)

	// Set (slice тільки з унікальними значеннями)
	dbClient.SetUniqueSlice(ctx, "keySet", []string{"1", "2", "1", "3", "2"})

	members, err := dbClient.GetUniqueSlice(ctx, "keySet")
	if err != nil {
		fmt.Println("Помилка отримання множини:", err)
	}
	fmt.Println("Множина:", members)

	/*
			age, err := strconv.Atoi(val)
		    if err != nil {
		        fmt.Println("Помилка конвертації:", err)
		        return
		    }

			/////////////

			user := User{Name: "Alice", Age: 30}

		    // Конвертуємо користувача в JSON
		    jsonData, err := json.Marshal(user)
		    if err != nil {
		        fmt.Println("Помилка конвертації в JSON:", err)
		        return
		    }

		    // Зберігаємо JSON як string
		    err = rdb.Set(ctx, "user:1", jsonData, 0).Err()
		    if err != nil {
		        fmt.Println("Помилка збереження:", err)
		        return
		    }

			// Отримуємо JSON як string
		    val, err := rdb.Get(ctx, "user:1").Result()
		    if err != nil {
		        fmt.Println("Помилка отримання:", err)
		        return
		    }

			byteStr := []byte(val)

		    // Конвертуємо JSON в структуру
		    var user User
		    err = json.Unmarshal(byteStr, &user)
		    if err != nil {
		        fmt.Println("Помилка конвертації з JSON:", err)
		        return
		    }

			// Конвертуємо string в []byte
		    data := []byte(val)

	*/
}

func (dc *DBClient) Set(ctx context.Context, key string, value any, expiration ...time.Duration) error {
	var exp time.Duration
	if len(expiration) > 0 {
		exp = expiration[0]
	}

	return dc.cache.Set(ctx, key, value, exp).Err()
}

func (dc *DBClient) Get(ctx context.Context, key string) (string, error) {
	return dc.cache.Get(ctx, key).Result()
}

func (dc *DBClient) Del(ctx context.Context, key string) error {
	return dc.cache.Del(ctx, key).Err()
}

// List (slice)
func (dc *DBClient) PushSlice(ctx context.Context, key string, values []string) {
	dc.cache.RPush(ctx, key, values)
}

func (dc *DBClient) RangeSlice(ctx context.Context, key string) ([]string, error) {
	return dc.cache.LRange(ctx, key, 0, -1).Result()
}

// Hash (map)
func (dc *DBClient) SetMap(ctx context.Context, key string, values map[string]any) {
	dc.cache.HSet(ctx, key, values)
}

func (dc *DBClient) GetMap(ctx context.Context, key string, keyValue string) (string, error) {
	return dc.cache.HGet(ctx, key, keyValue).Result()
}

// Set (slice тільки з унікальними значеннями)
func (dc *DBClient) SetUniqueSlice(ctx context.Context, key string, values []string) {
	dc.cache.SAdd(ctx, key, values)
}

func (dc *DBClient) GetUniqueSlice(ctx context.Context, key string) ([]string, error) {
	return dc.cache.SMembers(ctx, key).Result()
}

/*
Публікація повідомлення:

err := rdb.Publish(ctx, "mychannel", "Hello, Redis!").Err()
if err != nil {
    fmt.Println("Помилка публікації:", err)
    return
}


Підписка на канал:

pubsub := rdb.Subscribe(ctx, "mychannel")
defer pubsub.Close()

ch := pubsub.Channel()
for msg := range ch {
    fmt.Println("Отримано повідомлення:", msg.Payload)
}

*/
