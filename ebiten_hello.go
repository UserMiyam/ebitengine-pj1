package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct{}

// Updateメソッドは状態更新、呼び出し１回を１ティックと呼ぶ
func (g *game) Update() error {
	return nil
}

// Drawメソッドは画面描画、呼び出し１回を１フレームと呼ぶ
func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout(レイアウト) メソッドで画面サイズ（幅と高さ）
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowTitle("Ebiten Hello World")
	ebiten.SetWindowSize(640, 480)

	g := &game{}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}

}
