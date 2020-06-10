package main

import "github.com/tadvi/winc"

func main() {

	var Playeur bool = true
	var play [10]bool
	var PlayeurD [10]int

	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("MorpioZiGame")

	/*WinWindow := winc.NewForm(nil) //
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("MorpioZiGame") */

	btn := winc.NewPushButton(mainWindow) // Bouton close
	btn.SetText("close")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)

	btn1 := winc.NewPushButton(mainWindow) // Bouton 1
	btn1.SetText("-")
	btn1.SetPos(0, 0)
	btn1.SetSize(40, 40)
	btn1.OnClick().Bind(func(e *winc.Event) {
		if play[1] == false {
			if Playeur == true {
				btn1.SetText("X")
				Playeur = false
				PlayeurD[1] = 1
			} else {
				btn1.SetText("O")
				Playeur = true
				PlayeurD[1] = 2
			}
			play[1] = true
		}
		if (PlayeurD[1] == 1 && PlayeurD[2] == 1 && PlayeurD[3] == 1) || (PlayeurD[1] == 1 && PlayeurD[5] == 1 && PlayeurD[9] == 1) || (PlayeurD[1] == 1 && PlayeurD[4] == 1 && PlayeurD[7] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[1] == 2 && PlayeurD[2] == 2 && PlayeurD[3] == 2) || (PlayeurD[1] == 2 && PlayeurD[5] == 2 && PlayeurD[9] == 2) || (PlayeurD[1] == 2 && PlayeurD[4] == 2 && PlayeurD[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn2 := winc.NewPushButton(mainWindow) // Bouton 2
	btn2.SetText("-")
	btn2.SetPos(40, 0)
	btn2.SetSize(40, 40)
	btn2.OnClick().Bind(func(e *winc.Event) {
		if play[2] == false {
			if Playeur == true {
				btn2.SetText("X")
				Playeur = false
				PlayeurD[2] = 1
			} else {
				btn2.SetText("O")
				Playeur = true
				PlayeurD[2] = 2
			}
			play[2] = true
		}
		if (PlayeurD[1] == 1 && PlayeurD[2] == 1 && PlayeurD[3] == 1) || (PlayeurD[2] == 1 && PlayeurD[5] == 1 && PlayeurD[8] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[1] == 2 && PlayeurD[2] == 2 && PlayeurD[3] == 2) || (PlayeurD[2] == 2 && PlayeurD[5] == 2 && PlayeurD[8] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn3 := winc.NewPushButton(mainWindow) // Bouton 3
	btn3.SetText("-")
	btn3.SetPos(80, 0)
	btn3.SetSize(40, 40)
	btn3.OnClick().Bind(func(e *winc.Event) {
		if play[3] == false {
			if Playeur == true {
				btn3.SetText("X")
				Playeur = false
				PlayeurD[3] = 1
			} else {
				btn3.SetText("O")
				Playeur = true
				PlayeurD[3] = 2
			}
			play[3] = true
		}
		if (PlayeurD[1] == 1 && PlayeurD[2] == 1 && PlayeurD[3] == 1) || (PlayeurD[3] == 1 && PlayeurD[5] == 1 && PlayeurD[7] == 1) || (PlayeurD[3] == 1 && PlayeurD[6] == 1 && PlayeurD[9] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[1] == 2 && PlayeurD[2] == 2 && PlayeurD[3] == 2) || (PlayeurD[3] == 2 && PlayeurD[5] == 2 && PlayeurD[7] == 2) || (PlayeurD[3] == 2 && PlayeurD[6] == 2 && PlayeurD[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn4 := winc.NewPushButton(mainWindow) // Bouton 4
	btn4.SetText("-")
	btn4.SetPos(0, 40)
	btn4.SetSize(40, 40)
	btn4.OnClick().Bind(func(e *winc.Event) {
		if play[4] == false {
			if Playeur == true {
				btn4.SetText("X")
				Playeur = false
				PlayeurD[4] = 1
			} else {
				btn4.SetText("O")
				Playeur = true
				PlayeurD[4] = 2
			}
			play[4] = true
		}
		if (PlayeurD[4] == 1 && PlayeurD[5] == 1 && PlayeurD[6] == 1) || (PlayeurD[1] == 1 && PlayeurD[4] == 1 && PlayeurD[7] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[4] == 2 && PlayeurD[5] == 2 && PlayeurD[6] == 2) || (PlayeurD[1] == 2 && PlayeurD[4] == 2 && PlayeurD[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn5 := winc.NewPushButton(mainWindow) // Bouton 5
	btn5.SetText("-")
	btn5.SetPos(40, 40)
	btn5.SetSize(40, 40)
	btn5.OnClick().Bind(func(e *winc.Event) {
		if play[5] == false {
			if Playeur == true {
				btn5.SetText("X")
				Playeur = false
				PlayeurD[5] = 1
			} else {
				btn5.SetText("O")
				Playeur = true
				PlayeurD[5] = 2
			}
			play[5] = true
		}
		if (PlayeurD[4] == 1 && PlayeurD[5] == 1 && PlayeurD[6] == 1) || (PlayeurD[1] == 1 && PlayeurD[5] == 1 && PlayeurD[9] == 1) || (PlayeurD[2] == 1 && PlayeurD[5] == 1 && PlayeurD[8] == 1) || (PlayeurD[3] == 1 && PlayeurD[5] == 1 && PlayeurD[7] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[4] == 2 && PlayeurD[5] == 2 && PlayeurD[6] == 2) || (PlayeurD[1] == 2 && PlayeurD[5] == 2 && PlayeurD[9] == 2) || (PlayeurD[2] == 2 && PlayeurD[5] == 2 && PlayeurD[8] == 2) || (PlayeurD[3] == 1 && PlayeurD[5] == 1 && PlayeurD[7] == 1) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn6 := winc.NewPushButton(mainWindow) // Bouton 6
	btn6.SetText("-")
	btn6.SetPos(80, 40)
	btn6.SetSize(40, 40)
	btn6.OnClick().Bind(func(e *winc.Event) {
		if play[6] == false {
			if Playeur == true {
				btn6.SetText("X")
				Playeur = false
				PlayeurD[6] = 1
			} else {
				btn6.SetText("O")
				Playeur = true
				PlayeurD[6] = 2
			}
			play[6] = true
		}
		if (PlayeurD[4] == 1 && PlayeurD[5] == 1 && PlayeurD[6] == 1) || (PlayeurD[3] == 1 && PlayeurD[6] == 1 && PlayeurD[9] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[4] == 2 && PlayeurD[5] == 2 && PlayeurD[6] == 2) || (PlayeurD[3] == 2 && PlayeurD[6] == 2 && PlayeurD[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn7 := winc.NewPushButton(mainWindow) // Bouton 7
	btn7.SetText("-")
	btn7.SetPos(0, 80)
	btn7.SetSize(40, 40)
	btn7.OnClick().Bind(func(e *winc.Event) {
		if play[7] == false {
			if Playeur == true {
				btn7.SetText("X")
				Playeur = false
				PlayeurD[7] = 1
			} else {
				btn7.SetText("O")
				Playeur = true
				PlayeurD[7] = 2
			}
			play[7] = true
		}
		if (PlayeurD[7] == 1 && PlayeurD[8] == 1 && PlayeurD[9] == 1) || (PlayeurD[7] == 1 && PlayeurD[5] == 1 && PlayeurD[3] == 1) || (PlayeurD[1] == 1 && PlayeurD[4] == 1 && PlayeurD[7] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[7] == 2 && PlayeurD[8] == 2 && PlayeurD[9] == 2) || (PlayeurD[3] == 2 && PlayeurD[5] == 2 && PlayeurD[7] == 2) || (PlayeurD[1] == 2 && PlayeurD[4] == 2 && PlayeurD[7] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn8 := winc.NewPushButton(mainWindow) // Bouton 8
	btn8.SetText("-")
	btn8.SetPos(40, 80)
	btn8.SetSize(40, 40)
	btn8.OnClick().Bind(func(e *winc.Event) {
		if play[8] == false {
			if Playeur == true {
				btn8.SetText("X")
				Playeur = false
				PlayeurD[8] = 1
			} else {
				btn8.SetText("O")
				Playeur = true
				PlayeurD[8] = 2
			}
			play[8] = true
		}
		if (PlayeurD[7] == 1 && PlayeurD[8] == 1 && PlayeurD[9] == 1) || (PlayeurD[2] == 1 && PlayeurD[5] == 1 && PlayeurD[8] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[7] == 2 && PlayeurD[8] == 2 && PlayeurD[9] == 2) || (PlayeurD[2] == 2 && PlayeurD[5] == 2 && PlayeurD[8] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	btn9 := winc.NewPushButton(mainWindow) // Bouton 9
	btn9.SetText("-")
	btn9.SetPos(80, 80)
	btn9.SetSize(40, 40)
	btn9.OnClick().Bind(func(e *winc.Event) {
		if play[9] == false {
			if Playeur == true {
				btn9.SetText("X")
				Playeur = false
				PlayeurD[9] = 1
			} else {
				btn9.SetText("O")
				Playeur = true
				PlayeurD[9] = 2
			}
			play[9] = true
		}
		if (PlayeurD[7] == 1 && PlayeurD[8] == 1 && PlayeurD[9] == 1) || (PlayeurD[1] == 1 && PlayeurD[5] == 1 && PlayeurD[9] == 1) || (PlayeurD[3] == 1 && PlayeurD[6] == 1 && PlayeurD[9] == 1) {
			mainWindow.Close()
			WinPlayeur()
		} else if (PlayeurD[7] == 2 && PlayeurD[8] == 2 && PlayeurD[9] == 2) || (PlayeurD[1] == 2 && PlayeurD[5] == 2 && PlayeurD[9] == 2) || (PlayeurD[3] == 2 && PlayeurD[6] == 2 && PlayeurD[9] == 2) {
			mainWindow.Close()
			WinJ2()
		}
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
}
func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func WinPlayeur() {
	WinWindow := winc.NewForm(nil) // Playeur 1 Gagne ez
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("Playeur 1 WIN")
	WinWindow.Show()
	WinWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

	btn := winc.NewPushButton(WinWindow) // Bouton pour close la fenêtre
	btn.SetText("close")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
}

func WinJ2() {
	WinWindow := winc.NewForm(nil) // Playeur 2 Gagne ez
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("Playeur 2 WIN")
	WinWindow.Show()
	WinWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

	btn := winc.NewPushButton(WinWindow) // Bouton pour close la fenêtre
	btn.SetText("close")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
}
