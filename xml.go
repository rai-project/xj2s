package xj2s

import (
	"regexp"
	"strings"

	"github.com/knq/snaker"
)

func XmlPath2SrtructLinesNoNesting(paths []string) (string, map[string]StructNode, map[string]map[string]StructNode) {
	var RootName string
	RootStruct := make(map[string]StructNode)
	RestStructs := make(map[string]map[string]StructNode)

	RootName = strings.Split(paths[0], ".")[0]

	deDuplicateMap := make(map[string]string)

	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	for _, path := range paths {
		path = removeNum.ReplaceAllString(path, "[]")
		Flods := strings.Count(path, "[")
		path = strings.Replace(path, "[]", "", -1)
		splitedPath := strings.Split(path, ".")
		last := splitedPath[len(splitedPath)-1]
		if strings.Index(last, "-") == 0 { //Attr
			if RootName == splitedPath[len(splitedPath)-2] { //RootAttr
				NodeName := snaker.SnakeToCamel(last[1:])
				xmlRoute := "`xml:" + `"` + last[1:] + `,attr"` + "`"
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = "Rss" + NodeName
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				RootStruct[xmlRoute] = StructLineAppend
			} else { //NoneRootAttr
				NodeName := snaker.SnakeToCamel(splitedPath[len(splitedPath)-2])
				xmlRoute := strings.Join(splitedPath[1:len(splitedPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = ""
						for _, v := range strings.Split(xmlRoute, ">") {
							NodeName = snaker.SnakeToCamel(v) + NodeName
						}
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				RootStruct[xmlRoute] = StructLineAppend

				LeafName := snaker.SnakeToCamel(last[1:])
				RsetStructLineAppend := StructNode{Name: LeafName, Type: "string", Path: "`xml:" + `"` + last[1:] + `,attr"` + "`"}

				// log.Println(NodeName, LeafName)
				if _, exist := RestStructs[NodeName]; exist {
					RestStructs[NodeName][LeafName] = RsetStructLineAppend
				} else {
					NewLeafStruct := make(map[string]StructNode)
					NewLeafStruct[LeafName] = RsetStructLineAppend
					RestStructs[NodeName] = NewLeafStruct
				}

			}
		} else if strings.Index(last, "#") == 0 { //chardata
			if RootName == splitedPath[len(splitedPath)-2] { //RootChartata
				NodeName := snaker.SnakeToCamel(last[1:])
				xmlRoute := "`xml:" + `",chardata"` + "`"
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = "Rss" + NodeName
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				RootStruct[xmlRoute] = StructLineAppend
			} else { //NonRootChardata
				NodeName := snaker.SnakeToCamel(splitedPath[len(splitedPath)-2])
				xmlRoute := strings.Join(splitedPath[1:len(splitedPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = ""
						for _, v := range strings.Split(xmlRoute, ">") {
							NodeName = snaker.SnakeToCamel(v) + NodeName
						}
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				RootStruct[xmlRoute] = StructLineAppend

				LeafName := snaker.SnakeToCamel(last[1:])
				RsetStructLineAppend := StructNode{Name: LeafName, Type: "string", Path: "`xml:" + `",chardata"` + "`"}

				if _, exist := RestStructs[NodeName]; exist {
					RestStructs[NodeName][LeafName] = RsetStructLineAppend
				} else {
					NewLeafStruct := make(map[string]StructNode)
					NewLeafStruct[LeafName] = RsetStructLineAppend
					RestStructs[NodeName] = NewLeafStruct
				}
			}
		} else { //NormalString
			NodeName := snaker.SnakeToCamel(splitedPath[len(splitedPath)-1])
			xmlRoute := strings.Join(splitedPath[1:], ">")
			xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
			Stype := "string"
			for i := 0; i < Flods; i++ {
				Stype = "[]" + Stype
			}
			if _, exist := deDuplicateMap[NodeName]; exist {
				if deDuplicateMap[NodeName] != xmlRoute {
					NodeName = ""
					for _, v := range strings.Split(xmlRoute, ">") {
						NodeName = snaker.SnakeToCamel(v) + NodeName
					}
					deDuplicateMap[NodeName] = xmlRoute
				}
			} else {
				deDuplicateMap[NodeName] = xmlRoute
			}
			StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
			RootStruct[xmlRoute] = StructLineAppend
		}
	}
	return snaker.SnakeToCamel(RootName), RootStruct, RestStructs
}
