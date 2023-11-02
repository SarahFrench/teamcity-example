package packagec

import "testing"

func TestAcc_HelloWorld_PackageC(t *testing.T) {
	want := "Hello from Package C"
	got := HelloWorld_PackageC()
	if want != got {
		t.Fatalf("wanted `%s`, got `%s`", want, got)
	}
}
