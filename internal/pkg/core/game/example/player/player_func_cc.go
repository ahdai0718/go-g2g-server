package player

// Observe .
func (player *Player) Observe() {

}

// CloseRoutine .
func (player *Player) CloseRoutine() {
	player.closeRoutineChannel <- true
}
