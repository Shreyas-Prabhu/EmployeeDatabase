package controller

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/Shreyas-Prabhu/EmployeeDatabase/helpers"
	"github.com/Shreyas-Prabhu/EmployeeDatabase/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

var mu sync.Mutex

func InsertController(c *gin.Context) {
	var employee models.Employee
	var checkEmployee models.Employee
	var message string
	var statusCode int
	err := c.BindJSON(&employee)
	if err != nil {
		log.Error(err)
		return
	}
	err = validator.New().Struct(employee)
	if err != nil {
		log.Errorf("Validation failed due to %v\n", err)
		statusCode = 422
		message = "Some fields are empty"
		c.JSON(statusCode, gin.H{"message": message})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	helpers.GetEmployee(&checkEmployee, employee.ID)
	if !reflect.ValueOf(checkEmployee).IsZero() {
		statusCode = 409
		message = fmt.Sprint("Employee already present with ID:", employee.ID)
	} else {
		statusCode = 201
		message = "Employee inserted"
		err := helpers.InsertEmployee(employee)
		if err != nil {
			statusCode = 500
			message = "Unable to insert employee record at the moment(Check Logs)"
			c.JSON(statusCode, gin.H{
				"message": message,
			})
			return
		}
	}

	c.JSON(statusCode, gin.H{
		"message": message,
	})
}

func GetController(c *gin.Context) {
	var employee models.Employee
	var message string
	var statusCode int
	var data interface{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	helpers.GetEmployee(&employee, id)
	if reflect.ValueOf(employee).IsZero() {
		statusCode = 404
		message = fmt.Sprint("No Employee record having ID:", id)
		data = nil
	} else {
		statusCode = 200
		message = "Employee Record"
		data = employee
	}
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func UpdateController(c *gin.Context) {
	var message string
	var statusCode int
	var data interface{}
	var employee models.Employee
	var updateEmployee models.Employee
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return
	}

	c.BindJSON(&updateEmployee)
	mu.Lock()
	defer mu.Unlock()
	helpers.GetEmployee(&employee, id)
	if reflect.ValueOf(employee).IsZero() {
		statusCode = 404
		message = fmt.Sprint("Employee does not exist with ID:", id)
		data = nil
	} else {
		err := helpers.UpdateEmployee(&employee, updateEmployee)
		if err != nil {
			statusCode = 500
			message = "Unable to update employee record at the moment(Check Logs)"
			c.JSON(statusCode, gin.H{
				"message": message,
			})
			return
		}
		statusCode = 200
		message = fmt.Sprint("Employee record updated with ID:", employee.ID)
		data = employee
	}

	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func DeleteController(c *gin.Context) {
	var employee models.Employee
	var checkEmployee models.Employee
	var message string
	var statusCode int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	helpers.GetEmployee(&checkEmployee, id)
	if reflect.ValueOf(checkEmployee).IsZero() {
		statusCode = 404
		message = fmt.Sprint("Employee does not exist with ID:", id)
	} else {
		err := helpers.DeleteEmployee(&employee, id)
		if err != nil {
			statusCode = 500
			message = "Unable to delete employee record at the moment(Check Logs)"
			c.JSON(statusCode, gin.H{
				"message": message,
			})
			return
		}
		statusCode = 200
		message = fmt.Sprint("Employee deleted with ID:", id)
	}

	c.JSON(statusCode, gin.H{
		"message": message,
	})
}

func PaginationController(c *gin.Context) {
	var message string
	var statusCode int
	var empArr []models.Employee
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "5"))
	if err != nil {
		size = 5
	}
	offset := (page - 1) * size

	err = helpers.GetEmployeePagination(&empArr, offset, size)
	if err != nil {
		statusCode = 500
		message = "Unable to get employee list at the moment(Check Logs)"
		c.JSON(statusCode, gin.H{
			"message": message,
		})
		return
	}
	totalRecords, err := helpers.GetEmployeeCount()
	if err != nil {
		statusCode = 500
		message = "Unable to get count of employees at the moment(Check Logs)"
		c.JSON(statusCode, gin.H{
			"message": message,
		})
		return
	}
	statusCode = 200
	message = "Paginated data"
	c.JSON(statusCode, gin.H{
		"message":      message,
		"page":         page,
		"size":         size,
		"totalRecords": totalRecords,
		"totalPages":   (int(totalRecords) + size - 1) / size,
		"data":         empArr,
	})
}
