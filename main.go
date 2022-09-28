package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

type playLog struct {
	createdAt string
	playerId  string
	score     int
}
type meanScoreRank struct {
	rank       int
	player_id  string
	mean_score int
}

func main() {
	inputFilePath := flag.String("filepath", "input_data/game_score.csv", "input file path") // CSVからJSOINに仕様が変わる想定もするなら、。。
	flag.Parse()
	rows := readCsvBodyAndConvertToSlice(inputFilePath)
	playLogs := convertToPlayLogs(rows)
	meanScores := calculateMeanScoreAndConvertToMeanScoreRank(playLogs)
	meanScoreRanks := sortOrderByMeanScoreAndRank(meanScores)
	topTenMeanScoreRanks := extractTopTenScoreRank(meanScoreRanks)
	outputMeanScoreRanksToCsv(topTenMeanScoreRanks)
}

func readCsvBodyAndConvertToSlice(filepath *string) [][]string {
	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	// ヘッダーはスキップ
	_, err = r.Read()
	if err != nil {
		log.Fatalln(err)
	}

	var rows [][]string
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		rows = append(rows, row)
	}

	return rows
}

func convertToPlayLogs(rows [][]string) []*playLog {
	playLogs := []*playLog{}
	for _, row := range rows {
		score, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatalln(err)
		}
		pl := playLog{createdAt: row[0], playerId: row[1], score: score}
		playLogs = append(playLogs, &pl)
	}
	return playLogs
}

func calculateMeanScoreAndConvertToMeanScoreRank(playLogs []*playLog) []meanScoreRank {
	scoreMap := make(map[string]int)
	countMap := make(map[string]int)

	for _, pl := range playLogs {
		scoreMap[pl.playerId] += pl.score
		countMap[pl.playerId]++
	}

	// 集計したscoreを足したscoreの数で割って平均値を算出する
	meanScoreRanks := []meanScoreRank{}
	for pid, pscore := range scoreMap {
		meanScore := int(math.Round(float64(pscore) / float64(countMap[pid])))
		msr := meanScoreRank{1, pid, meanScore}
		meanScoreRanks = append(meanScoreRanks, msr)
	}

	return meanScoreRanks
}

// ここ計算量多そうだからちょっとアルゴリズム考え直した方がいいかも。
func sortOrderByMeanScoreAndRank(msr []meanScoreRank) []meanScoreRank {
	sort.SliceStable(msr, func(i, j int) bool {
		return msr[i].mean_score > msr[j].mean_score
	})

	// 計算量O(n^2)か、、、だけど、nはlen(msr)なので、ユーザー上限10000がmax
	// for i := 0; i < len(msr); i++ {
	// 	for j := i; j < len(msr); j++ {
	// 		if msr[i].mean_score > msr[j].mean_score {
	// 			msr[j].rank += 1
	// 		}
	// 	}
	// }
	rank(msr)

	return msr
}

// 計算量的にはO(n)となる。ただし、スコア上限99999回は施工するため、
// ユーザー上限10000を考慮すると、ユーザーが316名を超えてくるなら、316^2 ≒ 100000　なので、こちらの方が効率がいい。
// []meanScoreRankを加工するだけなので、メソッドにできそう。
func rank(msr []meanScoreRank) {
	// 既にsortされているものとする

	// mean_scoreの最大を99999とする。
	// 100行までなら総当たりで10000回
	// 100000回を上回ってくるならこちらの方法の方が試行回数が少ない
	num := len(msr)
	const max = 99999

	// インデックスが0から10までの配列作るならarr[11]
	// インデックスが0から100000までの配列作るならarr[100001]
	var ranking [max + 2]int

	meanScores := []int{}
	for _, v := range msr {
		meanScores = append(meanScores, v.mean_score)
	}

	// 全ランクを0で初期化
	for i := 0; i <= max; i++ {
		ranking[i] = 0
	}

	// 点数をインデックスとして+1
	for i := 0; i < num; i++ {
		ranking[meanScores[i]]++
	}

	// ranking[]の一番右端の余分な部分に1を入れる
	ranking[max+1] = 1

	// 一つ右隣の要素を足す
	// 右から順にやっていくので、iはmaxから始めて、0まで
	for i := max; i >= 0; i-- {
		ranking[i] += ranking[i+1]
	}

	// 最高平均点が90000だった場合、ranking[90000+1] == 1
	// 2位が80000だった場合、ranking[80000+1] == 2
	// のように、一つ右隣に順位が書いてある
	for i := 0; i < num; i++ {
		msr[i].rank = ranking[meanScores[i]+1]
	}
}

/*
// メソッドにする。
func extractTopTenScoreRank(srArr []meanScoreRank) []meanScoreRank {
	if len(srArr) == 0 {
		return srArr
	}
	result := []meanScoreRank{}
	for i, v := range srArr {
		// 必ず上から10個目までは抽出
		if i < 10 {
			result = append(result, v)
		}
		// 10個以上の出力になるときは必ず、11個目以降が10個目のランクと同じとき
		if i >= 10 && v.rank == srArr[9].rank {
			result = append(result, v)
		}
	}

	return result
}
*/

func extractTopTenScoreRank(srArr []meanScoreRank) []meanScoreRank {
	// スライスを新しく定義して、それに必要なものだけ移し替えるのが一つ上の関数のやり方。
	// 今回は、srArrはそのままに、必要ない要素を削除するような方法
	if len(srArr) == 0 {
		log.Fatalln("meanScoreRanks is empty")
		return srArr
	}

	for i, v := range srArr {
		if i >= 10 && v.rank != srArr[9].rank {
			srArr = srArr[:i]
			break
		}
	}
	return srArr
}

func outputMeanScoreRanksToCsv(meanScoreRanks []meanScoreRank) {
	var records [][]string
	csvHeader := []string{"rank", "player_id", "mean_score"}
	records = append(records, csvHeader)

	for _, v := range meanScoreRanks {
		strRank := strconv.Itoa(v.rank)
		strMeanScore := strconv.Itoa(v.mean_score)
		record := []string{strRank, v.player_id, strMeanScore}
		records = append(records, record)
	}

	timestamp := time.Now().Format("2006-01-02T15:04:05Z07:00")
	// TODO 出力先が変わったらどうする？その時のために、環境変数か何かに切り出しておく？
	file, err := os.Create("./result/get_ranking game_score_log_" + timestamp + ".csv")
	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(file)

	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
}
