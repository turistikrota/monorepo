package query

type Item struct {
	Key    string
	Values []interface{}
	Skip   bool
}

func Build(items []Item) (string, []interface{}) {
	var query string
	var values []interface{}
	for _, item := range items {
		if item.Skip {
			continue
		}
		if len(item.Values) > 0 && item.Values[0] != "" {
			query += item.Key + " AND "
			values = append(values, item.Values...)
		}
	}

	if len(query) == 0 {
		return "", nil
	}

	return query[:len(query)-5], values
}
