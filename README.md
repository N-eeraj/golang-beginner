# Go Bill - A Simple Billing System

Go Bill is a simple command-line billing system built using Go. The application allows users to create, view, update, and remove items from a bill. It also includes functionality to save the bill in text or CSV format.

## Features

- **Create New Bill**: Start a new bill with a unique invoice number.
- **Add Items**: Add available dishes to the bill with quantity.
- **Update Item Quantity**: Update the quantity of an item already added to the bill.
- **Remove Item**: Remove an item from the bill.
- **Save Bill**: Save the bill as a text file or CSV file in the "bills" directory.
- **View Saved Bills**: View all saved bills and their details.
- **Print Bill**: Print the bill in selected format (CSV or Text).

## Requirements

- Go 1.18 or later

## Setup

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Open a terminal in the project directory.
4. Run the following command to start the application:

```bash
go run main.go
```

## How to Use

1. **Start a New Bill**: 
   - Select the `New Bill` option from the main menu.
   - Add items by selecting from the dish menu and specifying the quantity.
   - Update item quantity or remove items as needed using the bill menu.
   - When done, you can save the bill as a `.txt` or `.csv` file.

2. **Show Saved Bills**: 
   - Select the `Show Bills` option from the main menu to view all previously saved bills and their details.
   - Choose a bill to view.

3. **Print Bill**:
   - After selecting a saved bill, you can print it to a file in your desired format (`.txt` or `.csv`).
   - The bill will be saved in the `bills/` directory as a file with the selected format.

4. **Exit**:
   - Select the `Exit` option from the main menu to close the application.

## File Storage

- **Directory**: Bills are stored in a folder named `bills/` in the project's root directory.
- **Saved Bills**: Once a bill is saved, it will be stored in the `bills/` folder with the following naming convention:
  - `GBI00001.txt` for a text file
  - `GBI00001.csv` for a CSV file
- **File Formats**: 
  - **Text File (.txt)**: This format presents the bill in a formatted tabular view.
  - **CSV File (.csv)**: This format stores the bill in a CSV format for easier data manipulation.
- **File Permissions**: The files will be created with the `0644` file permission, meaning they are readable and writable by the owner, and only readable by others.
- **Directory Creation**: If the `bills/` directory doesn't exist, the program will automatically create it when you save a bill.

## Example Output

**Text Format**:
```txt
------------------------------------------------------------------
| Bill No: GBI00001                                              |
------------------------------------------------------------------
| Sl.| Dish                     |   Price | Quantity |     Total |
------------------------------------------------------------------
|  1 | Cheeseburger             | $  7.99 |        2 | $   15.98 |
|  2 | Fried Chicken            | $  9.49 |        1 | $    9.49 |
------------------------------------------------------------------
| Grand Total                                        | $   25.47 |
------------------------------------------------------------------
```

**CSV Format**:
```csv
Bill No:,GBI00001,,,
Sl.,Dish,Price,Quantity,Total
1,Cheeseburger,$7.99,2,$15.98
2,Fried Chicken,$9.49,1,$9.49
,,,Grand Total,$25.47
```
<br />

> The other files in this repo are a part of my Go learning journey, with each one covering a few of the fundamental concepts of the language.
