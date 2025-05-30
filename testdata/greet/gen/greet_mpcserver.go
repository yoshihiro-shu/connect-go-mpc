// Code generated by connect-go-mpc DO NOT EDIT.
package greetv1_connectgompc

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	connectgompc "github.com/yoshihiro-shu/connect-go-mpc"
)

// NewMCPServerWithTools は設定済みの GreetService MCP サーバーを生成して返します
func NewGreetServiceMCPServer(baseURL string, opts ...connectgompc.ClientOption) *mcp.Server {
	server := server.NewMCPServer("GreetService", "1.0.0")

	toolHandler := connectgompc.NewToolHandler(baseURL, opts...)
	server.AddTool(
		mcp.NewTool("挨拶",
			mcp.WithDescription("挨拶リクエスト"),
			mcp.WithString("name"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return toolHandler.Handle(ctx, req, "Greet")
		},
	)

	server.AddTool(
		mcp.NewTool("ピン",
			mcp.WithDescription("ピンリクエスト"),
			mcp.WithString("message"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return toolHandler.Handle(ctx, req, "Ping")
		},
	)

	return server
}
