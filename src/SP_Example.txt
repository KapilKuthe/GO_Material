package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Employee represents the Employee table
type Employee struct {
	EmployeeID   int
	EmployeeName string
	DepartmentID int
}

func initDB() (*sql.DB, error) {
	// Replace with your actual database connection details
	server := "10.112.31.91"
	port := 1433
	user := "admin"
	password := "AdM!n@91"
	database := "test"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return nil, err
	}

	return db, nil
}

func getEmployeeByID(db *sql.DB, employeeID int) (Employee, error) {
	var emp Employee

	// Prepare the stored procedure call
	stmt, err := db.Prepare("EXEC SP_GetEmployee @EmployeeID = @p1")
	if err != nil {
		return emp, err
	}
	defer stmt.Close()

	// Execute the stored procedure
	err = stmt.QueryRow(employeeID).Scan(&emp.EmployeeID, &emp.EmployeeName, &emp.DepartmentID)
	if err != nil {
		return emp, err
	}

	return emp, nil
}

func main() {
	// Initialize the database connection
	db, err := initDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
		return
	}
	defer db.Close()

	// Replace with the employee ID you want to retrieve
	employeeID := 4

	// Call the stored procedure
	employee, err := getEmployeeByID(db, employeeID)
	if err != nil {
		log.Fatal("Error calling stored procedure: ", err)
		return
	}

	// Display the retrieved employee information
	fmt.Printf("Employee ID: %d\nEmployee Name: %s\nDepartment ID: %d\n",
		employee.EmployeeID, employee.EmployeeName, employee.DepartmentID)
}
