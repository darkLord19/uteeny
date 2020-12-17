package main

import (
	"crypto/sha1"
	"encoding/json"
	"time"
)

type url struct {
	hash     string
	original string
}

func (u *url) calculateHash() string {
	h := sha1.New()
	h.Write([]byte(u.original))
	h.Write([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	bs := h.Sum(nil)
	return string(bs)[:8]
}

func (u *url) toJSON() []byte {
	data, err := json.Marshal(u)
	if err != nil {
		return []byte("")
	}
	return data
}
