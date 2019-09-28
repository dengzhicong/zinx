package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//IServer接口实现
type Server struct {
	//服务器名称
	Name string
	// ip版本
	IPVersion string
	//服务器监听的IP
	IP   string
	Port int
}

func (s *Server) Start() {
	fmt.Println("start server Listener at ip:%s,port %d,is startting", s.IP, s.Port)

	go func() {
		//获取一个tcp的ip
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
		}
		//监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listener error: ", s.IPVersion, "err ", err)
		}

		fmt.Println("start zinx server success, ", s.Name, " success,listening...")
		//阻塞等待客户端链接，处理客户端链接业务
		for {
			//如果有客户端连接过俩，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error: ", err)
				continue
			}

			//已经与客户端建立链接，做一些业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, e := conn.Read(buf)
					if e != nil {
						fmt.Println("rec buf err", err)
						continue
					}

					//回写客户端
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	//TODO 将服务器资源，状态，链接信息进行停止
}

func (s *Server) Serve() {
	//启动服务
	s.Start()

	//TODO 做一些启动访问之外的业务

	//阻塞状态
	select {}
}

//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
