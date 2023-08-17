package main

type EntryType string
type ExecHandle uint32

type Entry struct {
	Type EntryType `json:"type"`
	Icon string    `json:"icon"`
}
