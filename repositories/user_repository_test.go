package repositories

import (
	"go-learning/todoApp/db"
	"go-learning/todoApp/models"
	"testing"
)

func TestTableCreate(t *testing.T) {
	tests := []struct {
		name              string
		email             string
		isErr             bool
		isCreateSameEmail bool
	}{
		{"RepositoryCreate_Success", "test@test.com", false, false},
		{"RepositoryCreate_SameEmail", "test@test.com", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db.Init()
			rep := UserRepository{}
			user := models.User{
				Name:     "test",
				Email:    tt.email,
				Password: "test",
			}

			db.DB.Where("email = ?", tt.email).Delete(&models.User{})

			if tt.isCreateSameEmail == true {
				_, err := rep.Create(user)
				if err != nil {
					t.Fatalf("エラー:%s", err)
				}
			}

			_, err := rep.Create(user)
			if (err != nil) != tt.isErr {
				t.Errorf(
					"%s: 想定=%v, 実際=%v",
					tt.name,
					tt.isErr,
					(err != nil),
				)
			}
		})
	}
}

func TestTableGetByEmail(t *testing.T) {
	tests := []struct {
		name        string
		createEmail string
		getEmail    string
		isErr       bool
	}{
		{"RepositoryGetEmail_Success", "test@test.com", "test@test.com", false},
		{"RepositoryGetEmail_OtherEmail", "test@test.com", "other@test.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db.Init()

			rep := UserRepository{}
			user := models.User{
				Name:     "test",
				Email:    tt.createEmail,
				Password: "test",
			}

			db.DB.Where("email = ?", tt.createEmail).Delete(&models.User{})

			_, err := rep.Create(user)
			if err != nil {
				t.Fatalf("エラー:%v", err)
			}

			_, err = rep.GetByEmail(tt.getEmail)
			if (err != nil) != tt.isErr {
				t.Errorf(
					"%s: 想定=%v, 実際=%v",
					tt.name,
					tt.isErr,
					(err != nil),
				)
			}
		})
	}
}
