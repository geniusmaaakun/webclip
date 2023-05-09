/*
import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
*/

import React from "react";
import { render, screen } from "@testing-library/react";
import App from "../App";

test("renders learn react link", () => {
  render(<App />);

  //デバッグ情報を表示
  screen.debug();
  expect(screen.getByText(/WebClip/i)).toBeInTheDocument();
  //存在するか
  //screen.debug(screen.getByText(/WebClip/i)); //h1タグを取得している
  expect(screen.getByText(/WebClip/i)).toBeTruthy();
});
