package parser

import (
	"bufio"
	"context"
	"fmt"
	"github.com/brandesign/arma-go-parser/command"
	"io"
)

type Handlers map[string]func() interface{}

type Subscriber interface {
	GetSubscriptions() []*Subscription
}

func NewParser(r io.Reader, handlers Handlers, subscribers ...Subscriber) (*Parser, error) {
	p := &Parser{
		reader:    r,
		listeners: map[string][]func(evt interface{}) error{},
		handlers:  handlers,
	}

	for _, subscriber := range subscribers {
		subs := subscriber.GetSubscriptions()
		for _, sub := range subs {
			if err := p.addListener(sub.name, sub.listener); err != nil {
				return nil, err
			}
		}
	}

	return p, nil
}

type Parser struct {
	reader    io.Reader
	listeners map[string][]func(evt interface{}) error
	handlers  Handlers
}

func (p *Parser) Run(ctx context.Context) {
	scanner := bufio.NewScanner(p.reader)
	for {
		select {
		case <-ctx.Done():
			command.Logf("context done")
			return
		default:
			if scanner.Scan() {
				l := NewLine(scanner.Text())
				if err := p.handleLine(l); err != nil {
					command.Logf("cannot handle Line: %s, error: %v\n", l, err)
				}
			}
		}
	}
}

func (p *Parser) addListener(evtName string, l func(evt interface{}) error) error {
	ls, ok := p.listeners[evtName]
	if !ok {
		p.listeners[evtName] = []func(evt interface{}) error{l}

		return command.Rawf("LADDERLOG_WRITE_%s 1", evtName)
	} else {
		p.listeners[evtName] = append(ls, l)
	}

	return nil
}

func (p *Parser) handleLine(l Line) error {
	listeners, ok := p.listeners[l.eventName]
	if !ok {
		return nil
	}

	evt, err := p.handlers.handleLine(l)
	if err != nil {
		return err
	}

	for _, ls := range listeners {
		if err := ls(evt); err != nil {
			return err
		}
	}

	return nil
}

func NewSubscription(evt string, l func(evt interface{}) error) *Subscription {
	return &Subscription{
		name:     evt,
		listener: l,
	}
}

type Subscription struct {
	name     string
	listener func(evt interface{}) error
}

func (h Handlers) handleLine(l Line) (interface{}, error) {
	f, ok := h[l.eventName]
	if !ok {
		return nil, fmt.Errorf("no event factory found for %s", l.eventName)
	}

	evt := f()
	if err := l.Scan(evt); err != nil {
		return nil, err
	}

	return evt, nil
}
