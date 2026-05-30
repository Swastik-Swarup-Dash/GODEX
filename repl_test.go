package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input: "Hello World",
            expected: []string{"hello", "world"},
        },
        {
            input: "Go is Great",
            expected: []string{"go", "is", "great"},
        },
        {
            input:    "",
            expected: []string{},
        },
        {
            input:    "   ",
            expected: []string{},
        },
        {
            input:    "\tGo\n\nIs\tGreat\t",
            expected: []string{"go", "is", "great"},
        },
        {
            input:    "  Multiple   spaces  here ",
            expected: []string{"multiple", "spaces", "here"},
        },
        {
            input:    "Hello, world!",
            expected: []string{"hello,", "world!"},
        },
        {
            input:    "Grüße München",
            expected: []string{"grüße", "münchen"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Expected %v, got %v", c.expected, actual)
            continue
        }
        for i := 0; i < len(c.expected); i++ {
            if actual[i] != c.expected[i] {
                t.Errorf("Expected %v, got %v", c.expected, actual)
                break
            }
        }
    }

}