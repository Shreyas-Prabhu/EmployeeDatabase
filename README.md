# EmployeeDatabase
CRUD operation using MySQL, GO, GIN, GORM

DESCRIPTION:
This application contains 5 endpoints as described below:
    1)POST /employee/insert - This endpoint is used to insert the employee record in the database
    To use:
    -Hit the endpoint - POST http://localhost:8000/employee/insert
    -pass data in json format, example,
        {
            "id":118,
            "name":"Newton Isaac",
            "position":"Golang developer III",
            "salary":1400000
        }
    
    2)GET /employee/get/:id - This endpoint is used to get the employee detail whose id is passed as path variabale
    To use:
    -Hit the endpoint - GET http://localhost:8000/employee/get/id -- Replace id with the employee id that you want to retrieve

    3)PUT /employee/update/:id -  This endpoint is used to update the employee details whose id is passed as path variable
    To use:
    -Hit the endpoint - PUT http://localhost:8000/employee/update/id -- Replace id with the employee id that you want to update
    -Pass the update details, example,
        {
            "name":"Newton Isaac",
            "position":"Senior Golang developer",
            "salary":1500000
        }

    4)DELETE /employee/delete/:id - This endpoint is used to delete the employee detail whose id is passed as path variable
    To use:
    -Hit the endpoint - GET http://localhost:8000/employee/delete/id -- Replace id with the employee id that you want to delete

    5)GET /employee/getList - This endpoint is used to retrieve the employee list from database by limiting the number of records per page to the provided size.
    -Hit the endpoint - GET http://localhost:8000/employee/getList?page=1&size=2
    The above means response is sent where the first page record is displayed containing 2 records.
    Default size of records is 5 and if the page number is not provided then the defult page is set to be 1.


CONFIGURATION:
1)The application used MySQL database. The connection string is present in .env file. Replace the database connection string with yours when running the application.
2)Port used is 8000 and is present in .env file.

UNIT TEST:
The unit test is present in file employee_test.go.
To run the unit test, use command: 
    go test -v
