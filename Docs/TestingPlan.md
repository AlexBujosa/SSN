# Testing Plan

---

| Escene    | Resume                          | Priority | Preconditions | Input        | Steps                                  | Expected Result                  | Obtained Result |
| --------- | ------------------------------- | -------- | ------------- | ------------ | -------------------------------------- | -------------------------------- | --------------- |
| Esc-1-1-1 | Validate social security number | High     | N/A           | 001-293-1829 | Open the program and then digit de SSN | Social Security Number valid     |                 |
| Esc-1-1-2 | Validate social security number | High     | N/A           | 000-34-2817  | Open the program and then digit de SSN | Social Security Number not valid |                 |
| Esc-1-1-3 | Validate social security number | High     | N/A           | 666-48-3899  | Open the program and then digit de SSN | Social Security Number not valid |                 |
| Esc-1-1-4 | Validate social security number | High     | N/A           | 901-23-2749  | Open the program and then digit de SSN | Social Security Number not valid |                 |
| Esc-1-1-5 | Validate social security number | High     | N/A           | 567-00-4837  | Open the program and then digit de SSN | Social Security Number valid     |                 |
| Esc-1-1-6 | Validate social security number | High     | N/A           | 567-99-0000  | Open the program and then digit de SSN | Social Security Number not valid |                 |
