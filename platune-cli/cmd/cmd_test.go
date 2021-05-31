package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/aschey/platune/cli/v2/test"
	"github.com/aschey/platune/cli/v2/utils"
	platune "github.com/aschey/platune/client"
	"github.com/c-bata/go-prompt"
	"github.com/golang/mock/gomock"
)

var originalArgs = os.Args

func runTest(t *testing.T, expected string,
	expectFunc func(expect *test.MockPlayerClientMockRecorder), args ...string) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := test.NewMockPlayerClient(ctrl)
	expectFunc(mock.EXPECT())
	utils.Client = utils.NewTestClient(mock)

	os.Args = append(originalArgs, args...)
	rescueStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	rootCmd.SetOut(wOut)
	os.Stdout = wOut

	if err := rootCmd.Execute(); err != nil {
		t.Errorf(err.Error())
	}
	wOut.Close()
	rootCmd.SetOut(rescueStdout)
	os.Stdout = rescueStdout
	var out, _ = ioutil.ReadAll(rOut)
	outStr := string(out)
	if outStr != expected {
		t.Errorf("Expected %s, Got %s", expected, outStr)
	}
}

func TestAddQueue(t *testing.T) {
	testSong := "test"
	runTest(t, "Added\n", func(expect *test.MockPlayerClientMockRecorder) {
		matcher := test.NewMatcher(func(arg interface{}) bool {
			return arg.(*platune.AddToQueueRequest).Song == testSong
		})
		expect.AddToQueue(gomock.Any(), matcher)
	}, "addQueue", testSong)
}

func TestSetQueue(t *testing.T) {
	testSong1 := "test1"
	testSong2 := "test2"
	runTest(t, "Queue Set\n", func(expect *test.MockPlayerClientMockRecorder) {
		matcher := test.NewMatcher(func(arg interface{}) bool {
			queue := arg.(*platune.QueueRequest).Queue
			return queue[0] == testSong1 && queue[1] == testSong2
		})
		expect.SetQueue(gomock.Any(), matcher)
	}, "setQueue", testSong1, testSong2)
}

func TestSeek(t *testing.T) {
	testCases := []struct {
		formatStr string
		expected  uint64
	}{
		{"30", 30000},
		{"2:30", 150000},
		{"3:05:30", 11130000},
	}

	for _, tc := range testCases {
		matcher := test.NewMatcher(func(arg interface{}) bool {
			return arg.(*platune.SeekRequest).Millis == tc.expected
		})
		runTest(t, fmt.Sprintf("Seeked to %s\n", tc.formatStr), func(expect *test.MockPlayerClientMockRecorder) {
			expect.Seek(gomock.Any(), matcher)
		}, "seek", tc.formatStr)
	}

}

func TestResume(t *testing.T) {
	runTest(t, "Resumed\n", func(expect *test.MockPlayerClientMockRecorder) {
		expect.Resume(gomock.Any(), gomock.Any())
	}, "resume")
}

func TestPause(t *testing.T) {
	runTest(t, "Paused\n", func(expect *test.MockPlayerClientMockRecorder) {
		expect.Pause(gomock.Any(), gomock.Any())
	}, "pause")
}

func TestNext(t *testing.T) {
	runTest(t, "Next\n", func(expect *test.MockPlayerClientMockRecorder) {
		expect.Next(gomock.Any(), gomock.Any())
	}, "next")
}

func TestPrevious(t *testing.T) {
	runTest(t, "Previous\n", func(expect *test.MockPlayerClientMockRecorder) {
		expect.Previous(gomock.Any(), gomock.Any())
	}, "previous")
}

func TestStop(t *testing.T) {
	runTest(t, "Stopped\n", func(expect *test.MockPlayerClientMockRecorder) {
		expect.Stop(gomock.Any(), gomock.Any())
	}, "stop")
}

func TestAddQueueCompleter(t *testing.T) {
	state := newCmdState()

	buf := prompt.NewBuffer()
	buf.InsertText("add-queue root", false, true)
	doc := buf.Document()

	results := state.completer(*doc)
	if len(results) != 1 {
		t.Error("Should've found one result")
	}
	if results[0].Text != "root.go" {
		t.Error("Result should be root.go")
	}
}

func TestSetQueueCompleter(t *testing.T) {
	state := newCmdState()
	state.isSetQueueMode = true

	buf := prompt.NewBuffer()
	buf.InsertText("root", false, true)
	doc := buf.Document()

	results := state.completer(*doc)
	if len(results) != 1 {
		t.Error("Should've found one result")
	}
}

func TestSetQueueExecutor(t *testing.T) {
	state := newCmdState()
	state.executor("set-queue")
	if state.livePrefix != "set-queue> " {
		t.Error("Live prefix should be set to set-queue> ")
	}
	state.executor("root.go")
	if len(state.currentQueue) != 1 {
		t.Error("Should've added an item to the queue")
	}
	if !strings.HasSuffix(state.currentQueue[0], "root.go") {
		t.Error("root.go should've been added to the queue")
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := test.NewMockPlayerClient(ctrl)
	matcher := test.NewMatcher(func(arg interface{}) bool {
		queue := arg.(*platune.QueueRequest).Queue
		return strings.HasSuffix(queue[0], "root.go")
	})
	mock.EXPECT().SetQueue(gomock.Any(), matcher)
	utils.Client = utils.NewTestClient(mock)
	state.executor("")
}
