# MCP-Server

**Source** (https://modelcontextprotocol.io/legacy/concepts/architecture)

## Overview

- **Hosts**: LLM applications that initiate connections
- **Clients**: Maintain 1:1 connections with servers, inside the host application
- **Servers**: Provide context, tools, and prompts to clients

## Core Components

- **Protocol Layer**: Handles message framing, request/response linking, and high-level communication patterns
  - Key classes: `Protocol`, `Client`, `Server`
- **Transport Layer**: Handles the actual communication between client and server. MCP supports multiple transport mechanisms (Stdio transport and streamable HTTP transport). For our use case, we will just support streamable HTTP transport. All transports use JSON-RPC 2.0 to exchange messages.
  - **Streamable HTTP transport**:
    - Uses HTTP with optional Server-Sent Events (SSE) for streaming (SSE is optional, but it's a way for the server to stream the response to the client)
    - HTTP POST for client-to-server messages

## Message Types

- **Request**: Expects a response from the other side
- **Result**: Successful responses to the request
- **Error**: Indicates the request failed
- **Notification**: One-way messages that don't expect a response

## Connection Lifecycle

1. **Initialization**
   - Client sends `initialize` request with protocol version and capabilities
   - Server responds with its protocol version and capabilities
   - Client sends `initialized` notification as acknowledgment
   - Normal message exchange begins

2. **Message Exchange**
   - After initialization, the following patterns are supported:
     - **Request-Response**: Client or server sends requests, the other responds
     - **Notifications**: Either party sends one-way messages

3. **Termination**
   - Either party can terminate the connection:
     - Clean shutdown via `close()`
     - Transport disconnection
     - Error conditions
