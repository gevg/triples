package triples

import "testing"

var ts []Triple = []Triple{
	{0, 0, 2}, {0, 0, 3}, {0, 1, 0}, {1, 0, 4}, {1, 2, 0},
	{1, 2, 1}, {2, 0, 2}, {2, 1, 0}, {3, 2, 1}, {3, 2, 2},
	{4, 2, 4},
}

func TestFrom(t *testing.T) {
	store, err := From(ts)
	if err != nil {
		t.Fatal(err)
	}

	calc := store.Triples()

	if len(calc) != len(ts) {
		t.Errorf(
			"The calculated number of triples is %d != %d the original number of triples.\n",
			len(calc), len(ts),
		)
		return
	}
	for i, want := range ts {
		if got := calc[i]; got != want {
			t.Errorf("Triple[%d] is different from the original triple.\n", i)
		}
	}
}

func BenchmarkFrom(b *testing.B) {

}
