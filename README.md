## To Do List

- [ ] Search Data
- [ ] Auth JWT
- [ ] Middleware
- [ ] Dockerisation
- [ ] Darkmode UI

## Learn Crud Golang


#golang-migration

firstime you need instal package golang-migrate

```bash
  brew install golang-migrate
```

#create migration

- original :

```bash
  migrate create -ext sql  -dir <path> -seq <migration_name>

```
example :

```bash
  migrate create -ext sql  -dir config/migrations -seq tags
```

- modif on your mac

1. nano ~/.zshrc
2. add alias 
```bash
  alias go-migrate="migrate create -ext sql"
```
3. and then create migration like this :
```bash
  go-migrate -dir <path> -seq <migration_name>
```
example :
```bash
  go-migrate -dir config/migrations -seq tags
```

# how to UP

```bash
  migrate -path <path> -database ‘<connection_string>’ -verbose up
```

example :
```bash
  migrate -path config/migrations/ -database 'mysql://user:password@/database_name' -verbose up
```

# how to DOWN

```bash
  migrate -path <path> -database ‘<connection_string>’ -verbose down
```

example :
```bash
  migrate -path config/migrations/ -database 'mysql://user:password@/database_name’ -verbose down
```
