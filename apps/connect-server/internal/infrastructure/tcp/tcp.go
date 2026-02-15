package tcp

import (
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/crypto"
	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/grpc"
	"github.com/antonioducs/wyd/connect-server/internal/session"
)

type TCPServer struct {
	Addr        string
	MaxConnChan chan struct{}
	Logger      *slog.Logger
}

type TCPServerOptions struct {
	Host    string
	Port    string
	MaxConn int
	Logger  *slog.Logger
}

func NewTCPServer(options TCPServerOptions) *TCPServer {
	return &TCPServer{
		Addr:        fmt.Sprintf("%s:%s", options.Host, options.Port),
		MaxConnChan: make(chan struct{}, options.MaxConn),
		Logger:      options.Logger,
	}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	s.Logger.Info("Servidor TCP iniciado", "addr", s.Addr, "limit", cap(s.MaxConnChan))

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.Logger.Error("Erro ao aceitar conexão", "error", err)
			continue
		}

		s.MaxConnChan <- struct{}{}

		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	sess := session.Global.Add(conn)

	s.Logger.Info("Nova conexão", "id", sess.ID, "remote", conn.RemoteAddr())

	defer func() {
		s.Logger.Info("Conexão encerrada", "id", sess.ID)
		session.Global.Remove(sess.ID)
		<-s.MaxConnChan
		conn.Close()
	}()

	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				s.Logger.Error("Erro na leitura", "error", err)
			}
			return
		}

		if n <= 4 {
			continue
		}

		decryptedData, err := crypto.Decrypt(buffer)
		if err != nil {
			s.Logger.Error("Erro ao descriptografar", "error", err)
			continue
		}

		if grpc.Hub != nil {
			grpc.Hub.SendToTimer(sess.ID, decryptedData)
		} else {
			s.Logger.Warn("Hub gRPC não iniciado, pacote perdido", "id", sess.ID)
		}
	}
}
