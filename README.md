# Week 9 Assignment: An Exercise in Automated Programming

## Project Summary
This week we revisit week 2 assignment, Testing Go for Statistics, to explore how different automated programming methods can be used to complete or improve the code.

### Automated Code Generation
The Go standard library has a tool called 'go generate' for automating code generation. It can be useful for automating repetive tasks in Go code. 'go generate' works by scanning for special comments (//go:generate ...) in Go source code that identify general commands to run. The specified command in the //go:generate comment can be any valid shell command or script. The specified command is run whenever the 'go generate' command is invoked.

Some benefits for using go generate:
1. Simplify your build process - when a project requires generating code based on some external specifications or configurations, go generate can automate this process.
2. Improve documentation - use a tool like godoc to generate HTML documentation for your project or for generating API documentation.
3. Enhance testing - generate test files using the [gotests](https://pkg.go.dev/github.com/cweill/gotests@v1.6.0#section-readme) tool. 
4. Increase code quality -  automatically run a linter on your codebase.
5. Streamline collaboration - enforce tools and version consistency.

There are a number of tools available that are designed to be run by 'go generate', such as [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer?utm_source=godoc), [jsonenums](https://github.com/campoy/jsonenums), and [schematyper](https://github.com/idubinskiy/schematyper).  

One way that 'go generate' could be used specifically to improve week 2's assignment, Testing Go for Statistics, would be to use [gotests](https://pkg.go.dev/github.com/cweill/gotests) to automatically generate unit tests. gotests analyzes Go code and generates corresponding test code. I am running into issues with getting gotests to work. Says it cannot find the package even after I installed it. Will continue to look into this.

#### References
[Generating Code](https://go.dev/blog/generate)  
[Introduction to 'go generate'](https://medium.com/@hsleep/introduction-to-go-generate-99a93f30dc35)

### AI-Assisted Programming
GitHub Copilot is a tool that assists developers by suggesting code completions based on the context of their code. It is an automated pair programming environment that allows developers to focus on coding decisions rather than the syntax itself. It's powered by OpenAI's Codex, which is a language model trained on a diverse range of publicly available code. GitHub Copilot can be a helpful tool for generating code snippets, especially for common patterns and tasks. A recent update added two new features, Github Copilot Chat and Github Copilot Voice. Copilot Chat is an IDE extension that allows chatting with AI about your code, ask questions, generate unit tests, fix bug, refactor, explain the parts in code you don’t understand etc. Copilot Voice is a voice assistant that allows writing code without typing.

The Testing Go for Statistics assignment was one of the first programs that I had written in Go. Having a tool such as GitHub Copilot could have assisted me in writing the code to ensure proper syntax and more. Looking back at my code now that I have a bit more experience in Go, I see ways that the code can be improved, for example, making it more modular. Because Copilot runs in your IDE, it has access to your whole codebase, so it tries to infer from your comments and existing code what you are trying to accomplish. Copilot was trained on a data set that includes a large concentration of public source code, including Go code, and could have helped get me to a better version of the code quicker. The Copilot Chat interface can provide explanations of what code blocks are intended to do, generate unit tests and propose fixes to bugs. This could really improve the learning curve for developers learning a new language.

#### References
[GitHub Copilot](https://github.com/features/copilot)  
[Introduction to GitHub Copilot](https://learn.microsoft.com/en-us/training/modules/introduction-to-github-copilot/)  
[Why GitHub Copilot Chat Is Better For Developers Than ChatGPT](https://medium.com/dare-to-be-better/why-github-copilot-chat-is-better-for-developers-than-chatgpt-0cd2930e3290)

### AI-Generated Code
I used the free version of ChatGPT to generate a version of the code for week 2's assignment. I approached this by trying to see if ChatGPT could really generate correct code from scratch, giving just the basic prompt to start with. 

Initial prompt:

> Help me write a program in Go using github.com/montanaflynn/stats package to estimate linear regression coefficients for the Anscombe quartet data. We are looking to replicate the following Python code in Go: [copy of Python starter code]

The first generated code did not work, however, provided a good starting point for building the code. The first issue that we ran into was that ChatGPT did not know the expected input type for the LinearRegression package. It attempted to send 2 slices with the X and Y coordinates, rather than the a 'stats.Series' type. In response, I provided ChatGPT with the error, as well as the telling ChatGPT what the expected input type for stats.LinearRegression should be based on the documentation. I continued with this approach of providing the errors until we reached a version of the code that compiled and ran (5 prompts). 

The first working code did not produce results that matched the output of Python and R. It took many more prompts to reach a version of the code that returned the correct results. I had to eventually provide a copy of the stats.LinearRegression function from the documentation to show ChatGPT that the function did not return the coeffients. I first tried to example that using my own words but ChatGPT did not understand until I provided an actual copy of the code.

Next, I prompted ChatGPT to create use the Go testing package to test and benchmark the code. The first response only creating test functions that needed some modifications. I do note that the tests that were produced were more thorough than what I wrote myself for the week 2 assignment.

Please see [this link](https://chat.openai.com/share/6a127960-2671-4eb0-baeb-76be2e5191b0) or gpt/script.txt for the complete conversation with the LLM agent.

Overall, it took some work to get ChatGPT to produce the expected results. It is important to keep in mind that ChatGPT was last updated in January 2022, so updates to packages that we use after that time may be the cause for some issues. That should not have been the case here since the LinearRegression function in the stats package was last updated 5 years ago. None-the-less, we still needed to provide a copy of the function so that ChatGPT could see what was expected. I intentionally started out somewhat vague to test what ChatGPT can do. Providing more guidance and background on the codebase we used may have led to 

#### References
[32 ChatGPT Tips for Beginners in 2023!](https://www.youtube.com/watch?v=dUjWMdR_Kw8)

## Conculsion
Although there are definite benefits to using automated programming, it ...

It is important to note that GitHub Copilot and ChatGPT may also be trained on "bad code." Professional developers will still be required on staff to guide the models and decide which suggestions to accept and which to discard.