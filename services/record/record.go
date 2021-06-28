package record

import (
	. "mo-for-desktop/model/record"
)

// local storage
// ../record/record_[encode(key)]  ---- value: encode(Record)

func Lists(groupKey, recordType int) RecordsList {
	if groupKey == 0 && recordType == 0 {
		return ListAll()
	}
	return nil
}

func ListAll() RecordsList {
	return nil
}
