package repository

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
)

func TestSetUpDatabaseConnection(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// {
		// 	name: "should return a connected database",
		// 	want: &gorm.DB{
		// 		Config:       &gorm.Config{},
		// 		Error:        nil,
		// 		RowsAffected: 0,
		// 		Statement:    &gorm.Statement{},
		// 	},
		// },
		// {
		// 	name: "should return an error if the database connection fails",
		// 	want: &gorm.DB{Error: errors.New("failed to connect to database")},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUpDatabaseConnection(); !cmp.Equal(got, tt.want, cmpopts.EquateErrors()) {
				t.Errorf("SetUpDatabaseConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
