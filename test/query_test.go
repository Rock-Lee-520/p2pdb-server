// Copyright 2020-2022 P2PDB, Inc.
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

package test

import (
	"context"
	"io"

	debug "github.com/favframework/debug"
	sqle "github.com/kkguan/p2pdb-server"
	"github.com/kkguan/p2pdb-store/memory"
	"github.com/kkguan/p2pdb-store/sql"
	"github.com/kkguan/p2pdb-store/sqlite"
)

func Select2() {
	debug.Dump("Example start====")

	// Create a test memory database and register it to the default engine.
	db := createTestMemory()
	e := sqle.NewDefault(sql.NewDatabaseProvider(db))

	ctx := sql.NewContext(context.Background()).WithCurrentDB("test")
	query := `select * from test.userinfo`
	//ctx.SetRawStatement(query)
	_, r, err := e.Query(ctx, query)

	debug.Dump("Example RawStatement")
	debug.Dump(ctx.RawStatement())
	// debug.Dump(r)

	checkIfError(err)

	// Iterate results and print them.
	for {
		row, err := r.Next(ctx)
		//debug.Dump(row)
		if err == io.EOF {
			break
		}
		checkIfError(err)
		debug.Dump(row)
	}
	debug.Dump("Example end====")
}

func Example() {
	debug.Dump("Example start====")

	// Create a test memory database and register it to the default engine.
	db := createTestDatabase()
	e := sqle.NewDefault(sql.NewDatabaseProvider(db))

	ctx := sql.NewContext(context.Background()).WithCurrentDB("test")
	// query := `SELECT name, count(*) FROM mytable
	// WHERE name = 'John Doe'
	// GROUP BY name`

	query := `SELECT count(*) as a,count(*) as n,name,email FROM mytable`
	//ctx.SetRawStatement(query)
	_, r, err := e.Query(ctx, query)

	// debug.Dump("Example RawStatement")
	// debug.Dump(ctx.RawStatement())
	// debug.Dump(r)

	checkIfError(err)

	// Iterate results and print them.

	debug.Dump(ctx.Query())
	debug.Dump("Example  Next start====")
	//for {
	row, err := r.Next(ctx)
	//debug.Dump(row)
	if err == io.EOF {
		debug.Dump(err)
		//break
	}
	checkIfError(err)

	debug.Dump(row)
	debug.Dump("Example  Next end====")
	// name := row[0]
	// count := row[1]
	// fmt.Println(name, count)
	//}

	debug.Dump("Example end====")

	// Output: John Doe 2
}

func checkIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func createTestDatabase() sql.Database {
	db := sqlite.NewDatabase("test")
	table := sqlite.NewTable("mytable", sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name", Type: sql.Text, Source: "mytable"},
		{Name: "email", Type: sql.Text, Source: "mytable"},
	}))
	db.AddTable("mytable", table)
	ctx := sql.NewEmptyContext()

	rows := []sql.Row{
		sql.NewRow("John Doe", "john@doe.com"),
		sql.NewRow("John Doe", "johnalt@doe.com"),
		sql.NewRow("Jane Doe", "jane@doe.com"),
		sql.NewRow("Evil Bob", "evilbob@gmail.com"),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}

func createTestMemory() sql.Database {
	db := memory.NewDatabase("test")
	table := memory.NewTable("mytable", sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "name", Type: sql.Text, Source: "mytable"},
		{Name: "email", Type: sql.Text, Source: "mytable"},
	}))
	db.AddTable("mytable", table)
	ctx := sql.NewEmptyContext()

	rows := []sql.Row{
		sql.NewRow("John Doe", "john@doe.com"),
		sql.NewRow("John Doe", "johnalt@doe.com"),
		sql.NewRow("Jane Doe", "jane@doe.com"),
		//	sql.NewRow("Evil Bob", "evilbob@gmail.com"),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}
