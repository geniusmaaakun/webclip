import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Search } from './components/page/Search';
import { MarkdownEditor }  from './components/page/Markdown';
import { MarkdownProvider } from './hooks/providers/useMarkdownsProvider';
import { BrowserRouter, BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { AppRouter} from "./router/Router"

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <AppRouter />
      </BrowserRouter>
    </div>
  );
}

export default App;
