// Copyright 2021-2022 P2PDB, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"log"

	sqle "github.com/Rock-liyi/p2pdb-server"
	"github.com/Rock-liyi/p2pdb-server/auth"
	"github.com/Rock-liyi/p2pdb-server/server"
	"github.com/Rock-liyi/p2pdb-store/memory"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sql/information_schema"
	"github.com/Rock-liyi/p2pdb-store/sqlite"
	conf "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	environment := conf.GetEnv()
	// do something here to set environment depending on an environment variable
	// or command-line flag
	if environment == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetLevel(log.DebugLevel)
	}
}

// Example of how to implement a MySQL server based on a Engine:
//
// ```
// > mysql --host=127.0.0.1 --port=5123 -u user -ppass db -e "SELECT * FROM mytable"
// +----------+-------------------+-------------------------------+---------------------+
// | name     | email             | phone_numbers                 | created_at          |
// +----------+-------------------+-------------------------------+---------------------+
// | John Doe | john@doe.com      | ["555-555-555"]               | 2018-04-18 09:41:13 |
// | John Doe | johnalt@doe.com   | []                            | 2018-04-18 09:41:13 |
// | Jane Doe | jane@doe.com      | []                            | 2018-04-18 09:41:13 |
// | Evil Bob | evilbob@gmail.com | ["555-666-555","666-666-666"] | 2018-04-18 09:41:13 |
// +----------+-------------------+-------------------------------+---------------------+
// ```
func main() {

	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			createSqliteDatabase(), //choose createMemoryDatabase or createSqliteDatabase
			information_schema.NewInformationSchemaDatabase(),
		))

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	log.Debug("main function finish =====")
	s.Start()
}

func createSqliteDatabase() *sqlite.Database {
	const (
		dbName    = "test"
		tableName = "p2pdbtest2"
	)

	db := sqlite.NewDatabase(dbName)
	table := sqlite.NewTable(tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "id", Type: sql.Int64, Nullable: false, Source: tableName},
		// {Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		// {Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	}))
	// debug.Dump(table)
	db.AddTable(tableName, table)
	// ctx := sql.NewEmptyContext()
	// table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 1))
	// table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 2))
	// table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
	// db.DropTable(ctx, tableName)
	return db
}

func createMemoryDatabase() *memory.Database {
	const (
		dbName    = "test"
		tableName = "test2"
	)

	db := memory.NewDatabase(dbName)
	table := memory.NewTable(tableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email2", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "id", Type: sql.Int64, Nullable: false, Source: tableName},
		// {Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		// {Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	}))
	// debug.Dump(table)
	db.AddTable(tableName, table)
	ctx := sql.NewEmptyContext()
	table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 1))
	table.Insert(ctx, sql.NewRow("John Doe", "john@doe.com", 2))
	// table.Insert(ctx, sql.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()))
	// table.Insert(ctx, sql.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
	// db.DropTable(ctx, tableName)
	return db
}
