package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	// Loop through and download each image
	for i, imageUrl := range urls {
		// Set filename based on position in loop
		filename := fmt.Sprintf("image_%d.jpg", i)
		err := downloadImage(imageUrl, filename)
		// check if download failed
		if err != nil {
			fmt.Printf("Download failed: %v\n", err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup

	// Use waitgroup to download iamges concurrently
	for i, imageUrl := range urls {
		// Add each picture to waitgroup
		wg.Add(1)

		go func(imageUrl string, i int) {
			// Decrement waitgroup when goroutine is done
			defer wg.Done()

			// Set filename based on position in loop
			filename := fmt.Sprintf("image_%d.jpg", i)
			err := downloadImage(imageUrl, filename)
			// check if download failed
			if err != nil {
				fmt.Printf("Download failed: %v\n", err)
			}
		}(imageUrl, i)
	}
	// Wait for all pictures to be downloaded
	wg.Wait()
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/0qHN3cX1M0U/download?ixid=M3wxMjA3fDB8MXxhbGx8Mjl8fHx8fHwyfHwxNzAxMTA5ODMzfA&force=true&w=640",
		"https://unsplash.com/photos/Stki0y7Pepk/download?ixid=M3wxMjA3fDB8MXxhbGx8MTgzfHx8fHx8Mnx8MTcwMTEwOTgyOXw&force=true&w=640",
		"https://cdn.pixabay.com/photo/2023/09/04/07/56/lemur-8232231_1280.png",
		"https://cdn.pixabay.com/photo/2023/11/21/04/12/chicken-8402334_1280.jpg",
		"https://images.pexels.com/photos/17850755/pexels-photo-17850755/free-photo-of-view-of-the-canal-between-traditional-buildings-in-venice-italy.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Get image from address
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the image from the response body to the file.
	_, err = io.Copy(file, resp.Body)
	return nil
}
