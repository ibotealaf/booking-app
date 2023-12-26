package main

import "testing"

func TestStats0(t *testing.T) {
	got := stats([]string{})
	expected := "No customers yet"

	if expected != got {
		t.Errorf("expected %q \ngot %q", expected, got)
	}
}

func TestStats1(t *testing.T) {
	got := stats([]string{"tobi"})
	expected := "We had a total of 1 customer(s) with the following names: tobi"

	if expected != got {
		t.Errorf("expected %q \ngot %q", expected, got)
	}
}

func TestStats2(t *testing.T) {
	got := stats([]string{"tobi", "tobi"})
	expected := "We had a total of 1 customer(s) with the following names: tobi"

	if expected != got {
		t.Errorf("expected %q \ngot %q", expected, got)
	}
}

func TestStats3(t *testing.T) {
	got := stats([]string{"tobi", "temi"})
	expected := "We had a total of 2 customer(s) with the following names: tobi & temi"

	if expected != got {
		t.Errorf("expected %q \ngot %q", expected, got)
	}
}

func TestStats4(t *testing.T) {
	got := stats([]string{"tobi", "temi", "brave"})
	expected := "We had a total of 3 customer(s) with the following names: tobi, temi & brave"

	if expected != got {
		t.Errorf("expected %q \ngot %q", expected, got)
	}
}
