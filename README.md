# Bolt2JSON

## 项目简介 / Project Introduction

一个简单的命令行工具，用于将 BoltDB 数据库中的内容导出为 JSON 文件，按照 Bucket 结构创建相应的文件夹。

A simple command-line tool to export contents from BoltDB database to JSON files, creating directories according to the Bucket structure.

## 功能特性 / Features

- 读取 BoltDB 数据库文件 / Read BoltDB database files
- 按照 Bucket 结构创建相应的文件夹 / Create directories according to Bucket structure
- 将每个键值对导出为单独的 JSON 文件 / Export each key-value pair as a separate JSON file
- 自动处理嵌套的 JSON 数据 / Automatically handle nested JSON data
- 支持自定义数据库文件路径和导出路径 / Support custom database file path and export path

## 安装 / Installation

```bash
# 克隆仓库 / Clone the repository
git clone <repository-url>
cd bolt2json

# 构建项目 / Build the project
go build -o bolt2json
```

## 使用方法 / Usage

```bash
# 使用默认参数（读取当前目录的 database.boltdb，导出到 ./export 目录）
# Use default parameters (read database.boltdb in current directory, export to ./export directory)
./bolt2json

# 指定数据库文件和导出路径 / Specify database file and export path
./bolt2json -db /path/to/your/database.boltdb -output /path/to/export/directory
```

## 命令行参数 / Command-line Arguments

- `-db`：BoltDB 数据库文件的路径（默认：database.boltdb） / BoltDB database file path (default: database.boltdb)
- `-output`：JSON 文件的导出目录（默认：./export） / Export directory for JSON files (default: ./export)
- `-h`, `--help`：显示帮助信息 / Show help information

## 输出结构 / Output Structure

导出后的目录结构将如下所示：

The exported directory structure will look like this:

```
export/
├── bucket1/
│   ├── key1.json
│   ├── key2.json
│   └── ...
├── bucket2/
│   ├── key3.json
│   ├── key4.json
│   └── ...
└── ...
```

每个 JSON 文件包含对应键的值，如果值本身是 JSON 格式，则会保持其结构；否则，将作为字符串保存。

Each JSON file contains the value of the corresponding key. If the value is in JSON format, its structure will be preserved; otherwise, it will be saved as a string.

## 依赖项 / Dependencies

- [go.etcd.io/bbolt](https://github.com/etcd-io/bbolt) - BoltDB 客户端库 / BoltDB client library

## 许可证 / License

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details