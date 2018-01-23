package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Name   string
	Weight int
	Parent string
	Subs   []string
}

func parseRow(s string) (n string, w int, rel []string) {
	if strings.Contains(s, "->") {
		rx := regexp.MustCompile(`(\w+),`)
		if subs := rx.FindAllString(fmt.Sprintf("%s,", s), -1); subs != nil {
			for _, sn := range subs {
				rel = append(rel, strings.Replace(strings.Trim(sn, ""), ",", "", 1))
			}
		}

	}
	nx := regexp.MustCompile(`^(\w+)\s`)
	n = strings.Trim(nx.FindString(s), " ")

	wx := regexp.MustCompile(`\s\((\d+)\)`)
	if i, err := strconv.Atoi(wx.FindStringSubmatch(s)[1]); err == nil {
		w = i
	}
	return
}

func part1(rows []string) string {
	nodes := map[string]Node{}
	for _, r := range rows {
		if r == "" {
			continue
		}
		n, w, rel := parseRow(r)
		nodes[n] = Node{
			Name:   n,
			Weight: w,
			Subs:   rel,
		}
	}
	for _, n := range nodes {
		for _, sn := range n.Subs {
			sub := nodes[sn]
			sub.Parent = n.Name
			nodes[sn] = sub
		}
	}
	for _, n := range nodes {
		if n.Parent == "" {
			return n.Name
		}
	}
	return ""
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	fmt.Printf("Part1: %s\n", part1(rows))
}
