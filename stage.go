package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct {
	// background
	bg_top    *ebiten.Image
	bg_bottom *ebiten.Image
}

// NewGameはゲームの初期化 background
func newGame() (*game, error) {
	g := &game{}

	// 背景画像を読み込む
	top, _, err := ebitenutil.NewImageFromFile("bg_top.png")
	if err != nil {
		return nil, err
	}
	g.bg_top = top

	bottom, _, err := ebitenutil.NewImageFromFile("bg_bottom.png")
	if err != nil {
		return nil, err
	}
	g.bg_bottom = bottom
	return g, nil
}

// Updateメソッドは状態更新、呼び出し１回を１ティックと呼ぶ
func (g *game) Update() error {
	return nil
}

// Drawメソッドは画面描画、呼び出し１回を１フレームと呼ぶ
func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	// screen上に画像を描画
	// sw = screen width(画面の横幅)
	// sh = screen height(画面の高さ)
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bannerH := sh / 3       // 上部バナー：全体の1/3
	bottomH := sh - bannerH // 下部：残り2/3

	// 上（バナー、screenで縮小）
	// tw, th = bg_topのtop画像のwidth/height
	opTop := &ebiten.DrawImageOptions{}
	tw, th := g.bg_top.Bounds().Dx(), g.bg_top.Bounds().Dy()
	opTop.GeoM.Scale(float64(sw)/float64(tw), float64(bannerH)/float64(th))
	screen.DrawImage(g.bg_top, opTop)

	// 下（残り2/3に縮小してから、バナー分下にTranslateで平行移動）
	//　bw, bh = bg_bottomのbottom画像のwidth/height
	opBottom := &ebiten.DrawImageOptions{}
	bw, bh := g.bg_bottom.Bounds().Dx(), g.bg_bottom.Bounds().Dy()
	opBottom.GeoM.Scale(float64(sw)/float64(bw), float64(bottomH)/float64(bh))
	opBottom.GeoM.Translate(0, float64(bannerH))
	screen.DrawImage(g.bg_bottom, opBottom)
}

// Layout(レイアウト) メソッドで画面サイズ（幅と高さ）
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("CYBER GUARD")
	ebiten.SetWindowSize(1280, 720)

	g, err := newGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
