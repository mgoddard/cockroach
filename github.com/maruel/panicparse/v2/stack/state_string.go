// Code generated by "stringer -type state"; DO NOT EDIT.

package stack

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[looking-0]
	_ = x[done-1]
	_ = x[betweenRoutine-2]
	_ = x[gotRoutineHeader-3]
	_ = x[gotFunc-4]
	_ = x[gotCreated-5]
	_ = x[gotFileFunc-6]
	_ = x[gotFileCreated-7]
	_ = x[gotUnavail-8]
	_ = x[gotRaceHeader1-9]
	_ = x[gotRaceHeader2-10]
	_ = x[gotRaceOperationHeader-11]
	_ = x[gotRaceOperationFunc-12]
	_ = x[gotRaceOperationFile-13]
	_ = x[betweenRaceOperations-14]
	_ = x[gotRaceGoroutineHeader-15]
	_ = x[gotRaceGoroutineFunc-16]
	_ = x[gotRaceGoroutineFile-17]
	_ = x[betweenRaceGoroutines-18]
}

const _state_name = "lookingdonebetweenRoutinegotRoutineHeadergotFuncgotCreatedgotFileFuncgotFileCreatedgotUnavailgotRaceHeader1gotRaceHeader2gotRaceOperationHeadergotRaceOperationFuncgotRaceOperationFilebetweenRaceOperationsgotRaceGoroutineHeadergotRaceGoroutineFuncgotRaceGoroutineFilebetweenRaceGoroutines"

var _state_index = [...]uint16{0, 7, 11, 25, 41, 48, 58, 69, 83, 93, 107, 121, 143, 163, 183, 204, 226, 246, 266, 287}

func (i state) String() string {
	if i < 0 || i >= state(len(_state_index)-1) {
		return "state(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _state_name[_state_index[i]:_state_index[i+1]]
}
