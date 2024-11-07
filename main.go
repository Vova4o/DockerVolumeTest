package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    // Убедитесь, что директория существует
    if _, err := os.Stat("/app/data"); os.IsNotExist(err) {
        os.MkdirAll("/app/data", 0755)
    }

    // Открываем файл, если он существует, или создаем новый
    file, err := os.OpenFile("/app/data/test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Получаем текущее время
    currentTime := time.Now()

    // Форматируем строку для записи
    testString := fmt.Sprintf("Hello, World! %s\n", currentTime)

    // Записываем строку в файл
    _, err = file.WriteString(testString)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Запись в файл выполнена успешно")
}