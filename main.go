package main

// Seat ..
type Seat int

// Bar ..
type Bar Seat

func main() {
}

// var maxCCU = 10
// var consumers = make(chan consumer)
// var hc = httpclient.NewClient()

// type consumer struct {
// 	resp chan httpclient.Response
// 	args httpclient.Args
// }

// func serveHTTPClient(consumers chan consumer) {
// 	for consum := range consumers {
// 		resp := hc.Do(context.TODO(), consum.args)
// 		consum.resp <- resp
// 	}
// }

// func init() {
// 	for i := 0; i < maxCCU; i++ {
// 		go serveHTTPClient(consumers)
// 	}
// }

// func main() {
// 	http.HandleFunc("/mem", func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(time.Second * 5)

// 		w.Write([]byte("mem"))
// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		now := time.Now()
// 		rs := []string{}
// 		l := make([]chan httpclient.Response, 0)
// 		for i := 0; i < 5; i++ {
// 			resp := make(chan httpclient.Response)
// 			args := httpclient.Args{
// 				URL:    "http://localhost:8001",
// 				Method: http.MethodGet,
// 			}
// 			consumers <- consumer{
// 				resp: resp,
// 				args: args,
// 			}
// 			l = append(l, resp)
// 		}
// 		for _, v := range l {
// 			b := <-v
// 			c := b.Body
// 			rs = append(rs, string(c))
// 		}

// 		fmt.Println("GET /:", time.Since(now))

// 		w.Write([]byte(strings.Join(rs, "\t")))
// 	})

// 	http.ListenAndServe(":8080", nil)
// }
