package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	workDuration := 1 * time.Minute  // Продолжительность работы
	breakDuration := 1 * time.Minute // Продолжительность перерыва
	numPomodoros := 2                // Количество повторений?

	for i := 0; i < numPomodoros; i++ {
		fmt.Println("Starting work session...")
		startTimer(breakDuration)
		playSound("beep.wav")

		fmt.Println("Starting break session...")
		startTimer(breakDuration)
		playSound("beep.wav")
	}

	fmt.Println("All done! Good job!")
}

func startTimer(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
	fmt.Println("Time's up!")
}

func playSound(soundFile string) {
	f, err := os.Open(soundFile)
	if err != nil {
		panic(err)
	}

	s, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- true
	})))
	<-done
}
