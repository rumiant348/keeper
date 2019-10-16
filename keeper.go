// Package provides functions to save data to a file
package keeper

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	Title   string
	Content string
}

type notesFile struct {
	File *os.File
}

func NewNotesFile(name string) (*notesFile, error) {
	f, err := getFile(name)
	if err != nil {
		return nil, err
	}
	return &notesFile{
		File: f,
	}, nil
}

func getFile(name string) (*os.File, error) {
	var f *os.File

	// check of file existing, is somewhat ugly
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			// file doesn't exist, creating
			f, err = os.Create(name)
			if err != nil {
				return nil, err
			}
		} else {
			//some error accessing the file
			return nil, err
		}
	} else {
		//file exists, opening
		f, err = os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func (n *notesFile) Save(s string) error {
	_, err := n.File.WriteString(fmt.Sprintf("%v\n", s))
	return err
}

func (n *notesFile) Close() error {
	err := n.File.Close()
	return err
}

func getLatest() {
	// TODO: implement
}

func getAll() {
	// TODO: implement
}

func (n *notesFile) Print() error {
	f, err := os.Open(n.File.Name())
	if err != nil {
		return err
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = f.Close()
	return err
}
