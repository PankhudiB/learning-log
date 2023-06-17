### Golang channel exercises    

Solving golang problems based on channels 

### Problem : Divide and Conquer to 2 goroutines to compute sum

Solving : https://go.dev/tour/concurrency/2

`spin 2 goroutines that compute each half of the input array
and total computaiton waits for the 2 goroutines for respective results on channel`

```go

func main() {
	input := []int{1, 1, 1, 1, 2, 2, 2, 2}

	result := make(chan int)
	go sum(input[:input[len(input)/2]], result)
	go sum(input[input[len(input)/2]:], result)

	oneHalfResult := <-result
	otherHalfResult := <-result

	fmt.Println("Total = ", oneHalfResult+otherHalfResult)
}

func sum(arr []int, result chan int) {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	result <- sum
}

```

### Problem: Equivalent Binary Trees

Solving : https://go.dev/tour/concurrency/8

initializing the tree :
```go 
func main() {
	t5 := &tree.Tree{Value: 5}
	t3 := &tree.Tree{Value: 3}
	t4 := &tree.Tree{Value: 4, Left: t3, Right: t5}
	t1 := &tree.Tree{Value: 1}
	t2 := &tree.Tree{Value: 2, Left: t1, Right: t4}

	fmt.Println("isSame : ", Same(t2, t2))
	
	fmt.Println("isSame : ", Same(t2, t4))
}
```

Walk -> similar to how one would write Binary Search Tree `InOrder` traversal 

The only difference here - instead of printing value -- keep sending them to channel: 

```go
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return  
	}
	if t.Left == nil {
		
		ch <- t.Value           //root if left is null 
		if t.Right != nil {     // go on to process right of such root
			Walk(t.Right, ch)
		}
		return
		
	} else {
	    Walk(t.Left, ch)
	}
	
	ch <- t.Value       //root
	
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}
```

Need a wrapper to close the channel post traversal completion - else runtime error of deadline!

```go
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}
```

The isSame func spins 2 goroutines to walk respective trees

```go
func Same(t1, t2 *tree.Tree) bool {
	var isSame bool     //for storing result

	c1 := make(chan int)        
	c2 := make(chan int)

	go Walker(t1, c1)
	go Walker(t2, c2)

    //a comparator goroutine that listens on both the channel and compares them 
    //- everytime both channel receive 1 1 item each
    
    //need done channel - to close of this comparator goroutine
    
	done := make(chan bool)
	go func() {             
		
		for i1 := range c1 {    // IMP
			i2 := <-c2          // IMP
			
			if i1 == i2 {
				isSame = true
			} else {
				isSame = false
				break
			}
		}
		
		done <- true
	}()
	
	<-done
	return isSame
}
```
