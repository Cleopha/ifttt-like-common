package common

import (
	"github.com/Cleopha/ifttt-like-common/protos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAction(t *testing.T) {
	testCases := []struct {
		name          string
		task          *protos.Task
		expected      [2]string
		expectedError error
	}{
		{
			name: "Valid google calendar task",
			task: &protos.Task{Action: protos.TaskAction_GOOGLE_CREATE_NEW_EVENT},
			expected: [2]string{
				"google",
				"CreateNewEvent",
			},
			expectedError: nil,
		},
		{
			name: "Valid google document task",
			task: &protos.Task{Action: protos.TaskAction_GOOGLE_CREATE_NEW_DOCUMENT},
			expected: [2]string{
				"google",
				"CreateNewDocument",
			},
			expectedError: nil,
		},
		{
			name: "Valid google sheet task",
			task: &protos.Task{Action: protos.TaskAction_GOOGLE_CREATE_NEW_SHEET},
			expected: [2]string{
				"google",
				"CreateNewSheet",
			},
			expectedError: nil,
		},
		{
			name: "Non existing task",
			task: &protos.Task{Action: 10},
			expected: [2]string{
				"",
				"",
			},
			expectedError: ErrInvalidTaskFormat,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			service, method, _, err := ParseAction(tt.task)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expected[0], service)
			assert.Equal(t, tt.expected[1], method)
		})
	}
}
