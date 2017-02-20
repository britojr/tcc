package characteristic

import (
	"reflect"
	"testing"

	"github.com/britojr/tcc/ktree"
)

var relabeledFig1A ktree.Ktree = ktree.Ktree{
	[][]int{
		{4, 6, 7, 8},
		{3, 8, 9, 10},
		{8, 9, 10},
		{1, 9, 10},
		{0, 6, 7, 8, 9},
		{7, 8, 10},
		{0, 4, 7},
		{0, 4, 5, 6, 8, 9, 10},
		{0, 1, 2, 4, 5, 7, 9, 10},
		{1, 2, 3, 4, 7, 8, 10},
		{1, 2, 3, 5, 7, 8, 9},
	},
	3,
}

var Rk ktree.RenyiKtree = ktree.RenyiKtree{
	&relabeledFig1A,
	[]int{1, 2, 8},
}

var Tk Tree = Tree{
	[]int{-1, 5, 0, 0, 2, 8, 8, 1, 0},
	[]int{-1, 2, -1, -1, 0, 2, 1, 2, -1},
}
var iphi = []int{0, 10, 9, 3, 4, 5, 6, 7, 1, 2, 8}
var n, k = 11, 3
var childrenTk = [][]int{
	{2, 3, 8},
	{7},
	{4},
	(nil),
	(nil),
	{1},
	(nil),
	(nil),
	{5, 6},
}
var adjTk = [][]int{
	{2, 3, 8},
	{7, 5},
	{4, 0},
	{0},
	{2},
	{1, 8},
	{8},
	{1},
	{5, 6, 0},
}
var KTk = [][]int{
	{9, 10, 11},
	{5, 8, 9},
	{9, 10, 11},
	{9, 10, 11},
	{2, 10, 11},
	{8, 9, 10},
	{8, 9, 11},
	{1, 5, 8},
	{9, 10, 11},
}
var cliquesTk = [][]int{
	[]int{1, 2, 8},
	[]int{0, 4, 7, 1},
	[]int{10, 1, 2, 8},
	[]int{9, 1, 2, 8},
	[]int{3, 10, 2, 8},
	[]int{4, 7, 1, 2},
	[]int{5, 7, 1, 8},
	[]int{6, 0, 4, 7},
	[]int{7, 1, 2, 8},
}

func TestTreeFrom(t *testing.T) {
	want := &Tk
	got := TreeFrom(&Rk)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TreeFrom(%v) = %v; want %v", Rk, got, want)
	}
}

func TestRenyiKtreeFrom(t *testing.T) {
	want := &Rk
	got := RenyiKtreeFrom(11, 3, []int{1, 2, 8}, &Tk)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RenyiKtreeFrom(%v) = %v; want %v", Tk, got, want)
	}
}

func TestChildrenList(t *testing.T) {
	want := childrenTk
	got := childrenList(&Tk)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("childrenList(%v) = %v; want %v", Tk, got, want)
	}
}

func TestExtractCliqueTree(t *testing.T) {
	wantclique, wantadj := cliquesTk, adjTk
	gotclique, gotadj := ExtractCliqueTree(&Tk, iphi)
	if !reflect.DeepEqual(gotclique, wantclique) {
		t.Errorf("Clique: got %v; want %v", gotclique, wantclique)
	}
	if !reflect.DeepEqual(gotadj, wantadj) {
		t.Errorf("Adj: got %v; want %v", gotadj, wantadj)
	}
}
