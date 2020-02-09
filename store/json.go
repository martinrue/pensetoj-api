package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/martinrue/pensetoj-api/logger"
)

type memory struct {
	Likes   map[string]int  `json:"likes"`
	Listens map[string]int  `json:"listens"`
	Sources map[string]bool `json:"sources"`
}

// JSON manages application state as a JSON file.
type JSON struct {
	mu sync.Mutex

	logger   *logger.Logger
	filepath string
	memory   *memory
}

// AddAction adds a new action to the store.
func (j *JSON) AddAction(t string, id string, ip string) bool {
	j.mu.Lock()
	defer j.mu.Unlock()

	sourceKey := fmt.Sprintf("%s:%s:%s", t, id, ip)

	if _, ok := j.memory.Sources[sourceKey]; ok {
		return false
	}

	j.memory.Sources[sourceKey] = true

	var data map[string]int

	if t == "like" {
		data = j.memory.Likes
	}

	if t == "listen" {
		data = j.memory.Listens
	}

	if curr, ok := data[id]; ok {
		data[id] = curr + 1
	} else {
		data[id] = 1
	}

	go j.save()
	return true
}

// GetSummary creates a summary of data held in the store.
func (j *JSON) GetSummary() *Summary {
	summary := &Summary{
		Likes: make(map[string]int),
	}

	for key, value := range j.memory.Likes {
		summary.LikesTotal += value
		summary.Likes[key] = value
	}

	for _, value := range j.memory.Listens {
		summary.ListensTotal += value
	}

	return summary
}

// Size returns a string containing store size info.
func (j *JSON) Size() (string, error) {
	fi, err := os.Stat(j.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return "---", nil
		}

		return "", err
	}

	return fmt.Sprintf("%d bytes", fi.Size()), nil
}

// Load attempts to load data from an existing JSON store file.
func (j *JSON) Load() error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if _, err := os.Stat(j.filepath); os.IsNotExist(err) {
		return nil
	}

	data, err := ioutil.ReadFile(j.filepath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, j.memory); err != nil {
		return err
	}

	return nil
}

// save writes the current in-memory structure to a disk file.
func (j *JSON) save() {
	j.mu.Lock()
	defer j.mu.Unlock()

	data, err := json.Marshal(j.memory)
	if err != nil {
		j.logger.System("store: save failed: %v", err)
		return
	}

	if err := ioutil.WriteFile(j.filepath, data, 0644); err != nil {
		j.logger.System("store: save failed: %v", err)
	}
}

// NewJSON loads or creates a new JSON store.
func NewJSON(filepath string, logger *logger.Logger) *JSON {
	return &JSON{
		logger:   logger,
		filepath: filepath,
		memory: &memory{
			Likes:   make(map[string]int),
			Listens: make(map[string]int),
			Sources: make(map[string]bool),
		},
	}
}
