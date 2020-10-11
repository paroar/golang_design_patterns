package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{0, 0},
		{3, 14},
		{5, 55},
	}

	var res int
	for _, test := range tableTest {
		res = LaunchPipelines(test[0])
		if res != test[1] {
			t.Logf("%d == %d\n", test[1], res)
			t.Fatal()
		}
	}
}
