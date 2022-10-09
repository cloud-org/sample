// author: ashing
// time: 2020/4/6 8:37 下午
// mail: axingfly@gmail.com
// Less is more.

package service

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}
