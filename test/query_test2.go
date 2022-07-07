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
	"fmt"
	"io"
	"testing"

	sqle "github.com/Rock-liyi/p2pdb-server"
	"github.com/Rock-liyi/p2pdb-store/sql"
	"github.com/Rock-liyi/p2pdb-store/sqlite"
	debug "github.com/favframework/debug"
	"github.com/stretchr/testify/require"
)

func TestDatabase_Name(t *testing.T) {
	require := require.New(t)
	db := sqlite.NewDatabase("test")
	debug.Dump(db.Name())
	require.Equal("test", db.Name())
}

func TestDatabase_DropTable(t *testing.T) {
	require := require.New(t)
	debug.Dump("Example start====")

	// Create a test memory database and register it to the default engine.
	db := createTestDatabase3()
	e := sqle.NewDefault(sql.NewDatabaseProvider(db))

	ctx := sql.NewContext(context.Background()).WithCurrentDB("test")
	query := `SELECT name, count(*) FROM mytable
	WHERE name = 'John Doe'
	GROUP BY name`
	//ctx.SetRawStatement(query)
	_, r, err := e.Query(ctx, query)

	debug.Dump("Example RawStatement")
	debug.Dump(ctx.RawStatement())
	// debug.Dump(r)

	checkIfError2(err)

	// Iterate results and print them.
	for {
		row, err := r.Next(ctx)
		//debug.Dump(row)
		if err == io.EOF {
			break
		}
		checkIfError2(err)

		name := row[0]
		count := row[1]

		fmt.Println(name, count)
	}
	debug.Dump("Example end====")
	require.True(true)
	// Output: John Doe 2
}

func checkIfError2(err error) {
	if err != nil {
		panic(err)
	}
}

func createTestDatabase3() sql.Database {
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
