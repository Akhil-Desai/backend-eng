import unittest
from protocol.client.llm import llm


class TestLLMCall(unittest.TestCase):

    def setUp(self):
        self.LLM_Client = llm.LLM()

    def test_response(self):

        self.LLM_Client.set_input("I'm testing that this API works, respond with Hey!")
        self.LLM_Client.set_instructions('')

        response = self.LLM_Client.build_request()

        print("I am broke and can't afford the OpenAI API (will buy $2.00 of credits to test soon)")
        self.assertIsNotNone(response)
        self.assertIsInstance(response,str)
