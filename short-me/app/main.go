package main

func main() {
	var (
		a *App
	)
	a = &App{
	}
	a.Initialize(getEnv())
	a.Run("127.0.0.1:8888")

}
