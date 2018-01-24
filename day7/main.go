package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Name      string
	Weight    int
	Parent    string
	SubWeight int
	Subs      []string
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

func parseTree(rows []string) map[string]Node {
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
	return nodes
}

func part1(rows []string) string {
	nodes := parseTree(rows)
	for _, n := range nodes {
		if n.Parent == "" {
			return n.Name
		}
	}
	return ""
}

func sumNode(n Node, nodes map[string]Node) int {
	sum := n.Weight
	for _, c := range nodes {
		if c.Parent == n.Name {
			sum += c.Weight
			if len(c.Subs) > 0 {
				for _, s := range c.Subs {
					sum += sumNode(nodes[s], nodes)
				}
			}
		}
	}
	return sum
}

func diffNode(pn string, nodes map[string]Node) (n string, diff int) {
	tw := 0
	for _, c := range nodes {
		if c.Parent == pn {
			if tw == 0 {
				tw = c.SubWeight
			}
			if c.SubWeight != tw {
				n = c.Name
				diff = tw - c.SubWeight
				// fmt.Printf("%s d:%d w:%d s:%d\n", n, diff, c.Weight, c.SubWeight)
				if sn, sd := diffNode(n, nodes); sd != 0 {
					n = sn
					diff = sd
					return
				}
			}
		}
	}
	return
}

func part2(rows []string) int {
	nodes := parseTree(rows)
	top := ""
	for _, n := range nodes {
		n.SubWeight += sumNode(n, nodes)
		nodes[n.Name] = n
		if n.Parent == "" {
			top = n.Name
		}
	}
	nn, diff := diffNode(top, nodes)
	n := nodes[nn]
	return n.Weight + diff
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	fmt.Printf("Part1: %s\n", part1(rows))
	fmt.Printf("Part2: %d\n", part2(rows))
}
