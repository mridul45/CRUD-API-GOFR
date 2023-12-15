# CRUD-API-GOFR

__OrderItem__
1. A customer can create an order , edit it , view it and detele it.
2. The table for it contains dishID connected via ForeignKey to Dish Table and Quantity

__Customer__
1. It contains name , age , phone and spending of the customer.

__Category__
1. Conatins the name. Like which category the dish belongs . eg(Chinese, South Indian)

__Dish__

1. Contains name,price,categoryId.
2. Connect to category tble via foreign key.

__Bill__

1. Contains CustomerId and and amount.
2. Connected to the customer table via foreignkey.
