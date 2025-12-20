package hashset

import "testing"

func TestHashSet(t *testing.T) {
	data := []string{
		"s1",
		"s2",
		"s3",
		"s4",
		"s5",
		"s6",
	}
	hs := New[string]()

	for _, s := range data {
		hs.Add(s)
	}

	if hs.Len() != len(data) {
		t.Logf("expected len %d, received %d", len(data), hs.Len())
		t.Fail()
	}

	for _, s := range data {
		if !hs.Has(s) {
			t.Logf("expected hs.Has(%s) to be true", s)
			t.Fail()
		}
	}

	for i := 0; i < len(data); i++ {
		hs.Del(data[i])
		if hs.Has(data[i]) {
			t.Logf("deleted item %s, but still exists", data[i])
			t.Fail()
		}
		added := data[i+1:]
		for _, s := range added {
			if !hs.Has(s) {
				t.Logf("item %s not deleted, but not exists", s)
				t.Fail()
			}
		}
	}

	hs2 := New[string]()
	for range 10 {
		hs2.Add("many")
	}

	if hs2.Len() != 1 {
		t.Logf("expected hs.Len 1, received %d", hs2.Len())
		t.Fail()
	}

	if !hs2.Has("many") {
		t.Log("item 'many' added but not exists")
		t.Fail()
	}

	hs2.Del("many")
	if hs2.Has("many") {
		t.Log("item 'many' deleted but still exists")
		t.Fail()
	}

}
