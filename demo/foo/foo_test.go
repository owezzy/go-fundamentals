package foo

import "testing"

func Test_IsOdd(t *testing.T) {
	if !IsOdd(1) {
		t.Fatal("1 should be odd")
	}

	if IsOdd(2) {
		t.Fatal("2 should not be odd")
	}

}

func Test_GetUser(t *testing.T) {
	u, err := GetUser(1)
	if err != nil {
		t.Error(err)
	}

	exp := "Owezzy"
	act := u.Name

	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}
