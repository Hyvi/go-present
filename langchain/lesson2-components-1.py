# langchain is 一个开发语言模型应用的框架
# langchain is a framework for developing language model applications
# advantage: 1. 它让你将外部数据，如文件、其他程序的输出、网络数据等，作为输入，以便你的程序可以处理这些数据
#         2. 通过代理功能，让你的语言模型能够与环境绩效你高兴交互，实现决策制定。


# 1. Langchain的组件

# 1.1 models
# 1.1.1 LanguageModel
# 1.1.2 Chat Model
# 1.1.3 Text Embedding Model

# 1.2 Prompts 
# 1.2.1 Prompt template, 
# 1.2.2 Example Selectors, 

from langchain.prompts.example_selector import SemanticSimilarityExampleSelector
from langchain.vectorstores import FAISS
# import OpenAI llms
from langchain.llms import OpenAI
from langchain.prompts import PromptTemplate, FewShotPromptTemplate
from langchain.embeddings import OpenAIEmbeddings
from langchain.output_parsers import ResponseSchema, StructuredOutputParser

llm = OpenAI(model_name="text-davinci-003")

#  example_prompt = PromptTemplate(
#          input_variables=["input_text", "output_text"],
#          template="Input: {input_text}\nOutput: {output_text}\n",
#          )
#
#  examples = [
#          {"input_text": "pirate", "output_text": "ship"},
#          {"input_text": "pilot", "output_text": "plane"},
#          {"input_text": "driver", "output_text": "car"},
#          {"input_text": "writer", "output_text": "book"},
#          {"input_text": "chef", "output_text": "restaurant"},
#          ]
#
#  example_selector = SemanticSimilarityExampleSelector.from_examples(
#          examples,
#          OpenAIEmbeddings(),
#          FAISS,
#          k=2
#          )
#  similarity_prompt = FewShotPromptTemplate(
#          example_selector=example_selector,
#          example_prompt=example_prompt,
#          prefix="Given the location an item is usually found in",
#          suffix="input: {noun}\noutput: ",
#          input_variables=["noun"],
#          )
#  my_noun = "student"
#  print(similarity_prompt.format(noun=my_noun))
#

#  print(llm(similarity_prompt.format(noun=my_noun), max_tokens=10))

# 1.2.3 Output Parser,用于格式化模型的输出

# 1.2.3.1 格式化指令，format instructions
# 1.2.3.2 解析器, parser


response_schemas = [
        ResponseSchema(name="bad_string", description="This is a poorly formatted user input string"),
        ResponseSchema(name="good_string", description="This is a well formatted response string"),
        ]

output_parser = StructuredOutputParser.from_response_schemas(response_schemas)
format_instruction = output_parser.get_format_instructions()
print(format_instruction)

template  = """
you will be given a poorly formatted string, and you will need to format it correctly.

{formatted_instruction}

% User Input: 
{bad_input}

Your answer: 
"""

prompt = PromptTemplate(
        input_variables=["bad_input"],
        partial_variables={"formatted_instruction": format_instruction},
        template=template,
        )

promptValue  = prompt.format(bad_input="welcome to china")

print(promptValue)

llm_output = llm(promptValue)
print(llm_output)

output_parser.parse(llm_output)


