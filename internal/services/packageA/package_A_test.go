package packagea

import "testing"

func TestHelloWorld_PackageA(t *testing.T) {
	want := "Hello from Package A"
	got := HelloWorld_PackageA()
	if want != got {
		t.Fatalf("wanted `%s`, got `%s`", want, got)
	}
}
