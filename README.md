# OpsEnv

Модуль предназначен для работы с системными переменными окружения в Linux.
Он предоставляет функции для проверки наличия переменной, получения, получения с возможностью добавить дефолтное значение, установки, удаления и получения списка всех переменных.

## Возможности
- Проверка существования переменной окружения (`HasEnv`).
- Получение значение переменной окружения или ошибку (`GetEnv`).
- Получение значение переменной окружения или дефолтное значение (`GetEnvDefault`).
- Установка переменной окружения (`SetEnv`).
- Удаление переменной окружения (`UnsetEnv`).
- Получение списка всех переменных окружения (`GetAllEnv`).

## Требования
- Go 1.20 или выше.
- Unix-подобная ОС (Linux, macOS и т.д.).

## Установка
Склонируйте репозиторий или добавьте модуль в ваш проект:
```shell
go get github.com/mosregdata/ops-env
```

## Использование
Пример использования модуля:
```go
package main

import (
    "fmt"
    opsenv "github.com/mosregdata/ops-env"
)

func main() {
    // Проверяем существование переменной окружения
    hme := "HOME does not exist"
    if opsenv.HasEnv("HOME") {
        hme = "HOME exists"
    }
    fmt.Println("Variable", hme)

    // Получаем значение переменной окружения или ошибку
    value, err := opsenv.GetEnv("HOME")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Variable HOME: %s\n", value)

    // Получаем значение переменной окружения или дефолтное
    defval := opsenv.GetEnvDefault("NONEXISTENT", "DefaultValue")
    fmt.Printf("Variable is: %s\n", defval)

    // Устанавливаем переменную окружения
    er := opsenv.SetEnv("MYVAR", "myvalue")
    if er != nil {
        fmt.Printf("Error: %v\n", er)
        return
    }
    fmt.Println("MYVAR set successfully")

    // Получаем переменную MYVAR
    myvar, err := opsenv.GetEnv("MYVAR")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Variable MYVAR: %s\n", myvar)

    // Удаляем переменную окружения
    e := opsenv.UnsetEnv("MYVAR")
    if e != nil {
        fmt.Printf("Error: %v\n", e)
        return
    }
    fmt.Println("MYVAR unset successfully")

	// Получаем значение переменной MYVAR, которую уже удалили
    defmvr := opsenv.GetEnvDefault("MYVAR", "Default MYVAR Value")
    fmt.Printf("Variable MYVAR is: %s\n", defmvr)

    // Перебираем все переменные окружения
    envVars := opsenv.GetAllEnv()
    for key, item := range envVars {
        fmt.Printf("%s=%s\n", key, item)
    }
}
```

## Результат
Код примера представленный выше вернет такой результат:
```
Variable HOME exists
Variable HOME: /home/dude
Variable is: DefaultValue
MYVAR set successfully
Variable MYVAR: myvalue
MYVAR unset successfully
Variable MYVAR is: Default MYVAR Value
...
USER=dude
PAPERSIZE=a4
SHELL=/bin/bash
M2_HOME=/opt/maven
GDMSESSION=ubuntu
...
```