from langchain.llms import OpenAI 
from langchain.chains import ConversationChain
from langchain.agents import initialize_agent
from langchain.agents import load_tools
from langchain.chains import LLMChain
from langchain.prompts import PromptTemplate

llm = OpenAI(model_name="text-davinci-003", temperature=0.9, max_tokens=100)

text = "What would be a good company name for a company that makes colorful socks?"

# print(llm(text))

# manage LLMs prompt templates

prompt = PromptTemplate(
        input_variables=["product"],
        template="What would be a good company name for a company that makes {product}?",
)

chain = LLMChain(llm=llm, prompt=prompt)
# print(chain.run(product="colorful socks"))

# Agents 
# 基于与用户输入动态地调用chains 

tools = load_tools(["llm-math"], llm=llm)
agent = initialize_agent(
        tools=tools,
        llm=llm,
        agent="zero-shot-react-description",
        verbose=True,)
text = "12 raised to the power of 3 and result raised to 2 power?"


#  print (agent.run(text))

# Memory


llm = OpenAI(temperature=0)

conversation = ConversationChain(llm=llm, verbose=True)

conversation.predict(input="Hi there, how are you?")
conversation.predict(input="I'm doing well, Just having a chat with you.")

