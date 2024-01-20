package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/nobonobo/wrc-codriver/codriver"
	"github.com/nobonobo/wrc-codriver/easportswrc"
	"github.com/nobonobo/wrc-codriver/logger"
)

type Config struct {
	Listen       string `json:"listen"`
	Forward      string `json:"forward"`
	LoggerEnable bool   `json:"logger_enable"`
}

var (
	config = Config{
		Listen:       "127.0.0.1:20777",
		Forward:      "",
		LoggerEnable: false,
	}
	dest *net.UDPAddr
)

func init() {
	flag.StringVar(&config.Listen, "listen", config.Listen, "listen address")
	flag.StringVar(&config.Forward, "forward", config.Forward, "forward address")
	flag.BoolVar(&config.LoggerEnable, "logger", config.LoggerEnable, "enable logger")
}

func receiveryProc(ctx context.Context) (<-chan *easportswrc.PacketEASportsWRC, error) {
	if config.Forward != "" {
		addr, err := net.ResolveUDPAddr("udp", config.Forward)
		if err != nil {
			return nil, err
		}
		dest = addr
	}
	conn, err := net.ListenPacket("udp", config.Listen)
	if err != nil {
		return nil, err
	}
	ch := make(chan *easportswrc.PacketEASportsWRC, 8)
	go func() {
		defer close(ch)
		buf := make([]byte, 4096)
		for {
			n, _, err := conn.ReadFrom(buf)
			if err != nil {
				log.Print(err)
				continue
			}
			if dest != nil {
				if _, err := conn.WriteTo(buf[:n], dest); err != nil {
					log.Print(err)
				}
			}
			if n != easportswrc.PacketEASportsWRCLength {
				continue
			}
			p := new(easportswrc.PacketEASportsWRC)
			p.UnmarshalBinary(buf[:n])
			ch <- p
		}
	}()
	go func() {
		<-ctx.Done()
		conn.Close()
	}()
	return ch, nil
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch, err := receiveryProc(ctx)
	if err != nil {
		log.Fatal(err)
	}
	funcs := []func(*easportswrc.PacketEASportsWRC) error{}
	funcs = append(funcs, codriver.Setup(ctx))
	if config.LoggerEnable {
		funcs = append(funcs, logger.Setup(ctx))
	}
	for {
		select {
		case <-ctx.Done():
			return
		case p := <-ch:
			for _, fn := range funcs {
				if err := fn(p); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
