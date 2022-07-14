package discord

import (
	"errors"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Discord struct {
	Logger *zap.Logger
	Token  string
	Guild  string // For development purpose

	s *discordgo.Session
}

func New(t string, opts ...Option) *Discord {
	discord := &Discord{
		Logger: zap.NewNop(),
		Token:  t,
	}

	for _, opt := range opts {
		opt(discord)
	}

	return discord
}

type Option func(*Discord)

func WithLogger(l *zap.Logger) Option {
	return func(d *Discord) {
		d.Logger = l
	}
}

func WithGuild(g string) Option {
	return func(d *Discord) {
		d.Guild = g
	}
}

func (d *Discord) init() error {
	if d.Token == "" {
		return errors.New("missing discord token")
	}

	s, err := discordgo.New("Bot " + d.Token)
	if err != nil {
		return err
	}

	d.s = s

	return nil
}

func (d *Discord) Run() error {

	// Create discord session
	if err := d.init(); err != nil {
		return err
	}

	// Register actions and commands to our discord bot
	if err := d.register(); err != nil {
		return err
	}

	// Open discord connection
	if err := d.open(); err != nil {
		return err
	}
	defer d.close()

	// TODO: API

	// Waiting to close the program
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	return nil
}

func (d *Discord) register() error {

	// TODO:

	return nil
}

func (d *Discord) open() error {
	return d.s.Open()
}

func (d *Discord) close() {
	d.s.Close()
}
