package query

type V []interface{}

type Item struct {
	Key    string
	Values V
	Skip   bool
}

func Build(items []Item) (string, V) {
	var query string
	var values V
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
