# Self learning of Go
This repo is for personnal use in order to learn go by building a demo API.


### Personal notes about used packages and tools used

- goose for migrations managment

helpfull commands : `goose -s create create_products sql`

- sqlc for Models generation (Generates automaticly go code from SQL files)

No direct edit on SQLC code in internal/adapters/sqlc code

- Other API related go packages such as Chi...