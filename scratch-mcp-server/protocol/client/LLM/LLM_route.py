import os
from types import Dict
from openai import OpenAI

class LLM_Endpoint():

    def __init__(self):
        api_key = os.getenv("OPENAI_API_KEY")
        if not api_key:
            return ValueError("The API key environment variable is not set 💥")
        self.client = OpenAI(api_key=api_key)
        self.__input = None
        self.__instructions = None
        self.__reasoning = None


    def set_input(self,input: str):
        self.__input = input

    def set_instructions(self, instructions: str):
        self.__instructions = instructions

    def set_reasoning(self, reasoning: Dict[str,str]):
        self.__reasoning = reasoning

    def build_request(self):
        response = self.client.responses.create(
            model= 'gpt-3.5-turbo',
            input=self.__input,
            instructions=self.__instructions,
            reasoning=self.__reasoning,
        )
        
        return response.output_text
