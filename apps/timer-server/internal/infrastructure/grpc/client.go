package grpc

import (
	"context"
	"io"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/antonioducs/wyd/pkg/proto/gateway"
)

type PacketHandler func(ctx context.Context, sessionID uint32, payload []byte)

type Client struct {
	hubAddress string
	logger     *slog.Logger
	handler    PacketHandler

	sendChan chan *pb.Packet
}

func NewClient(addr string, logger *slog.Logger, handler PacketHandler) *Client {
	return &Client{
		hubAddress: addr,
		logger:     logger,
		handler:    handler,
		sendChan:   make(chan *pb.Packet, 4096),
	}
}

func (c *Client) SetHandler(handler PacketHandler) {
	c.handler = handler
}

func (c *Client) Start() {
	for {
		c.logger.Info("Tentando conectar ao Hub...", "addr", c.hubAddress)
		err := c.connectAndServe()
		if err != nil {
			c.logger.Error("Conexão com Hub perdida/falhou", "error", err)
		}

		time.Sleep(3 * time.Second)
	}
}

func (c *Client) Send(sessionID uint32, payload []byte) {
	c.sendChan <- &pb.Packet{
		SessionId: sessionID,
		Type:      pb.EventType_DATA,
		Payload:   payload,
	}
}

func (c *Client) SendKick(sessionID uint32) {
	c.sendChan <- &pb.Packet{
		SessionId: sessionID,
		Type:      pb.EventType_DISCONNECT,
	}
}

func (c *Client) connectAndServe() error {
	conn, err := grpc.NewClient(c.hubAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewGameGateClient(conn)
	stream, err := client.StreamPackets(context.Background())
	if err != nil {
		return err
	}

	c.logger.Info("✅ Conectado ao Hub com sucesso!")

	errChan := make(chan error, 1)
	go func() {
		for pkt := range c.sendChan {
			if err := stream.Send(pkt); err != nil {
				c.logger.Error("Erro ao enviar pacote para Hub", "error", err, "session", pkt.SessionId)
				errChan <- err
				return
			}
		}
	}()

	for {
		select {
		case err := <-errChan:
			return err
		default:
		}

		pkt, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch pkt.Type {
		case pb.EventType_CONNECT:
			c.logger.Info("Novo Jogador Conectado", "session", pkt.SessionId)
		case pb.EventType_DISCONNECT:
			c.logger.Info("Jogador Desconectou", "session", pkt.SessionId)
		case pb.EventType_DATA:
			c.handler(stream.Context(), pkt.SessionId, pkt.Payload)
		}
	}
}
