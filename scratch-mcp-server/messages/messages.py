from dataclasses import dataclass
from typing import Optinal, Any, Dict, Literal
from .types import ResponeType


@dataclass
class Response:
    jsonrpc: str
    id: int | str
    type: ResponeType

@dataclass
class Request:
    jsonrpc: str
    id: int | str
    required: Response


@dataclass
class Notification:
    jsonrpc: str
    id: int | str
    required: None
