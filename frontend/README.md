# axios jest
npm i -D --exact jest-watch-typeahead@0.6.5

```
package.json

"test": "react-scripts test --transformIgnorePatterns \"node_modules/(?!axios)/\"",

"jest": {
"transformIgnorePatterns": ["node_modules/(?!axios)/"]
}
```

```
node:internal/process/promises:227
          triggerUncaughtException(err, true /* fromPromise */);
          ^

[UnhandledPromiseRejection: This error originated either by throwing inside of an async function without a catch block, or by rejecting a promise which was not handled with .catch(). The promise rejected with the reason "CanceledError: canceled".] {
  code: 'ERR_UNHANDLED_REJECTION'
}
ChatGPT
このエラーは、非同期関数内でキャッチされていない例外がスローされたり、処理されていないプロミスが拒否されたりしたために発生します。このエラーは、CanceledErrorがキャッチされていないことを示しています。

エラーがどこで発生しているか特定するには、テストコードを確認し、非同期関数の呼び出しやプロミスのチェーンに対して適切なエラーハンドリングが実装されていることを確認してください。

例えば、async関数内でエラーがスローされる場合、try-catchブロックを使用して例外をキャッチし、適切に処理する必要があります。

javascript
Copy code
async function someAsyncFunction() {
  try {
    // 非同期処理を行うコード
  } catch (error) {
    // エラー処理を行うコード
  }
}
また、プロミスチェーンでエラーが発生する場合は、.catch()を使用してエラーを処理する必要があります。

javascript
Copy code
somePromiseFunction()
  .then(result => {
    // 成功時の処理
  })
  .catch(error => {
    // エラー処理
  });
これらのエラーハンドリング手法を適切に実装して、未処理の例外やプロミスの拒否が発生しないようにしてください。それでもエラーが解決しない場合は、問題が発生しているコードの詳細をお知らせください。


```

# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).
