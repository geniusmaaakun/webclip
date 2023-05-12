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
import { MarkdownProvider } from "../hooks/providers/useMarkdownsProvider";
import { Editor } from "../components/parts/Editor";

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

/*
jest.mockはJestのモック機能を提供する関数で、指定したモジュールの挙動を書き換えるために使用されます。これにより、テスト中にコントロールが難しい外部依存関係や副作用を排除し、テスト対象のコードの挙動をより純粋に検証できます。

例えば、次のようなコードがあるとします：

jsx
Copy code
import externalLibrary from 'external-library';

export function myFunction() {
  return externalLibrary.someFunction();
}
上記のmyFunctionはexternal-libraryに依存しています。external-library.someFunctionの挙動をコントロールできない場合、myFunctionのテストは難しくなるでしょう。

ここでjest.mockを使うと、以下のようにexternal-libraryをモック化し、someFunctionの挙動をテストの中で自由に設定できます：

jsx
Copy code
jest.mock('external-library', () => ({
  someFunction: jest.fn().mockReturnValue('mocked value'),
}));

import { myFunction } from './myFunction';

it('returns the mocked value', () => {
  expect(myFunction()).toBe('mocked value');
});
このようにjest.mockを使うことで、テスト対象のコードが依存する外部のモジュールをモック化し、その挙動をコントロールすることができます。これによりテストの信頼性と再現性が向上します。
*/
jest.mock('react-simplemde-editor', () => (props) => (
    <textarea data-testid="mock-editor" onChange={(e) => props.onChange(e.target.value)} value={props.value} />
));


describe('<Editor />', () => {
    it('It renders the recipe', async () => {
        render(
            <Editor id={"1"}/>, { wrapper: MarkdownProvider }
        );
        //screen.debug();
        
        expect(await screen.findByText(/## h2/));
        
    });
});