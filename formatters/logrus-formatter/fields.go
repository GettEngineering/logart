package logrusformatter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

type OrderedFields []string

type field struct {
	key string
	val string
}

func logFields(e *logrus.Entry, o FormatOptions) []field {
	return convertFieldsToPairs(e.Data, o)
}

func buildFields(pairs []field) string {
	if len(pairs) == 0 {
		return ""
	}

	var builtPairs []string
	for _, p := range pairs {
		var sb strings.Builder
		sb.WriteString(p.key)
		sb.WriteString(":")
		sb.WriteString(p.val)
		builtPairs = append(builtPairs, sb.String())
	}

	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString(strings.Join(builtPairs, ", "))
	sb.WriteString("}")
	return sb.String()
}

func convertFieldsToPairs(fields map[string]interface{}, o FormatOptions) []field {
	ff := clone(fields) // clone original fields - to not affect original data

	firstEverPrintedFields := fill(o.FirstEverPrintedFields, ff)
	lastEverPrintedFields := fill(o.LastEverPrintedFields, ff)

	printedInMiddleFields := make([]field, 0, len(ff))
	for k, v := range ff {
		addTo(&printedInMiddleFields, k, v)
	}

	sort.Slice(printedInMiddleFields, func(i, j int) bool {
		return printedInMiddleFields[i].key < printedInMiddleFields[j].key
	})

	res := append(firstEverPrintedFields, printedInMiddleFields...)
	res = append(res, lastEverPrintedFields...)

	for i := range res {
		res[i].val = removeWhiteSpaces(res[i].val)
	}
	return res
}

func fill(fields OrderedFields, ff map[string]interface{}) []field {
	firstEverPrintedFields := make([]field, 0, len(fields))
	for _, key := range fields {
		moveIfExist(key, ff, &firstEverPrintedFields)
	}
	return firstEverPrintedFields
}

func clone(fields map[string]interface{}) map[string]interface{} {
	ff := map[string]interface{}{}
	for k, v := range fields {
		ff[k] = v
	}
	return ff
}

func addTo(ff *[]field, k string, v interface{}) {
	*ff = append(*ff, field{key: k, val: fmt.Sprintf("%+v", v)})
}

func moveIfExist(key string, ff map[string]interface{}, lastEverPrintedFields *[]field) {
	if v, ok := ff[key]; ok {
		addTo(lastEverPrintedFields, key, v)
		delete(ff, key)
	}

}

func removeWhiteSpaces(v string) string {
	// TODO: make this smarter
	v = strings.Replace(v, "\n", " ", -1)
	v = strings.Replace(v, "\t", " ", -1)

	// remove spaces (limited to 20 to be safe)
	for i := 0; i < 20 && strings.Contains(v, "  "); i++ {
		v = strings.Replace(v, "  ", " ", -1)
	}

	return v
}
