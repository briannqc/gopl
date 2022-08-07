package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	apikey := flag.String("apikey", "", "Required: apikey used to call http://www.omdbapi.com APIs")
	title := flag.String("t", "", "Required: movie title")
	flag.Parse()

	if err := SearchAndDownloadPoster(*apikey, *title); err != nil {
		log.Fatalln(err)
	}
}

type Movie struct {
	IMDBID string
	Title  string
	Poster string
}

func SearchAndDownloadPoster(apikey, title string) error {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apikey, title)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got non-ok response, status: %v", resp.Status)
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return err
	}

	titleSnake := strings.ReplaceAll(strings.ToLower(movie.Title), " ", "_")
	ext := path.Ext(movie.Poster)
	outFileName := titleSnake + ext

	return downloadPoster(movie.Poster, outFileName)
}

func downloadPoster(posterURL string, outFileName string) error {
	outFile, err := os.Create(outFileName)
	if err != nil {
		return fmt.Errorf("create output file failed, err: %w", err)
	}
	defer outFile.Close()

	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get poster non-ok, status: %v", resp.Status)
	}

	if _, err := io.Copy(outFile, resp.Body); err != nil {
		return err
	}
	return nil
}
