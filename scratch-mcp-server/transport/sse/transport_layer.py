import httpx
import json
from typing import AsyncIterator

class SSETransport():

    def __init__(self, server_addr: str):
        self.server_addr = server_addr
        self.client = httpx.AsyncClient(timeout=None)


    async def send_request(self, data: str) -> str:

        """
        Send JSON-RPC 2.0 request to server via POST.
        Used for client -> server communication.
        """

        response = await self.client.post(
            f"{self.server_addr}/request",
            content = data,
            headers={"Content-Type":"application/json"},
        )

        return response.text

    async def listen_events(self) -> AsyncIterator[str]:

        """
        Listen to server events via long-lived HTTP connection (SSE).
        Used for server -> client communication.
        """

        async with self.client.stream('GET', f"{self.server_addr}/sse") as response:
            async for line in response.aiter_lines():
                if line.startswith("data: "):
                    yield line[6:]

    async def close(self):
        """
        Close the HTTP Client
        """
        await self.client.aclose()
