package uuid

import "testing"

/* UNIT TESTS FOR UUID */

func UUIDTest(t *testing.T) {
	id1 := UUID()
	id2 := UUID()

	if id1 == id2 {
		t.Errorf("IDs incorrectly match; id1 is %s and id2 is %s; want no equivalency", id1, id2)
	}
}
