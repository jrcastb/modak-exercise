package main

func main() {
	srv := NewServer(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
