import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Search } from './components/page/Search';
import MarkdownEditor  from './components/page/Markdown';

function App() {
  return (
    <div className="App">
      <Search />
      <MarkdownEditor />
    </div>
  );
}

export default App;
