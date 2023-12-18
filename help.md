# Generating Migrations
atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://mysql/8/ent"

# Apply migration

atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://root:pass@localhost:3306/example"


# Create a schema
go run -mod=mod entgo.io/ent/cmd/ent new User