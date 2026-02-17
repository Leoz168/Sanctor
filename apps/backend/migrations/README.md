# Database Migrations

This directory contains database migration files.

## Structure

Migrations should be named with a timestamp prefix:
```
001_create_users_table.sql
002_create_posts_table.sql
003_add_indexes.sql
```

## Running Migrations

```bash
# TODO: Add migration commands
# For now, migrations can be run manually
```

## Creating a New Migration

1. Create a new file with the next sequential number
2. Write UP migration (apply changes)
3. Write DOWN migration (rollback changes)
4. Test both directions

## Tools

Consider using:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)
- [atlas](https://atlasgo.io/)
