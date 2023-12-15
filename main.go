package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/example/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Importing the PostgreSQL dialect here
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

// Defining Database Models

// Defining Database models Ends

var db *gorm.DB

func initDB() {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Log SQL queries in development mode
	db.LogMode(true)

	// Auto Migrate the Customer model
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Dish{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.OrderItem{})
	db.AutoMigrate(&models.Bill{})
}

func main() {
	// Initialise gofr object
	app := gofr.New()

	// Initialize the database connection
	initDB()
	defer db.Close()

	//<--------------------------------------- GET REQUESTS // ------------------------------------------>

	app.GET("/customers", func(ctx *gofr.Context) (interface{}, error) {
		var customers []models.Customer

		// Getting the customers from the database using GORM
		err := db.Find(&customers).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return customers, nil
	})

	app.GET("/dishes", func(ctx *gofr.Context) (interface{}, error) {
		var dishes []models.Dish

		// Preload the 'Category' relationship when retrieving dishes
		err := db.Preload("Category").Find(&dishes).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return dishes, nil
	})

	app.GET("/categories", func(ctx *gofr.Context) (interface{}, error) {

		var categories []models.Category

		err := db.Find(&categories).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return categories, nil
	})

	// Assuming you have a route for retrieving all orders

	app.GET("/bills", func(ctx *gofr.Context) (interface{}, error) {
		var bills []models.Bill

		// Retrieve all orders from the database
		err := db.Preload("Customer").Find(&bills).Error
		if err != nil {
			return nil, err
		}

		return bills, nil
	})

	//<--------------------------------------- GET Requests End ----------------------------------------------->

	//<--------------------------------------- POST Requests End ----------------------------------------------->

	app.POST("/customers", func(ctx *gofr.Context) (interface{}, error) {
		// Decode JSON from the request body
		var newCustomer models.Customer
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&newCustomer); err != nil {
			return nil, err
		}

		// Create a new customer in the database
		err := db.Create(&newCustomer).Error
		if err != nil {
			return nil, err
		}

		return newCustomer, nil
	})

	app.POST("/dishes", func(ctx *gofr.Context) (interface{}, error) {
		var newDish models.Dish
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&newDish); err != nil {
			return nil, err
		}

		err := db.Create(&newDish).Error
		if err != nil {
			return nil, err
		}

		return newDish, nil
	})

	app.POST("/categories", func(ctx *gofr.Context) (interface{}, error) {
		var newCategory models.Category
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&newCategory); err != nil {
			return nil, err
		}

		err := db.Create(&newCategory).Error
		if err != nil {
			return nil, err
		}

		return newCategory, nil
	})

	app.POST("/bills", func(ctx *gofr.Context) (interface{}, error) {
		var newBill models.Bill
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&newBill); err != nil {
			return nil, err
		}

		err := db.Create(&newBill).Error
		if err != nil {
			return nil, err
		}

		return newBill, nil
	})
	//<--------------------------------------- POST Requests End ----------------------------------------------->

	//<--------------------------------------- PUT Requests End ----------------------------------------------->

	// PUT request to update a customer
	app.PUT("/customers/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract customer ID from the URL parameters
		customerID := ctx.PathParam("id")

		// Decode JSON from the request body
		var updatedCustomer models.Customer
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&updatedCustomer); err != nil {
			return nil, err
		}

		// Fetch the existing customer from the database
		var existingCustomer models.Customer
		err := db.First(&existingCustomer, "id = ?", customerID).Error
		if err != nil {
			return nil, err
		}

		// Update the existing customer with the new data
		existingCustomer.Name = updatedCustomer.Name
		existingCustomer.Age = updatedCustomer.Age
		existingCustomer.Phone = updatedCustomer.Phone
		existingCustomer.Spending = updatedCustomer.Spending

		// Save the updated customer to the database
		err = db.Save(&existingCustomer).Error
		if err != nil {
			return nil, err
		}

		// Return the updated customer data
		return existingCustomer, nil
	})

	app.PUT("/categories/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract customer ID from the URL parameters
		categoryID := ctx.PathParam("id")

		// Decode JSON from the request body
		var updatedCategory models.Category
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&updatedCategory); err != nil {
			return nil, err
		}

		// Fetch the existing customer from the database
		var existingCategory models.Category
		err := db.First(&existingCategory, "id = ?", categoryID).Error
		if err != nil {
			return nil, err
		}

		// Update the existing customer with the new data
		existingCategory.Name = updatedCategory.Name

		// Save the updated customer to the database
		err = db.Save(&existingCategory).Error
		if err != nil {
			return nil, err
		}

		// Return the updated customer data
		return existingCategory, nil
	})

	// PUT request to update a dish
	app.PUT("/dishes/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract dish ID from the URL parameters
		dishID := ctx.PathParam("id")

		// Decode JSON from the request body
		var updatedDish models.Dish
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&updatedDish); err != nil {
			return nil, err
		}

		// Fetch the existing dish from the database
		var existingDish models.Dish
		err := db.First(&existingDish, "id = ?", dishID).Error
		if err != nil {
			return nil, err
		}

		// Update the existing dish with the new data
		existingDish.Name = updatedDish.Name
		existingDish.Price = updatedDish.Price
		existingDish.CategoryID = updatedDish.CategoryID

		// Save the updated dish to the database
		err = db.Save(&existingDish).Error
		if err != nil {
			return nil, err
		}

		// Return the updated dish data
		return existingDish, nil
	})

	// PUT request to update a bill
	app.PUT("/bills/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract bill ID from the URL parameters
		billID := ctx.PathParam("id")

		// Decode JSON from the request body
		var updatedBill models.Bill
		decoder := json.NewDecoder(ctx.Request().Body)
		if err := decoder.Decode(&updatedBill); err != nil {
			return nil, err
		}

		// Fetch the existing bill from the database
		var existingBill models.Bill
		err := db.First(&existingBill, "id = ?", billID).Error
		if err != nil {
			return nil, err
		}

		// Update the existing bill with the new data
		existingBill.CustomerID = updatedBill.CustomerID
		existingBill.Amount = updatedBill.Amount

		// Save the updated bill to the database
		err = db.Save(&existingBill).Error
		if err != nil {
			return nil, err
		}

		// Return the updated bill data
		return existingBill, nil
	})

	// PUT request to update an order

	// <----------------------------------- DELETE REQUESTS ----------------------------------->

	// DELETE request to delete a customer
	app.DELETE("/customers/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract customer ID from the URL parameters
		customerID := ctx.PathParam("id")

		// Validate the customer ID
		id, err := strconv.Atoi(customerID)
		if err != nil {
			return nil, errors.InvalidParam{Param: []string{"id"}}
		}

		// Fetch the existing customer from the database
		var existingCustomer models.Customer
		err = db.First(&existingCustomer, id).Error
		if err != nil {
			return nil, errors.EntityNotFound{
				Entity: "customer",
				ID:     strconv.Itoa(id),
			}
		}

		// Delete the customer from the database
		err = db.Delete(&existingCustomer).Error
		if err != nil {
			return nil, err
		}

		// Return a success message or any relevant response
		return map[string]string{"message": "Customer deleted successfully"}, nil
	})

	// DELETE request to delete a category
	app.DELETE("/categories/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract category ID from the URL parameters
		categoryID := ctx.PathParam("id")

		// Validate the category ID
		id, err := strconv.Atoi(categoryID)
		if err != nil {
			return nil, errors.InvalidParam{Param: []string{"id"}}
		}

		// Fetch the existing category from the database
		var existingCategory models.Category
		err = db.First(&existingCategory, id).Error
		if err != nil {
			return nil, errors.EntityNotFound{
				Entity: "category",
				ID:     strconv.Itoa(id),
			}
		}

		// Delete the category from the database
		err = db.Delete(&existingCategory).Error
		if err != nil {
			return nil, err
		}

		// Return a success message or any relevant response
		return map[string]string{"message": "Category deleted successfully"}, nil
	})

	// DELETE request to delete a dish
	app.DELETE("/dishes/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract dish ID from the URL parameters
		dishID := ctx.PathParam("id")

		// Validate the dish ID
		id, err := strconv.Atoi(dishID)
		if err != nil {
			return nil, errors.InvalidParam{Param: []string{"id"}}
		}

		// Fetch the existing dish from the database
		var existingDish models.Dish
		err = db.First(&existingDish, id).Error
		if err != nil {
			return nil, errors.EntityNotFound{
				Entity: "dish",
				ID:     strconv.Itoa(id),
			}
		}

		// Delete the dish from the database
		err = db.Delete(&existingDish).Error
		if err != nil {
			return nil, err
		}

		// Return a success message or any relevant response
		return map[string]string{"message": "Dish deleted successfully"}, nil
	})

	// DELETE request to delete a bill
	app.DELETE("/bills/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract bill ID from the URL parameters
		billID := ctx.PathParam("id")

		// Validate the bill ID
		id, err := strconv.Atoi(billID)
		if err != nil {
			return nil, errors.InvalidParam{Param: []string{"id"}}
		}

		// Fetch the existing bill from the database
		var existingBill models.Bill
		err = db.First(&existingBill, id).Error
		if err != nil {
			return nil, errors.EntityNotFound{
				Entity: "bill",
				ID:     strconv.Itoa(id),
			}
		}

		// Delete the bill from the database
		err = db.Delete(&existingBill).Error
		if err != nil {
			return nil, err
		}

		// Return a success message or any relevant response
		return map[string]string{"message": "Bill deleted successfully"}, nil
	})

	// Starts the server, it will listen on the default port 8000.
	// It can be overridden through configs
	app.Start()
}

//<--------------------------------------- PUT Requests End ----------------------------------------------->
