package main
// Duck typing: 오리처럼 행동하면 오리라고 간주한다.

// type interface_name interface { method_name(), method_name(parm_type, ...)}
// interface는 규격이다.
// 업케스팅, 다운케스팅
import (
	"week13/gadget"
)

type Player interface {
	// abstract method
	Play(string)
	Stop()
}
func TryOut(player Player) {
	player.Play("데모 트랙")
	player.Stop()

	// Player interface에 Record method가 구현되어 있지 않다.
	// player.Record()

	// 인터페이스는 기본적으로 일반 형변환을 지원하지 않는다.
	// recorder := gadget.TapePlayer(player)
	// recorder.Record()

	// 인터페이스만의 형변환 문법을 적용
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})

	//p1 := gadget.TapeRecorder{}
	//p2 := gadget.Mp3Player{}
	//
	//mix_tape1 := []string{"난 알아요", "거짓말", "으르렁"}
	//mix_tape2 := []string{"라캄파넬라", "G선상의 아리아"}
	//
	//playList(p1, mix_tape1)
	//
	//playList(p2, mix_tape2)
}
