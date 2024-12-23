package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

type TableInfo struct {
    TableName    string
    RowCount     int
    ColumnCount  int
    Columns      []ColumnInfo
}

type ColumnInfo struct {
    Name     string
    DataType string
}

func main() {
    // Connection parameters - update these with your actual credentials
    connStr := "host=localhost port=5432 user=postgres password=PostgresAtWork dbname=users_db sslmode=disable"
    
    fmt.Println("Attempting to connect to database...")
    
    // Connect to database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error opening database connection:", err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    fmt.Println("Successfully connected to database")

    // Get tables in public schema
    fmt.Println("Fetching table information...")
    tables, err := getTableInfo(db)
    if err != nil {
        log.Fatal("Error getting table info:", err)
    }

    if len(tables) == 0 {
        fmt.Println("No tables found in public schema")
        return
    }

    // Print information
    fmt.Printf("\nFound %d tables\n", len(tables))
    for _, table := range tables {
        fmt.Printf("\nTable: %s\n", table.TableName)
        fmt.Printf("Row count: %d\n", table.RowCount)
        fmt.Printf("Column count: %d\n", table.ColumnCount)
        fmt.Println("Columns:")
        for _, col := range table.Columns {
            fmt.Printf("  - %s (%s)\n", col.Name, col.DataType)
        }
        fmt.Println("------------------------")
    }
}

func getTableInfo(db *sql.DB) ([]TableInfo, error) {
    // Query to get table names
    tableQuery := `
        SELECT table_name 
        FROM information_schema.tables 
        WHERE table_schema = 'public'
        AND table_type = 'BASE TABLE'
    `
    
    rows, err := db.Query(tableQuery)
    if err != nil {
        return nil, fmt.Errorf("error querying tables: %v", err)
    }
    defer rows.Close()

    var tables []TableInfo
    
    // Process each table
    for rows.Next() {
        var table TableInfo
        err := rows.Scan(&table.TableName)
        if err != nil {
            return nil, fmt.Errorf("error scanning table name: %v", err)
        }

        fmt.Printf("Processing table: %s\n", table.TableName)

        // Get row count
        countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", table.TableName)
        err = db.QueryRow(countQuery).Scan(&table.RowCount)
        if err != nil {
            return nil, fmt.Errorf("error getting row count for %s: %v", table.TableName, err)
        }

        // Get column information
        columnQuery := `
            SELECT column_name, data_type 
            FROM information_schema.columns 
            WHERE table_schema = 'public' 
            AND table_name = $1
            ORDER BY ordinal_position
        `
        
        colRows, err := db.Query(columnQuery, table.TableName)
        if err != nil {
            return nil, fmt.Errorf("error getting columns for %s: %v", table.TableName, err)
        }

        for colRows.Next() {
            var col ColumnInfo
            err := colRows.Scan(&col.Name, &col.DataType)
            if err != nil {
                colRows.Close()
                return nil, fmt.Errorf("error scanning column info: %v", err)
            }
            table.Columns = append(table.Columns, col)
        }
        colRows.Close()
        
        table.ColumnCount = len(table.Columns)
        tables = append(tables, table)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating tables: %v", err)
    }

    return tables, nil
}