package util

import "testing"

const succeed = "\u2713"
const failed = "\u2717"

func TestContainsString(t *testing.T) {
	playerNames := []string{"David Bechkam", "Nathaniel Chalobah", "Jasper Cillessen", "Cristiano Ronaldo"}

	t.Log("Given the need to test ContainsString.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for ContainsString %q", playerNames, "David Bechkam")
		{
			output := ContainsString(playerNames, "David Bechkam")

			t.Logf("\t%s\tShould be able to make the ContainsString call.", succeed)
			if output == true {
				t.Logf("\t%s\tShould get true as output", succeed)
			} else {
				t.Errorf("\t%s\tShould get true as output : false", failed)
			}
		}
	}
}
