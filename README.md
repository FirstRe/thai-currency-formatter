# Thai Currency Formatter

This program takes a decimal value and converts it into a Thai currency format.

## Requirements
- Decimal library requires Go version >=1.10

## Functionality
- **Input**: A decimal value
- **Output**: The amount in Thai text format:
  - If the value has no fractional part, it will append "ถ้วน".
  - If the value has a fractional part, the result will be followed by "สตางค์" for the fractional part.

## Examples

1. **Input**: `1234`
   **Output**: `"หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"`

2. **Input**: `33333.75`
   **Output**: `"สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"`

## Code Structure
1. **`textThai` Map**: A dictionary maps numbers to their Thai text.
2. **`numberToThai(num int64) string`**: Convert an integer into Thai text.
3. **`formatAmount(input decimal.Decimal) string`**: Main function that processes the decimal value to format both the integer parts and fractional parts.

## Usage

1. Install

    ```bash
    go mod tidy
    ```

2. Run the program:

    ```bash
    go run main.go
    ```
    
3. Run test program:

    ```bash
    go test ./testing
    ```

    The program will output the Thai currency formatted string.
