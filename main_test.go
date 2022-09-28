package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readCsvBodyAndConvertToSlice(t *testing.T) {
	cases := []struct {
		msg      string
		a        string
		expected [][]string
	}{
		{
			msg: "normal pattern",
			a:   "test_data/game_score_test.csv",
			expected: [][]string{
				{"2021/01/01 12:00", "player0001", "12345"},
				{"2021/01/01 12:00", "player0002", "10000"},
			},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			got := readCsvBodyAndConvertToSlice(&v.a)
			if !reflect.DeepEqual(got, v.expected) {
				t.Errorf("\nwant:\n %v, \ngot:\n %v", v.expected, got)
			}
		})
	}
}

func Test_calculateMeanScoreAndConvertToMeanScoreRank(t *testing.T) {
	cases := []struct {
		msg      string
		a        []*playLog
		expected []meanScoreRank
	}{
		{
			msg: "normal pattern",
			a: []*playLog{
				{"2021/01/01 12:00", "player0001", 10000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
			},
		},
		{
			msg: "float pattern",
			a: []*playLog{
				{"2021/01/01 12:00", "player0001", 10000},
				{"2021/01/01 12:00", "player0001", 10001},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10001},
			},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			got := calculateMeanScoreAndConvertToMeanScoreRank(v.a)
			if !reflect.DeepEqual(got, v.expected) {
				t.Errorf("\nwant:\n %v, \ngot:\n %v", v.expected, got)
			}
		})
	}
}

func Test_sortOrderByMeanScoreAndRank(t *testing.T) {
	cases := []struct {
		msg      string
		a        []meanScoreRank
		expected []meanScoreRank
	}{
		{
			msg: "normal pattern",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 9000},
				{1, "player0003", 8000},
				{1, "player0004", 7000},
				{1, "player0005", 6000},
				{1, "player0006", 5000},
				{1, "player0007", 4000},
				{1, "player0008", 3000},
				{1, "player0009", 2000},
				{1, "player0010", 1000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
			},
		},
		{
			msg: "same rank pattern",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 10000},
				{1, "player0003", 10000},
				{1, "player0004", 7000},
				{1, "player0005", 6000},
				{1, "player0006", 5000},
				{1, "player0007", 5000},
				{1, "player0008", 5000},
				{1, "player0009", 5000},
				{1, "player0010", 1000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 10000},
				{1, "player0003", 10000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{6, "player0007", 5000},
				{6, "player0008", 5000},
				{6, "player0009", 5000},
				{10, "player0010", 1000},
			},
		},
		{
			msg: "random pattern",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 5000},
				{1, "player0003", 7000},
				{1, "player0004", 10000},
				{1, "player0005", 6000},
				{1, "player0006", 5000},
				{1, "player0007", 1000},
				{1, "player0008", 5000},
				{1, "player0009", 10000},
				{1, "player0010", 5000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0004", 10000},
				{1, "player0009", 10000},
				{4, "player0003", 7000},
				{5, "player0005", 6000},
				{6, "player0002", 5000},
				{6, "player0006", 5000},
				{6, "player0008", 5000},
				{6, "player0010", 5000},
				{10, "player0007", 1000},
			},
		},
		{
			msg: "more ten record pattern",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 5000},
				{1, "player0003", 7000},
				{1, "player0004", 10000},
				{1, "player0005", 6000},
				{1, "player0006", 5000},
				{1, "player0007", 1000},
				{1, "player0008", 5000},
				{1, "player0009", 10000},
				{1, "player0010", 5000},
				{1, "player0011", 900},
				{1, "player0012", 800},
				{1, "player0013", 700},
				{1, "player0014", 600},
				{1, "player0015", 500},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0004", 10000},
				{1, "player0009", 10000},
				{4, "player0003", 7000},
				{5, "player0005", 6000},
				{6, "player0002", 5000},
				{6, "player0006", 5000},
				{6, "player0008", 5000},
				{6, "player0010", 5000},
				{10, "player0007", 1000},
				{11, "player0011", 900},
				{12, "player0012", 800},
				{13, "player0013", 700},
				{14, "player0014", 600},
				{15, "player0015", 500},
			},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			got := sortOrderByMeanScoreAndRank(v.a)
			if !reflect.DeepEqual(got, v.expected) {
				t.Errorf("\nwant:\n %v, \ngot:\n %v", v.expected, got)
			}
		})
	}
}

func Test_extractTopTenScoreRank(t *testing.T) {
	cases := []struct {
		msg      string
		a        []meanScoreRank
		expected []meanScoreRank
	}{
		{
			msg: "normal case",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
			},
		},
		{
			msg:      "empty pattern",
			a:        []meanScoreRank{},
			expected: []meanScoreRank{},
		},
		{
			msg: "one record pattern",
			a: []meanScoreRank{
				{1, "player0001", 10000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
			},
		},
		{
			msg: "more ten record case",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
				{11, "player0011", 900},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
			},
		},
		{
			msg: "more ten record case all tie score",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 10000},
				{1, "player0003", 10000},
				{1, "player0004", 10000},
				{1, "player0005", 10000},
				{1, "player0006", 10000},
				{1, "player0007", 10000},
				{1, "player0008", 10000},
				{1, "player0009", 10000},
				{1, "player0010", 10000},
				{1, "player0011", 10000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{1, "player0002", 10000},
				{1, "player0003", 10000},
				{1, "player0004", 10000},
				{1, "player0005", 10000},
				{1, "player0006", 10000},
				{1, "player0007", 10000},
				{1, "player0008", 10000},
				{1, "player0009", 10000},
				{1, "player0010", 10000},
				{1, "player0011", 10000},
			},
		},
		{
			msg: "output more ten rank case",
			a: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
				{10, "player0011", 1000},
			},
			expected: []meanScoreRank{
				{1, "player0001", 10000},
				{2, "player0002", 9000},
				{3, "player0003", 8000},
				{4, "player0004", 7000},
				{5, "player0005", 6000},
				{6, "player0006", 5000},
				{7, "player0007", 4000},
				{8, "player0008", 3000},
				{9, "player0009", 2000},
				{10, "player0010", 1000},
				{10, "player0011", 1000},
			},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			got := extractTopTenScoreRank(v.a)
			if !reflect.DeepEqual(got, v.expected) {
				t.Errorf("\nwant:\n %v, \ngot:\n %v", v.expected, got)
			}
		})
	}

}

func Test_convertToPlayLogs(t *testing.T) {
	type args struct {
		rows [][]string
	}
	cases := []struct {
		msg      string
		args     args
		expected []*playLog
	}{
		{
			msg: "normal pattern",
			args: args{[][]string{
				{"2021/01/01 12:00", "player0001", "10000"},
				{"2021/01/01 12:00", "player0002", "9000"},
			}},
			expected: []*playLog{
				{"2021/01/01 12:00", "player0001", 10000},
				{"2021/01/01 12:00", "player0002", 9000},
			},
		},
	}
	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			if got := convertToPlayLogs(v.args.rows); !reflect.DeepEqual(got, v.expected) {
				t.Errorf("\nwant:\n %v, \ngot:\n %v", v.expected, got)
			}
		})
	}
}
