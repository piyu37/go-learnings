package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ApiResponse struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	Name              string  `json:"name"`
	RuntimeOfSeries   string  `json:"runtime_of_series"`
	Certificate       string  `json:"certificate"`
	RuntimeOfEpisodes string  `json:"runtime_of_episodes"`
	Genre             string  `json:"genre"`
	ImdbRating        float64 `json:"imdb_rating"`
	Overview          string  `json:"overview"`
	NoOfVotes         int     `json:"no_of_votes"`
	ID                int     `json:"id"`
}

func bestInGenre(genre string) string {
	baseURL := "https://jsonmock.hackerrank.com/api/tvseries"
	params := url.Values{}
	// if this doesn't work, then remove genre from params & iterate all data one by one
	params.Add("genre", genre)
	params.Add("page", "")

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	page := 1

	bestRating := -1.0
	bestShowName := ""

	for {
		fullUrl := fmt.Sprintf("%s%d", url, page)
		resp, err := http.Get(fullUrl)
		if err != nil {
			return ""
		}

		var result ApiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		resp.Body.Close()

		if err != nil {
			return ""
		}

		for _, show := range result.Data {
			if show.ImdbRating > bestRating || (show.ImdbRating == bestRating && show.Name < bestShowName) {
				bestRating = show.ImdbRating
				bestShowName = show.Name
			}
		}

		// if we can't filter out based on genres, then we need to check all shows that the input genre is matching or not
		// for _, show := range result.Data {
		// 	match := false

		// 	showGenres := strings.SplitSeq(show.Genre, ",")

		// 	for sg := range showGenres {
		// 		if strings.TrimSpace(sg) == genre {
		// 			match = true
		// 			break
		// 		}
		// 	}

		// 	if !match {
		// 		continue
		// 	}

		// 	if show.ImdbRating > bestRating || (show.ImdbRating == bestRating && show.Name < bestShowName) {
		// 		bestRating = show.ImdbRating
		// 		bestShowName = show.Name
		// 	}
		// }

		page++

		if page > result.TotalPages {
			break
		}
	}

	return bestShowName
}

// API Fetch - Best TV Show in Genre(Agoda)
// Problem Statement:
// Utilize the HTTP GET method to retrieve details about recent TV shows.
// Query [https://jsonmock.hackerrank.com/api/tvseries](https://jsonmock.hackerrank.com/api/tvseries) to locate all shows
// within a specific genre. The results are paginated. To access more pages, add ?page={num} to the URL where {num} represents
// the page number.

// The response is a JSON object with the following 5 fields:

// page: the current page of the results (Number)

// per_page: the maximum number of results returned per page (Number)

// total: the total number of results (Number)

// total_pages: the total number of pages with results (Number)

// data: an array of TV series records

// Example of a data array object:

// JSON
// {
//   "name": "Game of Thrones",
//   "runtime_of_series": "(2011-2019)",
//   "certificate": "A",
//   "runtime_of_episodes": "57 min",
//   "genre": "Action, Adventure, Drama",
//   "imdb_rating": 9.3,
//   "overview": "Nine noble families fight for control over the lands of Westeros, while an ancient enemy returns after being dormant for millennia.",
//   "no_of_votes": 1773458,
//   "id": 1
// }

// In data, each TV series has the following schema:

// name: (String)

// runtime_of_series: years with a new season (String)

// certificate: rating (String)

// runtime_of_episodes: average length per episode in minutes (String)

// genre: genre (String)

// imdb_rating: average viewer rating (Number)

// overview: short description (String)

// no_of_votes: how many votes were cast at IMDb (Number)

// id: unique id (Number)

// Given a genre, find the series with the highest imdb_rating. If there is a tie, return the alphabetically lower name.

// Function Description:
// Complete the function bestInGenre in the editor with the following parameter:

// string genre: the genre to search

// Return:

// string: the highest-rated show in the genre, with the lowest name alphabetically if there is a tie.

// Sample Input 0:
// Action

// Sample Output 0:
// Game of Thrones

// Explanation: The 4 highest-rated shows in the 'Action' genre are shown. 'Game of Thrones', 9.3; 'Avatar: The Last Airbender',
// 9.2; 'Hagane no renkinjutsushi', 9.1; 'Shingeki no kyojin', 8.9.

// Sample Input 1:
// Animation

// Sample Output 1:
// Avatar: The Last Airbender

// Explanation: The 4 highest-rated shows in the 'Animation' genre are shown. 'Avatar: The Last Airbender', 9.2; 'Rick and Morty',
// 9.2; 'Hagane no renkinjutsushi', 9.1; 'Death Note: Desu noto', 9. 'Avatar' is tied with 'Rick and Morty' and is
// lower alphabetically.
func fetchBestTvShow() {
	genre := "Action"
	fmt.Println(bestInGenre(genre))
}
