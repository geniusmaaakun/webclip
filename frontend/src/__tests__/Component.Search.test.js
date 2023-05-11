import React from "react";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import { rest } from "msw";
import { setupServer } from "msw/node";
import { useMarkdowns } from "../hooks/providers/useMarkdownsProvider";
import { useLoadMarkdown, useLoadMarkdowns } from "../hooks/markdowns/useMarkdowns";
import { Search } from '../components/page/Search';
import { MarkdownProvider } from "../hooks/providers/useMarkdownsProvider";
//useEffectのテスト

const server = setupServer(
    //サーバーURLを構築
    rest.get("http://localhost:8080/api/markdowns", (req, res, ctx) => {
        return res(ctx.status(200), ctx.json([{id: 1, title: "React dummy", content: "## h2",path: "test/README.md", srcurl: "http://test.com", created_at: "2021-01-01"}]));
    })
);

/////////
//前処理
//モックサーバー起動
beforeAll(() => server.listen());

//テストケース後の後処理
//テストケースが終わるたびに実行される
//リセット
afterEach(() => {
  server.resetHandlers();
});
//後処理
//全て終わったら閉じる リソース解放
afterAll(() => server.close());
/////////

const mockedUsedNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
   ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate,
}));
  

describe("Search Component", () => {
    it("should be able to parse markdown", async() => {
        //providerのテスト
        render(<Search />, { wrapper: MarkdownProvider });
        //stateの確認
        expect(screen.queryByText(/React dummy/)).toBeNull();
        expect(await screen.findByText(/React dummy/)).toBeInTheDocument();

        screen.debug();        
    });
});
