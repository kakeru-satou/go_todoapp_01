package services

import "testing"

func UserSetup() *UserService {
	rep := NewFakeUserRepository()
	ser := NewUserService(rep, rep)

	return ser
}

func TestTableSignup(t *testing.T) {
	tests := []struct {
		name              string
		email             string
		isErr             bool
		isCreateSameEmail bool
	}{
		{"ServiceSignup_Success", "test@test.com", false, false},
		{"ServiceSignup_SameEmail", "test@test.com", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ser := UserSetup()
			userName := "test"
			password := "test"

			if tt.isCreateSameEmail {
				_, _, err := ser.Signup(userName, tt.email, password)
				if err != nil {
					t.Fatalf("エラー:%v", err)
				}
			}

			_, _, err := ser.Signup(userName, tt.email, password)
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

func TestTableSignin(t *testing.T) {
	tests := []struct {
		name           string
		createEmail    string
		createPassword string
		getEmail       string
		getPassword    string
		isErr          bool
	}{
		{"ServiceSignin_Success", "test@test.com", "test", "test@test.com", "test", false},
		{"ServiceSignin_InvalidEmail", "test@test.com", "test", "invalid@test.com", "test", true},
		{"ServiceSignin_InvalidPassword", "test@test.com", "test", "test@test.com", "invalid", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ser := UserSetup()

			userName := "test"
			_, _, err := ser.Signup(userName, tt.createEmail, tt.createPassword)
			if err != nil {
				t.Fatalf("エラー:%v", err)
			}

			_, err = ser.Signin(tt.getEmail, tt.getPassword)
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
