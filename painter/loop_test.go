package painter

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"
	"runtime"
	"testing"

	"golang.org/x/exp/shiny/screen"
)

func TestLoopPostAndUpdate(t *testing.T) {
	var (
		l  Loop
		tr testReceiver
	)
	l.Receiver = &tr

	l.Start(mockScreen{})
	l.Post(&WhiteFill{})
	l.Post(&GreenFill{})
	l.Post(UpdateOp)
	l.StopAndWait()

	if tr.lastTexture == nil {
		t.Fatal("Texture was not updated")
	}
	mt, ok := tr.lastTexture.(*mockTexture)
	if !ok {
		t.Fatal("Unexpected texture", tr.lastTexture)
	}
	if mt.Colors[0] != color.White {
		t.Error("First color is not white:", mt.Colors)
	}
	if len(mt.Colors) != 2 {
		t.Error("Unexpected number of colors:", mt.Colors)
	}
}

func TestLoopQueueWait(t *testing.T) {
	var (
		l  Loop
		tr testReceiver
	)
	l.Receiver = &tr

	l.Start(mockScreen{})
	runtime.Gosched()
	l.StopAndWait()
}

func TestLoopQueueSeq(t *testing.T) {
	var (
		l      Loop
		tr     testReceiver
		gotSeq []string
	)
	l.Receiver = &tr

	l.Start(mockScreen{})
	l.Post(OperationFunc(func(screen.Texture) {
		gotSeq = append(gotSeq, "Operation 1")
		l.Post(OperationFunc(func(screen.Texture) {
			gotSeq = append(gotSeq, "Operation 3")
		}))
	}))
	l.Post(OperationFunc(func(screen.Texture) {
		gotSeq = append(gotSeq, "Operation 2")
	}))
	l.Post(UpdateOp)
	l.StopAndWait()

	wantSeq := []string{"Operation 1", "Operation 2", "Operation 3"}
	if !reflect.DeepEqual(gotSeq, wantSeq) {
		t.Errorf("Operation sequence mismatch:\n got=%v\n want=%v", gotSeq, wantSeq)
	}
}

type testReceiver struct {
	lastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.lastTexture = t
}

type mockScreen struct{}

func (m mockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	panic("implement me")
}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	panic("implement me")
}

type mockTexture struct {
	Colors []color.Color
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point { return size }

func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: m.Size()}
}

func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {}

func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.Colors = append(m.Colors, src)
}
