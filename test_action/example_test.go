package test_action

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHu Actions. Dev.to is awesome" {
		t.Errorf("Greet() = %s; want Hello GitHu Actions. Dev.to is awesome", result)
	}
}
