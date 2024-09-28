package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

const (
	dbUser     = "root"
	dbPassword = "123ttk"
	dbHost     = "47.97.249.37"
	dbName     = "reservoir_data"
	excelFile  = "./水利局202408---2018-2024以来气象资料.xlsx"
)
func main() {
    // 连接 MySQL 数据库
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", dbUser, dbPassword, dbHost, dbName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 打开 Excel 文件
    f, err := excelize.OpenFile(excelFile)
    if err != nil {
        log.Fatal(err)
    }

    // 获取工作表名称
    sheetName := f.GetSheetName(0)

    // 读取行数据
    rows, err := f.Rows(sheetName)
    if err != nil {
        log.Fatal(err)
    }

    // 跳过表头行
    rows.Next()

    for rows.Next() {
        row, err := rows.Columns()
        if err != nil {
            log.Fatal(err)
        }

        // 确保有足够的列
        if len(row) < 6 {
            log.Println("行数据不足")
            continue
        }

        // 解析数据
        date := row[0]
        avgTemp := parseFloat(row[1])
        rainfall := parseFloat(row[2])
        avgHumidity := parseFloat(row[3])
        avgWindSpeed := parseFloat(row[4])
        avgPressure := parseFloat(row[5])

        // 插入数据到数据库
        _, err = db.Exec(`INSERT INTO weather_info (
            weather_condition, average_temperature, rainfall,
            humidity, wind_speed, air_pressure, date,
            created_at, updated_at, deleted_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
            "未知",                       // weather_condition
            avgTemp,                     // average_temperature
            rainfall,                    // rainfall
            avgHumidity,                 // humidity
            avgWindSpeed,                // wind_speed
            avgPressure,                 // air_pressure
            date,                        // date
            time.Now(),                  // created_at
            time.Now(),                  // updated_at
            nil,                         // deleted_at
        )
        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("数据成功插入到数据库")
}

// parseFloat 将字符串转换为浮点数
func parseFloat(s string) float64 {
    var value float64
    _, err := fmt.Sscan(s, &value)
    if err != nil {
        return 0.0 // 返回 0.0 如果转换失败
    }
    return value
}