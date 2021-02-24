# Details

Date : 2021-02-22 22:03:33

Directory /home/bagus2x/Documents/workspace/go/src/github.com/bagus2x/fainmi

Total : 61 files,  3544 codes, 134 comments, 714 blanks, all 4392 lines

[summary](results.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Makefile](/Makefile) | Makefile | 16 | 0 | 3 | 19 |
| [README.md](/README.md) | Markdown | 0 | 0 | 1 | 1 |
| [api/app.go](/api/app.go) | Go | 73 | 9 | 15 | 97 |
| [api/middleware/auth.go](/api/middleware/auth.go) | Go | 34 | 3 | 9 | 46 |
| [api/routes/background.go](/api/routes/background.go) | Go | 117 | 1 | 20 | 138 |
| [api/routes/button.go](/api/routes/button.go) | Go | 114 | 1 | 20 | 135 |
| [api/routes/common.go](/api/routes/common.go) | Go | 21 | 1 | 6 | 28 |
| [api/routes/font.go](/api/routes/font.go) | Go | 114 | 1 | 20 | 135 |
| [api/routes/like.go](/api/routes/like.go) | Go | 94 | 1 | 19 | 114 |
| [api/routes/link.go](/api/routes/link.go) | Go | 188 | 1 | 37 | 226 |
| [api/routes/profile.go](/api/routes/profile.go) | Go | 109 | 1 | 20 | 130 |
| [api/routes/style.go](/api/routes/style.go) | Go | 101 | 1 | 21 | 123 |
| [api/routes/test.go](/api/routes/test.go) | Go | 15 | 1 | 5 | 21 |
| [config/postgres.go](/config/postgres.go) | Go | 13 | 1 | 5 | 19 |
| [go.mod](/go.mod) | XML | 11 | 0 | 3 | 14 |
| [migrations/1_create_tables.down.sql](/migrations/1_create_tables.down.sql) | SQL | 7 | 0 | 0 | 7 |
| [migrations/1_create_tables.up.sql](/migrations/1_create_tables.up.sql) | SQL | 75 | 0 | 7 | 82 |
| [pkg/background/repository.go](/pkg/background/repository.go) | Go | 81 | 2 | 18 | 101 |
| [pkg/background/repository_test.go](/pkg/background/repository_test.go) | Go | 57 | 0 | 11 | 68 |
| [pkg/background/service.go](/pkg/background/service.go) | Go | 91 | 2 | 18 | 111 |
| [pkg/background/service_test.go](/pkg/background/service_test.go) | Go | 43 | 0 | 10 | 53 |
| [pkg/button/repository.go](/pkg/button/repository.go) | Go | 81 | 2 | 17 | 100 |
| [pkg/button/repository_test.go](/pkg/button/repository_test.go) | Go | 57 | 0 | 11 | 68 |
| [pkg/button/service.go](/pkg/button/service.go) | Go | 91 | 2 | 18 | 111 |
| [pkg/button/service_test.go](/pkg/button/service_test.go) | Go | 43 | 0 | 10 | 53 |
| [pkg/entities/background.go](/pkg/entities/background.go) | Go | 9 | 1 | 3 | 13 |
| [pkg/entities/button.go](/pkg/entities/button.go) | Go | 9 | 1 | 3 | 13 |
| [pkg/entities/font.go](/pkg/entities/font.go) | Go | 9 | 1 | 3 | 13 |
| [pkg/entities/like.go](/pkg/entities/like.go) | Go | 6 | 1 | 2 | 9 |
| [pkg/entities/link.go](/pkg/entities/link.go) | Go | 16 | 2 | 4 | 22 |
| [pkg/entities/profile.go](/pkg/entities/profile.go) | Go | 16 | 2 | 4 | 22 |
| [pkg/entities/style.go](/pkg/entities/style.go) | Go | 16 | 5 | 4 | 25 |
| [pkg/font/repository.go](/pkg/font/repository.go) | Go | 81 | 2 | 19 | 102 |
| [pkg/font/repository_test.go](/pkg/font/repository_test.go) | Go | 57 | 0 | 11 | 68 |
| [pkg/font/service.go](/pkg/font/service.go) | Go | 91 | 2 | 18 | 111 |
| [pkg/font/service_test.go](/pkg/font/service_test.go) | Go | 43 | 0 | 10 | 53 |
| [pkg/like/repository.go](/pkg/like/repository.go) | Go | 58 | 2 | 15 | 75 |
| [pkg/like/repository_test.go](/pkg/like/repository_test.go) | Go | 52 | 0 | 10 | 62 |
| [pkg/like/service.go](/pkg/like/service.go) | Go | 55 | 3 | 15 | 73 |
| [pkg/like/service_test.go](/pkg/like/service_test.go) | Go | 31 | 0 | 8 | 39 |
| [pkg/link/repository.go](/pkg/link/repository.go) | Go | 155 | 4 | 27 | 186 |
| [pkg/link/repository_test.go](/pkg/link/repository_test.go) | Go | 96 | 1 | 16 | 113 |
| [pkg/link/service.go](/pkg/link/service.go) | Go | 155 | 2 | 23 | 180 |
| [pkg/link/service_test.go](/pkg/link/service_test.go) | Go | 53 | 0 | 10 | 63 |
| [pkg/models/background.go](/pkg/models/background.go) | Go | 14 | 4 | 5 | 23 |
| [pkg/models/button.go](/pkg/models/button.go) | Go | 14 | 4 | 5 | 23 |
| [pkg/models/errors/errors.go](/pkg/models/errors/errors.go) | Go | 56 | 26 | 4 | 86 |
| [pkg/models/font.go](/pkg/models/font.go) | Go | 14 | 4 | 5 | 23 |
| [pkg/models/like.go](/pkg/models/like.go) | Go | 4 | 2 | 4 | 10 |
| [pkg/models/link.go](/pkg/models/link.go) | Go | 32 | 8 | 10 | 50 |
| [pkg/models/profile.go](/pkg/models/profile.go) | Go | 40 | 10 | 12 | 62 |
| [pkg/models/style.go](/pkg/models/style.go) | Go | 18 | 6 | 4 | 28 |
| [pkg/models/web.go](/pkg/models/web.go) | Go | 6 | 1 | 2 | 9 |
| [pkg/profile/repository.go](/pkg/profile/repository.go) | Go | 106 | 3 | 18 | 127 |
| [pkg/profile/repository_test.go](/pkg/profile/repository_test.go) | Go | 55 | 0 | 10 | 65 |
| [pkg/profile/service.go](/pkg/profile/service.go) | Go | 196 | 2 | 41 | 239 |
| [pkg/profile/service_test.go](/pkg/profile/service_test.go) | Go | 47 | 0 | 9 | 56 |
| [pkg/style/repository.go](/pkg/style/repository.go) | Go | 92 | 3 | 16 | 111 |
| [pkg/style/repository_test.go](/pkg/style/repository_test.go) | Go | 60 | 0 | 11 | 71 |
| [pkg/style/service.go](/pkg/style/service.go) | Go | 88 | 2 | 20 | 110 |
| [pkg/style/service_test.go](/pkg/style/service_test.go) | Go | 48 | 0 | 9 | 57 |

[summary](results.md)