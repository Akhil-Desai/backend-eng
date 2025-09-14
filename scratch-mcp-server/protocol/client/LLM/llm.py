import os
from typing import Dict
from openai import OpenAI
from dotenv import load_dotenv

load_dotenv()

class LLM():

    def __init__(self):
        api_key = os.getenv("OPENAI_API_KEY")
        if not api_key:
            raise ValueError("The API key environment variable is not set 💥")
        self.client = OpenAI(api_key=api_key)
        self.__input = None
        self.__instructions = None
        self.__reasoning = None


    def set_input(self,input: str):
        self.__input = input

    def set_instructions(self, instructions: str):
        self.__instructions = instructions

    def set_reasoning(self, reasoning: Dict[str,str]):
        """
        Reasoning is only supported in certain models, the base gpt 3.5 model does not support please do not set
        """
        self.__reasoning = reasoning

    def build_request(self):

        response = self.client.responses.create(
            model= 'gpt-3.5-turbo',
            input=self.__input,
            instructions=self.__instructions,

        )

        return response.output_text
