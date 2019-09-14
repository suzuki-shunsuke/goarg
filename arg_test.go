package goarg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	NewHandler(func(args ...string) (interface{}, error) {
		return nil, nil
	}, "foo")
}

func TestNewParser(t *testing.T) {
	NewParser(
		NewHandler(func(args ...string) (interface{}, error) {
			return nil, nil
		}, "foo"))
}

func TestHandlerMatch(t *testing.T) {
	h := NewHandler(func(args ...string) (interface{}, error) {
		return nil, nil
	}, "foo", "bar")
	s, b := h.Match("bar")
	require.Nil(t, s)
	require.False(t, b)
	s, b = h.Match("foo", "zoo")
	require.Nil(t, s)
	require.False(t, b)
	s, b = h.Match("foo", "bar", "zoo")
	require.Equal(t, s, []string{"zoo"})
	require.True(t, b)
}

func TestParserAdd(t *testing.T) {
	p := NewParser(
		NewHandler(func(args ...string) (interface{}, error) {
			return nil, nil
		}, "foo"))
	require.Equal(t, 1, len(p.handlers))
	q := p.Add(func(args ...string) (interface{}, error) {
		return nil, nil
	}, "foo")
	require.Equal(t, &p, q)
	var r *Parser
	r.Add(func(args ...string) (interface{}, error) {
		return nil, nil
	}, "foo")
	require.Nil(t, r)
}

func TestParserParse(t *testing.T) {
	p := NewParser(
		NewHandler(func(args ...string) (interface{}, error) {
			if len(args) == 0 {
				return nil, fmt.Errorf("args is required")
			}
			return nil, nil
		}, "git", "add"))
	_, b, err := p.Parse("git")
	require.Nil(t, err)
	require.False(t, b)
	_, b, err = p.Parse("git", "add", "foo.txt")
	require.Nil(t, err)
	require.True(t, b)
	var q *Parser
	_, b, err = q.Parse("git", "add", "foo.txt")
	require.Nil(t, err)
	require.False(t, b)
}
