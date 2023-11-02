package packagea

import "testing"

func TestAcc_HelloWorld_PackageA(t *testing.T) {
	want := "Hello from Package A"
	got := HelloWorld_PackageA()
	if want != got {
		t.Fatalf("wanted `%s`, got `%s`", want, got)
	}
}

func TestAcc_HelloWorld_PackageA_copy(t *testing.T) {
	want := "Hello from Package A"
	got := HelloWorld_PackageA()
	if want != got {
		t.Fatalf("wanted `%s`, got `%s`", want, got)
	}
}
