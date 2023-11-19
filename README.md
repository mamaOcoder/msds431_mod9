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
I used the free version of ChatGPT to generate a version of the code week 2's assignment. Some issues that I ran into was that ChatGPT did not seem to know some of the syntax for the Go stat's package that we were using. This could be due to updates to the package that was made after ChatGPT was last trained. The results of the initial code that was created was also creating incorrect results. Continuing conversation with GPT to see if we can achieve better results. Below is my initial prompt:

> Help me write a program in Go using github.com/montanaflynn/stats package to estimate linear regression coefficients for the Anscombe quartet data. We are looking to replicate the following Python code in Go: [copy of Python starter code]

Please see gpt/script.txt for the complete conversation with the LLM agent.

#### References
[32 ChatGPT Tips for Beginners in 2023!](https://www.youtube.com/watch?v=dUjWMdR_Kw8)

## Conculsion
Although there are definite benefits to using automated programming, it ...

It is important to note that GitHub Copilot and ChatGPT may also be trained on "bad code." Professional developers will still be required on staff to guide the models and decide which suggestions to accept and which to discard.