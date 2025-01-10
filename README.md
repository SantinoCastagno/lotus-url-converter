# Lotus Url Converter

## Execute

Execute docker container
```bash
docker compose up -d
```

Test database connection with psql
```bash
psql -h localhost -U myuser -d mydatabase -W -p 5431 # then put 'mypassword'
```

```sql
SELECT * FROM translations;
REMOVE FROM translations;
```

Stop docker container execution
```bash
docker compose down --volumes
```
