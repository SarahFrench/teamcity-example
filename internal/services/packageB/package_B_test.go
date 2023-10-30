package packageb

import "testing"

func TestHelloWorld_PackageB(t *testing.T) {
	want := "Hello from Package B"
	got := HelloWorld_PackageB()
	if want != got {
		t.Fatalf("wanted `%s`, got `%s`", want, got)
	}
}
