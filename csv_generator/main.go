package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	stringInput := strStdin()
	n, err := strconv.Atoi(stringInput)
	if err != nil {
		log.Fatalf("cannot convert to int: %v\n", err)
	}

	rows := generateCsvRows(n)

	timestamp := time.Now().Format("200601021504")
	file, err := os.Create("./result/game_score_" + timestamp + ".csv")
	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(file)

	if err := w.WriteAll(rows); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
}

func strStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func generateCsvRows(n int) [][]string {

	rows := [][]string{}

	header := []string{"create_timestamp", "player_id", "score"}
	rows = append(rows, header)

	for i := 0; i < n; i++ {
		timestamp := time.Now().Format("2006/01/02 15:04")

		rand.Seed(time.Now().UnixNano())
		id := fmt.Sprintf("%04d", rand.Intn(9999))
		playerId := "player" + id

		rand.Seed(time.Now().UnixNano())
		score := strconv.Itoa(rand.Intn(99999))

		rows = append(rows, []string{timestamp, playerId, score})
	}

	return rows
}
