package helpers

import (
	"github.com/Shreyas-Prabhu/EmployeeDatabase/config"
	"github.com/Shreyas-Prabhu/EmployeeDatabase/models"
	log "github.com/sirupsen/logrus"
)

var DB = config.MysqlConnect()

func init() {
	log.Info("Automigrating DB...")
	DB.AutoMigrate(&models.Employee{})
}

func InsertEmployee(emp models.Employee) error {
	log.Info("Inserting the employee detail having id:", emp.ID)
	err := DB.Create(&emp).Error
	if err != nil {
		log.Error("Error inserting data for ID:", emp.ID, err)
		return err
	}
	return nil

}

func GetEmployee(emp *models.Employee, id int) {
	log.Info("Getting the employee detail having id:", id)
	_ = DB.First(emp, id)
}

func UpdateEmployee(emp *models.Employee, updateEmp models.Employee) error {
	log.Info("Updating the employee with ID:", emp.ID)
	err := DB.Model(emp).Updates(models.Employee{Name: updateEmp.Name, Position: updateEmp.Position, Salary: updateEmp.Salary}).Error
	if err != nil {
		log.Error("Error Updating data for ID:", emp.ID, err)
		return err
	}
	return nil
}

func DeleteEmployee(emp *models.Employee, id int) error {
	log.Info("Deleting the employee with ID:", id)
	err := DB.Delete(emp, id).Error
	if err != nil {
		log.Error("Error deleting the data for ID:", emp.ID, err)
		return err
	}
	return nil
}

func GetEmployeePagination(empArr *[]models.Employee, offset int, size int) error {
	log.Info("Get list of all employees using pagination method")
	err := DB.Offset(offset).Limit(size).Find(empArr).Error
	if err != nil {
		log.Error("Error getting data the data for ID:", err)
		return err
	}
	return nil
}

func GetEmployeeCount() (int64, error) {
	var count int64
	err := DB.Model(&models.Employee{}).Count(&count).Error
	if err != nil {
		log.Error("Error getting  the count", err)
		return 0, err
	}
	return count, nil
}
