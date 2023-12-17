package db

import (
	"context"
	"crypto/tls"
	"log"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-sql-driver/mysql"
	"github.com/ramnath.1998/vyaaparjunction/ent"
	"github.com/ramnath.1998/vyaaparjunction/ent/migrate"
)

var (
	DbClient = ent.Client{}
)

func DBConnect() {

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.ap-southeast-1.prod.aws.tidbcloud.com",
	})
	client, err := ent.Open("mysql", "455GNSfHbboubjm.root:1fdYPVLgkhC7aZT1@tcp(gateway01.ap-southeast-1.prod.aws.tidbcloud.com:4000)/VyaaparJunctionDev?tls=tidb")
	if err != nil {
		log.Fatalf("failed opening connection to tidb: %v", err)
	}

	if err := client.Schema.Create(context.Background(),
		schema.WithAtlas(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	DbClient = *client
}
