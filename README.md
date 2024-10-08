# Linear Regression Analysis in Go

This project implements linear regression analysis on a dataset provided as a file. It calculates the slope and intercept of the best-fit line and computes the Pearson correlation coefficient for the data.

## Features

- Calculate the slope and intercept of a linear regression line.
- Compute the Pearson correlation coefficient to measure the strength of the relationship between two variables.
- Handle data input from a specified file.

## Requirements

- Go programming language
- A data file containing numerical values.

## Installation

1. Clone this repository:
```bash
git clone https://learn.zone01kisumu.ke/git/doonyango/linear-stats.git
cd linear-stats
```

## Usage
To run the program, use the following command in your terminal:

```bash
go run main.go data.txt
```
## Data Format
The program expects a data file containing numerical values, one value per line. For example:
```
1.0
2.5
3.0
4.5
5.0
```

## How It Works

1. **Argument Parsing**: 

    The program checks for command-line arguments to get the path of the data file.

2. **Data Reading**: 

   It uses a function from the linears/handleData package to read data from the specified file.

3. **Calculations**:

   - It calculates sums of x-values, y-values, x-squared values, and xy products.
   - Computes slope and intercept for the linear regression line using least squares method.
   - Calculates means for x and y values.
   - Computes Pearson correlation coefficient to assess linear relationship strength.

   For example:

   ```js
   // Linear Regression 
   slope = (n*xySum - xSum*ySum) / (n*xSquaredSum - xSum*xSum)
	intercept = (ySum - slope*xSum) / n
   ```

   ```js
   // Pearson Correlation Coefficient
   correlation =  numerator / math.Sqrt(xDenominator*yDenominator)
   ```

4. **Output**: 

   Finally, it prints out the linear regression equation and Pearson correlation coefficient.

## Contribution
This project is open for contibution,  create a pull request if need arises.

## Author
This project is maintained by:
[Doreen Onyango](https://github.com/Doreen-Onyango)
