package waitgrouputil

var done = make(chan struct{})

func (w waitGroupWrapper) Wait() {
	go func() {
		w.waitGroup.Wait()
		done <- struct{}{}
	}()
}

func MainWait() {

	WaitGroup.Wait()

	// select {
	// case <-done:

	// case <-time.After(10 * time.Second):
	// 	fmt.Println("wait group 超时")
	// }

	<-done

}
