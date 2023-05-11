/*
import { MarkdownEditor } from '../components/page/Markdown';
import { BrowserRouter } from "react-router-dom";
import userEvent from "@testing-library/user-event";
import { rest } from "msw";
import { setupServer } from "msw/node";
import { render, screen } from "@testing-library/react";
import { MarkdownProvider } from "../hooks/providers/useMarkdownsProvider";
import { Router, MemoryRouter, Route, Routes } from 'react-router-dom';
import { createMemoryHistory } from 'history';
import { useLoadMarkdowns } from '../hooks/markdowns/useMarkdowns';

const handler = [
    rest.get("http://localhost:8080/api/markdowns", (req, res, ctx) => {
        return res(ctx.status(200), ctx.json([{id: 1, title: "React dummy", content: "## h2",path: "test/README.md", srcurl: "http://test.com", created_at: "2021-01-01"}]));
    }),
    rest.get("http://localhost:8080/api/markdowns/1", (req, res, ctx) => {
        return res(ctx.status(200), ctx.json({id: 1, title: "React dummy", content: "## h2",path: "test/README.md", srcurl: "http://test.com", created_at: "2021-01-01"}));
    })
]

const server = setupServer(
    ...handler
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


// jest.mock('react-router-dom', () => ({
//    ...jest.requireActual('react-router-dom'),
//   useParams: () => ({
//     id: '1'
//   })
// }));

describe('<MarkdownEditor />', () => {
    it('It renders the recipe', async () => {
        const { loadMarkdowns } = useLoadMarkdowns();
        loadMarkdowns();
        const history = createMemoryHistory();
        //const id = '1';
        const route = `/markdowns/1`;
        history.push(route);

        render(
            <Router location={history.location} navigator={history}>
                <Routes>
                    <Route path="/markdowns/:id" element={<MarkdownEditor />} />
                </Routes>
            </Router>, { wrapper: MarkdownProvider }
        );
        //await waitForElementToBeRemoved(() => screen.queryByText('Loading...'));

        //await screen.debug();
        
        expect(await screen.findByText(/h2/));
        
    });
});

*/

import { rest } from "msw";
import { setupServer } from "msw/node";
import { AppRouter } from "../router/Router";
import { render, screen } from "@testing-library/react";
import { BrowserRouter, Router, MemoryRouter, Route, Routes } from 'react-router-dom';


const handler = [
    rest.get("http://localhost:8080/api/markdowns", (req, res, ctx) => {
        return res(ctx.status(200), ctx.json([{id: 1, title: "React dummy", content: "## h2",path: "test/README.md", srcurl: "http://test.com", created_at: "2021-01-01"}]));
    }),
    rest.get("http://localhost:8080/api/markdowns/1", (req, res, ctx) => {
        return res(ctx.status(200), ctx.json({id: 1, title: "React dummy", content: "## h2",path: "test/README.md", srcurl: "http://test.com", created_at: "2021-01-01"}));
    })
]

const server = setupServer(
    ...handler
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


describe('<AppRouter />', () => {
    it('It renders the recipe', async () => {
        render(
            <MemoryRouter initialEntries={[`/markdowns/1`]}><AppRouter /></MemoryRouter>
        );
        //await waitForElementToBeRemoved(() => screen.queryByText('Loading...'));

        //await screen.debug();
        
        expect(await screen.findByText(/h2/));
        
    });
});