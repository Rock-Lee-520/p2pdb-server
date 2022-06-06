// Copyright 2020-2021 Dolthub, Inc.
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

package enginetest

import (
	"testing"

	debug "github.com/favframework/debug"
	"github.com/kkguan/p2pdb-store/sql"
	"github.com/kkguan/p2pdb-store/sql/expression"
	"github.com/kkguan/p2pdb-store/sqlite"
)

// This file is for validating both the engine itself and the in-sqlite database implementation in the sqlite package.
// Any engine test that relies on the correct implementation of the in-sqlite database belongs here. All test logic and
// queries are declared in the exported enginetest package to make them usable by integrators, to validate the engine
// against their own implementation.

var numPartitionsVals = []int{
	1,
	testNumPartitions,
}

var parallelVals = []int{
	1,
	2,
}

func TestReadOnlyDatabasess(t *testing.T) {
	debug.Dump("TestReadOnlyDatabasess start====")
	var dbname = "test"
	connection, err := NewSQLITEHarness(dbname + "db")
	if err != nil {

		debug.Dump(err)
	}

	//debug.Dump(connection)
	ctx := connection.NewContext()
	// if connection.shim.HasDatabase(dbname){
	// 	connection.shim.DropDatabase(ctx,dbname)
	// }

	err = connection.shim.CreateDatabase(ctx, dbname+"db")
	if err != nil {
		debug.Dump("into CreateDatabase err")
		debug.Dump(err.Error())
	}
	_, err = connection.shim.Query("", "select * from mytable3")
	if err != nil {
		debug.Dump("into Query err")
		debug.Dump(err.Error())
	}

	err = connection.shim.Exec("", "CREATE TABLE   IF NOT EXISTS `mytable3`  (`name` text NOT NULL,`email` text NOT NULL,`phone_numbers` json NOT NULL,`created_at` timestamp NOT NULL)")
	if err != nil {
		debug.Dump("into Exec err")
		debug.Dump(err.Error())
	}

	err = connection.shim.Exec("", "INSERT INTO mytable3(name, email, phone_numbers, created_at) VALUES('Evil Bob', 'evilbob@gmail.com', 123, '2022-01-02 12:28:26.024723000');")
	if err != nil {
		debug.Dump("into Exec err")
		debug.Dump(err.Error())
	}

	rows, err := connection.shim.QueryRows("", "select * from mytable3")
	if err != nil {
		debug.Dump("into Query err")
		debug.Dump(err.Error())
	}

	debug.Dump(rows)
	// ctx.SetCurrentDatabase("test")
	// db, err := connection.shim.Database("test")
	// query := connection.shim.Query(db, "select * from mytable")
	// // name, err := db.GetTableNames(connection.NewContext())

	//	debug.Dump(db.Name())
	//debug.Dump(db)
	// sql, err := connection.shim.Query("test", "select * from mytable;")

	// if err != nil {
	// 	debug.Dump(err)
	// }
	// debug.Dump(sql)
	debug.Dump("TestReadOnlyDatabasess end====")

}

func mergableIndexDriver(dbs []sql.Database) sql.IndexDriver {
	return sqlite.NewIndexDriver("mydb", map[string][]sql.DriverIndex{
		"mytable": {
			newMergableIndex(dbs, "mytable",
				expression.NewGetFieldWithTable(0, sql.Int64, "mytable", "i", false)),
			newMergableIndex(dbs, "mytable",
				expression.NewGetFieldWithTable(1, sql.Text, "mytable", "s", false)),
			newMergableIndex(dbs, "mytable",
				expression.NewGetFieldWithTable(0, sql.Int64, "mytable", "i", false),
				expression.NewGetFieldWithTable(1, sql.Text, "mytable", "s", false)),
		},
		"othertable": {
			newMergableIndex(dbs, "othertable",
				expression.NewGetFieldWithTable(0, sql.Text, "othertable", "s2", false)),
			newMergableIndex(dbs, "othertable",
				expression.NewGetFieldWithTable(1, sql.Text, "othertable", "i2", false)),
			newMergableIndex(dbs, "othertable",
				expression.NewGetFieldWithTable(0, sql.Text, "othertable", "s2", false),
				expression.NewGetFieldWithTable(1, sql.Text, "othertable", "i2", false)),
		},
		"bigtable": {
			newMergableIndex(dbs, "bigtable",
				expression.NewGetFieldWithTable(0, sql.Text, "bigtable", "t", false)),
		},
		"floattable": {
			newMergableIndex(dbs, "floattable",
				expression.NewGetFieldWithTable(2, sql.Text, "floattable", "f64", false)),
		},
		"niltable": {
			newMergableIndex(dbs, "niltable",
				expression.NewGetFieldWithTable(0, sql.Int64, "niltable", "i", false)),
			newMergableIndex(dbs, "niltable",
				expression.NewGetFieldWithTable(1, sql.Int64, "niltable", "i2", true)),
		},
		"one_pk": {
			newMergableIndex(dbs, "one_pk",
				expression.NewGetFieldWithTable(0, sql.Int8, "one_pk", "pk", false)),
		},
		"two_pk": {
			newMergableIndex(dbs, "two_pk",
				expression.NewGetFieldWithTable(0, sql.Int8, "two_pk", "pk1", false),
				expression.NewGetFieldWithTable(1, sql.Int8, "two_pk", "pk2", false),
			),
		},
	})
}

func newMergableIndex(dbs []sql.Database, tableName string, exprs ...sql.Expression) *sqlite.Index {
	db, table := findTable(dbs, tableName)
	if db == nil {
		return nil
	}
	return &sqlite.Index{
		DB:         db.Name(),
		DriverName: sqlite.IndexDriverId,
		TableName:  tableName,
		Tbl:        table.(*sqlite.Table),
		Exprs:      exprs,
	}
}

func findTable(dbs []sql.Database, tableName string) (sql.Database, sql.Table) {
	for _, db := range dbs {
		names, err := db.GetTableNames(sql.NewEmptyContext())
		if err != nil {
			panic(err)
		}
		for _, name := range names {
			if name == tableName {
				table, _, _ := db.GetTableInsensitive(sql.NewEmptyContext(), name)
				return db, table
			}
		}
	}
	return nil, nil
}
