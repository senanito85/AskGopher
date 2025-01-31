### 🦫 AskGopher: Chat with Google’s Gemini AI in Your Terminal

AskGopher is a minimalist CLI tool that brings the power of Google’s Gemini AI directly to your terminal. Whether you’re debugging, brainstorming, or seeking intelligent answers, AskGopher lets you stay in the flow of your work—no browser required.

### ✨ Key Features
	•	Minimalist Interface: Clean and distraction-free experience.
	•	Powered by Google Gemini AI: Tap into state-of-the-art AI capabilities.
	•	Instant Responses: Get answers to your queries in real time.
	•	Browser-Free: No need to switch between applications—everything happens in the terminal.
	•	Blazing Fast: Built with Go for optimized performance.

___

🚀 Getting Started

##### 1. Prerequisites

Before you begin, ensure you have:
	•	Go (1.23)
	•	A Google Cloud account
	•	A Google API key with the Gemini API enabled

#####  2. Clone the Repository

```
git clone https://github.com/nsahil992/AskGopher.git  
cd AskGopher  
```

#####  3. Install Dependencies
```
go mod tidy  
```

#####  4. Set Up Google API Key
```
Create a .env file in the root directory:

touch .env  
```

2.	Add your Google API key to the .env file:

```
GOOGLE_API_KEY=your-google-api-key  
```

##### 5. Build and Run

Build the CLI tool:
```
go build -o askgopher  
```

Run the application:
```
./askgopher  
```
___

##### 🛠️ How to Use
	1.	Run the askgopher binary in your terminal.
	2.	Type your questions when prompted, and get instant responses from the Gemini AI!
	3.	Type quit to exit the application.

Example:

Ask Gopher! 🦫 
------------------  
💭 Ask something (or 'quit' to end): What is Go programming language?  

🤖 Answering:  
Go, also known as Golang, is an open-source programming language created by Google...  

___ 

⭐ Support the Project

If you find AskGopher helpful, consider giving it a ⭐ on GitHub! Your support keeps me motivated to build more innovative tools and share exciting projects.

##### 📝 Blog Coming Soon

I’ll be publishing a detailed blog on AskGopher. Stay tuned for insights!

#####  📧 Stay Connected
	•	Follow me on GitHub for more projects.
	•	Connect on LinkedIn.

##### ⚠️ Disclaimer

AskGopher requires a valid Google API key with access to the Gemini API. Please ensure you adhere to Google’s API usage policies to avoid issues.
