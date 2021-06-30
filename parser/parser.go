package parser

import (
	"bufio"
	"context"
	"fmt"
	"github.com/brandesign/arma-go-parser/command"
	"os"
)

type Handlers map[string]func() interface{}
type Listener func(evt interface{}) error

func NewParser(handlers Handlers) *Parser {
	p := &Parser{
		listeners: map[string][]Listener{},
		handlers:  handlers,
	}

	return p
}

type Parser struct {
	listeners map[string][]Listener
	handlers  Handlers
}

func (p *Parser) Run(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			if scanner.Scan() {
				l := NewLine(scanner.Text())
				fmt.Println("Got Line:", l.String())
				if err := p.handleLine(l); err != nil {
					fmt.Printf("cannot handle Line: %s, error: %v\n", l, err)
				}
			}
		}
	}
}

func (p *Parser) Subscribe(subscriptions ...*Subscription) error {
	for _, s := range subscriptions {
		if err := p.addListener(s.name, s.listener); err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) addListener(evtName string, l Listener) error {
	ls, ok := p.listeners[evtName]
	if !ok {
		p.listeners[evtName] = []Listener{l}

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

func NewSubscription(evt string, l Listener) *Subscription {
	return &Subscription{
		name:     evt,
		listener: l,
	}
}

type Subscription struct {
	name     string
	listener Listener
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
