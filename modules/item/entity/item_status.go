package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{"Doing", "Done", "Deleted"}

func (item ItemStatus) String() string {
	return allItemStatus[item]
}
func (item *ItemStatus) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("failed to scan ItemStatus: value is not a string or []byte")
	}

	str = strings.ToLower(str)
	switch str {
	case "doing":
		*item = ItemStatusDoing
	case "done":
		*item = ItemStatusDone
	case "deleted":
		*item = ItemStatusDeleted
	default:
		return fmt.Errorf("invalid status: %s", str)
	}
	return nil
}

func (item ItemStatus) Value() (driver.Value, error) {
	switch item {
	case ItemStatusDoing:
		return "Doing", nil
	case ItemStatusDone:
		return "Done", nil
	case ItemStatusDeleted:
		return "Deleted", nil
	default:
		return nil, fmt.Errorf("invalid status: %d", item)
	}
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	var s string
	// Giải mã JSON vào biến string tạm thời
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	s = strings.ToLower(s)
	switch s {
	case "doing":
		*item = ItemStatusDoing
	case "done":
		*item = ItemStatusDone
	case "deleted":
		*item = ItemStatusDeleted
	default:
		return fmt.Errorf("invalid status: %s", s)
	}
	return nil
}
