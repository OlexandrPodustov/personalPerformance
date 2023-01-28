package main

import (
	"testing"
)

func TestMultiply1(t *testing.T) {
	tcs := getCases()

	for _, tc := range tcs {
		got := multiply(tc.x, tc.y)
		if tc.want != got {
			t.Errorf("for x: %v, y: %v - wanted: %v, got: %v", tc.x, tc.y, tc.want, got)
			t.FailNow()
		}

	}
}

type tc struct {
	x, y uint64
	want uint64
}

func getCases() []tc {
	return []tc{
		{x: 10, y: 22, want: 220},
		{x: 11, y: 22, want: 242},
		{x: 1234, y: 5678, want: 7006652},
		// {x: 1000, y: 2200, want: 22000000},
		// {x: 21625695688898558125310188636840316594920403182768, y: 13306827740879180856696800391510469038934180115260, want: 22000000},
		// {x: 1685287499328328297814655639278583667919355849391453456921116729, y: 7114192848577754587969744626558571536728983167954552999895348492, want: 0},
		// {x: 49823261, y: 44269423, want: 2205647016448403},
	}
}

// X: 49823261
// Y: 44269423
// Z: 2205647016448403
// Value counter
// 8 -> [5, 4, 3]
// 7 -> [4, 2, 4]

// X: 54761293
// Y: 65394884
// Z: 3581108403425012
// Value counter
// 5 -> [4, 2, 2]

// X: 9313685456934674
// Y: 7658898761837539
// Z: 71332574014261268360454523927286
// Value counter
// 18 -> [5, 1, 2]
// 9 -> [7, 3, 3]
// 8 -> [8, 7, 5]

// X: 3957322621234423
// Y: 7748313756335578
// Z: 30662577304368647842211393201494
// Value counter
// 14 -> [5, 1, 1]
// 9 -> [6, 2, 5]
// 8 -> [13, 8, 7]

// X: 34215432964249374812219364786397
// Y: 94541964835273822784327848699719
// Z: 3234794260129733170788831535430575611379062580407060392628922443
// Value counter
// 45 -> [11, 3, 0]
// 9 -> [28, 19, 14]
// 8 -> [17, 11, 10]

// X: 71611955325935479159397713213124
// Y: 93567726499788166681348352945366
// Z: 6700567850052179472481148730882040129649508491917840721086183384
// Value counter
// 40 -> [15, 4, 2]
// 36 -> [11, 2, 2]
// 10 -> [11, 4, 6]

// X: 8436939677358274975644341226884162349535836199962392872868456892
// Y: 3864264464372346883776335161325428226997417338413342945574123327
// Z: 32602566183268675582196165592691544162522778809155901895617284287276672563976841699892789718741377833554298336397153444191119684
// Value counter
// 72 -> [14, 4, 5]
// 69 -> [12, 3, 0]
// 67 -> [14, 1, 3]

// X: 8711129198194917883527844183686122989894424943636426448417394566
// Y: 2924825637132661199799711722273977411715641477832758942277358764
// Z: 25478534007255378799894857247961445544397925869179138904636157575535921570058983065006369481295619500406669960288667484926076424
// Value counter
// 64 -> [20, 5, 3]
// 60 -> [12, 3, 4]
// 58 -> [12, 5, 2]
