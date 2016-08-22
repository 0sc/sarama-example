package main

var fakeDB string

func saveMessage(text string) { fakeDB = text }

func getMessage() string { return fakeDB }
