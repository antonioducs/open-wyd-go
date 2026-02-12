package tcp

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"unsafe"

	"github.com/antonioducs/wyd/protocol"
	"github.com/antonioducs/wyd/protocol/crypto"
	"github.com/antonioducs/wyd/protocol/incoming"
	"github.com/antonioducs/wyd/protocol/outgoing"
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
	defer func() {
		<-s.MaxConnChan
		conn.Close()
		s.Logger.Info("Conexão encerrada", "remote", conn.RemoteAddr())
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

		currentData := buffer[:n]

		if n < 4 {
			continue
		}

		if n == 120 {
			currentData = currentData[4:120]
		}

		decryptedData, err := crypto.Decrypt(currentData)
		if err != nil {
			s.Logger.Error("Erro ao descriptografar pacote", "error", err)
			continue
		}

		header := (*protocol.PacketHeader)(unsafe.Pointer(&decryptedData[0]))

		switch header.PacketID {
		case protocol.PackageIDLogin:
			login := (*incoming.Login)(unsafe.Pointer(&decryptedData[0]))
			s.Logger.Info("Pacote de login recebido",
				"remote", conn.RemoteAddr(),
				"packetId", fmt.Sprintf("%x", login.Header.PacketID))
			s.Logger.Info("Senha", "password", login.GetPassword())
			s.Logger.Info("Usuário", "username", login.GetUsername())
			message := outgoing.NewMessage("Hello, world!")
			conn.Write(message.Header.PrepareToSend())
		}
	}
}
