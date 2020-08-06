// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// GalaxiesColumns holds the columns for the "galaxies" table.
	GalaxiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"barred_spiral", "elliptical", "irregular", "spiral"}},
	}
	// GalaxiesTable holds the schema information for the "galaxies" table.
	GalaxiesTable = &schema.Table{
		Name:        "galaxies",
		Columns:     GalaxiesColumns,
		PrimaryKey:  []*schema.Column{GalaxiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PlanetsColumns holds the columns for the "planets" table.
	PlanetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeUint, Nullable: true},
		{Name: "galaxy_planets", Type: field.TypeInt, Nullable: true},
	}
	// PlanetsTable holds the schema information for the "planets" table.
	PlanetsTable = &schema.Table{
		Name:       "planets",
		Columns:    PlanetsColumns,
		PrimaryKey: []*schema.Column{PlanetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "planets_galaxies_planets",
				Columns: []*schema.Column{PlanetsColumns[3]},

				RefColumns: []*schema.Column{GalaxiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PlanetNeighborsColumns holds the columns for the "planet_neighbors" table.
	PlanetNeighborsColumns = []*schema.Column{
		{Name: "planet_id", Type: field.TypeInt},
		{Name: "neighbor_id", Type: field.TypeInt},
	}
	// PlanetNeighborsTable holds the schema information for the "planet_neighbors" table.
	PlanetNeighborsTable = &schema.Table{
		Name:       "planet_neighbors",
		Columns:    PlanetNeighborsColumns,
		PrimaryKey: []*schema.Column{PlanetNeighborsColumns[0], PlanetNeighborsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "planet_neighbors_planet_id",
				Columns: []*schema.Column{PlanetNeighborsColumns[0]},

				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "planet_neighbors_neighbor_id",
				Columns: []*schema.Column{PlanetNeighborsColumns[1]},

				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GalaxiesTable,
		PlanetsTable,
		PlanetNeighborsTable,
	}
)

func init() {
	PlanetsTable.ForeignKeys[0].RefTable = GalaxiesTable
	PlanetNeighborsTable.ForeignKeys[0].RefTable = PlanetsTable
	PlanetNeighborsTable.ForeignKeys[1].RefTable = PlanetsTable
}
