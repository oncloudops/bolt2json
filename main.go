package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go.etcd.io/bbolt"
)

func showHelp() {
	fmt.Println("Usage: bolt2json [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	// 解析命令行参数
	dbPath := flag.String("db", "database.boltdb", "Path to the BoltDB file")
	exportPath := flag.String("output", "./export", "Path to export JSON files")
	help := flag.Bool("h", false, "Show help information")
	flag.BoolVar(help, "help", false, "Show help information")
	flag.Parse()

	// 如果用户请求帮助，显示帮助信息并退出
	if *help {
		showHelp()
		os.Exit(0)
	}

	// 检查数据库文件是否存在
	if _, err := os.Stat(*dbPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Database file '%s' does not exist\n\n", *dbPath)
		showHelp()
		os.Exit(1)
	}

	// 打开BoltDB文件
	db, err := bbolt.Open(*dbPath, 0600, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// 确保导出目录存在
	mkdirErr := os.MkdirAll(*exportPath, 0755)
	if mkdirErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to create export directory: %v\n", mkdirErr)
		os.Exit(1)
	}

	// 遍历所有bucket并导出数据
	err = db.View(func(tx *bbolt.Tx) error {
		// 首先获取主bucket
		return tx.ForEach(func(name []byte, b *bbolt.Bucket) error {
			bucketName := string(name)
			fmt.Printf("Processing bucket: %s\n", bucketName)

			// 创建与bucket同名的目录
			bucketDir := filepath.Join(*exportPath, bucketName)
			if mkdirErr := os.MkdirAll(bucketDir, 0755); mkdirErr != nil {
					return fmt.Errorf("failed to create bucket directory: %v", mkdirErr)
				}

			// 导出bucket中的所有键值对
			return b.ForEach(func(k, v []byte) error {
				key := string(k)
				// 处理值，假设值可能是JSON格式或普通数据
				var data any
				unmarshalErr := json.Unmarshal(v, &data)
				if unmarshalErr != nil {
					// 如果不是有效的JSON，将其作为原始字节数据处理
					data = string(v)
				}

				// 创建JSON文件
				fileName := filepath.Join(bucketDir, key+".json")
				fileData, marshalErr := json.MarshalIndent(data, "", "  ")
				if marshalErr != nil {
					return fmt.Errorf("failed to marshal data: %v", marshalErr)
				}

				// 写入文件
				if writeErr := os.WriteFile(fileName, fileData, 0644); writeErr != nil {
					return fmt.Errorf("failed to write file %s: %v", fileName, writeErr)
				}

				fmt.Printf("Exported: %s\n", fileName)
				return nil
			})
		})
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Export completed successfully!")
}
