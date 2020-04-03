package hash

import "testing"

/* UNIT TESTS FOR HASH */

func HashTest(t *testing.T) {
	passHash1 := Hash("password")
	passHash2 := Hash("password")

	testVal := passHash1 == passHash2
	if !testVal {
		t.Errorf("Hashes did not match; hash1 is %s and hash2 is %s want equivalency", passHash1, passHash2)
	}

	passHash1 = Hash("password")
	passHash2 = Hash("passw0rd")
	testVal = passHash1 == passHash2
	if testVal {
		t.Errorf("Hashes incorrectly match; hash1 is %s and hash2 is %s want no equivalency", passHash1, passHash2)
	}

	passHash1 = Hash("password")
	password := "password"
	testVal = UnHashAndCompare(passHash1, password)
	if testVal {
		t.Errorf("Hash does not match password; hash is %s and password is %s want equivalency", passHash1, password)
	}
}
