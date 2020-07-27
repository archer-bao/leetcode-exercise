func TestConvert(t *testing.T) {
	if convert("LEETCODEISHIRING", 3) != "LCIRETOESIIGEDHN" {
		t.Fatal()
	}

	if convert("AB", 3) != "AB" {
		t.Fatal()
	}

}
