from langchain.prompts import PromptTemplate,FewShotPromptTemplate

# 第一个简单样本直接展示了问题和答案，没有思考过程。
# 第二个样本还使用了 Chain of Thoughts 技巧，给出了找术语的过程。
examples = [
    {
        "question": "\"提出的问题越舒服，得到的答案就越有价值，\
这一点使用过 ChatGPT 的读者一定有感受。\"\
这段文字中有哪些术语？",
        "answer": "\"ChatGPT\""
    },
    {
        "question": "\"GPT 提示工程（Prompt Engineering）是一门较新的学科\"\
这段文字中有哪些术语？",
        "answer":
        """这个问题可以如下分析找到答案：
1. 找到这段文字中"（）"中间的英文词组，如"Prompt Engineering"，这个词组就是术语；
2. 找到这段文字中间的英文缩写，如"GPT"，这个缩写就是术语；
3. 术语不包含任何中文字符，"提示工程（Prompt Engineering）"不是术语，"学科"也不是术语。
最终的答案就是："Prompt Engineering"、"GPT"
        """
    }
]

# 把样本组合成提示语一部分的模板 
example_prompt = PromptTemplate(
    input_variables=["question", "answer"],
    template="Question: {question}\n{answer}")

# FewShot 完整的提示语模板
fewshot_prompt = FewShotPromptTemplate(
    example_prompt=example_prompt, # 组合样本
    examples=examples, # 样本
    suffix="Question: \"{input} \"这段文字中有哪些术语？", # 替换真正要找出术语的文本
    input_variables=["input"]
)

# 打印出最终组合出来的提示语
print(fewshot_prompt.format(
    input="比如 OpenAI 2020 年发布的 gpt-3 用了 570GB 文本，\
它能生成的知识就冻结（Frozen）在了发布的 2020 年，2021 年之后发生的事一概不知，\
也只知道这 570GB 文本中的知识")
)
