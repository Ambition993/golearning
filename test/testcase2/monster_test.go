package monster

import "testing"

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "niha",
		Age:   33,
		Skill: "play",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("error")
	}
	t.Logf("ok")
	res = monster.Restore()
	if !res {
		t.Fatalf("error")
	}
	t.Logf("ok")
}
