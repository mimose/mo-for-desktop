package record

import (
	. "mo-for-desktop/model/record"
)

// local storage
// ../record/record_key  ---- value: Encrypted(Record)

func Lists(spaceKey string, recordType int) RecordsList {
	if spaceKey == "" && recordType == 0 {
		return ListAll()
	}
	return nil
}

func ListAll() RecordsList {
	return nil
}
