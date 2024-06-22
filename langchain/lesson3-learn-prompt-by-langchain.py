## 受《跟着LangChain 学Prompt》启发，自己动手写一个简单的学习Prompt的程序
## link: https://www.qinyu.info/post/learn-prompt-from-langchain/

from langchain.prompts import (
    ChatPromptTemplate,
    SystemMessagePromptTemplate,
    HumanMessagePromptTemplate,
)

# 给 System （即 AI）角色的指导
# 注意不是 AIMessagePromptTemplate（这是 AI 返回的消息）
system_prompt = SystemMessagePromptTemplate.from_template("\
				You are a experienced Tech Lead of a agile team. \
        If you don't know the answer, just say that you don't know. \
        Don't try to make up an answer.")
# 人要提的问题和要求，这里有模板参数，最后会被替换成真正的问题
human_prompt = HumanMessagePromptTemplate.from_template("\
        I have just joined your team and I am not familiar with agile practices. \
        Please tell me {question}.", 
        input_variables=["question"])
# 和 ChatOpenAI 交互的 ChatMessage，就是一组消息组合。
chat_prompt = ChatPromptTemplate.from_messages([
    system_prompt,
    human_prompt
])

# 最终组合后的提示词
prompt = chat_prompt.format_prompt(question="when to refactor my code")
print(prompt.to_string()) # Message 也可以转换成字符串消息
