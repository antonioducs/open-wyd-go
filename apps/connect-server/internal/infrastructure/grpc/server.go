package grpc

import (
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/crypto"
	"github.com/antonioducs/wyd/connect-server/internal/session"
	pb "github.com/antonioducs/wyd/pkg/proto/gateway"
)

type Server struct {
	pb.UnimplementedGameGateServer

	toTimerChan chan *pb.Packet

	mu             sync.RWMutex
	timerConnected bool
}

var Hub *Server

func Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("❌ Erro fatal no gRPC Listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	Hub = &Server{
		toTimerChan: make(chan *pb.Packet, 10000),
	}

	pb.RegisterGameGateServer(grpcServer, Hub)

	log.Printf("🚀 gRPC Hub ouvindo em %s (Esperando Timer-Server...)", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Erro fatal no gRPC Serve: %v", err)
	}
}

func (s *Server) StreamPackets(stream pb.GameGate_StreamPacketsServer) error {
	log.Println("✅ Timer-Server CONECTADO e pronto para trabalhar!")

	s.mu.Lock()
	s.timerConnected = true
	s.mu.Unlock()

	defer func() {
		log.Println("⚠️ Timer-Server DESCONECTADO! O Hub continua de pé.")
		s.mu.Lock()
		s.timerConnected = false
		s.mu.Unlock()
	}()

	go func() {
		for pkt := range s.toTimerChan {
			if err := stream.Send(pkt); err != nil {
				log.Printf("Erro ao enviar para Timer: %v", err)
				return
			}
		}
	}()

	for {
		pkt, err := stream.Recv()
		if err != nil {
			return err
		}

		handlePacketFromTimer(pkt)
	}
}

func (s *Server) SendToTimer(sessionID uint32, payload []byte) {
	s.mu.RLock()
	connected := s.timerConnected
	s.mu.RUnlock()

	if !connected {
		return
	}

	s.toTimerChan <- &pb.Packet{
		SessionId: sessionID,
		Type:      pb.EventType_DATA,
		Payload:   payload,
	}
}

func handlePacketFromTimer(pkt *pb.Packet) {
	log.Printf("📥 Pacote recebido do Timer: session=%d type=%v", pkt.SessionId, pkt.Type)

	sess := session.Global.Get(pkt.SessionId)
	if sess == nil {
		log.Printf("⚠️ Sessão %d não encontrada - pacote ignorado", pkt.SessionId)
		return
	}

	if pkt.Type == pb.EventType_DISCONNECT {
		log.Printf("🔌 Timer pediu Kick da Sessão %d", pkt.SessionId)
		sess.Conn.Close()
		return
	}

	dataToSend := pkt.Payload
	encryptedData := crypto.Encrypt(dataToSend)

	_, err := sess.Conn.Write(encryptedData)
	if err != nil {
		log.Printf("Erro ao escrever no TCP da sessão %d: %v", pkt.SessionId, err)
	}
}
