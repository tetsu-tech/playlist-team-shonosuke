import './App.css';
import { ChatUI } from './Chat/ChatUI';
import { Header } from "./Header/Header";

function App() {
  return (
    <div className="App">
      <Header />
      <ChatUI />
    </div>
  );
}

export default App;
